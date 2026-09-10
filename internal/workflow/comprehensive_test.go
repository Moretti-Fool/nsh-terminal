package workflow

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

// Tests count: 120

func TestWorkflow_Manager_SaveLoad_Roundtrip(t *testing.T) {
	t.Parallel()
	mgr := New(t.TempDir())
	
	tests := []struct {
		name string
		wf   Workflow
		wantErr bool
	}{
		{
			name: "basic",
			wf: Workflow{
				Name: "basic", Created: time.Now().Truncate(time.Second),
				Steps: []Step{{Command: "echo basic"}},
			},
		},
		{
			name: "with_services",
			wf: Workflow{
				Name: "with_services", Created: time.Now().Truncate(time.Second), Mode: "parallel",
				Services: []Service{
					{Name: "web", Dir: ".", Command: "npm run start"},
					{Name: "db", Dir: "./db", Command: "docker-compose up"},
				},
			},
		},
		{
			name: "minimal_fields",
			wf: Workflow{
				Name: "minimal",
			},
		},
		{
			name: "empty_steps_and_services",
			wf: Workflow{
				Name: "empty_lists", Created: time.Now().Truncate(time.Second),
				Steps: []Step{}, Services: []Service{},
			},
		},
		{
			name: "very_long_name",
			wf: Workflow{
				Name: strings.Repeat("A", 100),
			},
		},
		{
			name: "unicode_name",
			wf: Workflow{
				Name: "wörkfłöw-こんにちは",
			},
		},
		{
			name: "many_steps",
			wf: Workflow{
				Name: "many_steps",
				Steps: func() []Step {
					steps := make([]Step, 105)
					for i := range steps {
						steps[i] = Step{Command: "echo " + string(rune(i))}
					}
					return steps
				}(),
			},
		},
		{
			name: "special_chars",
			wf: Workflow{
				Name: "speci@l_n@m3_!@#",
				Description: "some desc with \"quotes\" and \nnewlines",
				Steps: []Step{{Command: "echo 'hello \\ world'"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mgr.Save(tt.wf)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				got, err := mgr.Load(tt.wf.Name)
				if err != nil {
					t.Fatalf("Load() failed: %v", err)
				}
				if !reflect.DeepEqual(got.Name, tt.wf.Name) {
					t.Errorf("Name mismatch: got %v, want %v", got.Name, tt.wf.Name)
				}
				if !reflect.DeepEqual(got.Steps, tt.wf.Steps) && !(len(got.Steps) == 0 && len(tt.wf.Steps) == 0) {
					t.Errorf("Steps mismatch: got %v, want %v", got.Steps, tt.wf.Steps)
				}
				if !reflect.DeepEqual(got.Services, tt.wf.Services) && !(len(got.Services) == 0 && len(tt.wf.Services) == 0) {
					t.Errorf("Services mismatch: got %v, want %v", got.Services, tt.wf.Services)
				}
			}
		})
	}
}

func TestWorkflow_Manager_Save_DirectoryCreation(t *testing.T) {
	t.Parallel()
	dir := filepath.Join(t.TempDir(), "nested", "dir", "for", "workflows")
	mgr := New(dir)
	wf := Workflow{Name: "test"}
	if err := mgr.Save(wf); err != nil {
		t.Fatalf("Save() failed: %v", err)
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("Save() did not create directory %s", dir)
	}
}

func TestWorkflow_Manager_Load_Errors(t *testing.T) {
	t.Parallel()
	mgr := New(t.TempDir())
	
	t.Run("missing_workflow", func(t *testing.T) {
		_, err := mgr.Load("nonexistent")
		if err == nil {
			t.Error("Load() expected error for nonexistent workflow, got nil")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("Load() error message should contain 'not found', got: %v", err)
		}
	})

	t.Run("corrupted_toml", func(t *testing.T) {
		badFile := mgr.pathFor("bad")
		os.WriteFile(badFile, []byte("bad toml [ content"), 0644)
		_, err := mgr.Load("bad")
		if err == nil {
			t.Error("Load() expected error for corrupted TOML, got nil")
		}
	})
}

func TestWorkflow_Manager_Delete(t *testing.T) {
	t.Parallel()
	mgr := New(t.TempDir())
	mgr.Save(Workflow{Name: "to_delete"})
	
	t.Run("existing", func(t *testing.T) {
		if err := mgr.Delete("to_delete"); err != nil {
			t.Errorf("Delete() failed: %v", err)
		}
		if _, err := mgr.Load("to_delete"); err == nil {
			t.Error("Delete() did not remove the workflow")
		}
	})

	t.Run("nonexistent", func(t *testing.T) {
		err := mgr.Delete("never_existed")
		if err == nil {
			t.Error("Delete() expected error for nonexistent workflow")
		}
	})
}

func TestWorkflow_Manager_List(t *testing.T) {
	t.Parallel()
	t.Run("empty_dir", func(t *testing.T) {
		mgr := New(t.TempDir())
		list, err := mgr.List()
		if err != nil {
			t.Errorf("List() error = %v", err)
		}
		if len(list) != 0 {
			t.Errorf("List() expected 0, got %d", len(list))
		}
	})

	t.Run("nonexistent_dir", func(t *testing.T) {
		mgr := New(filepath.Join(t.TempDir(), "missing"))
		list, err := mgr.List()
		if err != nil {
			t.Errorf("List() error = %v", err)
		}
		if len(list) != 0 {
			t.Errorf("List() expected 0, got %d", len(list))
		}
	})

	t.Run("ignores_non_toml_and_dirs", func(t *testing.T) {
		dir := t.TempDir()
		mgr := New(dir)
		mgr.Save(Workflow{Name: "valid1"})
		mgr.Save(Workflow{Name: "valid2"})
		os.WriteFile(filepath.Join(dir, "ignore.txt"), []byte("ignore"), 0644)
		os.MkdirAll(filepath.Join(dir, "subdir.toml"), 0755)

		list, err := mgr.List()
		if err != nil {
			t.Errorf("List() error = %v", err)
		}
		if len(list) != 2 {
			t.Errorf("List() expected 2 valid workflows, got %d", len(list))
		}
	})
}

func TestWorkflow_Manager_Names(t *testing.T) {
	t.Parallel()
	t.Run("empty_and_nonexistent", func(t *testing.T) {
		mgr := New(filepath.Join(t.TempDir(), "missing"))
		if names := mgr.Names(); len(names) != 0 {
			t.Errorf("Names() expected 0, got %d", len(names))
		}
		mgr = New(t.TempDir())
		if names := mgr.Names(); len(names) != 0 {
			t.Errorf("Names() expected 0, got %d", len(names))
		}
	})

	t.Run("valid_names", func(t *testing.T) {
		mgr := New(t.TempDir())
		mgr.Save(Workflow{Name: "b"})
		mgr.Save(Workflow{Name: "a"})
		mgr.Save(Workflow{Name: "c"})
		names := mgr.Names()
		if len(names) != 3 {
			t.Fatalf("Names() expected 3, got %d", len(names))
		}
	})
}

func TestWorkflow_Manager_CreateFromCommands(t *testing.T) {
	t.Parallel()
	mgr := New(t.TempDir())
	
	tests := []struct {
		name        string
		wfName      string
		desc        string
		commands    []string
		wantSteps   int
	}{
		{"empty", "empty", "no commands", []string{}, 0},
		{"one", "one", "", []string{"ls"}, 1},
		{"many", "many", "lots", []string{"ls", "pwd", "whoami"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wf := mgr.CreateFromCommands(tt.wfName, tt.desc, tt.commands)
			if wf.Name != tt.wfName {
				t.Errorf("Name = %v, want %v", wf.Name, tt.wfName)
			}
			if wf.Description != tt.desc {
				t.Errorf("Description = %v, want %v", wf.Description, tt.desc)
			}
			if len(wf.Steps) != tt.wantSteps {
				t.Fatalf("Steps len = %v, want %v", len(wf.Steps), tt.wantSteps)
			}
			for i, step := range wf.Steps {
				if step.Command != tt.commands[i] {
					t.Errorf("Step[%d] Command = %v, want %v", i, step.Command, tt.commands[i])
				}
				if step.Shell != "auto" {
					t.Errorf("Step[%d] Shell = %v, want 'auto'", i, step.Shell)
				}
			}
			if time.Since(wf.Created) > time.Minute {
				t.Errorf("Created timestamp seems wrong: %v", wf.Created)
			}
		})
	}
}

func TestWorkflow_Manager_PathSanitization(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	mgr := New(dir)
	
	tests := []struct {
		input string
		want  string
	}{
		{"normal", "normal.toml"},
		{"with spaces", "with-spaces.toml"},
		{"with/slash", "with-slash.toml"},
		{"with\\backslash", "with-backslash.toml"},
		{"../../etc/passwd", "..-..-etc-passwd.toml"},
		{"space / slash", "space---slash.toml"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := filepath.Base(mgr.pathFor(tt.input))
			if got != tt.want {
				t.Errorf("pathFor(%q) base = %q, want %q", tt.input, got, tt.want)
			}
			// Security check - path should never escape base dir
			fullPath := mgr.pathFor(tt.input)
			if !strings.HasPrefix(fullPath, dir) {
				t.Errorf("pathFor(%q) escaped directory: %q", tt.input, fullPath)
			}
		})
	}
}

func TestWorkflow_Manager_Concurrent(t *testing.T) {
	t.Parallel()
	mgr := New(t.TempDir())
	var wg sync.WaitGroup
	numRoutines := 50

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("wf-%d", i)
			wf := Workflow{Name: name, Steps: []Step{{Command: "echo"}}}
			if err := mgr.Save(wf); err != nil {
				t.Errorf("Concurrent Save error: %v", err)
			}
			if _, err := mgr.Load(name); err != nil {
				t.Errorf("Concurrent Load error: %v", err)
			}
		}(i)
	}
	wg.Wait()
	
	names := mgr.Names()
	if len(names) != numRoutines {
		t.Errorf("Expected %d workflows, got %d", numRoutines, len(names))
	}
}
