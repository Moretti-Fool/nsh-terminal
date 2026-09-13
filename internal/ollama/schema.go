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
    "label": {"type": "string", "enum": ["COMMAND", "NL", "HISTORY"]}
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
	Commands  any            `json:"commands"`
	Command   string         `json:"command"`
	Dialect   string         `json:"dialect"`
	Name      string         `json:"name"`
	Arguments any            `json:"arguments"`
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
		var cmds []string
		if cList, ok := dto.Commands.([]any); ok {
			for _, cItem := range cList {
				if s, ok := cItem.(string); ok {
					cmds = append(cmds, s)
				} else if m, ok := cItem.(map[string]any); ok {
					if s, ok := m["command"].(string); ok {
						cmds = append(cmds, s)
					}
				}
			}
		}
		if len(cmds) == 0 && strings.TrimSpace(dto.Command) != "" {
			cmds = []string{dto.Command}
		}
		if len(cmds) == 0 && dto.Name != "" {
			cmdStr := ""
			if dto.Name == "run_command" {
				if argMap, ok := dto.Arguments.(map[string]any); ok {
					if c, ok := argMap["command"].(string); ok {
						cmdStr = c
					}
				} else if argStr, ok := dto.Arguments.(string); ok {
					var m map[string]any
					if json.Unmarshal([]byte(argStr), &m) == nil {
						if c, ok := m["command"].(string); ok {
							cmdStr = c
						}
					}
				}
			} else {
				cmdStr = dto.Name
			}
			if cmdStr != "" {
				cmds = []string{cmdStr}
			}
		}

		expanded := expandCommandLines(cmds)
		if len(expanded) > 0 {
			dialect := strings.ToLower(strings.TrimSpace(dto.Dialect))
			if dialect == "" {
				dialect = "powershell"
			}
			return Plan{Commands: expanded, Dialect: dialect}, nil
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
