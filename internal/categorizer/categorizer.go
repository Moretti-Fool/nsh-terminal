package categorizer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/nsh-terminal/nsh/internal/config"
	"github.com/nsh-terminal/nsh/internal/ollama"
)

type Categorizer struct {
	client     *ollama.Client
	configDir  string
	categories []string
	history    map[string]string // command -> category
	mu         sync.RWMutex
	workerChan chan cmdTask
}

type cmdTask struct {
	Command   string
	Generated string
	Output    string
}

func New(client *ollama.Client) *Categorizer {
	c := &Categorizer{
		client:     client,
		configDir:  config.Dir(),
		history:    make(map[string]string),
		workerChan: make(chan cmdTask, 100),
	}
	c.load()
	go c.worker()
	return c
}

func (c *Categorizer) load() {
	path := filepath.Join(c.configDir, "categories.json")
	data, err := os.ReadFile(path)
	if err == nil {
		var state struct {
			Categories []string          `json:"categories"`
			History    map[string]string `json:"history"`
		}
		if json.Unmarshal(data, &state) == nil {
			c.categories = state.Categories
			c.history = state.History
			if c.history == nil {
				c.history = make(map[string]string)
			}
		}
	}
}

func (c *Categorizer) save() {
	path := filepath.Join(c.configDir, "categories.json")
	state := struct {
		Categories []string          `json:"categories"`
		History    map[string]string `json:"history"`
	}{
		Categories: c.categories,
		History:    c.history,
	}
	b, _ := json.MarshalIndent(state, "", "  ")
	os.WriteFile(path, b, 0644)
}

func (c *Categorizer) CategorizeAsync(command, generated, output string) {
	if strings.TrimSpace(command) == "" {
		return
	}
	// Limit output to prevent huge context windows
	if len(output) > 500 {
		output = output[:500]
	}
	select {
	case c.workerChan <- cmdTask{Command: command, Generated: generated, Output: output}:
	default:
		// Queue full, drop it so we don't block the REPL
	}
}

func (c *Categorizer) worker() {
	for task := range c.workerChan {
		c.mu.RLock()
		if _, exists := c.history[task.Command]; exists {
			c.mu.RUnlock()
			continue // Already categorized
		}
		catList, _ := json.Marshal(c.categories)
		c.mu.RUnlock()

		execContext := ""
		if task.Generated != "" {
			execContext = fmt.Sprintf("\nExecuted Under The Hood As: %s", task.Generated)
		}

		prompt := fmt.Sprintf(`You categorize terminal commands based on their output.
Existing categories: %s

Command: %s%s
Command Output: %s

Only reuse an existing category if it is highly specific and perfectly matches the command's core purpose.
Do NOT dump commands into overly broad categories like 'Shell' or 'Misc'.
If no highly specific category exists, invent ONE new broad category name (max 2 words, Title Case).
Respond in STRICT JSON format: {"category": "Name"}`, string(catList), task.Command, execContext, task.Output)

		// Call the LLM
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		resp, err := c.client.DoGenerateRaw(ctx, prompt, json.RawMessage(`"json"`))
		cancel()

		if err == nil && resp != "" {
			var result struct {
				Category string `json:"category"`
			}
			if json.Unmarshal([]byte(resp), &result) == nil && result.Category != "" && result.Category != "Unknown" {
				cat := strings.TrimSpace(result.Category)
				c.mu.Lock()
				c.history[task.Command] = cat
				found := false
				for _, ext := range c.categories {
					if strings.EqualFold(ext, cat) {
						found = true
						break
					}
				}
				if !found {
					c.categories = append(c.categories, cat)
				}
				c.save()
				c.mu.Unlock()
			}
		}
	}
}

// GetDomainList returns the list of unique domains/categories
func (c *Categorizer) GetDomainList() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make([]string, len(c.categories))
	copy(out, c.categories)
	return out
}

// GetCommandsInDomain returns all commands categorized under a specific domain (uses fuzzy matching)
func (c *Categorizer) GetCommandsInDomain(domain string) ([]string, string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	// First, find the closest matching category
	bestMatch := ""
	bestScore := -1
	lowerQuery := strings.ToLower(domain)
	
	for _, cat := range c.categories {
		lowerCat := strings.ToLower(cat)
		if lowerCat == lowerQuery {
			bestMatch = cat
			break
		}
		// Basic fuzzy matching: check if one contains the other
		if strings.Contains(lowerCat, lowerQuery) || strings.Contains(lowerQuery, lowerCat) {
			if bestScore < 10 {
				bestScore = 10
				bestMatch = cat
			}
		}
		// Token overlap for typos like "version controll" -> "version control"
		overlap := 0
		for _, w1 := range strings.Fields(lowerQuery) {
			for _, w2 := range strings.Fields(lowerCat) {
				if w1 == w2 || strings.HasPrefix(w1, w2) || strings.HasPrefix(w2, w1) {
					overlap++
				}
			}
		}
		if overlap > bestScore {
			bestScore = overlap
			bestMatch = cat
		}
	}
	
	if bestMatch == "" {
		return nil, domain
	}

	var out []string
	for cmd, cat := range c.history {
		if strings.EqualFold(cat, bestMatch) {
			out = append(out, cmd)
		}
	}
	return out, bestMatch
}
