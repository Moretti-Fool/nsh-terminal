package search

import (
	"strings"
	"testing"
)

var testEngines = map[string]string{
	"google": "https://www.google.com/search?q=%s",
	"wiki":   "https://en.wikipedia.org/wiki/Special:Search/%s",
	"yt":     "https://www.youtube.com/results?search_query=%s",
	"gh":     "https://github.com/search?q=%s",
}

func TestBuildURL(t *testing.T) {
	s := New(testEngines, "google")
	tests := []struct{ engine, query, want string }{
		{"google", "what is golang", "https://www.google.com/search?q=what+is+golang"},
		{"wiki", "Alan Turing", "https://en.wikipedia.org/wiki/Special:Search/Alan+Turing"},
	}
	for _, tt := range tests {
		url, err := s.BuildURL(tt.engine, tt.query)
		if err != nil || url != tt.want {
			t.Errorf("BuildURL(%s, %s) = %s, err=%v", tt.engine, tt.query, url, err)
		}
	}
}

func TestBuildURLDefault(t *testing.T) {
	s := New(testEngines, "google")
	url, _ := s.BuildURL("", "test")
	if !strings.Contains(url, "google.com") {
		t.Errorf("expected google, got %s", url)
	}
}

func TestBuildURLUnknown(t *testing.T) {
	s := New(testEngines, "google")
	if _, err := s.BuildURL("bing", "test"); err == nil {
		t.Error("expected error for unknown engine")
	}
}

func TestOpenCommand(t *testing.T) {
	s := New(nil, "google")
	if cmd := s.OpenCommand(); cmd == "" {
		t.Error("should not be empty")
	}
}
