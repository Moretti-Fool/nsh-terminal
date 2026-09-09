package classifier

import (
	"testing"
)

var testPathLookup = func(name string) bool {
	known := map[string]bool{
		"git": true, "docker": true, "ls": true, "cd": true,
		"npm": true, "go": true, "python": true, "cat": true,
		"mkdir": true, "rm": true, "curl": true, "ping": true,
		"find": true, "sort": true, "kill": true,
	}
	return known[name]
}

func TestClassifyCommands(t *testing.T) {
	c := New(testPathLookup, nil)
	for _, input := range []string{
		"git status", "docker ps", "ls -la", "npm run dev",
		"./script.sh", "/usr/bin/python",
		"cat file.txt | grep error", "echo hello > out.txt",
		"curl --silent https://example.com",
	} {
		if r := c.Classify(input); r.Type != Command {
			t.Errorf("Classify(%q) = %v, want Command", input, r.Type)
		}
	}
}

func TestClassifyNL(t *testing.T) {
	c := New(testPathLookup, nil)
	for _, input := range []string{
		"show me all running containers",
		"list all files in this directory",
		"find the largest file here",
		"what processes are using port 8080",
		"how much disk space is left",
		"compress this folder into a zip",
	} {
		if r := c.Classify(input); r.Type != NaturalLanguage {
			t.Errorf("Classify(%q) = %v, want NL", input, r.Type)
		}
	}
}

func TestClassifyBuiltin(t *testing.T) {
	c := New(nil, nil)
	for _, input := range []string{"nsh help", "nsh record start", "nsh workflows", "nsh history"} {
		if r := c.Classify(input); r.Type != Builtin {
			t.Errorf("Classify(%q) = %v, want Builtin", input, r.Type)
		}
	}
}

func TestClassifySearch(t *testing.T) {
	c := New(nil, nil)
	for _, input := range []string{
		"google what is golang", "search python tutorial",
		"wiki Alan Turing", "yt go patterns", "gh nsh terminal",
	} {
		if r := c.Classify(input); r.Type != Search {
			t.Errorf("Classify(%q) = %v, want Search", input, r.Type)
		}
	}
}

func TestClassifyWorkflow(t *testing.T) {
	c := New(func(string) bool { return false }, []string{"run-backend", "deploy-staging"})
	for _, input := range []string{"run-backend", "deploy-staging"} {
		if r := c.Classify(input); r.Type != Workflow {
			t.Errorf("Classify(%q) = %v, want Workflow", input, r.Type)
		}
	}
}

func TestClassifyEmpty(t *testing.T) {
	c := New(nil, nil)
	if r := c.Classify(""); r.Type != Empty {
		t.Errorf("got %v, want Empty", r.Type)
	}
	if r := c.Classify("   "); r.Type != Empty {
		t.Errorf("got %v, want Empty", r.Type)
	}
}

func TestClassifyWindowsPath(t *testing.T) {
	c := New(nil, nil)
	if r := c.Classify(`C:\Users\test\script.bat`); r.Type != Command {
		t.Errorf("got %v, want Command", r.Type)
	}
}

func TestClassifyAmbiguousNL(t *testing.T) {
	c := New(testPathLookup, nil)
	for _, input := range []string{
		"sort these items by name",
		"kill all background tasks",
	} {
		r := c.Classify(input)
		if r.Type == Command {
			t.Errorf("Classify(%q) = Command, want Ambiguous or NL", input)
		}
	}
}

func TestClassifyRealCommandsStayCommand(t *testing.T) {
	c := New(testPathLookup, nil)
	for _, input := range []string{
		"find . -name *.go",
		"sort -n file.txt",
		"kill -9 1234",
		"git status",
		"npm run dev",
		"docker compose up",
	} {
		r := c.Classify(input)
		if r.Type != Command {
			t.Errorf("Classify(%q) = %v, want Command", input, r.Type)
		}
	}
}

func TestClassifySearchModes(t *testing.T) {
	c := New(nil, nil)
	for _, input := range []string{
		"ask what is kubernetes",
		"google! explain docker",
		"search! latest golang news",
	} {
		r := c.Classify(input)
		if r.Type != Search {
			t.Errorf("Classify(%q) = %v, want Search", input, r.Type)
		}
	}
}
