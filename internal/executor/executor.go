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
	"sort"
	"strings"
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
}

func New(shellPref string, wslDistro string) *Executor {
	return &Executor{shellPref: shellPref, wslDistro: wslDistro}
}

func (e *Executor) DetectShell() (string, []string) {
	if e.shellPref != "auto" && e.shellPref != "" {
		return e.resolveExplicit(e.shellPref)
	}
	if runtime.GOOS == "windows" {
		if path, err := exec.LookPath("pwsh"); err == nil {
			return path, []string{"-NoProfile", "-Command"}
		}
		return "powershell.exe", []string{"-NoProfile", "-Command"}
	}
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	return shell, []string{"-c"}
}

func (e *Executor) fastShell(command string) (string, []string) {
	if runtime.GOOS != "windows" {
		return e.DetectShell()
	}
	tokens := strings.Fields(command)
	if len(tokens) == 0 {
		return e.DetectShell()
	}

	if hasPowerShellSyntax(command) {
		return e.DetectShell()
	}

	if _, err := exec.LookPath(tokens[0]); err == nil {
		return "cmd.exe", []string{"/C"}
	}
	return e.DetectShell()
}

func hasPowerShellSyntax(cmd string) bool {
	if strings.Contains(cmd, "$") {
		return true
	}
	if strings.Contains(cmd, "@(") || strings.Contains(cmd, "$(") {
		return true
	}
	tokens := strings.Fields(cmd)
	if len(tokens) > 0 && strings.Contains(tokens[0], "-") {
		parts := strings.SplitN(tokens[0], "-", 2)
		if len(parts) == 2 {
			verb := strings.ToLower(parts[0])
			for _, v := range []string{"get", "set", "new", "remove", "invoke", "start", "stop", "test", "write", "read", "select", "where", "foreach", "format", "out", "export", "import", "convertto", "convertfrom", "add", "clear", "copy", "move", "rename"} {
				if verb == v {
					return true
				}
			}
		}
	}
	return false
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

func (e *Executor) NeedsWSL(command string) bool {
	if runtime.GOOS != "windows" {
		return false
	}
	if _, err := exec.LookPath("wsl.exe"); err != nil {
		return false
	}
	tokens := strings.Fields(command)
	if len(tokens) == 0 {
		return false
	}
	cmd := strings.ToLower(tokens[0])

	unixOnly := map[string]bool{
		"grep": true, "awk": true, "sed": true, "tail": true, "head": true,
		"wc": true, "sort": true, "uniq": true, "tr": true, "cut": true,
		"find": true, "xargs": true, "chmod": true, "chown": true,
		"tar": true, "gzip": true, "gunzip": true, "df": true, "du": true,
		"top": true, "htop": true, "kill": true, "ps": true, "man": true,
		"touch": true, "ln": true, "diff": true, "patch": true,
		"ssh": true, "scp": true, "rsync": true, "wget": true,
	}
	if unixOnly[cmd] {
		return true
	}

	dualCommands := map[string]bool{
		"ls": true, "cat": true, "cp": true, "mv": true, "rm": true,
		"mkdir": true, "echo": true,
	}
	if dualCommands[cmd] && len(tokens) > 1 {
		for _, tok := range tokens[1:] {
			if len(tok) >= 2 && tok[0] == '-' && tok[1] != '-' {
				allLower := true
				for _, ch := range tok[1:] {
					if ch < 'a' || ch > 'z' {
						allLower = false
						break
					}
				}
				if allLower {
					return true
				}
			}
		}
	}

	return false
}

func (e *Executor) wslShell() (string, []string) {
	return "wsl.exe", []string{"--", "sh", "-c"}
}

func (e *Executor) TryBuiltin(command string) (RunResult, bool) {
	tokens := strings.Fields(command)
	if len(tokens) == 0 {
		return RunResult{}, false
	}

	switch strings.ToLower(tokens[0]) {
	case "ls":
		if len(tokens) == 1 {
			return builtinLs(".")
		}
		if len(tokens) == 2 && !strings.HasPrefix(tokens[1], "-") {
			return builtinLs(tokens[1])
		}
	case "dir":
		if len(tokens) == 1 {
			return builtinLs(".")
		}
		if len(tokens) == 2 && !strings.HasPrefix(tokens[1], "-") && !strings.HasPrefix(tokens[1], "/") {
			return builtinLs(tokens[1])
		}
	case "pwd":
		if len(tokens) == 1 {
			return builtinPwd()
		}
	case "clear", "cls":
		if len(tokens) == 1 {
			return builtinClear()
		}
	}
	return RunResult{}, false
}

func builtinLs(dir string) (RunResult, bool) {
	start := time.Now()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return RunResult{Output: err.Error() + "\n", ExitCode: 1, DurationMs: time.Since(start).Milliseconds()}, true
	}

	var dirs []string
	var files []string
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if e.IsDir() {
			dirs = append(dirs, name+"/")
		} else {
			files = append(files, name)
		}
	}
	sort.Strings(dirs)
	sort.Strings(files)

	var buf strings.Builder
	all := append(dirs, files...)
	if len(all) == 0 {
		return RunResult{DurationMs: time.Since(start).Milliseconds()}, true
	}

	maxLen := 0
	for _, name := range all {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	termWidth := 80
	colWidth := maxLen + 2
	cols := termWidth / colWidth
	if cols < 1 {
		cols = 1
	}

	for i, name := range all {
		if i > 0 && i%cols == 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(fmt.Sprintf("%-*s", colWidth, name))
	}
	buf.WriteByte('\n')

	output := buf.String()
	fmt.Print(output)
	return RunResult{Output: output, DurationMs: time.Since(start).Milliseconds()}, true
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

func (e *Executor) Run(command string) (RunResult, error) {
	if result, handled := e.TryBuiltin(command); handled {
		return result, nil
	}

	var shell string
	var args []string
	if e.NeedsWSL(command) {
		shell, args = e.wslShell()
	} else {
		shell, args = e.fastShell(command)
	}
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
