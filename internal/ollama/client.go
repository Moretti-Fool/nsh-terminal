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
	shellName := "bash"
	if runtime.GOOS == "windows" {
		shellName = "powershell"
	}
	return fmt.Sprintf(`You are nsh, a terminal command translator.
OS: %s
Shell: %s
CWD: %s

Respond with ONLY the shell command(s) to execute.
One command per line. No explanations. No markdown. No code fences.
If multiple commands are needed, separate with newlines.`, runtime.GOOS, shellName, cwd)
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
	return strings.TrimSpace(resp.Response), nil
}

func (c *Client) ClassifyInput(ctx context.Context, input string) (string, error) {
	resp, err := c.doGenerate(ctx, generateRequest{
		Model:  c.classifierModel,
		Prompt: input,
		System: `Classify the following input as one of: COMMAND, NL
COMMAND = a direct shell command the user wants to run as-is
NL = natural language describing what the user wants to do
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

func (c *Client) GenerateStream(ctx context.Context, input string, cwd string, onToken func(string)) (string, error) {
	body := generateRequest{
		Model:  c.generationModel,
		Prompt: fmt.Sprintf("User wants to: %s", input),
		System: c.BuildSystemPrompt(cwd),
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
	return strings.TrimSpace(full.String()), nil
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
