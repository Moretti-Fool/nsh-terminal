package executor

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/Moretti-Fool/nsh-terminal/internal/config"
	"github.com/Moretti-Fool/nsh-terminal/internal/ground"
	"github.com/Moretti-Fool/nsh-terminal/internal/ollama"
)

type nlCase struct {
	input      string
	execute    bool
	mustHint   []string // optional substrings we hope to see in generated command (case-insensitive)
	forbidHint []string
}

func TestLiveNLScenarios(t *testing.T) {
	if os.Getenv("NSH_LIVE_NL") != "1" {
		t.Skip("set NSH_LIVE_NL=1 to run live Ollama eval")
	}
	if runtime.GOOS != "windows" {
		t.Skip("windows NL/PowerShell eval")
	}
	cfg, err := config.Load(filepath.Join(config.Dir(), "config.toml"))
	if err != nil {
		cfg = config.Default()
	}
	client := ollama.New(cfg.Ollama.URL, cfg.Ollama.ClassifierModel, cfg.Ollama.GenerationModel, 45*time.Second)
	if !client.CheckHealth() {
		t.Skip("Ollama not running — skip live NL eval")
	}
	models, err := client.ListModels()
	if err != nil || len(models) == 0 {
		t.Skip("no Ollama models")
	}
	model := ollama.PickGenerationModel(cfg.Ollama.GenerationModel, models)
	if model == "" {
		t.Skip("no generate-capable Ollama model (only embeddings?)")
	}
	client.SetGenerationModel(model)
	client.SetClassifierModel(model)

	e := New("auto", "")
	defer e.Close()
	client.SetShellHint(e.NLShellName())
	shell := e.NLShellName()
	cwd, _ := os.Getwd()
	snap := ground.Capture(cwd, e.PathExists)

	cases := liveNLCases()
	if len(cases) < 50 {
		t.Fatalf("need 50+ cases, got %d", len(cases))
	}

	var pass, fail, genOnly int
	fmt.Printf("\n=== live NL eval (%d cases, model %s) ===\n", len(cases), client.GenerationModel())

	for i, tc := range cases {
		ctx := context.Background()
		genStart := time.Now()
		plan, msgs, err := client.Translate(ctx, ollama.TranslateRequest{
			Input: tc.input,
			Env: ollama.Env{
				CWD:     snap.CWD,
				Shell:   shell,
				Listing: snap.Listing,
				Present: snap.Present,
			},
			RunTool: func(name string, args map[string]any) string {
				return ground.ExecTool(name, args, cwd, e.PathExists)
			},
		})
		genMs := time.Since(genStart).Milliseconds()
		if err != nil {
			fail++
			t.Logf("[%02d] FAIL generate %q: %v (%dms)", i+1, tc.input, err, genMs)
			continue
		}
		if vErr := ollama.ValidatePlan(plan, shell); vErr != nil {
			repaired, newMsgs, rerr := client.Repair(ctx, msgs, vErr.Error())
			if rerr == nil && ollama.ValidatePlan(repaired, shell) == nil {
				plan = repaired
				msgs = newMsgs
			}
		}
		generated := plan.Join()
		if generated == "" {
			fail++
			t.Logf("[%02d] FAIL empty command for %q (%dms)", i+1, tc.input, genMs)
			continue
		}

		low := strings.ToLower(generated)
		hintOK := true
		for _, h := range tc.mustHint {
			if !strings.Contains(low, strings.ToLower(h)) {
				hintOK = false
			}
		}
		for _, h := range tc.forbidHint {
			if strings.Contains(low, strings.ToLower(h)) {
				hintOK = false
			}
		}

		execMs := int64(0)
		status := "GEN"
		if tc.execute && !e.IsDestructive(generated) {
			runStart := time.Now()
			var last RunResult
			ok := true
			for _, line := range strings.Split(generated, "\n") {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				last, err = e.RunGenerated(line)
				if err != nil || last.ExitCode != 0 {
					ok = false
					reason := last.Output
					if reason == "" {
						reason = fmt.Sprintf("exit %d", last.ExitCode)
					}
					repaired, _, rerr := client.Repair(ctx, msgs, reason)
					if rerr == nil && repaired.Join() != "" && repaired.Join() != generated && ollama.ValidatePlan(repaired, shell) == nil {
						generated = repaired.Join()
						last, err = e.RunGenerated(generated)
						ok = err == nil && last.ExitCode == 0
					}
					break
				}
			}
			execMs = time.Since(runStart).Milliseconds()
			if ok {
				status = "OK"
				pass++
			} else {
				status = "FAIL"
				fail++
			}
		} else {
			status = "GEN"
			genOnly++
			if hintOK {
				pass++
			} else {
				fail++
				status = "HINT"
			}
		}
		snip := generated
		if len(snip) > 120 {
			snip = snip[:120] + "…"
		}
		fmt.Printf("[%02d] %s gen=%4dms exec=%4dms  %q\n     → %s\n", i+1, status, genMs, execMs, tc.input, snip)
	}
	fmt.Printf("=== results: pass=%d fail=%d gen-only-counted=%d total=%d ===\n", pass, fail, genOnly, len(cases))
	if fail > len(cases)/2 {
		t.Fatalf("too many failures: %d/%d", fail, len(cases))
	}
}

