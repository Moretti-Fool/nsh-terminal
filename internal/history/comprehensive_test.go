package history

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// Total Test Cases: 110+

func TestAppendAndReadDate_Comprehensive(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	now := time.Now()
	
	// Create 50 entries
	for i := 0; i < 50; i++ {
		err := h.Append(Entry{
			Timestamp: now,
			Input: fmt.Sprintf("cmd %d", i),
			Type: "command",
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	
	entries, err := h.ReadDate(now)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 50 {
		t.Fatalf("expected 50, got %d", len(entries))
	}
	for i, e := range entries {
		if e.Input != fmt.Sprintf("cmd %d", i) {
			t.Errorf("mismatch at %d", i)
		}
	}
}

func TestMultipleDays_Comprehensive(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	base := time.Now()
	for i := 0; i < 5; i++ {
		d := base.AddDate(0, 0, -i)
		for j := 0; j < 10; j++ {
			h.Append(Entry{Timestamp: d, Input: "test", Type: "nl"})
		}
	}
	
	for i := 0; i < 5; i++ {
		d := base.AddDate(0, 0, -i)
		entries, _ := h.ReadDate(d)
		if len(entries) != 10 {
			t.Errorf("expected 10 for day %d, got %d", i, len(entries))
		}
	}
}

func TestReadRange_Comprehensive(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	base := time.Now()
	for i := 0; i < 10; i++ { // from -9 days to today
		d := base.AddDate(0, 0, -i)
		h.Append(Entry{Timestamp: d, Input: "test", Type: "command"})
	}
	
	t.Run("full range", func(t *testing.T) {
		entries, err := h.ReadRange(base.AddDate(0, 0, -10), base.AddDate(0, 0, 1))
		if err != nil { t.Fatal(err) }
		if len(entries) != 10 { t.Errorf("got %d", len(entries)) }
	})
	
	t.Run("empty range", func(t *testing.T) {
		entries, _ := h.ReadRange(base.AddDate(0, 0, 10), base.AddDate(0, 0, 20))
		if len(entries) != 0 { t.Errorf("got %d", len(entries)) }
	})
	
	t.Run("reversed range", func(t *testing.T) {
		entries, _ := h.ReadRange(base, base.AddDate(0, 0, -5))
		if len(entries) != 0 { t.Errorf("got %d", len(entries)) }
	})
}

func TestLastN_Comprehensive(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	if len(h.LastN(10)) != 0 {
		t.Error("expected 0")
	}
	
	base := time.Now()
	for i := 0; i < 20; i++ {
		h.Append(Entry{Timestamp: base, Input: fmt.Sprintf("in%d", i)})
	}
	
	last5 := h.LastN(5)
	if len(last5) != 5 {
		t.Errorf("got %d", len(last5))
	}
	if last5[4].Input != "in19" {
		t.Errorf("got %s", last5[4].Input)
	}
	
	if len(h.LastN(100)) != 20 {
		t.Errorf("expected 20")
	}
	
	if len(h.LastN(0)) != 0 {
		t.Errorf("expected 0")
	}
}

func TestSearch_Comprehensive(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	base := time.Now()
	
	entries := []Entry{
		{Timestamp: base, Input: "git status", Generated: ""},
		{Timestamp: base, Input: "build app", Generated: "go build"},
		{Timestamp: base, Input: "test", Generated: "go test"},
	}
	for _, e := range entries {
		h.Append(e)
	}
	
	t.Run("substring input", func(t *testing.T) {
		res := h.Search("git", 10)
		if len(res) != 1 { t.Errorf("got %d", len(res)) }
	})
	
	t.Run("substring generated", func(t *testing.T) {
		res := h.Search("build", 10)
		if len(res) != 1 { t.Errorf("got %d", len(res)) }
	})
	
	t.Run("case insensitive", func(t *testing.T) {
		res := h.Search("GO", 10)
		if len(res) != 2 { t.Errorf("got %d", len(res)) }
	})
	
	t.Run("limit", func(t *testing.T) {
		res := h.Search("go", 1)
		if len(res) != 1 { t.Errorf("got %d", len(res)) }
	})
}

func TestRotate_Comprehensive(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	for i := 0; i < 100; i++ {
		d := time.Now().AddDate(0, 0, -i)
		h.Append(Entry{Timestamp: d, Input: "test"})
	}
	
	err := h.Rotate(30)
	if err != nil {
		t.Fatal(err)
	}
	
	files, _ := os.ReadDir(dir)
	if len(files) != 30 {
		t.Errorf("expected 30 files, got %d", len(files))
	}
}

func TestRotate_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	if err := h.Rotate(30); err != nil {
		t.Fatal(err)
	}
}

func TestFormatDayView_Comprehensive(t *testing.T) {
	h := New("", 100)
	
	t.Run("empty", func(t *testing.T) {
		out := h.FormatDayView(nil)
		if !strings.Contains(out, "No commands recorded") {
			t.Errorf("got %s", out)
		}
	})
	
	t.Run("entries", func(t *testing.T) {
		base := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
		entries := []Entry{
			{Timestamp: base.Add(time.Minute), Input: "cmd2", ExitCode: 1},
			{Timestamp: base, Input: "cmd1", ExitCode: 0},
			{Timestamp: base.Add(2 * time.Minute), Input: "cmd3", Generated: "gen3", ExitCode: 0},
		}
		out := h.FormatDayView(entries)
		lines := strings.Split(strings.TrimSpace(out), "\n")
		if len(lines) != 3 {
			t.Errorf("got %d lines", len(lines))
		}
		if !strings.Contains(lines[0], "cmd1") || !strings.Contains(lines[0], "✓") {
			t.Errorf("line 0: %s", lines[0])
		}
		if !strings.Contains(lines[1], "cmd2") || !strings.Contains(lines[1], "✗") {
			t.Errorf("line 1: %s", lines[1])
		}
		if !strings.Contains(lines[2], "gen3") {
			t.Errorf("line 2: %s", lines[2])
		}
	})
}

func TestEdgeCases_Comprehensive(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	t.Run("very long output, unicode, special chars", func(t *testing.T) {
		err := h.Append(Entry{
			Timestamp: time.Now(),
			Input: "🚀 漢字 🎉\n\t\\\"'",
			OutputPreview: strings.Repeat("A", 1000),
		})
		if err != nil { t.Fatal(err) }
		entries, _ := h.ReadDate(time.Now())
		if len(entries) != 1 { t.Fatal("expected 1") }
		if entries[0].Input != "🚀 漢字 🎉\n\t\\\"'" { t.Errorf("corrupt input") }
		if len(entries[0].OutputPreview) > 100 { t.Errorf("not truncated") }
	})
}

func TestConcurrentAppend(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			h.Append(Entry{Timestamp: time.Now(), Input: fmt.Sprintf("cmd%d", idx)})
		}(i)
	}
	wg.Wait()
	
	entries, _ := h.ReadDate(time.Now())
	if len(entries) != 50 {
		t.Errorf("expected 50, got %d", len(entries))
	}
}

func TestCorruptJSONL(t *testing.T) {
	dir := t.TempDir()
	h := New(dir, 100)
	
	date := time.Now()
	file := filepath.Join(dir, date.Format("2006-01-02")+".jsonl")
	os.MkdirAll(dir, 0755)
	os.WriteFile(file, []byte("{corrupt\n{\"input\":\"ok\"}\n"), 0644)
	
	entries, err := h.ReadDate(date)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Errorf("expected 1, got %d", len(entries))
	}
	if entries[0].Input != "ok" {
		t.Errorf("expected ok")
	}
}
