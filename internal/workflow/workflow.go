package workflow

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Step struct {
	Command string `toml:"command"`
	Shell   string `toml:"shell,omitempty"`
}

type Workflow struct {
	Name        string    `toml:"name"`
	Description string    `toml:"description,omitempty"`
	Created     time.Time `toml:"created"`
	Steps       []Step    `toml:"steps"`
}

type Manager struct {
	dir string
}

func New(dir string) *Manager {
	return &Manager{dir: dir}
}

func (m *Manager) Save(wf Workflow) error {
	if err := os.MkdirAll(m.dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(m.pathFor(wf.Name))
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(wf)
}

func (m *Manager) Load(name string) (Workflow, error) {
	data, err := os.ReadFile(m.pathFor(name))
	if err != nil {
		return Workflow{}, fmt.Errorf("workflow %q not found: %w", name, err)
	}
	var wf Workflow
	err = toml.Unmarshal(data, &wf)
	return wf, err
}

func (m *Manager) Delete(name string) error {
	return os.Remove(m.pathFor(name))
}

func (m *Manager) List() ([]Workflow, error) {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var workflows []Workflow
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".toml") {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".toml")
		wf, err := m.Load(name)
		if err != nil {
			continue
		}
		workflows = append(workflows, wf)
	}
	return workflows, nil
}

func (m *Manager) Names() []string {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		return nil
	}
	var names []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".toml") {
			continue
		}
		names = append(names, strings.TrimSuffix(entry.Name(), ".toml"))
	}
	return names
}

func (m *Manager) CreateFromCommands(name, description string, commands []string) Workflow {
	steps := make([]Step, len(commands))
	for i, cmd := range commands {
		steps[i] = Step{Command: cmd, Shell: "auto"}
	}
	return Workflow{
		Name:        name,
		Description: description,
		Created:     time.Now(),
		Steps:       steps,
	}
}

func (m *Manager) pathFor(name string) string {
	safe := strings.NewReplacer(" ", "-", "/", "-", "\\", "-").Replace(name)
	return filepath.Join(m.dir, safe+".toml")
}
