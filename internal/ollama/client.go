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
	"time"
)

type Client struct {
	baseURL         string
	classifierModel string
	generationModel string
	shellHint       string
	availableBins   []string
	timeout         time.Duration
	httpClient      *http.Client
}

func New(baseURL, classifierModel, generationModel string, timeout time.Duration) *Client {
	return &Client{
		baseURL:         strings.TrimRight(baseURL, "/"),
		classifierModel: classifierModel,
		generationModel: generationModel,
		timeout:         timeout,
		// Request context enforces timeout so a tool round can use two chats.
		httpClient: &http.Client{},
	}
}

type generateRequest struct {
	Model     string          `json:"model"`
	Prompt    string          `json:"prompt"`
	System    string          `json:"system,omitempty"`
	Stream    bool            `json:"stream"`
	Format    json.RawMessage `json:"format,omitempty"`
	Options   map[string]any  `json:"options,omitempty"`
	KeepAlive string          `json:"keep_alive,omitempty"`
}

type generateResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func (c *Client) ShellName() string {
	if c.shellHint != "" {
		return c.shellHint
	}
	if runtime.GOOS == "windows" {
		return "powershell"
	}
	return "bash"
}

func (c *Client) BuildSystemPrompt(cwd string) string {
	return c.buildTranslatorPrompt(Env{CWD: cwd, Shell: c.ShellName()}, nil)
}

func SanitizeGenerated(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "```powershell")
	s = strings.TrimPrefix(s, "```powershell\n")
	s = strings.TrimPrefix(s, "```pwsh")
	s = strings.TrimPrefix(s, "```cmd")
	s = strings.TrimPrefix(s, "```bash")
	s = strings.TrimPrefix(s, "```sh")
	s = strings.TrimPrefix(s, "```ps1")
	s = strings.TrimPrefix(s, "```json")
	s = strings.TrimPrefix(s, "```")
	s = strings.TrimSuffix(s, "```")
	s = strings.TrimSpace(s)

	var lines []string
	for _, line := range strings.Split(s, "\n") {
		trim := strings.TrimSpace(line)
		if trim == "" || trim == "```" {
			continue
		}
		if strings.HasPrefix(trim, "```") {
			continue
		}
		lower := strings.ToLower(trim)
		if strings.HasPrefix(lower, "here is") || strings.HasPrefix(lower, "sure,") {
			continue
		}
		lines = append(lines, trim)
	}
	return strings.Join(lines, "\n")
}

