package ollama

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestBuildSystemPrompt(t *testing.T) {
	c := New("http://localhost:11434", "qwen2.5-coder:7b", "qwen2.5-coder:7b", 10*time.Second)
	c.SetShellHint("powershell")
	c.SetAvailableBins([]string{"ipconfig", "git"})
	prompt := c.BuildSystemPrompt("/home/user/project")
	if !strings.Contains(prompt, runtime.GOOS) {
		t.Error("should contain OS")
	}
	if !strings.Contains(prompt, "/home/user/project") {
		t.Error("should contain CWD")
	}
	if !strings.Contains(prompt, "powershell") {
		t.Error("should contain shell")
	}
	if !strings.Contains(prompt, "git=yes") && !strings.Contains(prompt, "On PATH: git") {
		t.Error("should mention git as present, not dump PATH")
	}
	if strings.Contains(prompt, "ipconfig") {
		t.Error("should not dump the full bin list")
	}
	if !strings.Contains(prompt, "Get-ChildItem") {
		t.Error("should mention PowerShell cmdlets")
	}
}

func TestSanitizeGenerated(t *testing.T) {
	in := "```powershell\nGet-ChildItem\n```"
	got := SanitizeGenerated(in)
	if got != "Get-ChildItem" {
		t.Fatalf("got %q", got)
	}
}

func TestLooksLikeBinDump(t *testing.T) {
	dump := "ipconfig\nnetstat\ntaskkill\ntasklist\nsc\nnet\nping"
	if !LooksLikeBinDump(dump) {
		t.Fatal("expected dump")
	}
	if LooksLikeBinDump("Get-ChildItem | Sort-Object Length") {
		t.Fatal("command is not a dump")
	}
}

func TestParsePlan(t *testing.T) {
	p, err := ParsePlan(`{"commands":["Get-ChildItem"],"dialect":"powershell"}`)
	if err != nil {
		t.Fatal(err)
	}
	if p.Join() != "Get-ChildItem" || p.Dialect != "powershell" {
		t.Fatalf("%+v", p)
	}
	p, err = ParsePlan("```json\n{\"command\":\"docker ps\",\"dialect\":\"bash\"}\n```")
	if err != nil || p.Join() != "docker ps" {
		t.Fatalf("got %+v err %v", p, err)
	}
	p, err = ParsePlan("Get-ChildItem -Name")
	if err != nil || p.Join() != "Get-ChildItem -Name" {
		t.Fatalf("plain fallback %+v %v", p, err)
	}
}

func TestValidatePlan(t *testing.T) {
	ok := Plan{Commands: []string{"Get-ChildItem"}, Dialect: "powershell"}
	if err := ValidatePlan(ok, "powershell"); err != nil {
		t.Fatal(err)
	}
	bad := Plan{Commands: []string{"dir /b"}, Dialect: "powershell"}
	if err := ValidatePlan(bad, "powershell"); err == nil {
		t.Fatal("expected dialect error")
	}
	dump := Plan{Commands: []string{"ipconfig", "netstat", "taskkill", "tasklist", "sc", "net", "ping"}, Dialect: "powershell"}
	if err := ValidatePlan(dump, "powershell"); err == nil {
		t.Fatal("expected dump error")
	}
}

func TestGenerate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/chat" {
			t.Errorf("expected /api/chat, got %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": map[string]interface{}{
				"role":    "assistant",
				"content": `{"commands":["docker ps"],"dialect":"bash"}`,
			},
			"done": true,
		})
	}))
	defer server.Close()

	c := New(server.URL, "phi3", "llama3.2", 5*time.Second)
	c.SetShellHint("bash")
	result, err := c.Generate(context.Background(), "show running containers", "/tmp")
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if !strings.Contains(result, "docker ps") {
		t.Errorf("got: %s", result)
	}
}

func TestTranslateToolRound(t *testing.T) {
	n := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n++
		if n == 1 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": map[string]interface{}{
					"role": "assistant",
					"tool_calls": []map[string]interface{}{
						{"function": map[string]interface{}{"name": "list_cwd", "arguments": map[string]any{}}},
					},
				},
				"done": true,
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": map[string]interface{}{
				"role":    "assistant",
				"content": `{"commands":["Get-ChildItem | Sort-Object Length -Descending | Select-Object -First 1"],"dialect":"powershell"}`,
			},
			"done": true,
		})
	}))
	defer server.Close()

	c := New(server.URL, "phi3", "llama3.2", 5*time.Second)
	c.SetShellHint("powershell")
	called := false
	plan, _, err := c.Translate(context.Background(), TranslateRequest{
		Input: "largest file",
		Env:   Env{CWD: "/tmp", Shell: "powershell"},
		RunTool: func(name string, args map[string]any) string {
			called = true
			if name != "list_cwd" {
				t.Fatalf("tool %s", name)
			}
			return "a.txt 10"
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !called {
		t.Fatal("expected tool call")
	}
	if !strings.Contains(plan.Join(), "Sort-Object") {
		t.Fatalf("got %s", plan.Join())
	}
	if n != 2 {
		t.Fatalf("expected 2 chat calls, got %d", n)
	}
}

func TestClassifyInput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"response": `{"label":"NL"}`, "done": true})
	}))
	defer server.Close()

	c := New(server.URL, "phi3", "llama3.2", 5*time.Second)
	result, _ := c.ClassifyInput(context.Background(), "unzip the archive")
	if result != "NL" {
		t.Errorf("got: %s", result)
	}
}

func TestCheckHealth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ollama is running"))
	}))
	defer server.Close()
	if !New(server.URL, "phi3", "llama3.2", 5*time.Second).CheckHealth() {
		t.Error("expected healthy")
	}
}

func TestCheckHealthDown(t *testing.T) {
	if New("http://localhost:99999", "phi3", "llama3.2", 1*time.Second).CheckHealth() {
		t.Error("expected unhealthy")
	}
}

func TestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
	}))
	defer server.Close()

	c := New(server.URL, "phi3", "llama3.2", 500*time.Millisecond)
	_, err := c.Generate(context.Background(), "test", "/tmp")
	if err == nil {
		t.Error("expected timeout error")
	}
}

func TestMatchWorkflow(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"response": "run-backend", "done": true})
	}))
	defer server.Close()

	c := New(server.URL, "phi3", "llama3.2", 5*time.Second)
	result, _ := c.MatchWorkflow(context.Background(), "start the backend", []string{"run-backend", "deploy"})
	if result != "run-backend" {
		t.Errorf("got: %s", result)
	}
}

func TestPickGenerationModel(t *testing.T) {
	models := []ModelInfo{{Name: "nomic-embed"}, {Name: "llama3.2:3b"}, {Name: "qwen2.5-coder:7b"}}
	if got := PickGenerationModel("qwen2.5-coder:7b", models); got != "qwen2.5-coder:7b" {
		t.Fatalf("got %s", got)
	}
	if got := PickGenerationModel("missing", models); got != "qwen2.5-coder:7b" {
		t.Fatalf("fallback got %s", got)
	}
}
