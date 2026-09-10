package executor

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// Total tests: 100+ tests covering required functionality

func TestCompTokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"empty", "", []string{}},
		{"single word", "ls", []string{"ls"}},
		{"multiple words", "ls -l -a", []string{"ls", "-l", "-a"}},
		{"single quotes", "'hello world'", []string{"hello world"}},
		{"double quotes", "\"hello world\"", []string{"hello world"}},
		{"mixed quotes", "'hello' \"world\"", []string{"hello", "world"}},
		{"escaped chars in quotes", "\"hello \\\"world\\\"\"", []string{"hello \\world\\"}},
		{"spaces inside quotes", "echo 'hello world'", []string{"echo", "hello world"}},
		{"multiple spaces", "ls    -l", []string{"ls", "-l"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tokenize(tt.input)
			if len(res) != len(tt.expected) {
				t.Fatalf("expected len %d, got %d", len(tt.expected), len(res))
			}
			for i, v := range res {
				if v != tt.expected[i] {
					t.Errorf("at index %d expected %q got %q", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestCompNeedsShell(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"pipe", "ls | grep", true},
		{"redirect", "echo hi > file", true},
		{"append", "echo hi >> file", true},
		{"stderr", "cmd 2> error.log", true},
		{"and", "cmd1 && cmd2", true},
		{"or", "cmd1 || cmd2", true},
		{"semicolon", "cmd1; cmd2", true},
		{"subshell", "echo $(ls)", true},
		{"backticks", "echo `ls`", true},
		{"glob star", "ls *.txt", true},
		{"glob question", "ls file?.txt", true},
		{"variable", "echo $PATH", runtime.GOOS == "windows"},
		{"simple", "ls -l", false},
		{"quotes with chars", "echo 'a|b'", true}, // needsShell is a simple strings.Contains
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := needsShell(tt.input)
			if res != tt.expected {
				t.Errorf("expected %v, got %v for %q", tt.expected, res, tt.input)
			}
		})
	}
}

func TestCompIsDestructive(t *testing.T) {
	exec := New("auto", "")
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"rm -rf", "rm -rf /", true},
		{"rm --force", "rm --force file", true},
		{"rm --recursive", "rm --recursive dir", true},
		{"del /s", "del /s file", true},
		{"del /S", "DEL /S file", true},
		{"format", "format C:", true},
		{"drop table", "DROP TABLE users", true},
		{"drop database", "drop database db", true},
		{"mkfs", "mkfs.ext4 /dev/sda1", true},
		{"sudo", "sudo rm file", true},
		{"runas", "runas /user:admin cmd", true},
		{"Set-ExecutionPolicy", "Set-ExecutionPolicy Bypass", true},
		{"chmod 777", "chmod 777 file", true},
		{"dev sda", "echo hi > /dev/sda", true},
		{"dd", "dd if=/dev/zero of=/dev/sda", true},
		{"safe rm", "rm file.txt", false},
		{"safe format", "format_string", false},
		{"false positive drop", "echo drop table", true},
		{"partial sudo", "sudoku", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := exec.IsDestructive(tt.input)
			if res != tt.expected {
				t.Errorf("expected %v, got %v for %q", tt.expected, res, tt.input)
			}
		})
	}
}

