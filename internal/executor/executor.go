package executor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

type RunResult struct {
	Output     string
	ExitCode   int
	DurationMs int64
}

type Executor struct {
	shellPref string
	wslDistro string
	pathCache sync.Map

	psMu       sync.Mutex
	ps         *psHost
	psErr      error
	psStarting bool
	psWait     chan struct{}
}

func New(shellPref string, wslDistro string) *Executor {
	return &Executor{shellPref: shellPref, wslDistro: wslDistro}
}

// WarmPSHost starts the persistent PowerShell process in the background so the
// first NL command does not pay a 2–3s cold start.
func (e *Executor) WarmPSHost() {
	if runtime.GOOS != "windows" {
		return
	}
	go func() { _ = e.ensurePSHost() }()
}

func (e *Executor) Close() {
	e.psMu.Lock()
	h := e.ps
	e.ps = nil
	e.psMu.Unlock()
	if h != nil {
		_ = h.close()
	}
}

func (e *Executor) ensurePSHost() error {
	if runtime.GOOS != "windows" {
		return nil
	}

	e.psMu.Lock()
	if e.ps != nil && e.ps.ready {
		e.psMu.Unlock()
		return nil
	}
	if e.psStarting {
		wait := e.psWait
		e.psMu.Unlock()
		<-wait
		e.psMu.Lock()
		err := e.psErr
		e.psMu.Unlock()
		return err
	}
	e.psStarting = true
	e.psWait = make(chan struct{})
	e.psMu.Unlock()

	h, err := startPSHost()

	e.psMu.Lock()
	e.ps = h
	e.psErr = err
	e.psStarting = false
	close(e.psWait)
	e.psMu.Unlock()
	return err
}

func (e *Executor) dropPSHost() {
	e.psMu.Lock()
	h := e.ps
	e.ps = nil
	e.psErr = nil
	e.psMu.Unlock()
	if h != nil {
		_ = h.close()
	}
}

// RunGenerated executes an LLM-produced command. On Windows this uses the
// warm PowerShell host so cmdlets work and startup cost is paid once.
func (e *Executor) RunGenerated(command string) (RunResult, error) {
	if runtime.GOOS != "windows" {
		return e.Run(command)
	}
	if err := e.ensurePSHost(); err != nil {
		return e.runPowerShellOnce(command)
	}
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	e.psMu.Lock()
	h := e.ps
	e.psMu.Unlock()
	if h == nil {
		return e.runPowerShellOnce(command)
	}
	result, err := h.run(command, cwd)
	if err != nil {
		e.dropPSHost()
		if retryErr := e.ensurePSHost(); retryErr == nil {
			e.psMu.Lock()
			h = e.ps
			e.psMu.Unlock()
			if h != nil {
				return h.run(command, cwd)
			}
		}
		return e.runPowerShellOnce(command)
	}
	return result, nil
}

func (e *Executor) runPowerShellOnce(command string) (RunResult, error) {
	exe := "powershell.exe"
	if p, err := e.CachedLookPath("pwsh"); err == nil {
		exe = p
	}
	start := time.Now()
	cmd := exec.Command(exe, "-NoProfile", "-NonInteractive", "-Command", command)
	cmd.Env = os.Environ()
	hideWindow(cmd)
	var outBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &outBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &outBuf)
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
			err = nil
		}
	}
	return RunResult{Output: outBuf.String(), ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, err
}

// NLShellName is the dialect the generator should target (PowerShell on Windows).
func (e *Executor) NLShellName() string {
	if runtime.GOOS == "windows" {
		return "powershell"
	}
	return e.ShellName()
}

func (e *Executor) AvailableBins() []string {
	candidates := []string{
		"ipconfig", "netstat", "taskkill", "tasklist", "sc", "net",
		"ping", "tracert", "nslookup", "curl", "git", "docker",
		"npm", "node", "python", "python3", "pip", "ssh", "where",
		"findstr", "robocopy", "xcopy", "msiexec", "wmic",
		"powershell", "pwsh", "cmd", "winget", "choco",
		"dotnet", "go", "cargo", "make", "kubectl",
	}
	var found []string
	for _, name := range candidates {
		if e.PathExists(name) {
			found = append(found, name)
		}
	}
	return found
}

func LooksLikeUnknownCommand(output string) bool {
	return looksLikeUnknownCommand(output)
}

