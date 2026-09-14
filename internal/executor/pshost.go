package executor

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"unicode/utf16"
)

const psReadyTimeout = 20 * time.Second
const psCommandTimeout = 2 * time.Minute

const psHostScript = `
$ErrorActionPreference = 'Continue'
$ProgressPreference = 'SilentlyContinue'
$ConfirmPreference = 'None'
try {
  [Console]::OutputEncoding = New-Object System.Text.UTF8Encoding $false
} catch {}
$OutputEncoding = [System.Text.Encoding]::UTF8
$inbox = $env:NSH_PS_INBOX

function Write-NshDone([string]$Id, [int]$Code) {
  [Console]::Out.WriteLine()
  [Console]::Out.WriteLine("NSH_DONE|$Id|$Code")
  [Console]::Out.Flush()
}

[Console]::Out.WriteLine("NSH_READY")
[Console]::Out.Flush()

while ($true) {
  if (-not $inbox) { break }
  if (-not [System.IO.File]::Exists($inbox)) {
    Start-Sleep -Milliseconds 50
    continue
  }
  try {
    $line = [System.IO.File]::ReadAllText($inbox)
    [System.IO.File]::Delete($inbox)
  } catch {
    Start-Sleep -Milliseconds 20
    continue
  }
  $line = $line.Trim()
  if ($line -eq 'NSH_QUIT') { break }
  if (-not $line.StartsWith('NSH_RUN|')) { continue }

  $parts = $line.Split('|', 4)
  if ($parts.Count -lt 4) {
    Write-NshDone "unknown" 1
    continue
  }
  $id = $parts[1]
  try {
    $cwd = [Text.Encoding]::UTF8.GetString([Convert]::FromBase64String($parts[2]))
    $command = [Text.Encoding]::UTF8.GetString([Convert]::FromBase64String($parts[3]))
  } catch {
    [Console]::Error.WriteLine($_.Exception.Message)
    Write-NshDone $id 1
    continue
  }

  try {
    Set-Location -LiteralPath $cwd
  } catch {
    [Console]::Error.WriteLine($_.Exception.Message)
    Write-NshDone $id 1
    continue
  }

  $global:LASTEXITCODE = 0
  try {
    Invoke-Expression $command 2>&1 | Out-Default
    $code = 0
    if (-not $?) { $code = 1 }
    if ($null -ne $global:LASTEXITCODE -and $global:LASTEXITCODE -ne 0) {
      $code = [int]$global:LASTEXITCODE
    }
    Write-NshDone $id $code
  } catch {
    [Console]::Error.WriteLine($_.Exception.Message)
    Write-NshDone $id 1
  }
}
`

type lineResult struct {
	line string
	err  error
}

type psHost struct {
	cmd    *exec.Cmd
	inbox  string
	lines  chan lineResult
	mu     sync.Mutex
	ready  bool
	seq    uint64
	errBuf *lockedBuffer
}

type lockedBuffer struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (l *lockedBuffer) Write(p []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	n, err := l.b.Write(p)
	return n, err
}

func (l *lockedBuffer) StringAndReset() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	s := l.b.String()
	l.b.Reset()
	return s
}

func encodeUTF16LEBase64(s string) string {
	u := utf16.Encode([]rune(s))
	raw := make([]byte, len(u)*2)
	for i, v := range u {
		raw[i*2] = byte(v)
		raw[i*2+1] = byte(v >> 8)
	}
	return base64.StdEncoding.EncodeToString(raw)
}

func startPSHost() (*psHost, error) {
	exe := "powershell.exe"
	if p, err := exec.LookPath("pwsh"); err == nil {
		exe = p
	} else if p, err := exec.LookPath("powershell.exe"); err == nil {
		exe = p
	}

	inbox := filepath.Join(os.TempDir(), fmt.Sprintf("nsh-ps-in-%d-%d.txt", os.Getpid(), time.Now().UnixNano()))
	_ = os.Remove(inbox)

	cmd := exec.Command(exe,
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-EncodedCommand", encodeUTF16LEBase64(psHostScript),
	)
	cmd.Env = append(os.Environ(), "NSH_PS_INBOX="+inbox)
	hideWindow(cmd)
	devNull, err := os.Open(os.DevNull)
	if err == nil {
		cmd.Stdin = devNull
		defer devNull.Close()
	}
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	errBuf := &lockedBuffer{}
	cmd.Stderr = errBuf

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	h := &psHost{
		cmd:    cmd,
		inbox:  inbox,
		lines:  make(chan lineResult, 64),
		errBuf: errBuf,
	}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 4*1024*1024)
		for scanner.Scan() {
			h.lines <- lineResult{line: scanner.Text()}
		}
		err := scanner.Err()
		if err == nil {
			err = io.EOF
		}
		h.lines <- lineResult{err: err}
		close(h.lines)
	}()

	timer := time.NewTimer(psReadyTimeout)
	defer timer.Stop()
	for {
		select {
		case got, ok := <-h.lines:
			if !ok || got.err != nil {
				_ = h.killUnlocked()
				if got.err != nil {
					return nil, fmt.Errorf("powershell host failed to start: %w", got.err)
				}
				return nil, fmt.Errorf("powershell host failed to start")
			}
			if strings.TrimSpace(got.line) == "NSH_READY" {
				h.ready = true
				return h, nil
			}
		case <-timer.C:
			_ = h.killUnlocked()
			return nil, fmt.Errorf("timeout waiting for powershell host")
		}
	}
}

