package inference

import (
	"context"

	"github.com/Moretti-Fool/nsh-terminal/internal/ollama"
)

// OllamaBackend implements the Backend interface using an ollama.Client.
type OllamaBackend struct {
	client *ollama.Client
}

// Ensure OllamaBackend implements Backend at compile time.
var _ Backend = (*OllamaBackend)(nil)

// NewOllamaBackend creates a new OllamaBackend.
func NewOllamaBackend(client *ollama.Client) *OllamaBackend {
	return &OllamaBackend{
		client: client,
	}
}

// Translate converts natural language to shell commands.
func (b *OllamaBackend) Translate(ctx context.Context, input string, cwd string, shell string) (TranslateResult, error) {
	req := ollama.TranslateRequest{
		Input: input,
		Env: ollama.Env{
			CWD:   cwd,
			Shell: shell,
		},
	}
	plan, _, err := b.client.Translate(ctx, req)
	if err != nil {
		return TranslateResult{}, err
	}
	return TranslateResult{
		Commands:    plan.Commands,
		Dialect:     plan.Dialect,
		Destructive: false,
	}, nil
}

// Repair takes a failed command and error output, returns a corrected command.
func (b *OllamaBackend) Repair(ctx context.Context, input string, failedCmd string, errOutput string, cwd string, shell string) (TranslateResult, error) {
	system := b.client.BuildSystemPrompt(cwd)
	userContent := "User wants to: " + input
	
	msgs := []ollama.ChatMessage{
		{Role: "system", Content: system},
		{Role: "user", Content: userContent},
		{Role: "assistant", Content: failedCmd},
	}
	
	plan, _, err := b.client.Repair(ctx, msgs, errOutput)
	if err != nil {
		return TranslateResult{}, err
	}
	
	return TranslateResult{
		Commands:    plan.Commands,
		Dialect:     plan.Dialect,
		Destructive: false,
	}, nil
}

// Classify determines if input is a COMMAND or NL.
func (b *OllamaBackend) Classify(ctx context.Context, input string) (string, error) {
	return b.client.ClassifyInput(ctx, input)
}

// GenerateStream produces streaming text output.
func (b *OllamaBackend) GenerateStream(ctx context.Context, input string, cwd string, onToken StreamToken) (string, error) {
	return b.client.GenerateStream(ctx, input, cwd, func(t string) {
		if onToken != nil {
			onToken(t)
		}
	})
}

// CheckHealth returns true if the backend is ready.
func (b *OllamaBackend) CheckHealth() bool {
	return b.client.CheckHealth()
}

// Name returns a human-readable name for this backend.
func (b *OllamaBackend) Name() string {
	return "Ollama HTTP"
}
