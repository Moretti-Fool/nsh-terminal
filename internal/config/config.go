package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Ollama  OllamaConfig  `toml:"ollama"`
	Shell   ShellConfig   `toml:"shell"`
	Search  SearchConfig  `toml:"search"`
	UI      UIConfig      `toml:"ui"`
	History HistoryConfig `toml:"history"`
}

type OllamaConfig struct {
	URL             string `toml:"url"`
	ClassifierModel string `toml:"classifier_model"`
	GenerationModel string `toml:"generation_model"`
	TimeoutMs       int    `toml:"timeout_ms"`
}

type ShellConfig struct {
	Default   string `toml:"default"`
	WSLDistro string `toml:"wsl_distro"`
}

type SearchConfig struct {
	DefaultEngine string            `toml:"default_engine"`
	OpenBrowser   bool              `toml:"open_browser"`
	AISummary     bool              `toml:"ai_summary"`
	Engines       map[string]string `toml:"engines"`
}

type UIConfig struct {
	Prompt               string `toml:"prompt"`
	ShowGeneratedCommand bool   `toml:"show_generated_command"`
	ConfirmDestructive   bool   `toml:"confirm_destructive"`
	Theme                string `toml:"theme"`
	ShowWelcome          bool   `toml:"show_welcome"`
}

type HistoryConfig struct {
	RetentionDays      int  `toml:"retention_days"`
	CaptureOutput      bool `toml:"capture_output"`
	OutputPreviewChars int  `toml:"output_preview_chars"`
}

func Default() Config {
	return Config{
		Ollama: OllamaConfig{
			URL:             "http://localhost:11434",
			ClassifierModel: "phi3",
			GenerationModel: "llama3.2",
			TimeoutMs:       10000,
		},
		Shell: ShellConfig{
			Default:   "auto",
			WSLDistro: "",
		},
		Search: SearchConfig{
			DefaultEngine: "google",
			OpenBrowser:   true,
			AISummary:     false,
			Engines: map[string]string{
				"google": "https://www.google.com/search?q=%s",
				"wiki":   "https://en.wikipedia.org/wiki/Special:Search/%s",
				"yt":     "https://www.youtube.com/results?search_query=%s",
				"gh":     "https://github.com/search?q=%s",
			},
		},
		UI: UIConfig{
			Prompt:               "nsh> ",
			ShowGeneratedCommand: true,
			ConfirmDestructive:   true,
			Theme:                "default",
			ShowWelcome:          true,
		},
		History: HistoryConfig{
			RetentionDays:      90,
			CaptureOutput:      true,
			OutputPreviewChars: 500,
		},
	}
}

func Dir() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "nsh")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "nsh")
}

func Load(path string) (Config, error) {
	cfg := Default()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}
	err = toml.Unmarshal(data, &cfg)
	return cfg, err
}

func Save(cfg Config, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(cfg)
}

func EnsureDefaults() (Config, error) {
	dir := Dir()
	path := filepath.Join(dir, "config.toml")
	cfg, err := Load(path)
	if err != nil {
		return cfg, err
	}
	if _, statErr := os.Stat(path); os.IsNotExist(statErr) {
		if saveErr := Save(cfg, path); saveErr != nil {
			return cfg, saveErr
		}
	}
	return cfg, nil
}
