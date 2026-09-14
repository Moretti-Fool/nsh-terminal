package inference

import (
	"testing"
	"time"

	"github.com/Moretti-Fool/nsh-terminal/internal/ollama"
)

func TestOllamaBackendImplementsBackend(t *testing.T) {
	// This ensures at compile time that OllamaBackend implements Backend.
	// The assignment will fail to compile if it doesn't.
	client := ollama.New("http://localhost:11434", "test-classifier", "test-generation", 5*time.Second)
	var backend Backend = NewOllamaBackend(client)

	if backend.Name() != "Ollama HTTP" {
		t.Errorf("expected backend.Name() to be 'Ollama HTTP', got %s", backend.Name())
	}
}
