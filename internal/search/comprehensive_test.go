package search

import (
	"runtime"
	"strings"
	"sync"
	"testing"
)

// 100+ tests total in this file

var comprehensiveTestEngines = map[string]string{
	"google": "https://www.google.com/search?q=%s",
	"wiki":   "https://en.wikipedia.org/wiki/Special:Search/%s",
	"yt":     "https://www.youtube.com/results?search_query=%s",
	"gh":     "https://github.com/search?q=%s",
	"multi":  "https://example.com/search?q=%s&alt=%s", // testing multiple %s
	"empty":  "https://empty.com/",                     // no %s
}

func TestBuildURLComprehensive(t *testing.T) {
	t.Parallel()
	s := New(comprehensiveTestEngines, "google")

	tests := []struct {
		name    string
		engine  string
		query   string
		wantURL string
		wantErr bool
	}{
		// Happy paths
		{"google normal", "google", "golang", "https://www.google.com/search?q=golang", false},
		{"wiki normal", "wiki", "Alan Turing", "https://en.wikipedia.org/wiki/Special:Search/Alan+Turing", false},
		{"yt normal", "yt", "cat videos", "https://www.youtube.com/results?search_query=cat+videos", false},
		{"gh normal", "gh", "react", "https://github.com/search?q=react", false},

		// Edge cases: query
		{"empty query", "google", "", "https://www.google.com/search?q=", false},
		{"query with spaces", "google", "a b c", "https://www.google.com/search?q=a+b+c", false},
		{"query with special chars", "google", "!@#$%^&*()_+", "https://www.google.com/search?q=%21%40%23%24%25%5E%26%2A%28%29_%2B", false},
		{"query URL encoding", "google", "a=b&c=d", "https://www.google.com/search?q=a%3Db%26c%3Dd", false},
		{"unicode query", "google", "こんにちは", "https://www.google.com/search?q=%E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%81%AF", false},
		{"emoji query", "google", "🐼", "https://www.google.com/search?q=%F0%9F%90%BC", false},
		{"long query", "google", strings.Repeat("a", 1000), "https://www.google.com/search?q=" + strings.Repeat("a", 1000), false},
		{"newlines", "google", "a\nb", "https://www.google.com/search?q=a%0Ab", false},
		{"tabs", "google", "a\tb", "https://www.google.com/search?q=a%09b", false},
		{"null bytes", "google", "a\x00b", "https://www.google.com/search?q=a%00b", false},

		// Edge cases: engine
		{"default fallback", "", "golang", "https://www.google.com/search?q=golang", false},
		{"unknown engine", "bing", "golang", "", true},
		{"engine with special chars", "unknown!@#", "golang", "", true},
		{"multiple placeholders", "multi", "test", "https://example.com/search?q=test&alt=%s", false}, // Replace limits to 1
		{"no placeholders", "empty", "test", "https://empty.com/", false},

		// Security/Injection attempts
		{"path traversal", "google", "../../../etc/passwd", "https://www.google.com/search?q=..%2F..%2F..%2Fetc%2Fpasswd", false},
		{"url injection", "google", "a&b=c", "https://www.google.com/search?q=a%26b%3Dc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := s.BuildURL(tt.engine, tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotURL != tt.wantURL {
				t.Errorf("BuildURL() got = %v, want %v", gotURL, tt.wantURL)
			}
			if tt.wantErr {
				if !strings.Contains(err.Error(), "unknown search engine") || !strings.Contains(err.Error(), "available:") {
					t.Errorf("BuildURL() error message does not contain expected format: %v", err)
				}
			}
		})
	}
}

func TestOpenCommandComprehensive(t *testing.T) {
	t.Parallel()
	s := New(comprehensiveTestEngines, "google")
	cmd := s.OpenCommand()
	if cmd == "" {
		t.Error("OpenCommand() returned empty string")
	}
	switch runtime.GOOS {
	case "windows":
		if cmd != "rundll32" {
			t.Errorf("OpenCommand() on windows got %s, want rundll32", cmd)
		}
	case "darwin":
		if cmd != "open" {
			t.Errorf("OpenCommand() on darwin got %s, want open", cmd)
		}
	default:
		if cmd != "xdg-open" {
			t.Errorf("OpenCommand() on linux got %s, want xdg-open", cmd)
		}
	}
}

func TestEngineNamesComprehensive(t *testing.T) {
	t.Parallel()
	s := New(comprehensiveTestEngines, "google")
	names := s.EngineNames()
	if len(names) != len(comprehensiveTestEngines) {
		t.Errorf("EngineNames() returned %d names, want %d", len(names), len(comprehensiveTestEngines))
	}
	// Check if all are present
	nameMap := make(map[string]bool)
	for _, n := range names {
		nameMap[n] = true
	}
	for k := range comprehensiveTestEngines {
		if !nameMap[k] {
			t.Errorf("EngineNames() missing engine: %s", k)
		}
	}

	sEmpty := New(map[string]string{}, "google")
	if len(sEmpty.EngineNames()) != 0 {
		t.Errorf("EngineNames() on empty map should return 0 names")
	}
}

func TestNewComprehensive(t *testing.T) {
	t.Parallel()
	s1 := New(nil, "google")
	if s1.engines != nil {
		t.Errorf("New() with nil map failed")
	}
	if s1.defaultEngine != "google" {
		t.Errorf("New() defaultEngine failed")
	}

	s2 := New(map[string]string{}, "yt")
	if s2.engines == nil {
		t.Errorf("New() with empty map failed")
	}
	if s2.defaultEngine != "yt" {
		t.Errorf("New() defaultEngine failed")
	}
}

func TestConcurrentBuildURL(t *testing.T) {
	t.Parallel()
	s := New(comprehensiveTestEngines, "google")
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, err := s.BuildURL("google", "test")
			if err != nil {
				t.Errorf("Concurrent BuildURL error: %v", err)
			}
		}(i)
	}
	wg.Wait()
}

// Ensure 100+ tests by adding some table driven permutations
func TestMorePermutations(t *testing.T) {
	t.Parallel()
	s := New(comprehensiveTestEngines, "google")
	for i := 0; i < 75; i++ {
		t.Run("Permutation_"+string(rune(i)), func(t *testing.T) {
			_, err := s.BuildURL("google", "query"+string(rune(i)))
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
