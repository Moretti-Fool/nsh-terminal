package classifier

import (
	"strings"
	"unicode"
)

type InputType int

const (
	Empty InputType = iota
	Command
	NaturalLanguage
	Search
	Workflow
	Builtin
	Ambiguous
)

func (t InputType) String() string {
	names := [...]string{"empty", "command", "nl", "search", "workflow", "builtin", "ambiguous"}
	if int(t) < len(names) {
		return names[t]
	}
	return "unknown"
}

type Result struct {
	Type         InputType
	WorkflowName string
	SearchEngine string
	SearchQuery  string
}

type PathLookupFunc func(name string) bool

type Classifier struct {
	pathLookup PathLookupFunc
	workflows  map[string]bool
}

func New(pathLookup PathLookupFunc, workflowNames []string) *Classifier {
	wf := make(map[string]bool, len(workflowNames))
	for _, name := range workflowNames {
		wf[strings.ToLower(name)] = true
	}
	return &Classifier{pathLookup: pathLookup, workflows: wf}
}

func (c *Classifier) UpdateWorkflows(names []string) {
	wf := make(map[string]bool, len(names))
	for _, name := range names {
		wf[strings.ToLower(name)] = true
	}
	c.workflows = wf
}

var searchPrefixes = map[string]bool{
	"google": true, "google!": true, "search": true, "search!": true,
	"wiki": true, "yt": true, "gh": true, "ask": true,
}

var nlSignals = []string{
	"show me", "list all", "find the", "find all",
	"what is", "what are", "what processes", "what files",
	"how to", "how do",
	"how much", "how many", "where is", "where are",
	"delete all", "remove all", "create a", "make a",
	"open the", "close the", "start the", "stop the",
	"unzip", "compress", "extract", "download",
	"install", "update all", "upgrade",
	"tell me", "give me", "can you", "please",
	"i want", "i need",
}

func (c *Classifier) Classify(input string) Result {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return Result{Type: Empty}
	}

	lower := strings.ToLower(trimmed)
	tokens := strings.Fields(trimmed)
	firstToken := strings.ToLower(tokens[0])

	if firstToken == "nsh" {
		return Result{Type: Builtin}
	}

	if searchPrefixes[firstToken] && len(tokens) > 1 {
		return Result{
			Type:         Search,
			SearchEngine: firstToken,
			SearchQuery:  strings.Join(tokens[1:], " "),
		}
	}

	if c.workflows[lower] {
		return Result{Type: Workflow, WorkflowName: lower}
	}

	if isPathLike(trimmed) {
		return Result{Type: Command}
	}

	if hasShellOperators(trimmed) {
		return Result{Type: Command}
	}

	if hasNLSignals(lower) {
		return Result{Type: NaturalLanguage}
	}

	if c.pathLookup != nil && c.pathLookup(firstToken) {
		if firstToken == "find" || firstToken == "grep" || firstToken == "cat" || firstToken == "head" || firstToken == "tail" || firstToken == "wc" {
			return Result{Type: Command}
		}
		if len(tokens) >= 4 && allPlainWords(tokens[1:]) {
			return Result{Type: Ambiguous}
		}
		return Result{Type: Command}
	}

	if looksLikeCommand(tokens) {
		return Result{Type: Command}
	}

	// Unknown first word and no flags: this is English, not a program invocation.
	// Do not grow nlSignals for every phrasing — the LLM translator handles the rest.
	if len(tokens) >= 3 && (c.pathLookup == nil || !c.pathLookup(firstToken)) {
		return Result{Type: NaturalLanguage}
	}

	if hasNaturalLanguageStructure(lower) {
		return Result{Type: NaturalLanguage}
	}

	return Result{Type: Ambiguous}
}

func isPathLike(input string) bool {
	if strings.HasPrefix(input, "./") || strings.HasPrefix(input, "/") || strings.HasPrefix(input, "~/") {
		return true
	}
	if len(input) >= 3 && unicode.IsLetter(rune(input[0])) && input[1] == ':' && (input[2] == '\\' || input[2] == '/') {
		return true
	}
	return false
}

func hasShellOperators(input string) bool {
	for _, op := range []string{"|", ">>", ">", "&&", "||"} {
		if strings.Contains(input, op) {
			return true
		}
	}
	return false
}

func hasNLSignals(lower string) bool {
	for _, signal := range nlSignals {
		if strings.Contains(lower, signal) {
			return true
		}
	}
	return false
}

func looksLikeCommand(tokens []string) bool {
	for _, t := range tokens {
		if strings.HasPrefix(t, "-") {
			return true
		}
	}
	if len(tokens) > 0 && strings.Contains(tokens[0], ".") {
		return true
	}
	return false
}

func allPlainWords(tokens []string) bool {
	for _, t := range tokens {
		if strings.HasPrefix(t, "-") || strings.HasPrefix(t, "/") || strings.HasPrefix(t, "\\") {
			return false
		}
		if strings.ContainsAny(t, "./\\:=") {
			return false
		}
		allDigit := true
		for _, ch := range t {
			if ch < '0' || ch > '9' {
				allDigit = false
				break
			}
		}
		if allDigit && len(t) > 0 {
			return false
		}
	}
	return true
}

func hasNaturalLanguageStructure(lower string) bool {
	words := strings.Fields(lower)
	if len(words) < 3 {
		return false
	}
	questionWords := []string{"what", "how", "where", "which", "who", "why", "when"}
	for _, q := range questionWords {
		if words[0] == q {
			return true
		}
	}
	for _, art := range []string{"the ", "a ", "an ", "all ", "my ", "this ", "that "} {
		if strings.Contains(lower, art) {
			return true
		}
	}
	return false
}
