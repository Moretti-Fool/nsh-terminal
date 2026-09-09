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
		httpClient:      &http.Client{Timeout: timeout},
	}
}

type generateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	System string `json:"system,omitempty"`
	Stream bool   `json:"stream"`
}

type generateResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func (c *Client) BuildSystemPrompt(cwd string) string {
	shellName := c.shellHint
	if shellName == "" {
		if runtime.GOOS == "windows" {
			shellName = "powershell"
		} else {
			shellName = "bash"
		}
	}
	bins := ""
	if len(c.availableBins) > 0 {
		bins = "\nAvailable: " + strings.Join(c.availableBins, ", ")
	}
	shellNote := ""
	switch shellName {
	case "powershell", "pwsh":
		shellNote = `
Use PowerShell syntax (cmdlets, pipelines, $_).
Prefer Get-ChildItem, Sort-Object, Select-Object, Get-Process, Get-Service, Get-NetTCPConnection, Restart-Service.
To search file contents, use Get-ChildItem -Recurse -File | Select-String -Pattern 'text' — do not Get-Content a guessed filename.
Do not use cmd.exe switches such as dir /b, dir /w, dir /o-s — those are not PowerShell.
Native binaries listed after Available: may be called when needed — never print that list.
Output at most 3 commands. No Out-GridView, Read-Host, pause, more, or other interactive UI.`
	case "cmd":
		shellNote = `
Use ONLY cmd.exe syntax. Do NOT use PowerShell cmdlets.
Use dir, sort, findstr, type, more, for, forfiles.
Example: dir /o-s`
	}
	return fmt.Sprintf(`You are nsh, a terminal command translator.
OS: %s
Shell: %s
CWD: %s%s

Respond with ONLY the shell command(s) to execute.
One command per line. No explanations. No markdown. No code fences.
If multiple commands are needed, separate with newlines.%s`, runtime.GOOS, shellName, cwd, bins, shellNote)
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
	resp, err := c.doGenerate(ctx, generateRequest{
		Model:  c.generationModel,
		Prompt: fmt.Sprintf("User wants to: %s", input),
		System: c.BuildSystemPrompt(cwd),
		Stream: false,
	})
	if err != nil {
		return "", err
	}
	return SanitizeGenerated(resp.Response), nil
}

func (c *Client) RepairCommand(ctx context.Context, input, cwd, failedCmd, errOutput string) (string, error) {
	system := c.BuildSystemPrompt(cwd) + `

The previous command failed. Emit a corrected command for this OS and shell.
Failed command: ` + failedCmd + `
Error:
` + trimForPrompt(errOutput, 800)
	resp, err := c.doGenerate(ctx, generateRequest{
		Model:  c.generationModel,
		Prompt: fmt.Sprintf("User still wants to: %s", input),
		System: system,
		Stream: false,
	})
	if err != nil {
		return "", err
	}
	return SanitizeGenerated(resp.Response), nil
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
		System: `Classify the following input as one of: COMMAND, NL
COMMAND = the user is invoking a program or builtin (first word is the executable), e.g. git status, ls -la, docker ps
NL = English describing a goal or asking a question, even if it does not use words like "please" or "show me"
If the first word is not a real command, classify as NL.
Respond with exactly one word: COMMAND or NL. Nothing else.`,
		Stream: false,
	})
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(strings.ToUpper(resp.Response))
	if result != "COMMAND" && result != "NL" {
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
		Stream: false,
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

func (c *Client) GeneratePython(ctx context.Context, input string, cwd string) (string, error) {
	system := fmt.Sprintf(`You are nsh, a Python script generator.
Generate a complete, runnable Python script for the user's request.
Include all necessary imports at the top.
Output ONLY the Python code. No explanations. No markdown. No code fences.
If the script works with files, use paths relative to the current directory.
CWD: %s`, cwd)

	resp, err := c.doGenerate(ctx, generateRequest{
		Model:  c.generationModel,
		Prompt: fmt.Sprintf("Write a Python script to: %s", input),
		System: system,
		Stream: false,
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
		Model:  c.generationModel,
		Prompt: prompt,
		System: system,
		Stream: true,
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
