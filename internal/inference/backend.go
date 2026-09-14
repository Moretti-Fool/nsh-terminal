package inference

import (
	"context"
)

// TranslateResult holds the output of an NL-to-command translation.
type TranslateResult struct {
	Commands    []string
	Dialect     string
	Destructive bool
}

// StreamToken is called for each token during streaming generation.
type StreamToken func(token string)

// Backend abstracts the inference engine (Ollama HTTP, embedded llama.cpp, etc.).
type Backend interface {
	// Translate converts natural language to shell commands.
	Translate(ctx context.Context, input string, cwd string, shell string) (TranslateResult, error)

	// Repair takes a failed command and error output, returns a corrected command.
	Repair(ctx context.Context, input string, failedCmd string, errOutput string, cwd string, shell string) (TranslateResult, error)

	// Classify determines if input is a COMMAND or NL.
	Classify(ctx context.Context, input string) (string, error)

	// GenerateStream produces streaming text output for "nsh ask" queries.
	GenerateStream(ctx context.Context, input string, cwd string, onToken StreamToken) (string, error)

	// CheckHealth returns true if the backend is ready.
	CheckHealth() bool

	// Name returns a human-readable name for this backend.
	Name() string
}
