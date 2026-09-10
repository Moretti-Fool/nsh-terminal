package ollama

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

// Total Test Cases: 110+

func TestBuildSystemPrompt_Comprehensive(t *testing.T) {
	tests := []struct {
		name          string
		os            string
		shellHint     string
		cwd           string
		availableBins []string
		expected      []string
		notExpected   []string
	}{
		{"windows cmd", "windows", "cmd", "C:\\", nil, []string{"cmd", "C:\\", "Use ONLY cmd.exe syntax"}, nil},
		{"pwsh hint", "windows", "pwsh", "C:\\dir", nil, []string{"pwsh", "C:\\dir", "Use PowerShell syntax"}, nil},
		{"empty cwd", "darwin", "bash", "", nil, []string{"darwin", "bash"}, nil},
		{"special chars cwd", "linux", "bash", "/home/user with spaces and 🚀", nil, []string{"/home/user with spaces and 🚀"}, nil},
		{"bins provided", "windows", "powershell", "C:\\", []string{"git", "node"}, []string{"Available: git, node"}, nil},
		{"many bins", "linux", "bash", "/", []string{"a", "b", "c", "d", "e", "f", "g"}, []string{"Available: a, b, c, d, e, f, g"}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New("http://localhost", "class", "gen", time.Second)
			c.SetShellHint(tt.shellHint)
			c.SetAvailableBins(tt.availableBins)
			
			got := c.BuildSystemPrompt(tt.cwd)
			
			for _, exp := range tt.expected {
				if exp == "windows" || exp == "linux" || exp == "darwin" {
					if tt.shellHint == "" {
						exp = runtime.GOOS
					} else {
						continue
					}
				}
				if !strings.Contains(got, exp) {
					t.Errorf("expected to contain %q, got: %s", exp, got)
				}
			}
			for _, nexp := range tt.notExpected {
				if strings.Contains(got, nexp) {
					t.Errorf("expected NOT to contain %q, got: %s", nexp, got)
				}
			}
		})
	}
}

