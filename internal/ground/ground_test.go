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
