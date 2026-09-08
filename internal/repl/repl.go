package repl

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/nsh-terminal/nsh/internal/classifier"
	"github.com/nsh-terminal/nsh/internal/config"
	"github.com/nsh-terminal/nsh/internal/executor"
	"github.com/nsh-terminal/nsh/internal/history"
	"github.com/nsh-terminal/nsh/internal/ollama"
	"github.com/nsh-terminal/nsh/internal/search"
	"github.com/nsh-terminal/nsh/internal/workflow"
)

const Version = "1.0.0"

type REPL struct {
	cfg        config.Config
	classifier *classifier.Classifier
	executor   *executor.Executor
	ollama     *ollama.Client
	history    *history.History
	workflows  *workflow.Manager
	search     *search.Handler
	recording  bool
	recorded   []string
	ollamaOK   bool
}

func New(cfg config.Config) *REPL {
	configDir := config.Dir()
	histDir := filepath.Join(configDir, "history")
	wfDir := filepath.Join(configDir, "workflows")

	wfMgr := workflow.New(wfDir)
	ollamaClient := ollama.New(
		cfg.Ollama.URL,
		cfg.Ollama.ClassifierModel,
		cfg.Ollama.GenerationModel,
		time.Duration(cfg.Ollama.TimeoutMs)*time.Millisecond,
	)

	pathLookup := func(name string) bool {
		_, err := exec.LookPath(name)
		return err == nil
	}

	hist := history.New(histDir, cfg.History.OutputPreviewChars)
	hist.Rotate(cfg.History.RetentionDays)

	r := &REPL{
		cfg:        cfg,
		classifier: classifier.New(pathLookup, wfMgr.Names()),
		executor:   executor.New(cfg.Shell.Default, cfg.Shell.WSLDistro),
		ollama:     ollamaClient,
		history:    hist,
		workflows:  wfMgr,
		search:     search.New(cfg.Search.Engines, cfg.Search.DefaultEngine),
		ollamaOK:   ollamaClient.CheckHealth(),
	}
	r.autoDetectModel()
	return r
}

func (r *REPL) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 64*1024), 64*1024)

	// Handle Ctrl+C gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		for range sigCh {
			if r.recording {
				r.recording = false
				r.recorded = nil
				fmt.Println("\n[nsh] Recording cancelled.")
			}
			fmt.Print("\n")
			r.printPrompt()
		}
	}()

	r.printWelcome()

	for {
		r.printPrompt()
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if input == "exit" || input == "quit" {
			break
		}

		r.handleInput(input)
	}
	return nil
}

func (r *REPL) printPrompt() {
	cwd, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	display := cwd
	if home != "" && strings.HasPrefix(cwd, home) {
		display = "~" + cwd[len(home):]
	}

	theme := r.cfg.UI.Theme
	cwdColor := "\033[32m"   // green
	promptColor := "\033[0m" // reset
	recColor := "\033[31m"   // red
	reset := "\033[0m"

	if theme == "minimal" {
		cwdColor = ""
		promptColor = ""
		recColor = ""
		reset = ""
	} else if theme == "blue" {
		cwdColor = "\033[34m"
	} else if theme == "cyan" {
		cwdColor = "\033[36m"
	} else if theme == "yellow" {
		cwdColor = "\033[33m"
	} else if theme == "magenta" {
		cwdColor = "\033[35m"
	}

	rec := ""
	if r.recording {
		rec = recColor + "[REC]" + reset + " "
	}
	fmt.Printf("%s%s%s%s %s", rec, cwdColor, display, reset, r.cfg.UI.Prompt)
	_ = promptColor
}

func (r *REPL) printWelcome() {
	if !r.cfg.UI.ShowWelcome {
		return
	}
	fmt.Printf("nsh v%s — the natural shell\n", Version)
	fmt.Println("Type commands normally, or use plain English.")
	if !r.ollamaOK {
		fmt.Println("[nsh] Ollama unavailable — NL features disabled, commands still work.")
	}
	fmt.Println("Run \"nsh help\" for more.")
	fmt.Println()
}

