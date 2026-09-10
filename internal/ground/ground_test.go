package ground

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestListCWD(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("hi"), 0644)
	os.Mkdir(filepath.Join(dir, "sub"), 0755)
	listing := ListCWD(dir, 10)
	if !strings.Contains(listing, "a.txt") {
		t.Fatalf("missing file: %s", listing)
	}
	if !strings.Contains(listing, "sub/") {
		t.Fatalf("missing dir: %s", listing)
	}
}

func TestExecToolWhich(t *testing.T) {
	exists := func(name string) bool { return name == "git" }
	if got := ExecTool("which", map[string]any{"name": "git"}, ".", exists); !strings.Contains(got, "git") {
		t.Fatalf("got %q", got)
	}
	if got := ExecTool("which", map[string]any{"name": "nope"}, ".", exists); !strings.Contains(got, "not on PATH") {
		t.Fatalf("got %q", got)
	}
	if got := ExecTool("nope", nil, ".", exists); !strings.Contains(got, "unknown tool") {
		t.Fatalf("got %q", got)
	}
}

func TestCapturePresent(t *testing.T) {
	dir := t.TempDir()
	s := Capture(dir, func(name string) bool { return name == "git" })
	if s.CWD != dir {
		t.Fatalf("cwd %s", s.CWD)
	}
	if len(s.Present) != 1 || s.Present[0] != "git" {
		t.Fatalf("present %+v", s.Present)
	}
}

func TestExecToolNetstat(t *testing.T) {
	got := ExecTool("netstat", nil, ".", nil)
	if strings.Contains(got, "unknown tool") {
		t.Fatalf("netstat should be implemented, got: %q", got)
	}
}

func TestExecToolEnvVar(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "test_value")
	defer os.Unsetenv("TEST_ENV_VAR")

	got := ExecTool("env_var", map[string]any{"name": "TEST_ENV_VAR"}, ".", nil)
	if got != "test_value" {
		t.Fatalf("expected test_value, got: %q", got)
	}

	gotEmpty := ExecTool("env_var", map[string]any{"name": "NON_EXISTENT_VAR_123"}, ".", nil)
	if gotEmpty != "(not set)" {
		t.Fatalf("expected (not set), got: %q", gotEmpty)
	}
}

func TestExecToolReadFileHead(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test.txt")
	content := "line1\nline2\nline3\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	got := ExecTool("read_file_head", map[string]any{"path": "test.txt"}, dir, nil)
	if !strings.Contains(got, "line1") {
		t.Fatalf("expected output to contain line1, got: %q", got)
	}
}
