package executor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
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
		return "wsl.exe", []string{"-d", e.wslDistro, "--", "bash", "-c"}
	case "bash":
		if runtime.GOOS == "windows" {
			return "wsl.exe", []string{"-d", e.wslDistro, "--", "bash", "-c"}
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

func (e *Executor) Run(command string) (RunResult, error) {
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
	info, err := os.Stat(dir)
	if err != nil {
		return "", fmt.Errorf("directory not found: %s", dir)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("not a directory: %s", dir)
	}
	return dir, nil
}
