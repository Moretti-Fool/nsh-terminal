package predict

import (
	"sync"
	"testing"
)

func TestRecord(t *testing.T) {
	p := New()
	p.Record("git status", "/home/user", 0)

	if len(p.session) != 1 {
		t.Fatalf("Expected session length 1, got %d", len(p.session))
	}

	if p.frequency["git status"] != 1 {
		t.Errorf("Expected freq 1, got %d", p.frequency["git status"])
	}

	p.Record("git commit -m 'test'", "/home/user", 0)

	if p.bigrams["git status"]["git commit"] != 1 {
		t.Errorf("Expected bigram count 1, got %d", p.bigrams["git status"]["git commit"])
	}
}

func TestSuggestAfterBigram(t *testing.T) {
	p := New()
	p.Record("git add .", "/test", 0)
	p.Record("git commit -m 'msg'", "/test", 0)
	p.Record("git add .", "/test", 0)
	p.Record("git status", "/test", 0)

	suggs := p.SuggestAfter("git add .", 2)
	if len(suggs) != 2 {
		t.Fatalf("Expected 2 suggestions, got %d", len(suggs))
	}
	if suggs[0].Command != "git commit" && suggs[0].Command != "git status" {
		t.Errorf("Unexpected top suggestion: %s", suggs[0].Command)
	}
}

func TestSuggestCWDAffinity(t *testing.T) {
	p := New()
	p.Record("npm install", "/project1", 0)
	p.Record("npm start", "/project1", 0)
	p.Record("go test", "/project2", 0)

	// Reset session so bigram doesn't dominate
	p.session = nil

	suggs := p.Suggest("/project2", 1)
	if len(suggs) != 1 {
		t.Fatalf("Expected 1 suggestion, got %d", len(suggs))
	}
	if suggs[0].Command != "go test" {
		t.Errorf("Expected go test, got %s", suggs[0].Command)
	}
}

func TestSuggestComposite(t *testing.T) {
	p := New()
	// Build freq
	for i := 0; i < 10; i++ {
		p.Record("ls", "/home", 0)
	}

	p.Record("git status", "/repo", 0)
	p.Record("git commit", "/repo", 0)

	suggs := p.Suggest("/repo", 5)
	if len(suggs) == 0 {
		t.Fatal("Expected suggestions")
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"ls -la", "ls"},
		{"git push origin main", "git push"},
		{"docker build .", "docker build"},
		{"npm run dev", "npm run"},
		{"kubectl get pods", "kubectl get"},
		{"go build -o app", "go build"},
		{"apt-get install tree", "apt-get install"},
		{"apt update", "apt update"},
		{"   echo hello   ", "echo"},
		{"", ""},
		{"   ", ""},
		{"git", "git"},
		{"docker", "docker"},
		{"ls", "ls"},
		{"cat file.txt", "cat"},
		{"grep -r main .", "grep"},
		{"mkdir -p a/b/c", "mkdir"},
		{"rm -rf /", "rm"},
		{"mv a b", "mv"},
		{"cp -r a b", "cp"},
	}

	for _, c := range cases {
		if got := normalize(c.in); got != c.out {
			t.Errorf("normalize(%q) = %q, want %q", c.in, got, c.out)
		}
	}
}

func TestSessionRingBuffer(t *testing.T) {
	p := New()
	for i := 0; i < 150; i++ {
		p.Record("ls", "/test", 0)
	}
	if len(p.session) != 100 {
		t.Errorf("Expected session length 100, got %d", len(p.session))
	}
}

func TestConcurrentAccess(t *testing.T) {
	p := New()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.Record("ls", "/test", 0)
			p.Suggest("/test", 5)
			p.SuggestAfter("ls", 2)
		}()
	}
	wg.Wait()
}

func TestEmptyPredictor(t *testing.T) {
	p := New()
	if suggs := p.Suggest("/test", 5); len(suggs) != 0 {
		t.Errorf("Expected no suggestions, got %v", suggs)
	}
	if suggs := p.SuggestAfter("ls", 5); len(suggs) != 0 {
		t.Errorf("Expected no suggestions, got %v", suggs)
	}
}
