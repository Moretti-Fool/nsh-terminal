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

func TestExecTool(t *testing.T) {
	got := ExecTool("nope", nil, ".", nil)
	if !strings.Contains(got, "error") && !strings.Contains(got, "not recognized") {
		t.Fatalf("got %q", got)
	}
}
