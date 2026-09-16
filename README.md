# nsh - The Natural Shell

**Version:** v1.2.1-lora  
**Developed by:** Sanchit

A cross-platform terminal that understands both shell commands and natural language. Type commands normally, or describe what you want in plain English — nsh figures out the rest using local LLMs via Ollama. Single Go binary, no dependencies.

## Features

- **Instant Execution** — Common commands (`ls`, `cat`, `grep`, `find`, `cp`, `mv`, `rm`, `mkdir`, `touch`, `head`, `tail`, `wc`, `echo`, `which`) run as Go-native builtins with zero shell overhead. Other binaries exec directly; shell is only used for pipes, redirects, and globs.
- **Smart Input Detection** — Automatically determines if your input is a command, natural language, a search query, or a workflow name. No prefixes needed; 90%+ of inputs are classified without any LLM call.
- **Natural Language to Commands** — Describe what you want ("show files sorted by size", "what's using port 8080") and nsh translates it into the right shell command using a local Ollama model.
- **Python Scratch Workspace** — Say "open sales.csv with pandas" or `nsh run <description>` and nsh generates a Python script, auto-creates a venv, installs dependencies, and runs it. Scripts saved for reuse.
- **Parallel Service Launcher** — Define your dev stack in a workflow file and launch everything with `nsh up`. Color-coded log streaming, one-command shutdown with `nsh down`.
- **Workflow Recording & Replay** — Record a sequence of commands, save it with a name, and replay it anytime. Ctrl+C interrupts the current command without cancelling the recording.
- **AI Search & Answers** — `ask what is kubernetes` gets an AI answer in the terminal. `google! <query>` gives both an AI answer and opens the browser.
- **Cross-Platform** — Works on Windows, macOS, and Linux. On Windows, typed commands still use fast builtins/`cmd.exe`; natural language runs through a warm PowerShell host so cmdlets work without a 2–3s cold start per command.
- **Structured History & Auto-Categorization Agent** — Background Auto-Categorization engine dynamically organizes command history by communicating with the local LLM async. Browse by day, search across domains, and replay past commands.
- **RAG OS/Shell Filtering** — Implements a lightweight, local in-memory vector database for chunk embeddings and semantic search. When retrieving context for AI features, nsh intelligently filters this vector store based on the active OS and shell environment, ensuring highly relevant responses.
- **Robust JSON Extraction** — Employs a resilient plan parsing logic that grabs the last valid JSON object in a response (often the final generation) and safely falls back to the first.
- **Semantic Caching & Verification (LLM Router)** — Uses a lightning-fast sub-billion parameter model (set via `nsh model judge <name>`, e.g., `qwen2.5:0.5b`) to judge intents. It acts as an instant Semantic Cache to bypass generation, verifies commands before execution, and automatically triggers fallbacks for "empty" shell outputs.
- **Double Ladder Fallback** — Set a fallback model via `nsh model fallback <name>` that nsh automatically switches to if the primary generation model fails or produces a completely blank output. The execution layer also utilizes a double ladder approach on Windows (warm host -> restart -> run once) for robustness.
- **Automated Fine-Tuning Pipeline** — Import `.jsonl` datasets with `nsh learn-import` to fine-tune the system's local memory and improve performance on custom tasks.
- **Destructive Command Safety** — Detects dangerous commands (`rm -rf`, `DROP TABLE`, etc.) and asks for confirmation.
- **Graceful Degradation** — Works as a normal shell even when Ollama isn't running. NL features simply become unavailable.
- **Customizable** — Choose your Ollama model, color theme, prompt style, scratch directory, and more via `config.toml`.

## Installation

### From Source (recommended)

