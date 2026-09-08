package history

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAppendAndRead(t *testing.T) {
	h := New(t.TempDir(), 500)
	err := h.Append(Entry{
		Timestamp: time.Now(), Input: "show running containers",
		Type: "nl", Generated: "docker ps", ExitCode: 0, CWD: "/tmp",
	})
	if err != nil {
		t.Fatalf("append: %v", err)
	}
	entries, err := h.ReadDate(time.Now())
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if len(entries) != 1 || entries[0].Generated != "docker ps" {
		t.Errorf("unexpected entries: %+v", entries)
	}
}

func TestMultiple(t *testing.T) {
	h := New(t.TempDir(), 500)
	for i := 0; i < 5; i++ {
		h.Append(Entry{Timestamp: time.Now(), Input: "cmd", Type: "command", CWD: "/tmp"})
	}
	entries, _ := h.ReadDate(time.Now())
	if len(entries) != 5 {
		t.Errorf("expected 5, got %d", len(entries))
	}
}

func TestTruncation(t *testing.T) {
	h := New(t.TempDir(), 20)
	h.Append(Entry{Timestamp: time.Now(), Input: "test", Type: "command",
		OutputPreview: "this is a very long output string that should be truncated", CWD: "/tmp"})
	entries, _ := h.ReadDate(time.Now())
	if len(entries[0].OutputPreview) > 20 {
		t.Errorf("not truncated: len=%d", len(entries[0].OutputPreview))
	}
}

func TestSearch(t *testing.T) {
	h := New(t.TempDir(), 500)
	h.Append(Entry{Timestamp: time.Now(), Input: "git status", Type: "command", CWD: "/tmp"})
	h.Append(Entry{Timestamp: time.Now(), Input: "show containers", Type: "nl", Generated: "docker ps", CWD: "/tmp"})
	if results := h.Search("docker", 10); len(results) != 1 {
		t.Errorf("expected 1, got %d", len(results))
	}
}

func TestLastN(t *testing.T) {
	h := New(t.TempDir(), 500)
	for i := 0; i < 10; i++ {
		h.Append(Entry{Timestamp: time.Now(), Input: "cmd", Type: "command", CWD: "/tmp"})
	}
	if last := h.LastN(3); len(last) != 3 {
		t.Errorf("expected 3, got %d", len(last))
	}
}

func TestRotation(t *testing.T) {
	tmp := t.TempDir()
	h := New(tmp, 500)
	oldDate := time.Now().AddDate(0, 0, -100)
	oldFile := filepath.Join(tmp, oldDate.Format("2006-01-02")+".jsonl")
	os.WriteFile(oldFile, []byte("{\"ts\":\"2020-01-01T00:00:00Z\",\"input\":\"old\"}\n"), 0644)
	recentFile := filepath.Join(tmp, time.Now().Format("2006-01-02")+".jsonl")
	os.WriteFile(recentFile, []byte("{\"ts\":\"2026-09-08T00:00:00Z\",\"input\":\"recent\"}\n"), 0644)

	h.Rotate(90)
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Error("old file should be deleted")
	}
	if _, err := os.Stat(recentFile); err != nil {
		t.Error("recent file should exist")
	}
}
