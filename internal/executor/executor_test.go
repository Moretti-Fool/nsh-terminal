package executor

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestDetectShell(t *testing.T) {
	e := New("auto", "")
	shell, args := e.DetectShell()
	if runtime.GOOS == "windows" {
		if !strings.Contains(strings.ToLower(shell), "cmd") {
			t.Errorf("expected cmd.exe, got %s", shell)
		}
		if args[0] != "/C" {
			t.Errorf("expected /C, got %v", args)
		}
	} else {
		if shell == "" {
			t.Error("empty shell")
		}
	}
}

func TestIsDestructive(t *testing.T) {
	e := New("auto", "")
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

func TestRunEcho(t *testing.T) {
	e := New("auto", "")
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

func TestRunLs(t *testing.T) {
	e := New("auto", "")
	result, err := e.Run("ls")
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.ExitCode != 0 {
		t.Errorf("exit code: %d", result.ExitCode)
	}
}

func TestRunPwd(t *testing.T) {
	e := New("auto", "")
	result, err := e.Run("pwd")
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.ExitCode != 0 {
		t.Errorf("exit code: %d", result.ExitCode)
	}
	cwd, _ := os.Getwd()
	if !strings.Contains(result.Output, filepath.Base(cwd)) {
		t.Errorf("expected cwd in output, got: %s", result.Output)
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

func TestTokenize(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"git status", []string{"git", "status"}},
		{`echo "hello world"`, []string{"echo", "hello world"}},
		{`echo 'hello world'`, []string{"echo", "hello world"}},
		{"ls -la /tmp", []string{"ls", "-la", "/tmp"}},
		{`grep "foo bar" file.txt`, []string{"grep", "foo bar", "file.txt"}},
	}
	for _, tt := range tests {
		got := tokenize(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("tokenize(%q) = %v, want %v", tt.input, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("tokenize(%q)[%d] = %q, want %q", tt.input, i, got[i], tt.want[i])
			}
		}
	}
}

func TestNeedsShell(t *testing.T) {
	shellNeeded := []string{
		"git log | grep fix",
		"echo hello > file.txt",
		"cmd1 && cmd2",
		"ls *.go",
	}
	noShell := []string{
		"git status",
		"npm install",
		"docker ps",
		"echo hello world",
	}
	for _, cmd := range shellNeeded {
		if !needsShell(cmd) {
			t.Errorf("expected shell needed: %s", cmd)
		}
	}
	for _, cmd := range noShell {
		if needsShell(cmd) {
			t.Errorf("expected no shell: %s", cmd)
		}
	}
}

func TestBuiltinTouch(t *testing.T) {
	e := New("auto", "")
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test.txt")
	result, err := e.Run("touch " + path)
	if err != nil {
		t.Fatalf("touch: %v", err)
	}
	if result.ExitCode != 0 {
		t.Errorf("exit code: %d", result.ExitCode)
	}
	if _, err := os.Stat(path); err != nil {
		t.Errorf("file not created: %v", err)
	}
}
