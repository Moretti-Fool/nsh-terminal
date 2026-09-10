package history

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

type Entry struct {
	Timestamp     time.Time `json:"ts"`
	Input         string    `json:"input"`
	Type          string    `json:"type"`
	Generated     string    `json:"generated,omitempty"`
	ExitCode      int       `json:"exit_code"`
	OutputPreview string    `json:"output_preview,omitempty"`
	CWD           string    `json:"cwd"`
	DurationMs    int64     `json:"duration_ms,omitempty"`
}

var scannerBufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024*1024)
	},
}

type History struct {
	dir             string
	maxPreviewChars int
}

func New(dir string, maxPreviewChars int) *History {
	return &History{dir: dir, maxPreviewChars: maxPreviewChars}
}

func (h *History) fileForDate(t time.Time) string {
	return filepath.Join(h.dir, t.Format("2006-01-02")+".jsonl")
}

func (h *History) Append(e Entry) error {
	if err := os.MkdirAll(h.dir, 0755); err != nil {
		return err
	}
	if h.maxPreviewChars > 0 && len(e.OutputPreview) > h.maxPreviewChars {
		e.OutputPreview = e.OutputPreview[:h.maxPreviewChars]
	}
	f, err := os.OpenFile(h.fileForDate(e.Timestamp), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	err = json.NewEncoder(f).Encode(e)
	if cerr := f.Close(); cerr != nil && err == nil {
		return cerr
	}
	return err
}

func (h *History) ReadDate(t time.Time) ([]Entry, error) {
	return readJSONLFile(h.fileForDate(t))
}

func (h *History) ReadRange(from, to time.Time) ([]Entry, error) {
	var all []Entry
	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		entries, err := h.ReadDate(d)
		if err != nil {
			return nil, err
		}
		all = append(all, entries...)
	}
	return all, nil
}

func (h *History) LastN(n int) []Entry {
	today := time.Now()
	var all []Entry
	for i := 0; i < 90 && len(all) < n; i++ {
		entries, err := h.ReadDate(today.AddDate(0, 0, -i))
		if err != nil && !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "[nsh] trace: failed to read history for %s: %v\n", today.AddDate(0, 0, -i).Format("2006-01-02"), err)
		}
		all = append(entries, all...)
	}
	if len(all) > n {
		all = all[len(all)-n:]
	}
	return all
}

func (h *History) Search(query string, limit int) []Entry {
	lower := strings.ToLower(query)
	today := time.Now()
	var results []Entry
	for i := 0; i < 90 && len(results) < limit; i++ {
		entries, err := h.ReadDate(today.AddDate(0, 0, -i))
		if err != nil && !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "[nsh] trace: failed to read history for %s: %v\n", today.AddDate(0, 0, -i).Format("2006-01-02"), err)
		}
		for j := len(entries) - 1; j >= 0 && len(results) < limit; j-- {
			e := entries[j]
			if strings.Contains(strings.ToLower(e.Input), lower) ||
				strings.Contains(strings.ToLower(e.Generated), lower) {
				results = append(results, e)
			}
		}
	}
	return results
}

func (h *History) Rotate(retentionDays int) error {
	entries, err := os.ReadDir(h.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	cutoff := time.Now().AddDate(0, 0, -retentionDays)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".jsonl") {
			continue
		}
		dateStr := strings.TrimSuffix(entry.Name(), ".jsonl")
		fileDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}
		if fileDate.Before(cutoff) {
			os.Remove(filepath.Join(h.dir, entry.Name()))
		}
	}
	return nil
}

func (h *History) FormatDayView(entries []Entry) string {
	if len(entries) == 0 {
		return "  No commands recorded.\n"
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp.Before(entries[j].Timestamp)
	})
	var b strings.Builder
	for _, e := range entries {
		timeStr := e.Timestamp.Format("15:04")
		status := "✓"
		if e.ExitCode != 0 {
			status = fmt.Sprintf("✗ (exit %d)", e.ExitCode)
		}
		if e.Generated != "" {
			b.WriteString(fmt.Sprintf("  %s  %q → %s  %s\n", timeStr, e.Input, e.Generated, status))
		} else {
			b.WriteString(fmt.Sprintf("  %s  %s  %s\n", timeStr, e.Input, status))
		}
	}
	return b.String()
}

func readJSONLFile(path string) ([]Entry, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()
	var entries []Entry
	scanner := bufio.NewScanner(f)
	
	bufInterface := scannerBufPool.Get()
	buf := bufInterface.([]byte)
	defer scannerBufPool.Put(bufInterface)
	
	scanner.Buffer(buf, len(buf))
	for scanner.Scan() {
		var e Entry
		if err := json.Unmarshal(scanner.Bytes(), &e); err != nil {
			continue
		}
		entries = append(entries, e)
	}
	return entries, scanner.Err()
}