Requires [Go 1.21+](https://go.dev/dl/).

```bash
go install github.com/Moretti-Fool/nsh-terminal@latest
```

### Build from Repository

```bash
git clone https://github.com/Moretti-Fool/nsh-terminal.git
cd nsh
go build -ldflags="-s -w" -o nsh .
```

The binary will be ~7MB. Move it to a directory in your PATH:

- **Windows:** Copy `nsh.exe` to a folder in your `%PATH%`
- **macOS/Linux:** `sudo mv nsh /usr/local/bin/`

### Binary Releases

Download pre-built binaries from the [Releases](https://github.com/Moretti-Fool/nsh-terminal/releases) page.

## Requirements

- **Required:** Go 1.21+ (to build)
- **Optional:** [Ollama](https://ollama.ai) (for natural language features and Python script generation)
- **Optional:** Python 3 (for scratch workspace)

### Setting up Ollama

```bash
# Install Ollama from https://ollama.ai
# Pull a model (smaller = faster, larger = smarter)
ollama pull llama3.2:3b    # Good balance of speed and quality
ollama pull phi3            # Lightweight, fast classifier
```

## Quick Start

```bash
# Launch nsh
nsh

# Regular commands work as expected — and they're fast (Go-native)
ls -la
find . -name "*.go"
grep -n "TODO" src/*.js
cat README.md | head -20

# Smart flag correction — forgot the dash?
ls ltr                       # treated as ls -ltr

# Natural language — just describe what you want
list all files larger than 10MB
what processes are using port 8080
show disk usage sorted by size

# AI answers in the terminal
ask what is kubernetes
ask explain TCP vs UDP

# Web search
google latest golang news
wiki quantum computing
yt docker tutorial
google! explain microservices  # AI answer + browser

# Python scratch workspace — mention a library, nsh handles the rest
open sales.csv with pandas
plot monthly revenue using matplotlib
nsh run merge all xlsx files in this folder

# Workflow recording
nsh record start
npm install
npm run build
npm test
nsh record stop "build-project"

# Run it later by name
build-project
```

## Parallel Service Launcher

Define your morning dev stack in a workflow file:

```bash
nsh edit morning-stack
```

```toml
name = "morning-stack"
mode = "parallel"

[[services]]
name = "backend"
dir = "~/tpro_backend"
command = "npm run dev"

[[services]]
name = "investment-api"
dir = "~/repo/investment"
setup = "venv/Scripts/activate.ps1"
command = "uvicorn main:app --reload"

[[services]]
name = "frontend"
dir = "~/tpro_unified"
command = "npm run dev"
```

```bash
# Launch all services
nsh up morning-stack

# Logs stream with color-coded prefixes:
# backend        | Server running on port 3000
# investment-api | Uvicorn running on http://127.0.0.1:8000
# frontend       | Compiled successfully

# Check what's running
nsh status

# Stop everything
nsh down
```

## Built-in Commands

### Shell Builtins (Go-native, instant)

| Command | Flags |
|---------|-------|
| `ls` / `dir` | `-l`, `-a`, `-t`, `-r`, `-h` (smart flag correction: `ls ltr` = `ls -ltr`) |
| `cat` / `type` | |
| `find` | `-name`, `-iname`, `-type` (keyword mode: `find nsh spec file`) |
| `grep` | `-i`, `-n` |
| `head` / `tail` | `-n` |
| `wc`, `touch`, `mkdir` | `-p` for mkdir |
| `cp` / `copy` | `-r` for recursive |
| `mv` / `move`, `rm` / `del` | `-r`, `-f` |
| `echo`, `pwd`, `clear` / `cls` | |
| `which` / `where` | cached PATH lookup |

### nsh Commands

| Command | Description |
|---------|-------------|
| `nsh help` | Show help |
| `nsh version` | Show version |
| `nsh models` | List available Ollama models |
| `nsh model <name>` | Set generation model |
| `nsh model classifier <name>` | Set classifier model |
| `nsh model fallback <name>` | Set fallback generation model |
| `nsh model judge <name>` | Set semantic judge model |
| `nsh learn-import <file>` | Import `.jsonl` dataset to fine-tune local memory |
| `nsh theme <name>` | Set color theme |
| `nsh run <description>` | Generate and run a Python script |
| `nsh scratch` | Show scratch directory location |
| `nsh record start` | Start recording commands |
| `nsh record stop "name"` | Save recorded commands as workflow |
| `nsh record cancel` | Cancel current recording |
| `nsh save-last N "name"` | Save last N commands as workflow |
| `nsh workflows` | List all saved workflows |
| `nsh edit <name>` | Edit a workflow file |
| `nsh delete <name>` | Delete a workflow |
| `nsh up <name>` | Launch parallel services from a workflow |
| `nsh down` | Stop all running services |
| `nsh status` | Show running services |
| `nsh history` | Show today's auto-categorized history |
| `nsh categories` | List auto-categorized command domains |
| `nsh categories <domain>` | Search commands within a domain |
| `nsh history yesterday` | Yesterday's history |
| `nsh history week` | Last 7 days |
| `nsh replay HH:MM [--run]` | Inspect or re-execute a past command |
| `nsh config` | Open config in your editor |
| `exit` / `quit` | Exit nsh |

## Configuration

Config file location:
- **Windows:** `%APPDATA%\nsh\config.toml`
- **macOS/Linux:** `~/.config/nsh/config.toml`

### Example config.toml

```toml
[ollama]
url = "http://localhost:11434"
classifier_model = "nsh-local"
generation_model = "nsh-local"
fallback_model = "qwen2.5-coder:3b"
judge_model = "qwen2.5:0.5b"
timeout_ms = 45000
max_retries = 3

[shell]
default = "auto"
wsl_distro = ""  # empty = use default WSL distro

[search]
default_engine = "google"
open_browser = true

[ui]
prompt = "nsh> "
show_generated_command = true
confirm_destructive = true
theme = "default"       # default, minimal, blue, cyan, yellow, magenta
show_welcome = true

[history]
retention_days = 90
capture_output = true
output_preview_chars = 500

[scratch]
dir = ""      # default: %APPDATA%\nsh\scratch (Windows) or ~/.config/nsh/scratch
python = ""   # default: auto-detect python3/python
```

### Themes

| Theme | Description |
|-------|-------------|
| `default` | Green prompt with ANSI colors |
| `minimal` | No colors, plain text |
| `blue` | Blue prompt |
| `cyan` | Cyan prompt |
| `yellow` | Yellow prompt |
| `magenta` | Magenta prompt |

Switch themes: `nsh theme cyan`

## How It Works

### 3-Tier Execution Pipeline

```
Input → Go-native builtin? → Run in-process (0ms overhead)
      → PowerShell cmdlet? → warm PowerShell host (~50–150ms after one cold start)
      → Binary in PATH?    → exec directly (no shell, ~50ms)
      → Needs pipes/globs? → cmd.exe /C fallback (~50ms)
```

Natural language is translated to **PowerShell** on Windows. A persistent `powershell.exe -NoProfile` process is started when nsh opens, so cmdlet-heavy NL commands do not pay a 2–3s startup on every request. Typed Unix-style commands (`ls`, `git`, …) still use builtins / direct exec.

### Input Classification

A fast heuristic classifier handles 90%+ of inputs without any LLM call:

1. **Builtins** — `nsh` prefix → builtin handler
2. **Search** — `google`, `ask`, `wiki`, `yt`, `gh` prefix → search/AI
3. **Workflows** — matches a saved workflow name → replay
4. **Path-like** — starts with `./`, `/`, `~/`, `C:\` → command
5. **Shell operators** — contains `|`, `>`, `&&` → command
6. **NL signals** — "show me", "how to", "what is" → natural language
7. **PATH lookup** — first token found in PATH → command (with plain-words guard for ambiguous cases like `find nsh design file`)
8. **NL structure** — question words, articles → natural language
9. **Ambiguous** → falls through to LLM classifier

### Python Scratch Workspace

When you mention a Python library (pandas, matplotlib, numpy, etc.) in natural language, nsh:

1. Sends your request to Ollama with a Python-specific prompt
2. Detects imports in the generated script
3. Maps imports to pip packages (50+ mappings: `sklearn`→`scikit-learn`, `PIL`→`Pillow`, `bs4`→`beautifulsoup4`, etc.)
4. Creates/reuses a shared venv, installs only missing packages
5. Runs the script and streams output
6. Saves the script with a timestamp to the scratch directory

## Architecture

```
nsh (single Go binary, ~7MB)
  main.go              - Entry point, --help/--version
  internal/
    classifier/        - Smart input routing (heuristic + LLM fallback)
    config/            - TOML config management
    executor/          - 3-tier execution pipeline + Go-native builtins
      builtins.go      - ls, cat, grep, find, cp, mv, rm, head, tail, etc.
      pshost.go        - Persistent PowerShell host for NL/cmdlets (Windows)
      services.go      - Parallel service launcher (nsh up/down/status)
    history/           - JSONL per-day history with search and background auto-categorization via LLM
    ollama/            - Ollama REST API client (command gen, Python gen, classify)
    repl/              - Main REPL loop
    scratch/           - Python scratch workspace (venv, deps, execution)
    search/            - Web search with multiple engines
    workflow/          - TOML workflow recording and replay
```

## License

MIT

## Contributing

Contributions welcome! Please open an issue first to discuss what you'd like to change.
