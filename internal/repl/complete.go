package repl

import (
	"os"
	"strings"
)

var nshSubcommands = []string{
	"help", "version", "models", "model", "theme", "run", "scratch",
	"record", "save-last", "workflows", "edit", "delete",
	"history", "replay", "config", "up", "down", "status",
}

var firstTokenCompletions = []string{
	"ls", "dir", "cd", "cat", "type", "find", "grep", "git", "go",
	"nsh", "pwd", "clear", "cls", "echo", "mkdir", "rm", "cp", "mv",
	"head", "tail", "wc", "which", "where", "exit", "quit",
}

type nshCompleter struct{}

func (nshCompleter) Do(line []rune, pos int) ([][]rune, int) {
	if pos > len(line) {
		pos = len(line)
	}
	prefix := string(line[:pos])
	word := currentWord(prefix)
	wordRunes := []rune(word)

	var names []string
	if isFirstToken(prefix) && !strings.ContainsAny(word, `/\:`) {
		names = append(names, firstTokenCompletions...)
	}
	if afterNshSubcommand(prefix) {
		names = append(names, nshSubcommands...)
	}
	names = append(names, pathCompletions(word)...)

	seen := map[string]bool{}
	var out [][]rune
	lowerWord := strings.ToLower(word)
	for _, name := range names {
		if seen[name] {
			continue
		}
		seen[name] = true
		nr := []rune(name)
		if !strings.HasPrefix(strings.ToLower(string(nr)), lowerWord) {
			continue
		}
		if len(nr) < len(wordRunes) {
			continue
		}
		out = append(out, nr[len(wordRunes):])
	}
	return out, len(wordRunes)
}

func currentWord(prefix string) string {
	i := len(prefix)
	for i > 0 {
		ch := prefix[i-1]
		if ch == ' ' || ch == '\t' {
			break
		}
		i--
	}
	return prefix[i:]
}

func isFirstToken(prefix string) bool {
	return !strings.ContainsAny(prefix, " \t")
}

func afterNshSubcommand(prefix string) bool {
	fields := strings.Fields(prefix)
	if len(fields) == 0 {
		return false
	}
	if !strings.EqualFold(fields[0], "nsh") {
		return false
	}
	if strings.HasSuffix(prefix, " ") && len(fields) == 1 {
		return true
	}
	return len(fields) == 2 && !strings.HasSuffix(prefix, " ")
}

func pathCompletions(word string) []string {
	expanded := word
	if strings.HasPrefix(word, "~") {
		if home, err := os.UserHomeDir(); err == nil {
			expanded = home + word[1:]
		}
	}

	dirPart, base := splitPath(expanded)
	readDir := dirPart
	if readDir == "" {
		readDir = "."
	}
	origDir, _ := splitPath(word)

	entries, err := os.ReadDir(readDir)
	if err != nil {
		return nil
	}
	baseLower := strings.ToLower(base)
	var out []string
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") && (base == "" || !strings.HasPrefix(base, ".")) {
			continue
		}
		if base != "" && !strings.HasPrefix(strings.ToLower(name), baseLower) {
			continue
		}
		full := origDir + name
		if e.IsDir() {
			sep := "/"
			if strings.Contains(word, `\`) || (origDir != "" && strings.Contains(origDir, `\`)) {
				sep = `\`
			}
			if os.PathSeparator == '\\' && !strings.Contains(word, "/") {
				sep = `\`
			}
			full += sep
		}
		out = append(out, full)
	}
	return out
}

func splitPath(p string) (dir, base string) {
	i := strings.LastIndexAny(p, `/\`)
	if i < 0 {
		return "", p
	}
	return p[:i+1], p[i+1:]
}