func (r *REPL) handleInput(input string) {
	if input == "--help" || input == "-h" || input == "help" {
		r.printHelp()
		return
	}
	if input == "--version" || input == "-v" || input == "version" {
		fmt.Printf("nsh v%s\n", Version)
		return
	}

	r.classifier.UpdateWorkflows(r.workflows.Names())
	result := r.classifier.Classify(input)

	switch result.Type {
	case classifier.Empty:
		return
	case classifier.Command:
		r.handleCommand(input)
	case classifier.NaturalLanguage:
		r.handleNL(input)
	case classifier.Search:
		r.handleSearch(result.SearchEngine, result.SearchQuery)
	case classifier.Workflow:
		r.handleWorkflow(result.WorkflowName)
	case classifier.Builtin:
		r.handleBuiltin(input)
	case classifier.Ambiguous:
		r.handleAmbiguous(input)
	}
}

func (r *REPL) handleCommand(input string) {
	if r.recording {
		r.recorded = append(r.recorded, input)
	}

	if input == "cd" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[nsh] %v\n", err)
			return
		}
		os.Chdir(home)
		r.saveHistory(input, "command", "", 0, "", 0)
		return
	}

	if strings.HasPrefix(input, "cd ") {
		dir := strings.TrimSpace(strings.TrimPrefix(input, "cd "))
		expanded, err := executor.CdExpand(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[nsh] %v\n", err)
			r.saveHistory(input, "command", "", 1, "", 0)
			return
		}
		os.Chdir(expanded)
		r.saveHistory(input, "command", "", 0, "", 0)
		return
	}

	result, err := r.executor.Run(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] execution error: %v\n", err)
		return
	}

	preview := result.Output
	if len(preview) > r.cfg.History.OutputPreviewChars {
		preview = preview[:r.cfg.History.OutputPreviewChars]
	}
	r.saveHistory(input, "command", "", result.ExitCode, preview, result.DurationMs)
}

func (r *REPL) handleNL(input string) {
	if !r.ollamaOK {
		r.ollamaOK = r.ollama.CheckHealth()
	}
	if !r.ollamaOK {
		fmt.Println("[nsh] Ollama not running — cannot process natural language.")
		fmt.Println("      Try: ollama serve")
		return
	}

	cwd, _ := os.Getwd()
	ctx := context.Background()

	fmt.Print("[nsh] thinking...")
	generated, err := r.ollama.Generate(ctx, input, cwd)
	fmt.Print("\r                \r")

	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] Ollama error: %v\n", err)
		return
	}
	if generated == "" {
		fmt.Println("[nsh] No command generated.")
		return
	}

	if r.cfg.UI.ShowGeneratedCommand {
		fmt.Printf("\033[36m> %s\033[0m\n", generated)
	}

	if r.cfg.UI.ConfirmDestructive && r.executor.IsDestructive(generated) {
		fmt.Print("[nsh] This looks destructive. Run it? [y/N] ")
		var confirm string
		fmt.Scanln(&confirm)
		if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
			fmt.Println("[nsh] Cancelled.")
			return
		}
	}

	if r.recording {
		r.recorded = append(r.recorded, generated)
	}

	commands := strings.Split(generated, "\n")
	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}
		if strings.HasPrefix(cmd, "cd ") {
			dir := strings.TrimSpace(strings.TrimPrefix(cmd, "cd "))
			expanded, err := executor.CdExpand(dir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[nsh] %v\n", err)
				r.saveHistory(input, "nl", generated, 1, "", 0)
				return
			}
			os.Chdir(expanded)
			continue
		}
		result, err := r.executor.Run(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[nsh] execution error: %v\n", err)
			r.saveHistory(input, "nl", generated, 1, "", 0)
			return
		}
		if result.ExitCode != 0 {
			r.saveHistory(input, "nl", generated, result.ExitCode, result.Output, result.DurationMs)
			return
		}
		r.saveHistory(input, "nl", generated, result.ExitCode, result.Output, result.DurationMs)
	}
}

func (r *REPL) handleSearch(engine, query string) {
	aiOnly := engine == "ask"
	aiBoth := strings.HasSuffix(engine, "!")
	browserEngine := strings.TrimSuffix(engine, "!")

	if aiOnly || aiBoth {
		r.searchAISummary(query)
	}

	if !aiOnly {
		if browserEngine == "ask" {
			browserEngine = "google"
		}
		searchURL, err := r.search.Open(browserEngine, query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[nsh] search error: %v\n", err)
			return
		}
		fmt.Printf("[nsh] Opened: %s\n", searchURL)
		r.saveHistory(engine+" "+query, "search", searchURL, 0, "", 0)
	}
}