func (e *Executor) CachedLookPath(name string) (string, error) {
	if v, ok := e.pathCache.Load(name); ok {
		if v == nil {
			return "", exec.ErrNotFound
		}
		return v.(string), nil
	}
	p, err := exec.LookPath(name)
	if err != nil {
		e.pathCache.Store(name, nil)
		return "", err
	}
	e.pathCache.Store(name, p)
	return p, nil
}

func (e *Executor) PathExists(name string) bool {
	_, err := e.CachedLookPath(name)
	return err == nil
}

func (e *Executor) Run(command string) (RunResult, error) {
	if runtime.GOOS == "windows" && looksLikePowerShell(command) {
		return e.RunGenerated(command)
	}

	// Commands with pipes, redirects, or globs need a real shell —
	// builtins can't handle these, so check before TryBuiltin.
	if needsShell(command) {
		return e.runViaShell(command)
	}

	if result, handled := e.TryBuiltin(command); handled {
		return result, nil
	}

	return e.runDirect(command)
}

func (e *Executor) TryBuiltin(command string) (RunResult, bool) {
	tokens := tokenize(command)
	if len(tokens) == 0 {
		return RunResult{}, false
	}

	switch strings.ToLower(tokens[0]) {
	case "ls", "dir":
		return builtinLsFull(tokens)
	case "cat", "type":
		if len(tokens) >= 2 {
			return builtinCat(tokens)
		}
	case "head":
		return builtinHead(tokens)
	case "tail":
		return builtinTail(tokens)
	case "wc":
		return builtinWc(tokens)
	case "touch":
		return builtinTouch(tokens)
	case "mkdir":
		return builtinMkdir(tokens)
	case "cp", "copy":
		return builtinCp(tokens)
	case "mv", "move":
		return builtinMv(tokens)
	case "rm", "del":
		return builtinRm(tokens)
	case "echo":
		return builtinEcho(tokens)
	case "find":
		return builtinFind(tokens)
	case "grep":
		if len(tokens) >= 3 {
			return builtinGrep(tokens)
		}
	case "pwd":
		if len(tokens) == 1 {
			return builtinPwd()
		}
	case "clear", "cls":
		if len(tokens) == 1 {
			return builtinClear()
		}
	case "which", "where":
		if len(tokens) >= 2 {
			return e.builtinWhich(tokens)
		}
	}
	return RunResult{}, false
}

func needsShell(command string) bool {
	for _, ch := range []string{"|", ">>", "&&", "||", ";", "$(", "`"} {
		if strings.Contains(command, ch) {
			return true
		}
	}
	for i, ch := range command {
		if ch == '>' && (i == 0 || command[i-1] != '>') && (i+1 >= len(command) || command[i+1] != '>') {
			if i > 0 && command[i-1] == '2' {
				return true
			}
			return true
		}
	}
	if strings.ContainsAny(command, "*?") {
		return true
	}
	if runtime.GOOS == "windows" && strings.Contains(command, "$") {
		return true
	}
	return false
}

func (e *Executor) runDirect(command string) (RunResult, error) {
	tokens := tokenize(command)
	if len(tokens) == 0 {
		return RunResult{}, nil
	}

	binary, err := e.CachedLookPath(tokens[0])
	if err != nil {
		return e.runViaShell(command)
	}

	start := time.Now()
	cmd := exec.Command(binary, tokens[1:]...)
	cmd.Env = os.Environ()

	var outBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &outBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &outBuf)
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	duration := time.Since(start).Milliseconds()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
			err = nil
		}
	}

	return RunResult{Output: outBuf.String(), ExitCode: exitCode, DurationMs: duration}, err
}

func (e *Executor) runViaShell(command string) (RunResult, error) {
	shell, args := e.DetectShell()
	cmdArgs := append(args, command)
	start := time.Now()

	cmd := exec.Command(shell, cmdArgs...)
	cmd.Env = os.Environ()

	var outBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &outBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &outBuf)
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	duration := time.Since(start).Milliseconds()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
			err = nil
		}
	}

	return RunResult{Output: outBuf.String(), ExitCode: exitCode, DurationMs: duration}, err
}

func (e *Executor) DetectShell() (string, []string) {
	if e.shellPref != "auto" && e.shellPref != "" {
		return e.resolveExplicit(e.shellPref)
	}
	if runtime.GOOS == "windows" {
		return "cmd.exe", []string{"/C"}
	}
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	return shell, []string{"-c"}
}

func (e *Executor) ShellName() string {
	shell, _ := e.DetectShell()
	base := strings.ToLower(filepath.Base(shell))
	base = strings.TrimSuffix(base, ".exe")
	switch base {
	case "powershell", "pwsh":
		return "powershell"
	case "wsl":
		return "bash"
	default:
		return base
	}
}