func LooksLikeBinDump(s string) bool {
	var lines []string
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	if len(lines) < 5 {
		return false
	}
	simple := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 1 && !strings.ContainsAny(fields[0], `/\`) {
			simple++
		}
	}
	return simple >= (len(lines)*3)/4
}

func (c *Client) Generate(ctx context.Context, input string, cwd string) (string, error) {
	plan, _, err := c.Translate(ctx, TranslateRequest{
		Input: input,
		Env:   Env{CWD: cwd, Shell: c.ShellName()},
	})
	if err != nil {
		return "", err
	}
	return plan.Join(), nil
}

func (c *Client) RepairCommand(ctx context.Context, input, cwd, failedCmd, errOutput string) (string, error) {
	env := Env{CWD: cwd, Shell: c.ShellName()}
	system := c.buildTranslatorPrompt(env, nil)
	msgs := []ChatMessage{
		{Role: "system", Content: system},
		{Role: "user", Content: buildUserMessage(input, env)},
		{Role: "assistant", Content: failedCmd},
	}
	plan, _, err := c.Repair(ctx, msgs, errOutput)
	if err != nil {
		return "", err
	}
	return plan.Join(), nil
}

func trimForPrompt(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func (c *Client) ClassifyInput(ctx context.Context, input string) (string, error) {
	resp, err := c.doGenerate(ctx, generateRequest{
		Model:  c.classifierModel,
		Prompt: input,
		System: `Classify the following input as one of: COMMAND, NL, HISTORY
COMMAND = the user is invoking a program or builtin (first word is the executable), e.g. git status, ls -la, docker ps
NL = English describing a goal or asking a question, even if it does not use words like "please" or "show me"
HISTORY = the user is asking to search or list past commands, workflows, or categories.
If the first word is not a real command, classify as NL or HISTORY.
Respond with JSON {"label":"COMMAND"}, {"label":"NL"}, or {"label":"HISTORY"}.`,
		Stream:    false,
		Format:    classifyFormat,
		Options:   genOptions(),
		KeepAlive: "10m",
	})
	if err != nil {
		return "", err
	}
	raw := strings.TrimSpace(resp.Response)
	var dto struct {
		Label string `json:"label"`
	}
	if json.Unmarshal([]byte(extractJSONObject(raw)), &dto) == nil {
		result := strings.ToUpper(strings.TrimSpace(dto.Label))
		if result == "COMMAND" || result == "NL" || result == "HISTORY" {
			return result, nil
		}
	}
	result := strings.TrimSpace(strings.ToUpper(raw))
	if strings.Contains(result, "HISTORY") {
		return "HISTORY", nil
	}
	if strings.Contains(result, "COMMAND") && !strings.Contains(result, "NL") {
		return "COMMAND", nil
	}
	if result != "COMMAND" && result != "NL" && result != "HISTORY" {
		return "NL", nil
	}
	return result, nil
}

func (c *Client) MatchWorkflow(ctx context.Context, input string, workflows []string) (string, error) {
	resp, err := c.doGenerate(ctx, generateRequest{
		Model:  c.classifierModel,
		Prompt: input,
		System: fmt.Sprintf(`You are matching user input to a saved workflow name.
Available workflows: %s
If the user's input matches one of these workflows, respond with EXACTLY the workflow name.
If none match, respond with NONE.
Respond with just the name or NONE. Nothing else.`, strings.Join(workflows, ", ")),
		Stream:    false,
		Options:   genOptions(),
		KeepAlive: "10m",
	})
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(resp.Response)
	for _, wf := range workflows {
		if strings.EqualFold(result, wf) {
			return wf, nil
		}
	}
	return "", nil
}

func (c *Client) ExtractDomain(ctx context.Context, input string) (string, error) {
	resp, err := c.doGenerate(ctx, generateRequest{
		Model:  c.classifierModel,
		Prompt: input,
		System: `Extract the requested category/domain from the user's input.
If the user says "commands under network", extract "network".
If the user says "list all git commands", extract "git".
Respond with just the domain name. Nothing else.`,
		Stream:    false,
		Options:   genOptions(),
		KeepAlive: "10m",
	})
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(resp.Response), nil
}

func (c *Client) GeneratePython(ctx context.Context, input string, cwd string) (string, error) {
	system := fmt.Sprintf(`You are nsh, a Python script generator.
Generate a complete, runnable Python script for the user's request.
Include all necessary imports at the top.
Output ONLY the Python code. No explanations. No markdown. No code fences.
If the script works with files, use paths relative to the current directory.
CWD: %s`, cwd)

	resp, err := c.doGenerate(ctx, generateRequest{
		Model:     c.generationModel,
		Prompt:    fmt.Sprintf("Write a Python script to: %s", input),
		System:    system,
		Stream:    false,
		Options:   genOptions(),
		KeepAlive: "10m",
	})
	if err != nil {
		return "", err
	}

	script := strings.TrimSpace(resp.Response)
	script = strings.TrimPrefix(script, "```python\n")
	script = strings.TrimPrefix(script, "```py\n")
	script = strings.TrimPrefix(script, "```\n")
	script = strings.TrimSuffix(script, "\n```")
	script = strings.TrimSuffix(script, "```")
	return strings.TrimSpace(script), nil
}

func (c *Client) GenerateStream(ctx context.Context, input string, cwd string, onToken func(string)) (string, error) {
	prompt := fmt.Sprintf("User wants to: %s", input)
	system := c.BuildSystemPrompt(cwd)
	if cwd == "" {
		prompt = input
		system = "You are a helpful assistant. Answer questions concisely and accurately."
	}
	body := generateRequest{
		Model:     c.generationModel,
		Prompt:    prompt,
		System:    system,
		Stream:    true,
		KeepAlive: "10m",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(reqCtx, "POST", c.baseURL+"/api/generate", bytes.NewReader(jsonBody))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var full strings.Builder
	decoder := json.NewDecoder(resp.Body)
	for {
		var chunk generateResponse
		if err := decoder.Decode(&chunk); err != nil {
			if err == io.EOF {
				break
			}
			return full.String(), err
		}
		full.WriteString(chunk.Response)
		if onToken != nil {
			onToken(chunk.Response)
		}
		if chunk.Done {
			break
		}
	}
	out := strings.TrimSpace(full.String())
	if cwd != "" {
		return SanitizeGenerated(out), nil
	}
	return out, nil
}

type ModelInfo struct {
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	ModifiedAt string `json:"modified_at"`
}

func (c *Client) ListModels() ([]ModelInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/api/tags", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Models []ModelInfo `json:"models"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Models, nil
}

func (c *Client) SetGenerationModel(model string) {
	c.generationModel = model
}

func (c *Client) SetClassifierModel(model string) {
	c.classifierModel = model
}

func (c *Client) SetShellHint(hint string) {
	c.shellHint = hint
}

func (c *Client) SetAvailableBins(bins []string) {
	c.availableBins = bins
}

func (c *Client) GenerationModel() string {
	return c.generationModel
}

func (c *Client) ClassifierModel() string {
	return c.classifierModel
}

func (c *Client) CheckHealth() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/", nil)
	if err != nil {
		return false
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func (c *Client) DoGenerateRaw(ctx context.Context, prompt string, format json.RawMessage) (string, error) {
	resp, err := c.doGenerate(ctx, generateRequest{
		Model:     c.classifierModel,
		Prompt:    prompt,
		Format:    format,
		Stream:    false,
		Options:   genOptions(),
		KeepAlive: "10m",
	})
	if err != nil {
		return "", err
	}
	return resp.Response, nil
}

func (c *Client) doGenerate(ctx context.Context, req generateRequest) (*generateResponse, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(reqCtx, "POST", c.baseURL+"/api/generate", bytes.NewReader(jsonBody))
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
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama returned %d: %s", resp.StatusCode, string(body))
	}

	var result generateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