func (r *REPL) searchAISummary(query string) {
	if !r.ollamaOK {
		r.ollamaOK = r.ollama.CheckHealth()
	}
	if !r.ollamaOK {
		fmt.Println("[nsh] Ollama not running — cannot generate summary.")
		return
	}

	ctx := context.Background()
	fmt.Print("[nsh] thinking...")

	prompt := fmt.Sprintf("Answer this question concisely in 3-5 sentences: %s", query)
	answer, err := r.ollama.GenerateStream(ctx, prompt, "", func(token string) {})
	fmt.Print("\r              \r")

	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] AI error: %v\n", err)
		return
	}

	fmt.Println("\033[36m" + answer + "\033[0m")
	fmt.Println()
	r.saveHistory("ask "+query, "ask", answer, 0, "", 0)
}

func (r *REPL) handleWorkflow(name string) {
	wf, err := r.workflows.Load(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] workflow error: %v\n", err)
		return
	}

	fmt.Printf("[nsh] Running workflow %q (%d steps)\n", wf.Name, len(wf.Steps))
	start := time.Now()

	for i, step := range wf.Steps {
		fmt.Printf("[nsh] step %d/%d: %s\n", i+1, len(wf.Steps), step.Command)

		if strings.HasPrefix(step.Command, "cd ") {
			dir := strings.TrimSpace(strings.TrimPrefix(step.Command, "cd "))
			expanded, err := executor.CdExpand(dir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[nsh] step %d failed: %v\n", i+1, err)
				r.saveHistory(name, "workflow", "", 1, fmt.Sprintf("failed at step %d", i+1), time.Since(start).Milliseconds())
				return
			}
			os.Chdir(expanded)
			continue
		}

		result, err := r.executor.Run(step.Command)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[nsh] step %d execution error: %v\n", i+1, err)
			r.saveHistory(name, "workflow", "", 1, "", time.Since(start).Milliseconds())
			return
		}
		if result.ExitCode != 0 {
			fmt.Fprintf(os.Stderr, "[nsh] step %d failed (exit %d)\n", i+1, result.ExitCode)
			r.saveHistory(name, "workflow", "", result.ExitCode, "", time.Since(start).Milliseconds())
			return
		}
	}

	duration := time.Since(start).Milliseconds()
	fmt.Printf("[nsh] Workflow %q completed (%dms)\n", wf.Name, duration)
	r.saveHistory(name, "workflow", "", 0, fmt.Sprintf("%d steps", len(wf.Steps)), duration)
}

func (r *REPL) handleBuiltin(input string) {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		r.printHelp()
		return
	}
	args := parts[2:]

	switch parts[1] {
	case "help", "--help", "-h":
		r.printHelp()
	case "version", "--version", "-v":
		fmt.Printf("nsh v%s\n", Version)
	case "record":
		r.handleRecord(args)
	case "save-last":
		r.handleSaveLast(args)
	case "workflows":
		r.handleListWorkflows()
	case "edit":
		r.handleEditWorkflow(args)
	case "delete":
		r.handleDeleteWorkflow(args)
	case "history":
		r.handleHistory(args)
	case "replay":
		r.handleReplay(args)
	case "config":
		r.handleConfig()
	case "models":
		r.handleListModels()
	case "model":
		r.handleSetModel(args)
	case "theme":
		r.handleTheme(args)
	default:
		fmt.Printf("[nsh] Unknown command: nsh %s\n", parts[1])
		fmt.Println("Run \"nsh help\" for available commands.")
	}
}

func (r *REPL) handleAmbiguous(input string) {
	if !r.ollamaOK {
		r.ollamaOK = r.ollama.CheckHealth()
	}
	if !r.ollamaOK {
		r.handleCommand(input)
		return
	}

	ctx := context.Background()
	classification, err := r.ollama.ClassifyInput(ctx, input)
	if err != nil {
		r.handleCommand(input)
		return
	}

	if classification == "COMMAND" {
		r.handleCommand(input)
		return
	}

	wfNames := r.workflows.Names()
	if len(wfNames) > 0 {
		match, err := r.ollama.MatchWorkflow(ctx, input, wfNames)
		if err == nil && match != "" {
			fmt.Printf("[nsh] Did you mean workflow %q? [Y/n] ", match)
			var confirm string
			fmt.Scanln(&confirm)
			confirm = strings.ToLower(strings.TrimSpace(confirm))
			if confirm == "" || confirm == "y" || confirm == "yes" {
				r.handleWorkflow(match)
				return
			}
		}
	}
	r.handleNL(input)
}

