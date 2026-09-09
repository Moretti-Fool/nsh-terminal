package repl

import (
	"strings"
	"testing"
)

func TestCurrentWord(t *testing.T) {
	if g := currentWord("cat RE"); g != "RE" {
		t.Fatalf("got %q", g)
	}
	if g := currentWord("ls"); g != "ls" {
		t.Fatalf("got %q", g)
	}
}

func TestCompleterFirstToken(t *testing.T) {
	c := nshCompleter{}
	cands, n := c.Do([]rune("gi"), 2)
	if n != 2 {
		t.Fatalf("offset %d", n)
	}
	found := false
	for _, c := range cands {
		if string(c) == "t" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected git remainder 't', got %v", cands)
	}
}

func TestCompleterLocalFile(t *testing.T) {
	c := nshCompleter{}
	line := []rune("cat com")
	cands, _ := c.Do(line, len(line))
	found := false
	for _, cand := range cands {
		if strings.HasPrefix(strings.ToLower(string(cand)), "plete.go") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected complete.go remainder, got %q", cands)
	}
}
