package classifier

import (
	"strings"
	"testing"
)

// 120+ tests
func TestComprehensiveClassifier(t *testing.T) {
	t.Parallel()

	testPathLookup := func(name string) bool {
		known := map[string]bool{
			"git": true, "docker": true, "ls": true, "cd": true,
			"npm": true, "go": true, "python": true, "cat": true,
			"mkdir": true, "rm": true, "curl": true, "ping": true,
			"find": true, "sort": true, "kill": true, "head": true,
			"tail": true, "wc": true, "grep": true,
		}
		return known[name]
	}

	c := New(testPathLookup, []string{"run-backend", "deploy-staging"})

	t.Run("InputType_String", func(t *testing.T) {
		tests := []struct {
			it   InputType
			want string
		}{
			{Empty, "empty"},
			{Command, "command"},
			{NaturalLanguage, "nl"},
			{Search, "search"},
			{Workflow, "workflow"},
			{Builtin, "builtin"},
			{Ambiguous, "ambiguous"},
			{InputType(999), "unknown"},
		}
		for _, tt := range tests {
			if got := tt.it.String(); got != tt.want {
				t.Errorf("InputType(%d).String() = %v, want %v", tt.it, got, tt.want)
			}
		}
	})

	t.Run("UpdateWorkflows", func(t *testing.T) {
		c2 := New(nil, []string{"old"})
		if r := c2.Classify("old"); r.Type != Workflow {
			t.Errorf("expected Workflow for old")
		}
		c2.UpdateWorkflows([]string{"new-wf", "ANOTHER"})
		if r := c2.Classify("old"); r.Type == Workflow {
			t.Errorf("did not expect Workflow for old")
		}
		if r := c2.Classify("new-wf"); r.Type != Workflow {
			t.Errorf("expected Workflow for new-wf")
		}
		if r := c2.Classify("another"); r.Type != Workflow {
			t.Errorf("expected Workflow for another (case insensitive)")
		}
	})

	t.Run("Classify_Empty_And_Whitespace", func(t *testing.T) {
		tests := []string{"", " ", "\t", "\n", "\r\n", "   \t  "}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Empty {
				t.Errorf("Classify(%q).Type = %v, want Empty", tt, got.Type)
			}
		}
	})

	t.Run("Classify_Builtin", func(t *testing.T) {
		tests := []string{
			"nsh",
			"NSH",
			"nsh help",
			"nsh record start",
			"nsh workflows",
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Builtin {
				t.Errorf("Classify(%q).Type = %v, want Builtin", tt, got.Type)
			}
		}
	})

	t.Run("Classify_Search", func(t *testing.T) {
		tests := []struct {
			input  string
			engine string
			query  string
		}{
			{"google how to code", "google", "how to code"},
			{"google! golang tricks", "google!", "golang tricks"},
			{"search local files", "search", "local files"},
			{"search! my computer", "search!", "my computer"},
			{"wiki Alan Turing", "wiki", "Alan Turing"},
			{"yt never gonna give you up", "yt", "never gonna give you up"},
			{"gh nsh terminal", "gh", "nsh terminal"},
			{"ask chatgpt about go", "ask", "chatgpt about go"},
			{"GOOGLE caps", "google", "caps"}, // First token gets lowercased only when matching search prefixes? Wait, search prefixes match exact `firstToken`, which is `strings.ToLower(tokens[0])`.
		}
		for _, tt := range tests {
			got := c.Classify(tt.input)
			if got.Type != Search {
				t.Errorf("Classify(%q).Type = %v, want Search", tt.input, got.Type)
			}
			if got.SearchEngine != tt.engine {
				t.Errorf("Classify(%q).SearchEngine = %v, want %v", tt.input, got.SearchEngine, tt.engine)
			}
			if got.SearchQuery != tt.query {
				t.Errorf("Classify(%q).SearchQuery = %v, want %v", tt.input, got.SearchQuery, tt.query)
			}
		}
	})

	t.Run("Classify_Search_SingleWord_NotSearch", func(t *testing.T) {
		// Single word "google" shouldn't trigger search mode (len(tokens) > 1 required)
		tests := []string{"google", "search", "wiki", "yt", "gh", "ask"}
		for _, tt := range tests {
			got := c.Classify(tt)
			if got.Type == Search {
				t.Errorf("Classify(%q).Type = Search, want something else", tt)
			}
		}
	})

	t.Run("Classify_Workflow", func(t *testing.T) {
		tests := []string{"run-backend", "deploy-staging", "RUN-BACKEND"}
		for _, tt := range tests {
			got := c.Classify(tt)
			if got.Type != Workflow {
				t.Errorf("Classify(%q).Type = %v, want Workflow", tt, got.Type)
			}
			if got.WorkflowName != strings.ToLower(tt) {
				t.Errorf("Classify(%q).WorkflowName = %v, want %v", tt, got.WorkflowName, strings.ToLower(tt))
			}
		}
	})

	t.Run("Classify_PathLike_Command", func(t *testing.T) {
		tests := []string{
			"./script.sh",
			"./script",
			"/usr/bin/env",
			"~/bin/program",
			"C:\\Users\\test\\file.exe",
			"D:/Games/Doom/doom.exe",
			"z:\\script.bat",
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Command {
				t.Errorf("Classify(%q).Type = %v, want Command", tt, got.Type)
			}
		}
	})

	t.Run("Classify_ShellOperators_Command", func(t *testing.T) {
		tests := []string{
			"echo hello | cat",
			"ls >> out.txt",
			"cat file > out2",
			"build && run",
			"fail || fallback",
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Command {
				t.Errorf("Classify(%q).Type = %v, want Command", tt, got.Type)
			}
		}
	})

	t.Run("Classify_NL_Signals", func(t *testing.T) {
		signals := []string{
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
		for _, sig := range signals {
			input := sig + " something"
			if got := c.Classify(input); got.Type != NaturalLanguage {
				t.Errorf("Classify(%q).Type = %v, want NL", input, got.Type)
			}
		}
	})

	t.Run("Classify_KnownCommand_Special", func(t *testing.T) {
		tests := []string{
			"find . -name *.go",
			"grep test file",
			"cat file",
			"head file",
			"tail file",
			"wc -l file",
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Command {
				t.Errorf("Classify(%q).Type = %v, want Command", tt, got.Type)
			}
		}
	})

	t.Run("Classify_KnownCommand_Ambiguous", func(t *testing.T) {
		// e.g. "git pull origin master" -> tokens[1:] are "pull", "origin", "master". all plain words -> Ambiguous
		tests := []string{
			"git pull origin master",
			"docker run image latest",
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Ambiguous {
				t.Errorf("Classify(%q).Type = %v, want Ambiguous", tt, got.Type)
			}
		}
	})

	t.Run("Classify_KnownCommand_Command", func(t *testing.T) {
		// tokens < 4 or has special symbols => Command
		tests := []string{
			"git status",
			"docker ps",
			"ls -l /tmp",
			"npm run dev",
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Command {
				t.Errorf("Classify(%q).Type = %v, want Command", tt, got.Type)
			}
		}
	})

	t.Run("Classify_LooksLikeCommand", func(t *testing.T) {
		tests := []string{
			"unknown -f flag",
			"unknown --flag",
			"unknown.py args",
			"script.exe start",
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != Command {
				t.Errorf("Classify(%q).Type = %v, want Command", tt, got.Type)
			}
		}
	})

	t.Run("Classify_UnknownFirstWord_English", func(t *testing.T) {
		tests := []string{
			"unknown word phrase",
			"random string here",
			"is this english", // length >= 3, first word not in path
		}
		for _, tt := range tests {
			if got := c.Classify(tt); got.Type != NaturalLanguage {
				t.Errorf("Classify(%q).Type = %v, want NL", tt, got.Type)
			}
		}
	})

	t.Run("Classify_NLStructure", func(t *testing.T) {
		// if length < 3, but matches structure? Structure requires len >= 3.
		// Testing structure specifically on non-path, len >= 3.
		// Wait, if pathLookup is true, it goes to Command/Ambiguous.
		// To hit hasNaturalLanguageStructure, it must bypass:
		// pathLookup == true, len >= 4 all plain words -> Ambiguous.
		// len < 3 -> goes to structure. But structure requires len >= 3.
		// Wait, let's see how to hit it.
		// "if len(tokens) >= 3 && (c.pathLookup == nil || !c.pathLookup(firstToken)) { NL }"
		// So if pathLookup is nil, len >= 3 goes to NL directly.
		// If len < 3, it skips that block.
		// Then it calls hasNaturalLanguageStructure, but that also requires len >= 3!
		// Actually, hasNaturalLanguageStructure will never match if len < 3.
		// Wait, what if pathLookup == true, len >= 3, and NOT allPlainWords? -> goes to Command!
		// Let's just test hasNaturalLanguageStructure directly.
		tests := []struct {
			input string
			want  bool
		}{
			{"what to do", true},
			{"how is it", true},
			{"where is it", true},
			{"which one is", true},
			{"who is this", true},
			{"why is it", true},
			{"when to go", true},
			{"do the thing", true},
			{"a test string", true},
			{"an apple here", true},
			{"all my files", true},
			{"my name is", true},
			{"this is a test", true},
			{"that was cool", true},
			{"short", false},
			{"two words", false},
			{"no special structure here", false},
		}
		for _, tt := range tests {
			if got := hasNaturalLanguageStructure(tt.input); got != tt.want {
				t.Errorf("hasNaturalLanguageStructure(%q) = %v, want %v", tt.input, got, tt.want)
			}
		}
	})

	t.Run("allPlainWords", func(t *testing.T) {
		tests := []struct {
			tokens []string
			want   bool
		}{
			{[]string{"hello", "world"}, true},
			{[]string{"-flag", "test"}, false},
			{[]string{"/path"}, false},
			{[]string{"\\windows"}, false},
			{[]string{"hello.txt"}, false},
			{[]string{"a/b"}, false},
			{[]string{"a\\b"}, false},
			{[]string{"key=val"}, false},
			{[]string{"123"}, false}, // all digits
			{[]string{"word123"}, true},
		}
		for _, tt := range tests {
			if got := allPlainWords(tt.tokens); got != tt.want {
				t.Errorf("allPlainWords(%v) = %v, want %v", tt.tokens, got, tt.want)
			}
		}
	})

	t.Run("EdgeCases", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"Unicode", "你好 世界"},
			{"Very Long", strings.Repeat("a ", 100)},
			{"Special Chars", "!@#$%^&*()"},
			{"Empty Quotes", `"" ''`},
		}
		for _, tt := range tests {
			// Just ensure it doesn't panic
			c.Classify(tt.input)
		}
	})
	
	t.Run("PathLookupNil", func(t *testing.T) {
		cNil := New(nil, nil)
		// Should not panic when pathLookup is nil
		res := cNil.Classify("some command here")
		if res.Type != NaturalLanguage {
			t.Errorf("Expected NL for unknown word when pathLookup is nil")
		}
	})
}
