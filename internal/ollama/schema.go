package ollama

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CommandFormat is the JSON schema passed to Ollama constrained decoding.
// The model still invents the command text; only the envelope is constrained.
var CommandFormat = json.RawMessage(`{
  "type": "object",
  "properties": {
    "commands": {
      "type": "array",
      "items": {"type": "string"}
    },
    "dialect": {
      "type": "string",
      "enum": ["powershell", "cmd", "bash"]
    }
  },
  "required": ["commands", "dialect"]
}`)

var classifyFormat = json.RawMessage(`{
  "type": "object",
  "properties": {
    "label": {"type": "string", "enum": ["COMMAND", "NL"]}
  },
  "required": ["label"]
}`)

// Plan is a free-form command list plus the dialect the model claims to have used.
type Plan struct {
	Commands []string `json:"commands"`
	Dialect  string   `json:"dialect"`
}

func (p Plan) Join() string {
	var lines []string
	for _, c := range p.Commands {
		c = strings.TrimSpace(c)
		if c != "" {
			lines = append(lines, c)
		}
	}
	return strings.Join(lines, "\n")
}

type planDTO struct {
	Commands []string `json:"commands"`
	Command  string   `json:"command"`
	Dialect  string   `json:"dialect"`
}

// ParsePlan reads constrained JSON, falling back to sanitized plain command text.
func ParsePlan(s string) (Plan, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return Plan{}, fmt.Errorf("empty model output")
	}

	raw := extractJSONObject(s)
	var dto planDTO
	if raw != "" && json.Unmarshal([]byte(raw), &dto) == nil {
		cmds := dto.Commands
		if len(cmds) == 0 && strings.TrimSpace(dto.Command) != "" {
			cmds = []string{dto.Command}
		}
		expanded := expandCommandLines(cmds)
		if len(expanded) > 0 {
			return Plan{Commands: expanded, Dialect: strings.ToLower(strings.TrimSpace(dto.Dialect))}, nil
		}
	}

	text := SanitizeGenerated(s)
	if text == "" {
		return Plan{}, fmt.Errorf("empty command")
	}
	return Plan{Commands: expandCommandLines([]string{text})}, nil
}

func extractJSONObject(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "```") {
		s = SanitizeGenerated(s)
	}
	start := strings.Index(s, "{")
	end := strings.LastIndex(s, "}")
	if start >= 0 && end > start {
		return s[start : end+1]
	}
	return ""
}

func expandCommandLines(cmds []string) []string {
	var out []string
	for _, c := range cmds {
		for _, line := range strings.Split(c, "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				out = append(out, line)
			}
		}
	}
	if len(out) > 3 {
		out = out[:3]
	}
	return out
}