func (r *REPL) handleRecord(args []string) {
	if len(args) == 0 {
		fmt.Println("[nsh] Usage: nsh record start | nsh record stop \"name\"")
		return
	}
	switch args[0] {
	case "start":
		r.recording = true
		r.recorded = nil
		fmt.Println("[nsh] Recording started. Run your commands, then: nsh record stop \"name\"")
	case "stop":
		if !r.recording {
			fmt.Println("[nsh] Not currently recording.")
			return
		}
		r.recording = false
		if len(args) < 2 {
			fmt.Println("[nsh] Usage: nsh record stop \"name\"")
			r.recorded = nil
			return
		}
		name := strings.Trim(strings.Join(args[1:], " "), "\"'")
		wf := r.workflows.CreateFromCommands(name, "", r.recorded)
		if err := r.workflows.Save(wf); err != nil {
			fmt.Fprintf(os.Stderr, "[nsh] save error: %v\n", err)
			return
		}
		fmt.Printf("[nsh] Workflow %q saved (%d commands)\n", name, len(r.recorded))
		r.recorded = nil
	default:
		fmt.Println("[nsh] Usage: nsh record start | nsh record stop \"name\"")
	}
}

func (r *REPL) handleSaveLast(args []string) {
	if len(args) < 2 {
		fmt.Println("[nsh] Usage: nsh save-last <count> \"name\"")
		return
	}
	var count int
	fmt.Sscanf(args[0], "%d", &count)
	if count <= 0 {
		fmt.Println("[nsh] Count must be a positive number.")
		return
	}
	name := strings.Trim(strings.Join(args[1:], " "), "\"'")
	entries := r.history.LastN(count)
	if len(entries) == 0 {
		fmt.Println("[nsh] No history entries found.")
		return
	}
	commands := make([]string, len(entries))
	for i, e := range entries {
		if e.Generated != "" {
			commands[i] = e.Generated
		} else {
			commands[i] = e.Input
		}
	}
	wf := r.workflows.CreateFromCommands(name, "", commands)
	if err := r.workflows.Save(wf); err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] save error: %v\n", err)
		return
	}
	fmt.Printf("[nsh] Workflow %q saved (%d commands)\n", name, len(commands))
}

func (r *REPL) handleListWorkflows() {
	list, err := r.workflows.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] error: %v\n", err)
		return
	}
	if len(list) == 0 {
		fmt.Println("[nsh] No workflows saved yet.")
		fmt.Println("      Use \"nsh record start\" to create one.")
		return
	}
	fmt.Println("[nsh] Saved workflows:")
	for _, wf := range list {
		desc := wf.Description
		if desc == "" {
			desc = fmt.Sprintf("%d steps", len(wf.Steps))
		}
		fmt.Printf("  %-20s %s\n", wf.Name, desc)
	}
}

func (r *REPL) handleEditWorkflow(args []string) {
	if len(args) == 0 {
		fmt.Println("[nsh] Usage: nsh edit <workflow-name>")
		return
	}
	path := filepath.Join(config.Dir(), "workflows", args[0]+".toml")
	editor := pickEditor()
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] editor error: %v\n", err)
	}
}

func (r *REPL) handleDeleteWorkflow(args []string) {
	if len(args) == 0 {
		fmt.Println("[nsh] Usage: nsh delete <workflow-name>")
		return
	}
	if err := r.workflows.Delete(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] delete error: %v\n", err)
		return
	}
	fmt.Printf("[nsh] Workflow %q deleted.\n", args[0])
}