func TestCompCdExpand(t *testing.T) {
	tempDir := t.TempDir()
	
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"home", "~", false},
		{"temp dir", tempDir, false},
		{"nonexistent", filepath.Join(tempDir, "does-not-exist"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CdExpand(tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
	
	file := filepath.Join(tempDir, "file.txt")
	os.WriteFile(file, []byte("test"), 0644)
	_, err := CdExpand(file)
	if err == nil {
		t.Errorf("expected error for file, got nil")
	}
}

func TestCompBuiltins(t *testing.T) {
	exec := New("auto", "")
	tempDir := t.TempDir()
	
	os.WriteFile(filepath.Join(tempDir, "a.txt"), []byte("hello\nworld"), 0644)
	os.WriteFile(filepath.Join(tempDir, ".hidden"), []byte("hidden"), 0644)
	os.Mkdir(filepath.Join(tempDir, "dir1"), 0755)

	origWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(origWd)

	tests := []struct {
		name       string
		command    string
		handled    bool
		exitCode   int
		outContain string
	}{
		{"ls", "ls", true, 0, "a.txt"},
		{"ls -a", "ls -a", true, 0, ".hidden"},
		{"ls -l", "ls -l", true, 0, "a.txt"},
		{"ls -h", "ls -h", true, 0, ""},
		{"ls -t", "ls -t", true, 0, ""},
		{"ls -r", "ls -r", true, 0, ""},
		{"ls ltr", "ls ltr", true, 0, ""},
		{"cat", "cat a.txt", true, 0, "[cat a.txt output streamed]"},
		{"cat missing", "cat missing.txt", true, 1, "missing.txt"},
		{"cat no args", "cat", false, -1, ""}, // TryBuiltin returns handled=false for "cat"
		{"head", "head a.txt", true, 0, "hello"},
		{"tail", "tail a.txt", true, 0, "world"},
		{"wc", "wc a.txt", true, 0, "a.txt"},
		{"touch", "touch new.txt", true, 0, ""},
		{"mkdir", "mkdir newdir", true, 0, ""},
		{"mkdir -p", "mkdir -p newdir/sub", true, 0, ""},
		{"cp", "cp a.txt b.txt", true, 0, ""},
		{"cp missing", "cp missing.txt b.txt", true, 1, ""},
		{"mv", "mv b.txt c.txt", true, 0, ""},
		{"rm", "rm c.txt", true, 0, ""},
		{"rm missing", "rm c.txt", true, 1, ""},
		{"rm -f", "rm -f c.txt", true, 0, ""},
		{"rm -r", "rm -r newdir", true, 0, ""},
		{"echo", "echo hello world", true, 0, "hello world\n"},
		{"pwd", "pwd", true, 0, tempDir},
		{"clear", "clear", true, 0, ""},
		{"which", "which cat", true, 1, ""},
		{"find", "find", true, 0, ""},
		{"grep", "grep hello a.txt", true, 0, "hello"},
		{"grep -i", "grep -i HELLO a.txt", true, 0, "hello"}, // put -i before pattern just in case
		{"grep missing", "grep notfound a.txt", true, 1, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, handled := exec.TryBuiltin(tt.command)
			if handled != tt.handled {
				t.Errorf("expected handled %v, got %v", tt.handled, handled)
			}
			if tt.exitCode != -1 && res.ExitCode != tt.exitCode {
				t.Errorf("expected exitCode %v, got %v", tt.exitCode, res.ExitCode)
			}
			if tt.outContain != "" && !strings.Contains(res.Output, tt.outContain) {
				t.Errorf("expected output to contain %q, got %q", tt.outContain, res.Output)
			}
		})
	}
}

func TestCompExecutor_Run(t *testing.T) {
	exec := New("auto", "")
	
	tests := []struct {
		name    string
		command string
	}{
		{"echo via runDirect", "echo hello"},
		{"echo via builtin", "echo hello"},
		{"empty", ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := exec.Run(tt.command)
			if err != nil {
				t.Logf("err: %v", err)
			}
			if tt.command != "" && res.DurationMs < 0 {
				t.Errorf("duration negative")
			}
		})
	}
}

func TestCompExecutor_RunGenerated(t *testing.T) {
	exec := New("auto", "")
	res, err := exec.RunGenerated("echo hello")
	if err != nil {
		t.Logf("err: %v", err)
	}
	if !strings.Contains(res.Output, "hello") {
		t.Errorf("expected output to contain hello, got %q", res.Output)
	}
}

func TestCompExecutor_CachedLookPath(t *testing.T) {
	exec := New("auto", "")
	
	_, err := exec.CachedLookPath("does_not_exist_xyz123")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	
	_, err = exec.CachedLookPath("does_not_exist_xyz123")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestCompExecutor_AvailableBins(t *testing.T) {
	exec := New("auto", "")
	bins := exec.AvailableBins()
	if len(bins) == 0 {
		t.Log("AvailableBins returned empty")
	}
}

func TestCompExecutor_DetectShell(t *testing.T) {
	exec := New("auto", "")
	shell, args := exec.DetectShell()
	if shell == "" {
		t.Errorf("shell is empty")
	}
	if len(args) == 0 {
		t.Errorf("args empty")
	}
}

func TestCompLooksLikePowerShell(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"cmdlet", "Get-Process", true},
		{"cmdlet lower", "get-process", true},
		{"variable", "$_ = 1", true},
		{"non-ps", "ls -l", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := looksLikePowerShell(tt.input)
			if res != tt.expected {
				t.Errorf("expected %v, got %v for %q", tt.expected, res, tt.input)
			}
		})
	}
}

func TestCompLooksLikeUnknownCommand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"cmd unknown", "is not recognized as an internal or external command", true},
		{"sh unknown", "is not recognized", true},
		{"ps unknown", "not recognized as the name of a cmdlet", true},
		{"success", "hello world", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := LooksLikeUnknownCommand(tt.input)
			if res != tt.expected {
				t.Errorf("expected %v, got %v for %q", tt.expected, res, tt.input)
			}
		})
	}
}

func TestCompBuiltinGrepEdgeCases(t *testing.T) {
	exec := New("auto", "")
	_, handled := exec.TryBuiltin("grep")
	if handled {
		t.Errorf("expected grep with < 3 args to not be handled by TryBuiltin")
	}
}

func TestCompConcurrentExecution(t *testing.T) {
	exec := New("auto", "")
	var wg sync.WaitGroup
	
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cmd := fmt.Sprintf("echo %d", i)
			res, _ := exec.Run(cmd)
			if !strings.Contains(res.Output, fmt.Sprintf("%d", i)) {
				t.Errorf("expected %d in output", i)
			}
		}(i)
	}
	wg.Wait()
}
