package main

import (
	"fmt"
	"os"

	"github.com/nsh-terminal/nsh/internal/config"
	"github.com/nsh-terminal/nsh/internal/repl"
)

func main() {
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
