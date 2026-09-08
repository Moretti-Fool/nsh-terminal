package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := Default()
	if cfg.Ollama.URL != "http://localhost:11434" {
		t.Errorf("got %s", cfg.Ollama.URL)
	}
	if cfg.Ollama.ClassifierModel != "phi3" {
		t.Errorf("got %s", cfg.Ollama.ClassifierModel)
	}
	if cfg.UI.Prompt != "nsh> " {
		t.Errorf("got %s", cfg.UI.Prompt)
	}
	if cfg.History.RetentionDays != 90 {
		t.Errorf("got %d", cfg.History.RetentionDays)
	}
	if _, ok := cfg.Search.Engines["google"]; !ok {
		t.Error("missing google engine")
	}
}

func TestConfigDir(t *testing.T) {
	dir := Dir()
	if dir == "" {
		t.Fatal("empty config dir")
	}
	if runtime.GOOS == "windows" {
		expected := filepath.Join(os.Getenv("APPDATA"), "nsh")
		if dir != expected {
			t.Errorf("expected %s, got %s", expected, dir)
		}
	}
}

func TestLoadAndSave(t *testing.T) {
	tmp := t.TempDir()
	cfg := Default()
	cfg.UI.Prompt = "test> "
	path := filepath.Join(tmp, "config.toml")

	if err := Save(cfg, path); err != nil {
		t.Fatalf("save: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if loaded.UI.Prompt != "test> " {
		t.Errorf("got %s", loaded.UI.Prompt)
	}
}

func TestLoadMissing(t *testing.T) {
	cfg, err := Load("/nonexistent/config.toml")
	if err != nil {
		t.Fatalf("should return defaults: %v", err)
	}
	if cfg.Ollama.URL != "http://localhost:11434" {
		t.Error("expected defaults")
	}
}