func (h *psHost) run(command, cwd string) (RunResult, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.cmd == nil || h.cmd.Process == nil || !h.ready {
		return RunResult{}, fmt.Errorf("powershell host not running")
	}

	start := time.Now()
	h.seq++
	id := fmt.Sprintf("%d-%d", h.cmd.Process.Pid, h.seq)

	msg := fmt.Sprintf("NSH_RUN|%s|%s|%s\n",
		id,
		base64.StdEncoding.EncodeToString([]byte(cwd)),
		base64.StdEncoding.EncodeToString([]byte(command)),
	)
	if err := h.writeInbox(msg); err != nil {
		return RunResult{}, err
	}

	marker := "NSH_DONE|" + id + "|"
	var outBuf strings.Builder
	timer := time.NewTimer(psCommandTimeout)
	defer timer.Stop()

	for {
		select {
		case got, ok := <-h.lines:
			if !ok {
				_ = h.killUnlocked()
				return RunResult{Output: outBuf.String() + h.errBuf.StringAndReset(), ExitCode: 1, DurationMs: time.Since(start).Milliseconds()}, fmt.Errorf("powershell host closed")
			}
			if got.err != nil {
				_ = h.killUnlocked()
				return RunResult{Output: outBuf.String() + h.errBuf.StringAndReset(), ExitCode: 1, DurationMs: time.Since(start).Milliseconds()}, got.err
			}
			if idx := strings.Index(got.line, marker); idx != -1 {
				codeStr := strings.TrimSpace(got.line[idx+len(marker):])
				code := 0
				fmt.Sscanf(codeStr, "%d", &code)
				stderr := h.errBuf.StringAndReset()
				output := outBuf.String()
				// Add whatever preceded the marker to output
				if idx > 0 {
					output += got.line[:idx] + "\n"
				}
				if stderr != "" {
					output += stderr
				}
				return RunResult{Output: output, ExitCode: code, DurationMs: time.Since(start).Milliseconds()}, nil
			}
			if got.line == "NSH_READY" {
				continue
			}
			if strings.HasPrefix(got.line, "#<") || strings.HasPrefix(got.line, "<Objs ") {
				continue
			}
			fmt.Fprintln(os.Stdout, got.line)
			outBuf.WriteString(got.line)
			outBuf.WriteByte('\n')
		case <-timer.C:
			_ = h.killUnlocked()
			return RunResult{Output: outBuf.String(), ExitCode: 1, DurationMs: time.Since(start).Milliseconds()}, fmt.Errorf("powershell command timed out")
		}
	}
}

func (h *psHost) writeInbox(msg string) error {
	if h.inbox == "" {
		return fmt.Errorf("powershell host inbox missing")
	}
	tmp := h.inbox + ".tmp"
	if err := os.WriteFile(tmp, []byte(msg), 0600); err != nil {
		return err
	}
	_ = os.Remove(h.inbox)
	return os.Rename(tmp, h.inbox)
}

func (h *psHost) killUnlocked() error {
	if h.inbox != "" {
		_ = os.Remove(h.inbox)
		_ = os.Remove(h.inbox + ".tmp")
	}
	if h.cmd != nil && h.cmd.Process != nil {
		_ = h.cmd.Process.Kill()
		_, _ = h.cmd.Process.Wait()
	}
	h.cmd = nil
	h.ready = false
	return nil
}

func (h *psHost) close() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.inbox != "" {
		_ = h.writeInbox("NSH_QUIT\n")
	}
	if h.cmd != nil && h.cmd.Process != nil {
		done := make(chan struct{})
		go func() {
			_, _ = h.cmd.Process.Wait()
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			_ = h.cmd.Process.Kill()
			<-done
		}
	}
	if h.inbox != "" {
		_ = os.Remove(h.inbox)
		_ = os.Remove(h.inbox + ".tmp")
	}
	h.cmd = nil
	h.ready = false
	return nil
}

func looksLikePowerShell(command string) bool {
	lower := strings.ToLower(command)
	if strings.Contains(command, "$_") || strings.Contains(command, "${") {
		return true
	}
	markers := []string{
		"get-childitem", "select-object", "sort-object", "where-object",
		"foreach-object", "format-table", "format-list", "measure-object",
		"get-process", "stop-process", "get-service", "restart-service",
		"get-nettcpconnection", "get-content", "set-content", "out-file",
		"invoke-webrequest", "invoke-restmethod", "convertto-json",
		"select-string", "expand-archive", "compress-archive",
		"get-netadapter", "get-date", "write-output", "write-host",
		"set-location", "get-location", "get-item", "remove-item",
		"copy-item", "move-item", "new-item", "test-path",
	}
	for _, m := range markers {
		if strings.Contains(lower, m) {
			return true
		}
	}
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return false
	}
	first := fields[0]
	if i := strings.IndexByte(first, '-'); i > 0 && i < len(first)-1 {
		verb := first[:i]
		switch strings.ToLower(verb) {
		case "get", "set", "new", "remove", "start", "stop", "restart",
			"out", "write", "select", "where", "sort", "format", "measure",
			"invoke", "convertto", "convertfrom", "import", "export",
			"add", "clear", "copy", "move", "rename", "test", "wait":
			return true
		}
	}
	return false
}

func looksLikeUnknownCommand(output string) bool {
	lower := strings.ToLower(output)
	needles := []string{
		"is not recognized",
		"commandnotfoundexception",
		"not recognized as the name of a cmdlet",
		"not recognized as an internal or external command",
		"could not find command",
		"the term '",
	}
	for _, n := range needles {
		if strings.Contains(lower, n) {
			return true
		}
	}
	return false
}
