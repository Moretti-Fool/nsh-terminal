package workflow

import (
	"testing"
	"time"
)

func TestSaveAndLoad(t *testing.T) {
	mgr := New(t.TempDir())
	wf := Workflow{
		Name: "run-backend", Created: time.Now(),
		Steps: []Step{{Command: "cd backend"}, {Command: "npm run dev"}},
	}
	if err := mgr.Save(wf); err != nil {
		t.Fatalf("save: %v", err)
	}
	loaded, err := mgr.Load("run-backend")
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if loaded.Name != "run-backend" || len(loaded.Steps) != 2 {
		t.Errorf("unexpected: %+v", loaded)
	}
}

func TestList(t *testing.T) {
	mgr := New(t.TempDir())
	for _, name := range []string{"alpha", "beta", "gamma"} {
		mgr.Save(Workflow{Name: name, Created: time.Now(), Steps: []Step{{Command: "echo"}}})
	}
	list, _ := mgr.List()
	if len(list) != 3 {
		t.Errorf("expected 3, got %d", len(list))
	}
}

func TestDelete(t *testing.T) {
	mgr := New(t.TempDir())
	mgr.Save(Workflow{Name: "temp", Created: time.Now(), Steps: []Step{{Command: "echo"}}})
	mgr.Delete("temp")
	if _, err := mgr.Load("temp"); err == nil {
		t.Error("should fail after delete")
	}
}

func TestNames(t *testing.T) {
	mgr := New(t.TempDir())
	mgr.Save(Workflow{Name: "a", Created: time.Now(), Steps: []Step{{Command: "echo"}}})
	mgr.Save(Workflow{Name: "b", Created: time.Now(), Steps: []Step{{Command: "echo"}}})
	if names := mgr.Names(); len(names) != 2 {
		t.Errorf("expected 2, got %d", len(names))
	}
}

func TestCreateFromCommands(t *testing.T) {
	mgr := New(t.TempDir())
	wf := mgr.CreateFromCommands("test", "desc", []string{"cmd1", "cmd2", "cmd3"})
	if len(wf.Steps) != 3 || wf.Name != "test" {
		t.Errorf("unexpected: %+v", wf)
	}
}
