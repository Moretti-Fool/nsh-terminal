package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func TestComprehensiveConfig(t *testing.T) {
	t.Run("Default_Values", func(t *testing.T) {
		cfg := Default()
		tests := []struct {
			name string
			got  interface{}
			want interface{}
		}{
			{"Ollama.URL", cfg.Ollama.URL, "http://localhost:11434"},
			{"Ollama.ClassifierModel", cfg.Ollama.ClassifierModel, "qwen2.5-coder:7b"},
			{"Ollama.GenerationModel", cfg.Ollama.GenerationModel, "qwen2.5-coder:7b"},
			{"Ollama.TimeoutMs", cfg.Ollama.TimeoutMs, 45000},
			{"Shell.Default", cfg.Shell.Default, "auto"},
			{"Shell.WSLDistro", cfg.Shell.WSLDistro, ""},
			{"Search.DefaultEngine", cfg.Search.DefaultEngine, "google"},
			{"Search.OpenBrowser", cfg.Search.OpenBrowser, true},
			{"Search.AISummary", cfg.Search.AISummary, false},
			{"UI.Prompt", cfg.UI.Prompt, "nsh> "},
			{"UI.ShowGeneratedCommand", cfg.UI.ShowGeneratedCommand, true},
			{"UI.ConfirmDestructive", cfg.UI.ConfirmDestructive, true},
			{"UI.Theme", cfg.UI.Theme, "default"},
			{"UI.ShowWelcome", cfg.UI.ShowWelcome, true},
			{"History.RetentionDays", cfg.History.RetentionDays, 90},
			{"History.CaptureOutput", cfg.History.CaptureOutput, true},
			{"History.OutputPreviewChars", cfg.History.OutputPreviewChars, 500},
			{"Scratch.Dir", cfg.Scratch.Dir, ""},
			{"Scratch.Python", cfg.Scratch.Python, ""},
		}

		for _, tt := range tests {
			if tt.got != tt.want {
				t.Errorf("Default %s = %v, want %v", tt.name, tt.got, tt.want)
			}
		}

		engines := map[string]string{
			"google": "https://www.google.com/search?q=%s",
			"wiki":   "https://en.wikipedia.org/wiki/Special:Search/%s",
			"yt":     "https://www.youtube.com/results?search_query=%s",
			"gh":     "https://github.com/search?q=%s",
		}
		for k, v := range engines {
			if cfg.Search.Engines[k] != v {
				t.Errorf("Default Search.Engines[%s] = %v, want %v", k, cfg.Search.Engines[k], v)
			}
		}
	})

	t.Run("Load_Save_RoundTrip", func(t *testing.T) {
		tmpDir := t.TempDir()
		
		// 50 iterations of roundtrips to test edge cases
		for i := 0; i < 50; i++ {
			cfg := Default()
			cfg.Ollama.TimeoutMs = i
			cfg.UI.Prompt = fmt.Sprintf("prompt%d> ", i)
			cfg.History.RetentionDays = -i
			cfg.UI.Theme = strings.Repeat("a", i)
			
			cfgFile := filepath.Join(tmpDir, fmt.Sprintf("config%d.toml", i))
			if err := Save(cfg, cfgFile); err != nil {
				t.Fatalf("Save failed: %v", err)
			}
			
			loaded, err := Load(cfgFile)
			if err != nil {
				t.Fatalf("Load failed: %v", err)
			}
			
			if loaded.Ollama.TimeoutMs != i {
				t.Errorf("Iter %d: Expected TimeoutMs %d, got %d", i, i, loaded.Ollama.TimeoutMs)
			}
			if loaded.UI.Prompt != cfg.UI.Prompt {
				t.Errorf("Iter %d: Expected Prompt %s, got %s", i, cfg.UI.Prompt, loaded.UI.Prompt)
			}
			if loaded.History.RetentionDays != -i {
				t.Errorf("Iter %d: Expected RetentionDays %d, got %d", i, -i, loaded.History.RetentionDays)
			}
		}
	})

	t.Run("Load_Malformed_TOML", func(t *testing.T) {
		malformed := []string{
			`bad toml`,
			`[ui] prompt = "missing quote`,
			`[ui` + "\n" + `prompt = "hello"`,
			`ollama = "string_instead_of_table"`,
			`[[ui]]`, // array of tables instead of table
			`ui.prompt = 123`, // type mismatch
		}
		tmpDir := t.TempDir()
		for i, m := range malformed {
			path := filepath.Join(tmpDir, fmt.Sprintf("bad%d.toml", i))
			os.WriteFile(path, []byte(m), 0644)
			_, err := Load(path)
			if err == nil {
				t.Errorf("Expected error for malformed TOML: %s", m)
			}
		}
	})

	t.Run("Load_Missing_Fields_Partial", func(t *testing.T) {
		partials := []struct{
			content string
			check func(Config) bool
		}{
			{`[ui]` + "\n" + `prompt = "hello> "`, func(c Config) bool { return c.UI.Prompt == "hello> " && c.Ollama.TimeoutMs == 45000 }},
			{`[ollama]` + "\n" + `timeout_ms = 50`, func(c Config) bool { return c.Ollama.TimeoutMs == 50 && c.UI.Theme == "default" }},
			{`[search]` + "\n" + `default_engine = "duckduckgo"`, func(c Config) bool { return c.Search.DefaultEngine == "duckduckgo" && c.Search.Engines["google"] != "" }},
		}
		
		tmpDir := t.TempDir()
		for i, p := range partials {
			path := filepath.Join(tmpDir, fmt.Sprintf("partial%d.toml", i))
			os.WriteFile(path, []byte(p.content), 0644)
			cfg, _ := Load(path)
			if !p.check(cfg) {
				t.Errorf("Partial check failed for %s", p.content)
			}
		}
	})

	t.Run("Load_Extra_Fields", func(t *testing.T) {
		content := `
[ui]
prompt = "nsh> "
unknown_field = true

[unknown_section]
foo = "bar"
`
		tmpDir := t.TempDir()
		path := filepath.Join(tmpDir, "extra.toml")
		os.WriteFile(path, []byte(content), 0644)
		
		// Should not error when extra fields exist
		cfg, err := Load(path)
		if err != nil {
			t.Errorf("Expected no error with extra fields, got %v", err)
		}
		if cfg.UI.Prompt != "nsh> " {
			t.Errorf("Expected UI.Prompt to be loaded properly")
		}
	})

	t.Run("EnsureDefaults", func(t *testing.T) {
		// Mock home dir
		tmpDir := t.TempDir()
		if runtime.GOOS == "windows" {
			orig := os.Getenv("APPDATA")
			defer os.Setenv("APPDATA", orig)
			os.Setenv("APPDATA", tmpDir)
		} else {
			orig := os.Getenv("HOME")
			defer os.Setenv("HOME", orig)
			os.Setenv("HOME", tmpDir)
		}

		// First call should create
		cfg1, err := EnsureDefaults()
		if err != nil {
			t.Fatalf("EnsureDefaults create failed: %v", err)
		}
		
		// Second call should load
		cfg2, err := EnsureDefaults()
		if err != nil {
			t.Fatalf("EnsureDefaults load failed: %v", err)
		}
		if cfg1.UI.Prompt != cfg2.UI.Prompt {
			t.Errorf("Config mismatch")
		}
	})

	t.Run("Dir_OS_Specific", func(t *testing.T) {
		d := Dir()
		if d == "" {
			t.Fatal("Dir() empty")
		}
		if runtime.GOOS == "windows" {
			if !strings.HasSuffix(d, "nsh") {
				t.Errorf("Windows Dir should end with nsh: %s", d)
			}
		} else {
			if !strings.Contains(d, ".config/nsh") {
				t.Errorf("Unix Dir should contain .config/nsh: %s", d)
			}
		}
	})

	t.Run("Save_EdgeCases", func(t *testing.T) {
		cfg := Default()
		cfg.UI.Prompt = "" // empty string
		cfg.Ollama.TimeoutMs = -100 // negative
		cfg.History.RetentionDays = 0 // zero
		cfg.Search.Engines["unicode"] = "🔍" // unicode
		cfg.UI.Theme = strings.Repeat("A", 45000) // very large string
		
		tmpDir := t.TempDir()
		path := filepath.Join(tmpDir, "edge.toml")
		if err := Save(cfg, path); err != nil {
			t.Fatalf("Save failed: %v", err)
		}
		
		loaded, err := Load(path)
		if err != nil {
			t.Fatalf("Load failed: %v", err)
		}
		if loaded.UI.Prompt != "" {
			t.Errorf("Expected empty prompt")
		}
		if loaded.Ollama.TimeoutMs != -100 {
			t.Errorf("Expected -100")
		}
		if loaded.History.RetentionDays != 0 {
			t.Errorf("Expected 0")
		}
		if loaded.Search.Engines["unicode"] != "🔍" {
			t.Errorf("Expected unicode engine")
		}
		if len(loaded.UI.Theme) != 45000 {
			t.Errorf("Expected large theme string")
		}
	})

	t.Run("Concurrent_LoadSave", func(t *testing.T) {
		tmpDir := t.TempDir()
		path := filepath.Join(tmpDir, "concurrent.toml")
		Save(Default(), path)
		
		var wg sync.WaitGroup
		for i := 0; i < 20; i++ {
			wg.Add(2)
			go func(val int) {
				defer wg.Done()
				cfg, _ := Load(path)
				cfg.Ollama.TimeoutMs = val
				// Ignore save errors as they can conflict, we just want to ensure it doesn't panic
				_ = Save(cfg, path) 
			}(i)
			go func() {
				defer wg.Done()
				Load(path)
			}()
		}
		wg.Wait()
	})
	
	t.Run("Spaces_In_Path", func(t *testing.T) {
		tmpDir := t.TempDir()
		spacePath := filepath.Join(tmpDir, "dir with spaces", "config.toml")
		err := Save(Default(), spacePath)
		if err != nil {
			t.Fatalf("Save failed: %v", err)
		}
		_, err = Load(spacePath)
		if err != nil {
			t.Fatalf("Load failed: %v", err)
		}
	})
}
