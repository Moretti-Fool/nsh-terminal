package executor

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type ServiceDef struct {
	Name    string
	Dir     string
	Setup   string
	Command string
}

type ServiceProcess struct {
	Name string
	Cmd  *exec.Cmd
	Done chan struct{}
}

type ServiceRunner struct {
	processes []*ServiceProcess
	mu        sync.Mutex
	colors    []string
}

func NewServiceRunner() *ServiceRunner {
	return &ServiceRunner{
		colors: []string{
			"\033[32m", // green
			"\033[33m", // yellow
			"\033[34m", // blue
			"\033[35m", // magenta
			"\033[36m", // cyan
			"\033[91m", // bright red
			"\033[92m", // bright green
			"\033[93m", // bright yellow
		},
	}
}

func (sr *ServiceRunner) LaunchAll(services []ServiceDef) error {
	maxNameLen := 0
	for _, s := range services {
		if len(s.Name) > maxNameLen {
			maxNameLen = len(s.Name)
		}
	}

	for i, svc := range services {
		color := sr.colors[i%len(sr.colors)]
		reset := "\033[0m"
		prefix := fmt.Sprintf("%s%-*s%s | ", color, maxNameLen, svc.Name, reset)

		dir := svc.Dir
		if strings.HasPrefix(dir, "~") {
			home, _ := os.UserHomeDir()
			dir = home + dir[1:]
		}
		absDir, err := filepath.Abs(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s[error] bad directory: %s\n", prefix, err)
			continue
		}
		if _, err := os.Stat(absDir); err != nil {
			fmt.Fprintf(os.Stderr, "%s[error] directory not found: %s\n", prefix, absDir)
			continue
		}

		fullCmd := svc.Command
		if svc.Setup != "" {
			if runtime.GOOS == "windows" {
				fullCmd = svc.Setup + " ; " + svc.Command
			} else {
				fullCmd = svc.Setup + " && " + svc.Command
			}
		}

		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", "-NoProfile", "-Command", fullCmd)
		} else {
			shell := os.Getenv("SHELL")
			if shell == "" {
				shell = "/bin/sh"
			}
			cmd = exec.Command(shell, "-c", fullCmd)
		}
		cmd.Dir = absDir
		cmd.Env = os.Environ()

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s[error] stdout pipe: %s\n", prefix, err)
			continue
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s[error] stderr pipe: %s\n", prefix, err)
			continue
		}

		done := make(chan struct{})
		sp := &ServiceProcess{Name: svc.Name, Cmd: cmd, Done: done}

		go func(p string) {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Printf("%s%s\n", p, scanner.Text())
			}
		}(prefix)

		go func(p string) {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				fmt.Printf("%s%s\n", p, scanner.Text())
			}
		}(prefix)

		if err := cmd.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "%s[error] start failed: %s\n", prefix, err)
			continue
		}

		fmt.Printf("%s[started] %s\n", prefix, svc.Command)

		go func(c *exec.Cmd, d chan struct{}, p string, name string) {
			err := c.Wait()
			if err != nil {
				fmt.Printf("%s[exited] %v\n", p, err)
			} else {
				fmt.Printf("%s[exited] clean\n", p)
			}
			close(d)
		}(cmd, done, prefix, svc.Name)

		sr.mu.Lock()
		sr.processes = append(sr.processes, sp)
		sr.mu.Unlock()
	}

	return nil
}

func (sr *ServiceRunner) Wait() {
	sr.mu.Lock()
	procs := make([]*ServiceProcess, len(sr.processes))
	copy(procs, sr.processes)
	sr.mu.Unlock()

	for _, p := range procs {
		<-p.Done
	}
}

func (sr *ServiceRunner) StopAll() {
	sr.mu.Lock()
	procs := make([]*ServiceProcess, len(sr.processes))
	copy(procs, sr.processes)
	sr.mu.Unlock()

	for _, p := range procs {
		if p.Cmd.Process != nil {
			if runtime.GOOS == "windows" {
				exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprintf("%d", p.Cmd.Process.Pid)).Run()
			} else {
				p.Cmd.Process.Kill()
			}
		}
	}
}

func (sr *ServiceRunner) Running() []string {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	var names []string
	for _, p := range sr.processes {
		select {
		case <-p.Done:
		default:
			names = append(names, p.Name)
		}
	}
	return names
}
