# nsh - The Natural Shell

**Version:** v1.0.0  
**Developed by:** Sanchit

A cross-platform terminal that understands both shell commands and natural language. Type commands normally, or describe what you want in plain English — nsh figures out the rest using local LLMs via Ollama.

## Features

- **Smart Input Detection** — Automatically determines if your input is a command, natural language, a search query, or a workflow name. No prefixes needed; 90%+ of inputs are classified without any LLM call.
- **Natural Language to Commands** — Describe what you want ("unzip file.zip", "list files sorted by size") and nsh translates it into the right shell command using a local Ollama model.
- **Workflow Recording & Replay** — Record a sequence of commands, save it with a name, and replay it anytime. Great for repetitive tasks like starting a dev server or deploying.
- **Cross-Platform** — Works on Windows (PowerShell + WSL), macOS, and Linux. Unix commands auto-route through WSL on Windows.
- **Web Search Integration** — Prefix your query with `google`, `wiki`, `yt`, or `gh` to search directly from the terminal.
- **Structured History** — Browse command history by day, search across days, replay past commands.
- **Destructive Command Safety** — Detects dangerous commands (rm -rf, DROP TABLE, etc.) and asks for confirmation before running.
- **Graceful Degradation** — Works as a normal shell even when Ollama isn't running. NL features simply become unavailable.
- **Customizable** — Choose your Ollama model, color theme, prompt style, and more via config.toml.

## Installation

### From Source (recommended)

Requires [Go 1.21+](https://go.dev/dl/).

```bash
go install github.com/nsh-terminal/nsh@latest
```

### Build from Repository

```bash
git clone https://github.com/nsh-terminal/nsh.git
cd nsh
go build -ldflags="-s -w" -o nsh .
```

The binary will be ~7MB. Move it to a directory in your PATH:

- **Windows:** Copy `nsh.exe` to a folder in your `%PATH%`
- **macOS/Linux:** `sudo mv nsh /usr/local/bin/`

### Binary Releases

Download pre-built binaries from the [Releases](https://github.com/nsh-terminal/nsh/releases) page.

## Requirements

- **Required:** Go 1.21+ (to build)
- **Optional:** [Ollama](https://ollama.ai) (for natural language features)
- **Optional:** WSL (for Unix command routing on Windows)

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

# Regular commands work as expected
ls -la
git status
docker ps

# Natural language - just describe what you want
list all files larger than 10MB
what processes are using port 8080
show disk usage sorted by size
unzip archive.tar.gz to output folder

# Web search
google latest news about LLMs
wiki quantum computing
yt golang tutorial
gh terminal emulator projects

# Workflows
nsh record start
npm install
npm run build
npm test
nsh record stop "build-project"

# Run it later by name
build-project
```

## Built-in Commands

| Command | Description |
|---------|-------------|
| `nsh help` | Show help |
| `nsh version` | Show version |
| `nsh models` | List available Ollama models |
| `nsh model <name>` | Set generation model |
| `nsh model classifier <name>` | Set classifier model |
| `nsh theme <name>` | Set color theme |
| `nsh record start` | Start recording commands |
| `nsh record stop "name"` | Save recorded commands as workflow |
| `nsh save-last N "name"` | Save last N commands as workflow |
| `nsh workflows` | List all saved workflows |
| `nsh edit <name>` | Edit a workflow file |
| `nsh delete <name>` | Delete a workflow |
| `nsh history` | Show today's history |
| `nsh history yesterday` | Show yesterday's history |
| `nsh history week` | Show last 7 days |
| `nsh history YYYY-MM-DD` | Show specific date |
| `nsh replay HH:MM` | Show details of a past command |
| `nsh replay HH:MM --run` | Re-execute a past command |
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
classifier_model = "phi3"
generation_model = "llama3.2:3b"
timeout_ms = 10000

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

1. **Input Classification** — A fast heuristic classifier checks if the input is a command (found in PATH, has flags/operators), natural language (contains NL signals like "show me", "how to"), a search query, or a workflow name. Only ambiguous inputs fall through to the LLM classifier.

2. **Command Execution** — Shell commands run directly through the detected shell. On Windows, Unix-style commands (grep, awk, ls -ltr) automatically route through WSL.

3. **NL Translation** — Natural language inputs are sent to your local Ollama model with system context (OS, shell, CWD). The model returns shell commands which are displayed and executed.

4. **Workflow Learning** — Record command sequences and save them as named workflows. Workflows are stored as TOML files and can be edited manually.

## Architecture

```
nsh (single Go binary)
  internal/
    classifier/   - Smart input routing (heuristic + LLM fallback)
    config/       - TOML config management
    executor/     - Cross-platform command execution + WSL routing
    history/      - JSONL per-day history with search
    ollama/       - Ollama REST API client
    repl/         - Main REPL loop wiring everything together
    search/       - Web search with multiple engines
    workflow/     - Named command sequence recording and replay
```

## License

MIT

## Contributing

Contributions welcome! Please open an issue first to discuss what you'd like to change.