func (r *REPL) handleHistory(args []string) {
	now := time.Now()
	if len(args) == 0 {
		r.printDayHistory(now, "Today")
		return
	}
	switch args[0] {
	case "yesterday":
		r.printDayHistory(now.AddDate(0, 0, -1), "Yesterday")
	case "week":
		entries, err := r.history.ReadRange(now.AddDate(0, 0, -7), now)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[nsh] error: %v\n", err)
			return
		}
		fmt.Println("━━━ Last 7 days ━━━━━━━━━━━━━━━━━━")
		fmt.Print(r.history.FormatDayView(entries))
	default:
		parsed, err := time.Parse("2006-01-02", args[0])
		if err != nil {
			fmt.Printf("[nsh] Invalid date: %s (use YYYY-MM-DD, 'yesterday', or 'week')\n", args[0])
			return
		}
		r.printDayHistory(parsed, parsed.Format("2006-01-02"))
	}
}

func (r *REPL) printDayHistory(date time.Time, label string) {
	entries, err := r.history.ReadDate(date)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] error: %v\n", err)
		return
	}
	fmt.Printf("━━━ %s (%s) ━━━━━━━━━━━━━━━━━━\n", label, date.Format("2006-01-02"))
	fmt.Print(r.history.FormatDayView(entries))
}

func (r *REPL) handleReplay(args []string) {
	if len(args) == 0 {
		fmt.Println("[nsh] Usage: nsh replay <HH:MM> [--run]")
		return
	}
	targetTime := args[0]
	shouldRun := len(args) > 1 && args[1] == "--run"

	entries, _ := r.history.ReadDate(time.Now())
	for _, e := range entries {
		if e.Timestamp.Format("15:04") == targetTime {
			fmt.Printf("Time:      %s\n", e.Timestamp.Format("15:04:05"))
			fmt.Printf("Input:     %s\n", e.Input)
			if e.Generated != "" {
				fmt.Printf("Generated: %s\n", e.Generated)
			}
			fmt.Printf("Type:      %s\n", e.Type)
			fmt.Printf("Exit:      %d\n", e.ExitCode)
			fmt.Printf("CWD:       %s\n", e.CWD)
			fmt.Printf("Duration:  %dms\n", e.DurationMs)
			if e.OutputPreview != "" {
				fmt.Printf("Output:    %s\n", e.OutputPreview)
			}
			if shouldRun {
				cmd := e.Input
				if e.Generated != "" {
					cmd = e.Generated
				}
				fmt.Printf("\n[nsh] Re-running: %s\n", cmd)
				r.handleCommand(cmd)
			}
			return
		}
	}
	fmt.Printf("[nsh] No entry found at %s today.\n", targetTime)
}

func (r *REPL) handleConfig() {
	path := filepath.Join(config.Dir(), "config.toml")
	fmt.Printf("[nsh] Config: %s\n", path)
	editor := pickEditor()
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func (r *REPL) handleListModels() {
	if !r.ollamaOK {
		r.ollamaOK = r.ollama.CheckHealth()
	}
	if !r.ollamaOK {
		fmt.Println("[nsh] Ollama not running. Start it with: ollama serve")
		return
	}
	models, err := r.ollama.ListModels()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] error listing models: %v\n", err)
		return
	}
	if len(models) == 0 {
		fmt.Println("[nsh] No models installed. Pull one with: ollama pull llama3.2:3b")
		return
	}
	fmt.Println("[nsh] Available Ollama models:")
	for _, m := range models {
		sizeMB := float64(m.Size) / 1024 / 1024
		marker := ""
		if m.Name == r.ollama.GenerationModel() {
			marker = " <- generation"
		}
		if m.Name == r.ollama.ClassifierModel() {
			marker += " <- classifier"
		}
		fmt.Printf("  %-30s %6.0f MB%s\n", m.Name, sizeMB, marker)
	}
	fmt.Println()
	fmt.Printf("  Current generation model:  %s\n", r.ollama.GenerationModel())
	fmt.Printf("  Current classifier model:  %s\n", r.ollama.ClassifierModel())
	fmt.Println()
	fmt.Println("  Switch with: nsh model <name>")
}

