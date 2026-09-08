package main

import (
	"fmt"
	"os"

	"github.com/nsh-terminal/nsh/internal/config"
	"github.com/nsh-terminal/nsh/internal/repl"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help", "-h":
			printHelp()
			return
		case "--version", "-v":
			fmt.Printf("nsh %s\n", repl.Version)
			return
		}
	}

	cfg, err := config.EnsureDefaults()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] config error: %v\n", err)
		os.Exit(1)
	}

	r := repl.New(cfg)
	if err := r.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[nsh] error: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Printf(`nsh — the natural shell (v%s)

Usage: nsh [OPTION]

A cross-platform terminal that understands both shell commands and natural
language. Type commands normally, or describe what you want in plain English.
Uses local Ollama LLMs for natural language translation.

Options:
  -h, --help       display this help and exit
  -v, --version    display version information and exit

Getting started:
  Just run 'nsh' with no arguments to start the interactive shell.
  Type any shell command as you normally would, or use plain English:

    $ ls -la                          (runs natively, instant)
    $ list files sorted by size       (translated via Ollama)
    $ google latest LLM news          (opens browser search)
    $ ask what is kubernetes           (AI answer in terminal)
    $ google! explain docker           (AI answer + open browser)

Interactive commands:
  nsh help                  show full command reference
  nsh models                list / switch Ollama models
  nsh record start          start recording a workflow
  nsh record stop "name"    save recorded workflow
  nsh workflows             list saved workflows
  nsh history               browse command history
  nsh theme <name>          switch color theme (default, blue, cyan, ...)
  nsh config                open config file in editor
  exit                      quit nsh

Configuration:
  Windows:     %%APPDATA%%\nsh\config.toml
  macOS/Linux: ~/.config/nsh/config.toml

Requirements:
  Ollama (optional) — install from https://ollama.ai for NL features.
  Without Ollama, nsh works as a regular shell with smart routing.

Developed by Sanchit
Report bugs at: https://github.com/nsh-terminal/nsh/issues
`, repl.Version)
}
