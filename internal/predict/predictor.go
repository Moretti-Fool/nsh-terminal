package predict

import (
	"sort"
	"strings"
	"sync"
)

// Predictor tracks command patterns within a session and across history
// to suggest likely next commands.
type Predictor struct {
	mu         sync.RWMutex
	session    []sessionEntry            // commands in current session
	bigrams    map[string]map[string]int // cmd_A -> cmd_B -> count
	cwdCmds    map[string]map[string]int // cwd -> cmd -> count
	frequency  map[string]int            // global command frequency
	maxSession int
}

type sessionEntry struct {
	Command  string
	CWD      string
	ExitCode int
}

// Suggestion represents a predicted command with a score and reason.
type Suggestion struct {
	Command string
	Score   float64
	Reason  string // e.g. "frequently follows 'git add'", "common in this directory"
}

// New creates a Predictor with default settings.
func New() *Predictor {
	return &Predictor{
		session:    make([]sessionEntry, 0, 100),
		bigrams:    make(map[string]map[string]int),
		cwdCmds:    make(map[string]map[string]int),
		frequency:  make(map[string]int),
		maxSession: 100,
	}
}

// normalize extracts the base command.
func normalize(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return ""
	}
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return ""
	}

	subCmds := map[string]bool{
		"git":     true,
		"docker":  true,
		"npm":     true,
		"kubectl": true,
		"go":      true,
		"apt":     true,
		"apt-get": true,
	}

	if subCmds[parts[0]] && len(parts) > 1 {
		return parts[0] + " " + parts[1]
	}
	return parts[0]
}

// Record adds a command execution to the predictor's memory.
func (p *Predictor) Record(command, cwd string, exitCode int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	normCmd := normalize(command)
	if normCmd == "" {
		return
	}

	// Update session ring buffer
	if len(p.session) >= p.maxSession {
		p.session = p.session[1:]
	}

	var lastCmd string
	if len(p.session) > 0 {
		lastCmd = normalize(p.session[len(p.session)-1].Command)
	}

	p.session = append(p.session, sessionEntry{
		Command:  command,
		CWD:      cwd,
		ExitCode: exitCode,
	})

	// Update bigrams
	if lastCmd != "" {
		if p.bigrams[lastCmd] == nil {
			p.bigrams[lastCmd] = make(map[string]int)
		}
		p.bigrams[lastCmd][normCmd]++
	}

	// Update cwdCmds
	if p.cwdCmds[cwd] == nil {
		p.cwdCmds[cwd] = make(map[string]int)
	}
	p.cwdCmds[cwd][normCmd]++

	// Update frequency
	p.frequency[normCmd]++
}

// Suggest returns up to `limit` predicted next commands.
func (p *Predictor) Suggest(cwd string, limit int) []Suggestion {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var lastCmd string
	if len(p.session) > 0 {
		lastCmd = normalize(p.session[len(p.session)-1].Command)
	}

	candidates := make(map[string]bool)

	totalBigramsFromLast := 0
	if p.bigrams[lastCmd] != nil {
		for c, count := range p.bigrams[lastCmd] {
			candidates[c] = true
			totalBigramsFromLast += count
		}
	}

	totalCmdsInCWD := 0
	if p.cwdCmds[cwd] != nil {
		for c, count := range p.cwdCmds[cwd] {
			candidates[c] = true
			totalCmdsInCWD += count
		}
	}

	maxFreq := 0
	for c, count := range p.frequency {
		candidates[c] = true
		if count > maxFreq {
			maxFreq = count
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	var suggestions []Suggestion
	for candidate := range candidates {
		bigramScore := 0.0
		if totalBigramsFromLast > 0 && p.bigrams[lastCmd] != nil {
			bigramScore = float64(p.bigrams[lastCmd][candidate]) / float64(totalBigramsFromLast)
		}

		cwdScore := 0.0
		if totalCmdsInCWD > 0 && p.cwdCmds[cwd] != nil {
			cwdScore = float64(p.cwdCmds[cwd][candidate]) / float64(totalCmdsInCWD)
		}

		freqScore := 0.0
		if maxFreq > 0 {
			freqScore = float64(p.frequency[candidate]) / float64(maxFreq)
		}

		finalScore := 0.5*bigramScore + 0.3*cwdScore + 0.2*freqScore

		var reason string
		if bigramScore >= cwdScore && bigramScore >= freqScore && bigramScore > 0 {
			reason = "frequently follows '" + lastCmd + "'"
		} else if cwdScore >= freqScore && cwdScore > 0 {
			reason = "common in this directory"
		} else {
			reason = "frequently used"
		}

		suggestions = append(suggestions, Suggestion{
			Command: candidate,
			Score:   finalScore,
			Reason:  reason,
		})
	}

	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].Score > suggestions[j].Score
	})

	if limit > 0 && len(suggestions) > limit {
		suggestions = suggestions[:limit]
	}

	return suggestions
}

// SuggestAfter returns predictions specifically based on what
// usually follows the given command.
func (p *Predictor) SuggestAfter(lastCmd string, limit int) []Suggestion {
	p.mu.RLock()
	defer p.mu.RUnlock()

	normLastCmd := normalize(lastCmd)
	if normLastCmd == "" || p.bigrams[normLastCmd] == nil {
		return nil
	}

	total := 0
	for _, count := range p.bigrams[normLastCmd] {
		total += count
	}

	if total == 0 {
		return nil
	}

	var suggestions []Suggestion
	for candidate, count := range p.bigrams[normLastCmd] {
		score := float64(count) / float64(total)
		suggestions = append(suggestions, Suggestion{
			Command: candidate,
			Score:   score,
			Reason:  "frequently follows '" + normLastCmd + "'",
		})
	}

	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].Score > suggestions[j].Score
	})

	if limit > 0 && len(suggestions) > limit {
		suggestions = suggestions[:limit]
	}

	return suggestions
}