func (e *Executor) resolveExplicit(name string) (string, []string) {
	switch strings.ToLower(name) {
	case "powershell":
		return "powershell.exe", []string{"-NoProfile", "-Command"}
	case "pwsh":
		if path, err := exec.LookPath("pwsh"); err == nil {
			return path, []string{"-NoProfile", "-Command"}
		}
		return "powershell.exe", []string{"-NoProfile", "-Command"}
	case "cmd":
		return "cmd.exe", []string{"/C"}
	case "wsl":
		return "wsl.exe", []string{"--", "sh", "-c"}
	case "bash":
		if runtime.GOOS == "windows" {
			return "wsl.exe", []string{"--", "sh", "-c"}
		}
		return "/bin/bash", []string{"-c"}
	case "zsh":
		return "/bin/zsh", []string{"-c"}
	default:
		return name, []string{"-c"}
	}
}

var destructivePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\brm\s+(-[a-z]*f|-[a-z]*r|--force|--recursive)`),
	regexp.MustCompile(`(?i)\brm\s+-rf\b`),
	regexp.MustCompile(`(?i)\bdel\s+/[sS]`),
	regexp.MustCompile(`(?i)\bformat\s+[A-Za-z]:`),
	regexp.MustCompile(`(?i)\bDROP\s+(TABLE|DATABASE)\b`),
	regexp.MustCompile(`(?i)\bmkfs\b`),
	regexp.MustCompile(`(?i)\bsudo\b`),
	regexp.MustCompile(`(?i)\brunas\b`),
	regexp.MustCompile(`(?i)\bSet-ExecutionPolicy\b`),
	regexp.MustCompile(`(?i)\bchmod\s+777\b`),
	regexp.MustCompile(`(?i)>\s*/dev/sd[a-z]`),
	regexp.MustCompile(`(?i)\bdd\s+.*of=/dev/`),
}

func (e *Executor) IsDestructive(cmd string) bool {
	for _, pat := range destructivePatterns {
		if pat.MatchString(cmd) {
			return true
		}
	}
	return false
}

func (e *Executor) RunSilent(command string) (RunResult, error) {
	shell, args := e.DetectShell()
	cmdArgs := append(args, command)
	start := time.Now()

	cmd := exec.Command(shell, cmdArgs...)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	duration := time.Since(start).Milliseconds()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
			err = nil
		}
	}

	return RunResult{Output: string(output), ExitCode: exitCode, DurationMs: duration}, err
}

func CdExpand(dir string) (string, error) {
	if strings.HasPrefix(dir, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dir = home + dir[1:]
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("invalid path: %s", dir)
	}
	info, err := os.Stat(abs)
	if err != nil {
		return "", fmt.Errorf("directory not found: %s", dir)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("not a directory: %s", dir)
	}
	return abs, nil
}

func builtinPwd() (RunResult, bool) {
	start := time.Now()
	cwd, err := os.Getwd()
	if err != nil {
		return RunResult{Output: err.Error() + "\n", ExitCode: 1, DurationMs: time.Since(start).Milliseconds()}, true
	}
	output := cwd + "\n"
	fmt.Print(output)
	return RunResult{Output: output, DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinClear() (RunResult, bool) {
	fmt.Print("\033[2J\033[H")
	return RunResult{}, true
}

func (e *Executor) builtinWhich(tokens []string) (RunResult, bool) {
	start := time.Now()
	var buf strings.Builder
	exitCode := 0
	for _, name := range tokens[1:] {
		p, err := e.CachedLookPath(name)
		if err != nil {
			msg := fmt.Sprintf("%s: not found\n", name)
			fmt.Print(msg)
			buf.WriteString(msg)
			exitCode = 1
		} else {
			line := p + "\n"
			fmt.Print(line)
			buf.WriteString(line)
		}
	}
	return RunResult{Output: buf.String(), ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}

func tokenize(command string) []string {
	var tokens []string
	var current strings.Builder
	inSingle := false
	inDouble := false

	for i := 0; i < len(command); i++ {
		ch := command[i]
		switch {
		case ch == '\'' && !inDouble:
			inSingle = !inSingle
		case ch == '"' && !inSingle:
			inDouble = !inDouble
		case ch == ' ' && !inSingle && !inDouble:
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		default:
			current.WriteByte(ch)
		}
	}
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	return tokens
}