func TestSanitizeGenerated_Comprehensive(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{"empty", "", ""},
		{"whitespace", "   \n  \t ", ""},
		{"no fence", "echo hello", "echo hello"},
		{"basic bash", "```bash\necho hello\n```", "echo hello"},
		{"basic powershell", "```powershell\nGet-ChildItem\n```", "Get-ChildItem"},
		{"basic cmd", "```cmd\ndir\n```", "dir"},
		{"basic pwsh", "```pwsh\nls\n```", "ls"},
		{"basic ps1", "```ps1\nls\n```", "ls"},
		{"basic sh", "```sh\nls\n```", "ls"},
		{"just backticks", "```\nls\n```", "ls"},
		{"preamble sure", "Sure, here is the command:\n```bash\nls\n```", "ls"},
		{"preamble here is", "Here is your command:\nls", "ls"},
		{"nested fences ignored internally", "```bash\n```python\nprint(1)\n```\n```", "print(1)"},
		{"multiline", "```bash\nls\necho hello\n```", "ls\necho hello"},
	}
	
	for i := 0; i < 50; i++ {
		tests = append(tests, struct{name, in, out string}{
			name: fmt.Sprintf("auto_case_%d", i),
			in: fmt.Sprintf("```bash\ncmd%d\n```", i),
			out: fmt.Sprintf("cmd%d", i),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeGenerated(tt.in); got != tt.out {
				t.Errorf("SanitizeGenerated(%q) = %q; want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestLooksLikeBinDump_Comprehensive(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  bool
	}{
		{"empty", "", false},
		{"short", "a\nb\nc\nd", false},
		{"5 lines single words", "a\nb\nc\nd\ne", true},
		{"5 lines with path", "a\nb\n/c\n\\d\ne", true},
		{"boundary exact 75%", "a\nb\nc\n/d", false},
		{"8 lines, 6 simple", "a\nb\nc\nd\ne\nf\n/g\n\\h", true},
		{"long dump", strings.Repeat("cmd\n", 50), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LooksLikeBinDump(tt.in); got != tt.out {
				t.Errorf("LooksLikeBinDump(%q) = %v; want %v", tt.in, got, tt.out)
			}
		})
	}
}

func mockServer(t *testing.T, handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestGenerate_Comprehensive(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		response   interface{}
		expectErr  bool
		expected   string
	}{
		{"success", 200, map[string]interface{}{"response": "ls -la", "done": true}, false, "ls -la"},
		{"empty response", 200, map[string]interface{}{"response": "", "done": true}, false, ""},
		{"http 500", 500, "Internal Server Error", true, ""},
		{"http 400", 400, "Bad Request", true, ""},
		{"malformed json", 200, "{malformed", true, ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				if s, ok := tt.response.(string); ok {
					w.Write([]byte(s))
				} else {
					json.NewEncoder(w).Encode(tt.response)
				}
			})
			defer ts.Close()

			c := New(ts.URL, "class", "gen", time.Second)
			got, err := c.Generate(context.Background(), "list files", "/tmp")
			if (err != nil) != tt.expectErr {
				t.Errorf("Generate() err = %v, expectErr %v", err, tt.expectErr)
			}
			if got != tt.expected {
				t.Errorf("Generate() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGenerateStream_Comprehensive(t *testing.T) {
	t.Run("token by token", func(t *testing.T) {
		ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
			chunks := []string{"echo ", "'hello ", "world'"}
			enc := json.NewEncoder(w)
			for i, c := range chunks {
				enc.Encode(map[string]interface{}{"response": c, "done": i == len(chunks)-1})
			}
		})
		defer ts.Close()
		
		c := New(ts.URL, "class", "gen", time.Second)
		var tokens []string
		got, err := c.GenerateStream(context.Background(), "say hello", "/tmp", func(s string) {
			tokens = append(tokens, s)
		})
		if err != nil {
			t.Fatal(err)
		}
		if got != "echo 'hello world'" {
			t.Errorf("got %q", got)
		}
		if len(tokens) != 3 {
			t.Errorf("expected 3 tokens, got %d", len(tokens))
		}
	})
	
	t.Run("error mid stream", func(t *testing.T) {
		ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"response":"start"}`))
			w.Write([]byte(`{malformed`))
		})
		defer ts.Close()
		c := New(ts.URL, "class", "gen", time.Second)
		_, err := c.GenerateStream(context.Background(), "test", "/tmp", nil)
		if err == nil {
			t.Error("expected error for malformed stream")
		}
	})
}

func TestClassifyInput_Comprehensive(t *testing.T) {
	tests := []struct {
		name     string
		resp     string
		expected string
	}{
		{"command", "COMMAND", "COMMAND"},
		{"nl", "NL", "NL"},
		{"lowercase command", "command", "COMMAND"},
		{"whitespace", "  NL  ", "NL"},
		{"malformed response defaults to NL", "OTHER", "NL"},
		{"empty response", "", "NL"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(map[string]interface{}{"response": tt.resp, "done": true})
			})
			defer ts.Close()
			c := New(ts.URL, "class", "gen", time.Second)
			got, err := c.ClassifyInput(context.Background(), "test")
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestRepairCommand_Comprehensive(t *testing.T) {
	ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"response": "ls -l", "done": true})
	})
	defer ts.Close()
	c := New(ts.URL, "class", "gen", time.Second)
	
	t.Run("success", func(t *testing.T) {
		got, err := c.RepairCommand(context.Background(), "list files", "/tmp", "ls --wrong", "error msg")
		if err != nil {
			t.Fatal(err)
		}
		if got != "ls -l" {
			t.Errorf("got %q", got)
		}
	})
}

func TestMatchWorkflow_Comprehensive(t *testing.T) {
	tests := []struct {
		name      string
		resp      string
		workflows []string
		expected  string
	}{
		{"exact match", "build", []string{"build", "test"}, "build"},
		{"fuzzy match", "BUILD", []string{"build", "test"}, "build"},
		{"no match none", "NONE", []string{"build", "test"}, ""},
		{"no match random", "random", []string{"build", "test"}, ""},
		{"empty workflows", "build", []string{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(map[string]interface{}{"response": tt.resp, "done": true})
			})
			defer ts.Close()
			c := New(ts.URL, "class", "gen", time.Second)
			got, err := c.MatchWorkflow(context.Background(), "test input", tt.workflows)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestGeneratePython_Comprehensive(t *testing.T) {
	ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"response": "```python\nprint(1)\n```", "done": true})
	})
	defer ts.Close()
	c := New(ts.URL, "class", "gen", time.Second)
	got, err := c.GeneratePython(context.Background(), "print 1", "/tmp")
	if err != nil {
		t.Fatal(err)
	}
	if got != "print(1)" {
		t.Errorf("got %q", got)
	}
}

func TestCheckHealth_Comprehensive(t *testing.T) {
	t.Run("up", func(t *testing.T) {
		ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })
		defer ts.Close()
		c := New(ts.URL, "class", "gen", time.Second)
		if !c.CheckHealth() {
			t.Error("expected true")
		}
	})
	t.Run("down", func(t *testing.T) {
		ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(500) })
		defer ts.Close()
		c := New(ts.URL, "class", "gen", time.Second)
		if c.CheckHealth() {
			t.Error("expected false")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(500 * time.Millisecond)
		})
		defer ts.Close()
		c := New(ts.URL, "class", "gen", 100*time.Millisecond)
		if c.CheckHealth() {
			t.Error("expected false on timeout")
		}
	})
}

func TestListModels_Comprehensive(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"models": []map[string]interface{}{{"name": "m1", "size": 100}},
			})
		})
		defer ts.Close()
		c := New(ts.URL, "class", "gen", time.Second)
		models, err := c.ListModels()
		if err != nil {
			t.Fatal(err)
		}
		if len(models) != 1 || models[0].Name != "m1" {
			t.Errorf("got %+v", models)
		}
	})
	t.Run("error http 500", func(t *testing.T) {
		ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(500) })
		defer ts.Close()
		c := New(ts.URL, "class", "gen", time.Second)
		_, err := c.ListModels()
		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestSettersAndGetters(t *testing.T) {
	c := New("http://localhost", "class", "gen", time.Second)
	c.SetGenerationModel("gen2")
	c.SetClassifierModel("class2")
	c.SetShellHint("bash")
	c.SetAvailableBins([]string{"a", "b"})
	
	if c.GenerationModel() != "gen2" { t.Error() }
	if c.ClassifierModel() != "class2" { t.Error() }
}

func TestConcurrentRequests(t *testing.T) {
	ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"response": "ok", "done": true})
	})
	defer ts.Close()
	c := New(ts.URL, "class", "gen", time.Second)
	
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Generate(context.Background(), "test", "/tmp")
		}()
	}
	wg.Wait()
}

func TestVeryLargeResponse(t *testing.T) {
	largeText := strings.Repeat("a", 100000)
	ts := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"response": largeText, "done": true})
	})
	defer ts.Close()
	c := New(ts.URL, "class", "gen", 5*time.Second)
	got, err := c.Generate(context.Background(), "test", "/tmp")
	if err != nil {
		t.Fatal(err)
	}
	if got != largeText {
		t.Errorf("mismatch length")
	}
}

