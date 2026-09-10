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
	c := New("http://localhost:11434", "phi3", "llama3.2", 10*time.Second)
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
	if !strings.Contains(prompt, "ipconfig") {
		t.Error("should contain available bins")
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

func TestGenerate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"response": "docker ps", "done": true,
		})
	}))
	defer server.Close()

	c := New(server.URL, "phi3", "llama3.2", 5*time.Second)
	result, err := c.Generate(context.Background(), "show running containers", "/tmp")
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if !strings.Contains(result, "docker ps") {
		t.Errorf("got: %s", result)
	}
}

func TestClassifyInput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"response": "NL", "done": true})
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