func (r *REPL) handleSetModel(args []string) {
	if len(args) == 0 {
		fmt.Println("[nsh] Usage:")
		fmt.Println("  nsh model <name>              Set generation model")
		fmt.Println("  nsh model classifier <name>   Set classifier model")
		fmt.Printf("\n  Current generation:  %s\n", r.ollama.GenerationModel())
		fmt.Printf("  Current classifier:  %s\n", r.ollama.ClassifierModel())
		return
	}
	if args[0] == "classifier" {
		if len(args) < 2 {
			fmt.Println("[nsh] Usage: nsh model classifier <name>")
			return
		}
		r.ollama.SetClassifierModel(args[1])
		r.cfg.Ollama.ClassifierModel = args[1]
		config.Save(r.cfg, filepath.Join(config.Dir(), "config.toml"))
		fmt.Printf("[nsh] Classifier model set to: %s\n", args[1])
		return
	}
	r.ollama.SetGenerationModel(args[0])
	r.cfg.Ollama.GenerationModel = args[0]
	config.Save(r.cfg, filepath.Join(config.Dir(), "config.toml"))
	fmt.Printf("[nsh] Generation model set to: %s\n", args[0])
}

func (r *REPL) handleTheme(args []string) {
	themes := []string{"default", "minimal", "blue", "cyan", "yellow", "magenta"}
	if len(args) == 0 {
		fmt.Println("[nsh] Available themes:")
		for _, t := range themes {
			marker := ""
			if t == r.cfg.UI.Theme {
				marker = " (active)"
			}
			fmt.Printf("  %s%s\n", t, marker)
		}
		fmt.Println("\n  Switch with: nsh theme <name>")
		return
	}
	r.cfg.UI.Theme = args[0]
	config.Save(r.cfg, filepath.Join(config.Dir(), "config.toml"))
	fmt.Printf("[nsh] Theme set to: %s\n", args[0])
}

func (r *REPL) autoDetectModel() {
	if !r.ollamaOK {
		return
	}
	models, err := r.ollama.ListModels()
	if err != nil || len(models) == 0 {
		return
	}

	genModel := r.ollama.GenerationModel()
	for _, m := range models {
		if m.Name == genModel {
			return
		}
	}

	fallback := models[0].Name
	r.ollama.SetGenerationModel(fallback)
	r.ollama.SetClassifierModel(fallback)
	fmt.Printf("[nsh] Model %q not found. Using %q instead.\n", genModel, fallback)
	fmt.Println("      Run \"nsh models\" to see available models, \"nsh model <name>\" to switch.")
}

func (r *REPL) printHelp() {
	fmt.Printf(`nsh v%s — the natural shell
Developed by Sanchit

Usage:
  Type commands normally, or use plain English.
  Built-in commands (ls, cat, grep, etc.) run natively — no shell needed.

Search & AI:
  google <query>              Open search in browser
  ask <query>                 Get AI answer in terminal (via Ollama)
  google! <query>             AI answer + open browser
  wiki/yt/gh <query>          Search Wikipedia, YouTube, GitHub

Built-in commands:
  nsh help                       Show this help
  nsh version                    Show version
  nsh models                     List available Ollama models
  nsh model <name>               Set generation model
  nsh model classifier <name>    Set classifier model
  nsh theme <name>               Set color theme
  nsh record start               Start recording a workflow
  nsh record stop "name"         Stop recording and save as workflow
  nsh save-last N "name"         Save last N commands as a workflow
  nsh workflows                  List saved workflows
  nsh edit <name>                Edit a workflow in your editor
  nsh delete <name>              Delete a workflow
  nsh history                    Show today's command history
  nsh history yesterday          Show yesterday's history
  nsh history week               Show last 7 days
  nsh history YYYY-MM-DD         Show specific date
  nsh replay HH:MM               Show details of a command at that time
  nsh replay HH:MM --run         Re-execute that command
  nsh config                     Open config in your editor
  exit / quit                    Exit nsh
`, Version)
}

func (r *REPL) saveHistory(input, inputType, generated string, exitCode int, output string, durationMs int64) {
	if !r.cfg.History.CaptureOutput {
		output = ""
	}
	cwd, _ := os.Getwd()
	r.history.Append(history.Entry{
		Timestamp:     time.Now(),
		Input:         input,
		Type:          inputType,
		Generated:     generated,
		ExitCode:      exitCode,
		OutputPreview: output,
		CWD:           cwd,
		DurationMs:    durationMs,
	})
}

func pickEditor() string {
	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}
	if _, err := exec.LookPath("code"); err == nil {
		return "code"
	}
	if _, err := exec.LookPath("notepad"); err == nil {
		return "notepad"
	}
	return "vi"
}
