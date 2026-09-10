package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
)

// FewShot is a successful past translation from this user, not a built-in recipe.
type FewShot struct {
	Input   string
	Command string
}

// Env is machine evidence injected into the translator.
type Env struct {
	CWD     string
	Shell   string
	Listing string
	Present []string
}

// ToolFunc executes a model-requested inspection tool. Return text for the next chat turn.
type ToolFunc func(name string, args map[string]any) string

// TranslateRequest is one NL-to-command attempt.
type TranslateRequest struct {
	Input    string
	Env      Env
	Examples []FewShot
	RunTool  ToolFunc
}

type chatRequest struct {
	Model     string          `json:"model"`
	Messages  []ChatMessage   `json:"messages"`
	Stream    bool            `json:"stream"`
	Tools     []chatTool      `json:"tools,omitempty"`
	Format    json.RawMessage `json:"format,omitempty"`
	Options   map[string]any  `json:"options,omitempty"`
	KeepAlive string          `json:"keep_alive,omitempty"`
}

type ChatMessage struct {
	Role      string     `json:"role"`
	Content   string     `json:"content,omitempty"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	ToolName  string     `json:"tool_name,omitempty"`
}

type ToolCall struct {
	Function ToolCallFn `json:"function"`
}

type ToolCallFn struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

func (f ToolCallFn) ArgsMap() map[string]any {
	if len(f.Arguments) == 0 {
		return map[string]any{}
	}
	var m map[string]any
	if err := json.Unmarshal(f.Arguments, &m); err == nil {
		return m
	}
	var s string
	if err := json.Unmarshal(f.Arguments, &s); err == nil {
		if json.Unmarshal([]byte(s), &m) == nil {
			return m
		}
	}
	return map[string]any{}
}

type chatTool struct {
	Type     string         `json:"type"`
	Function chatToolSchema `json:"function"`
}

type chatToolSchema struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}

type chatResponse struct {
	Message ChatMessage `json:"message"`
	Done    bool        `json:"done"`
}

func translatorTools() []chatTool {
	emptyObj := map[string]any{"type": "object", "properties": map[string]any{}}
	return []chatTool{
		{
			Type: "function",
			Function: chatToolSchema{
				Name:        "list_cwd",
				Description: "List files and directories in the current working directory with sizes. Use when the request depends on what is actually in this folder.",
				Parameters:  emptyObj,
			},
		},
		{
			Type: "function",
			Function: chatToolSchema{
				Name:        "which",
				Description: "Check whether a program name exists on PATH.",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"name": map[string]any{"type": "string", "description": "Executable name, e.g. git or rg"},
					},
					"required": []string{"name"},
				},
			},
		},
		{
			Type: "function",
			Function: chatToolSchema{
				Name:        "git_status",
				Description: "Show git status --short for the current directory if this is a git repo.",
				Parameters:  emptyObj,
			},
		},
		{
			Type: "function",
			Function: chatToolSchema{
				Name:        "netstat",
				Description: "Show listening TCP ports with their owning process IDs. Use when the request involves ports or network connections.",
				Parameters:  emptyObj,
			},
		},
		{
			Type: "function",
			Function: chatToolSchema{
				Name:        "env_var",
				Description: "Read the value of an environment variable.",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"name": map[string]any{"type": "string"},
					},
					"required": []string{"name"},
				},
			},
		},
		{
			Type: "function",
			Function: chatToolSchema{
				Name:        "read_file_head",
				Description: "Read the first 20 lines of a file to inspect its contents.",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"path": map[string]any{"type": "string"},
					},
					"required": []string{"path"},
				},
			},
		},
	}
}

func genOptions() map[string]any {
	return map[string]any{
		"temperature": 0,
		"num_predict": 256,
	}
}

// Translate turns English into shell commands using chat, JSON schema, and at most one tool round.
func (c *Client) Translate(ctx context.Context, req TranslateRequest) (Plan, []ChatMessage, error) {
	if req.Env.Shell == "" {
		req.Env.Shell = c.ShellName()
	}
	system := c.buildTranslatorPrompt(req.Env, req.Examples)
	user := buildUserMessage(req.Input, req.Env)
	msgs := []ChatMessage{
		{Role: "system", Content: system},
		{Role: "user", Content: user},
	}

	var tools []chatTool
	if req.RunTool != nil {
		tools = translatorTools()
	}

	resp, err := c.doChat(ctx, msgs, tools, CommandFormat)
	if err != nil {
		return Plan{}, msgs, err
	}
	msgs = append(msgs, resp.Message)

	if len(resp.Message.ToolCalls) > 0 && req.RunTool != nil {
		for i, tc := range resp.Message.ToolCalls {
			if i >= 5 {
				break
			}
			out := req.RunTool(tc.Function.Name, tc.Function.ArgsMap())
			msgs = append(msgs, ChatMessage{
				Role:     "tool",
				ToolName: tc.Function.Name,
				Content:  trimForPrompt(out, 2000),
			})
		}
		resp, err = c.doChat(ctx, msgs, nil, CommandFormat)
		if err != nil {
			return Plan{}, msgs, err
		}
		msgs = append(msgs, resp.Message)
		if len(resp.Message.ToolCalls) > 0 {
			msgs = append(msgs, ChatMessage{Role: "user", Content: "Do not call tools. Emit the JSON command plan now."})
			resp, err = c.doChat(ctx, msgs, nil, CommandFormat)
			if err != nil {
				return Plan{}, msgs, err
			}
			msgs = append(msgs, resp.Message)
		}
	}

	plan, err := ParsePlan(resp.Message.Content)
	if err != nil {
		return Plan{}, msgs, err
	}
	if plan.Dialect == "" {
		plan.Dialect = NormalizeShell(req.Env.Shell)
	}
	return plan, msgs, nil
}

// Repair continues the same chat with execution or validation evidence.
func (c *Client) Repair(ctx context.Context, msgs []ChatMessage, reason string) (Plan, []ChatMessage, error) {
	if len(msgs) == 0 {
		return Plan{}, nil, fmt.Errorf("no conversation to repair")
	}
	msgs = append(msgs, ChatMessage{
		Role: "user",
		Content: "The previous command failed with this error:\n" + trimForPrompt(reason, 800) +
			"\nEmit a corrected JSON plan for the original request. Same OS and shell.",
	})
	resp, err := c.doChat(ctx, msgs, nil, CommandFormat)
	if err != nil {
		return Plan{}, msgs, err
	}
	msgs = append(msgs, resp.Message)
	plan, err := ParsePlan(resp.Message.Content)
	if err != nil {
		return Plan{}, msgs, err
	}
	return plan, msgs, nil
}

func (c *Client) doChat(ctx context.Context, msgs []ChatMessage, tools []chatTool, format json.RawMessage) (*chatResponse, error) {
	body := chatRequest{
		Model:     c.generationModel,
		Messages:  msgs,
		Stream:    false,
		Tools:     tools,
		Format:    format,
		Options:   genOptions(),
		KeepAlive: "10m",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(reqCtx, "POST", c.baseURL+"/api/chat", bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama chat returned %d: %s", resp.StatusCode, string(b))
	}

	var result chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func buildUserMessage(input string, env Env) string {
	var b strings.Builder
	b.WriteString("User wants to: ")
	b.WriteString(input)
	if env.Listing != "" {
		b.WriteString("\n\nCWD listing (dirs end with /):\n")
		b.WriteString(env.Listing)
	}
	return b.String()
}

func (c *Client) buildTranslatorPrompt(env Env, examples []FewShot) string {
	shell := env.Shell
	if shell == "" {
		shell = c.ShellName()
	}
	present := ""
	if len(env.Present) > 0 {
		present = "\nOn PATH: " + strings.Join(env.Present, ", ")
	} else if bits := presentBits(c.availableBins); bits != "" {
		present = "\n" + bits
	}

	shellNote := translatorShellNote(shell)
	var ex strings.Builder
	if len(examples) > 0 {
		ex.WriteString("\nRecent successful translations on this machine:\n")
		for _, e := range examples {
			fmt.Fprintf(&ex, "- %q → %s\n", e.Input, e.Command)
		}
	}

	return fmt.Sprintf(`You are nsh, a terminal command translator.
OS: %s
Shell: %s
CWD: %s%s
%s
Respond with JSON only: {"commands":["..."],"dialect":"%s"}
dialect must be exactly %s.
At most 3 commands. No markdown. No explanations.
You may call list_cwd, which, or git_status at most once if you need facts from this machine, then emit the JSON plan.
Do not invent filenames that are not in the listing or tool results.
%s`, runtime.GOOS, shell, env.CWD, present, shellNote, NormalizeShell(shell), NormalizeShell(shell), ex.String())
}

func translatorShellNote(shell string) string {
	switch NormalizeShell(shell) {
	case "powershell":
		return `Use PowerShell (cmdlets, pipelines, $_).
Prefer Get-ChildItem, Sort-Object, Select-Object, Get-Process, Get-Service, Get-NetTCPConnection, Select-String.
To search file contents: Get-ChildItem -Recurse -File | Select-String -Pattern 'text'
Do not use cmd.exe switches such as dir /b, dir /w, dir /o-s.
No Out-GridView, Read-Host, pause, more, or other interactive UI.`
	case "cmd":
		return `Use ONLY cmd.exe syntax. Do NOT use PowerShell cmdlets. Use dir, sort, findstr, type, for, forfiles.`
	default:
		return `Use POSIX shell syntax for this OS.`
	}
}

func presentBits(bins []string) string {
	if len(bins) == 0 {
		return ""
	}
	set := map[string]bool{}
	for _, b := range bins {
		set[strings.ToLower(b)] = true
	}
	var parts []string
	for _, name := range []string{"git", "docker", "rg", "python", "python3"} {
		if set[name] {
			parts = append(parts, name+"=yes")
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return "On PATH: " + strings.Join(parts, " ")
}
