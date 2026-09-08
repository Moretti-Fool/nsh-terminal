package executor

import (
	"runtime"
	"strings"
	"testing"
)

func TestDetectShell(t *testing.T) {
	e := New("auto", "Ubuntu")
	shell, args := e.DetectShell()
	if runtime.GOOS == "windows" {
		if !strings.Contains(shell, "powershell") && !strings.Contains(shell, "pwsh") {
			t.Errorf("expected powershell, got %s", shell)
		}
		if args[0] != "-NoProfile" {
			t.Errorf("expected -NoProfile, got %v", args)
		}
	} else {
		if shell == "" {
			t.Error("empty shell")
		}
	}
}

func TestIsDestructive(t *testing.T) {
	e := New("auto", "Ubuntu")
	destructive := []string{"rm -rf /", "del /s /q folder", "format C:", "DROP TABLE users", "sudo rm file"}
	safe := []string{"ls -la", "git status", "docker ps", "cat file.txt"}

	for _, cmd := range destructive {
		if !e.IsDestructive(cmd) {
			t.Errorf("expected destructive: %s", cmd)
		}
	}
	for _, cmd := range safe {
		if e.IsDestructive(cmd) {
			t.Errorf("expected safe: %s", cmd)
		}
	}
}

func TestRunSimple(t *testing.T) {
	e := New("auto", "Ubuntu")
	result, err := e.Run("echo hello")
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.ExitCode != 0 {
		t.Errorf("exit code: %d", result.ExitCode)
	}
	if !strings.Contains(result.Output, "hello") {
		t.Errorf("output: %s", result.Output)
	}
}

func TestRunFailure(t *testing.T) {
	e := New("auto", "Ubuntu")
	result, err := e.Run("exit 1")
	if err != nil {
		t.Fatalf("should not error: %v", err)
	}
	if result.ExitCode == 0 {
		t.Error("expected non-zero exit")
	}
}

func TestCdExpand(t *testing.T) {
	_, err := CdExpand(".")
	if err != nil {
		t.Errorf("current dir should work: %v", err)
	}
	_, err = CdExpand("/nonexistent_dir_xyz")
	if err == nil {
		t.Error("should fail for nonexistent dir")
	}
}
