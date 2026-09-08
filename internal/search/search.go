package search

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

type Handler struct {
	engines       map[string]string
	defaultEngine string
}

func New(engines map[string]string, defaultEngine string) *Handler {
	return &Handler{engines: engines, defaultEngine: defaultEngine}
}

func (h *Handler) BuildURL(engine, query string) (string, error) {
	if engine == "" {
		engine = h.defaultEngine
	}
	tmpl, ok := h.engines[engine]
	if !ok {
		return "", fmt.Errorf("unknown search engine: %s (available: %s)",
			engine, strings.Join(h.EngineNames(), ", "))
	}
	encoded := strings.ReplaceAll(url.QueryEscape(query), "%20", "+")
	return strings.Replace(tmpl, "%s", encoded, 1), nil
}

func (h *Handler) Open(engine, query string) (string, error) {
	searchURL, err := h.BuildURL(engine, query)
	if err != nil {
		return "", err
	}
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", searchURL)
	case "darwin":
		cmd = exec.Command("open", searchURL)
	default:
		cmd = exec.Command("xdg-open", searchURL)
	}
	return searchURL, cmd.Start()
}

func (h *Handler) OpenCommand() string {
	switch runtime.GOOS {
	case "windows":
		return "rundll32"
	case "darwin":
		return "open"
	default:
		return "xdg-open"
	}
}

func (h *Handler) EngineNames() []string {
	names := make([]string, 0, len(h.engines))
	for k := range h.engines {
		names = append(names, k)
	}
	return names
}