func liveNLCases() []nlCase {
	return []nlCase{
		{input: "whats the biggest file in this folder", execute: true},
		{input: "whats the largest file in this folder", execute: true},
		{input: "whats the largest file", execute: true},
		{input: "list the files with their size", execute: true},
		{input: "list the files with their size and timestamp in desc order", execute: true},
		{input: "show files sorted by size", execute: true},
		{input: "list files sorted by name", execute: true},
		{input: "show hidden files", execute: true},
		{input: "count the files in this folder", execute: true},
		{input: "how many files are here", execute: true},
		{input: "list only directories", execute: true},
		{input: "list only files not folders", execute: true},
		{input: "show the newest file", execute: true},
		{input: "show the oldest file", execute: true},
		{input: "list go files", execute: true},
		{input: "find files named README.md", execute: true},
		{input: "print the current directory", execute: true},
		{input: "what is my cwd", execute: true},
		{input: "show today's date", execute: true},
		{input: "what time is it", execute: true},
		{input: "print my username", execute: true},
		{input: "show hostname", execute: true},
		{input: "print environment variable PATH first line only", execute: true},
		{input: "show git status", execute: true},
		{input: "show recent git commits", execute: true},
		{input: "what branch am i on", execute: true},
		{input: "list environment variables starting with USER", execute: true},
		{input: "show disk drives", execute: true},
		{input: "get powershell version", execute: true},
		{input: "list running processes named powershell", execute: true},
		{input: "show my ip config summary", execute: true, mustHint: []string{"ipconfig"}},
		{input: "flush my DNS cache", execute: false, mustHint: []string{"flush"}},
		{input: "kill the process using port 8080", execute: false},
		{input: "what process is using port 8080", execute: true},
		{input: "restart the print spooler service", execute: false, mustHint: []string{"spooler"}},
		{input: "list windows services that are running", execute: true},
		{input: "show the print spooler service status", execute: true},
		{input: "ping localhost once", execute: true},
		{input: "display dns cache first few entries", execute: true},
		{input: "list files larger than 1MB", execute: true},
		{input: "recursively find go files", execute: true},
		{input: "show the first 5 lines of README.md", execute: true},
		{input: "count lines in go.mod", execute: true},
		{input: "check if go.mod exists", execute: true},
		{input: "get the size of README.md", execute: true},
		{input: "list files modified today", execute: true},
		{input: "show system uptime or last boot time", execute: true},
		{input: "list local users is too heavy skip to whoami", execute: true},
		{input: "who am i", execute: true},
		{input: "show tcp connections on port 443", execute: true},
		{input: "get child items as a table with name and length", execute: true},
		{input: "sort files by last write time descending", execute: true},
		{input: "show the 3 largest files", execute: true},
		{input: "echo hello from nsh eval", execute: true},
		{input: "list modules currently loaded in powershell first 5", execute: true},
		{input: "what is the computer name", execute: true},
		{input: "show os caption or windows version", execute: true},
		{input: "list files with extension .md", execute: true},
		{input: "find todo comments in go files", execute: true},
		{input: "show free disk space on C", execute: true},
	}
}
