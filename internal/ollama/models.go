package ollama

import "strings"

func isEmbeddingOrTiny(name string) bool {
	n := strings.ToLower(name)
	return strings.Contains(n, "embed") || strings.Contains(n, "nomic") ||
		strings.Contains(n, "minilm") || strings.Contains(n, "moondream") ||
		strings.Contains(n, "0.5b")
}

// PickGenerationModel prefers an installed coder/instruct model when the configured name is missing.
func PickGenerationModel(preferred string, models []ModelInfo) string {
	match := func(want string) string {
		if want == "" || isEmbeddingOrTiny(want) {
			return ""
		}
		for _, m := range models {
			if isEmbeddingOrTiny(m.Name) {
				continue
			}
			if strings.EqualFold(m.Name, want) || strings.HasPrefix(strings.ToLower(m.Name), strings.ToLower(want)) {
				return m.Name
			}
		}
		return ""
	}
	if got := match(preferred); got != "" {
		return got
	}
	order := []string{
		"qwen3-coder", "qwen2.5-coder", "qwen3", "qwen2.5",
		"gemma4", "gemma3", "llama3.2", "mistral", "phi3",
	}
	for _, want := range order {
		if got := match(want); got != "" {
			return got
		}
	}
	for _, m := range models {
		if !isEmbeddingOrTiny(m.Name) {
			return m.Name
		}
	}
	return ""
}
