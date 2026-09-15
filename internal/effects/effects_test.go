package effects

import (
	"testing"
)

func TestPredict(t *testing.T) {
	tests := []struct {
		name       string
		command    string
		wantCat    string
		wantRisk   string
		wantWrite  bool
		wantRevers bool
	}{
		// Filesystem destructive
		{"rm", "rm -rf node_modules", "filesystem", "dangerous", true, false},
		{"del", "del /f /s /q *", "filesystem", "dangerous", true, false},
		{"Remove-Item", "Remove-Item -Recurse -Force .", "filesystem", "dangerous", true, false},
		{"rmdir", "rmdir /s /q dist", "filesystem", "dangerous", true, false},
		
		// Filesystem create/copy
		{"mkdir", "mkdir -p src/components", "filesystem", "safe", true, true},
		{"touch", "touch index.js", "filesystem", "safe", true, true},
		{"copy", "copy a.txt b.txt", "filesystem", "safe", true, true},
		
		// Filesystem move/rename
		{"mv", "mv old.txt new.txt", "filesystem", "moderate", true, true},
		{"rename", "Rename-Item old.txt new.txt", "filesystem", "moderate", true, true},
		
		// File content modification
		{"sed", "sed -i 's/foo/bar/g' file.txt", "filesystem", "moderate", true, false},
		{"overwrite", "echo hello > file.txt", "filesystem", "moderate", true, false},
		{"append", "echo hello >> file.txt", "filesystem", "safe", true, true},
		
		// Filesystem read
		{"ls", "ls -la", "filesystem", "safe", false, true},
		{"cat", "cat file.txt", "filesystem", "safe", false, true},
		{"grep", "grep -r 'TODO' .", "filesystem", "safe", false, true},
		
		// Git operations
		{"git reset", "git reset --hard HEAD~1", "git", "dangerous", true, false},
		{"git push", "git push origin main", "git", "moderate", true, true},
		{"git commit", "git commit -m 'msg'", "git", "safe", true, true},
		{"git status", "git status", "git", "safe", false, true},
		
		// Process management
		{"kill", "kill -9 1234", "process", "dangerous", true, false},
		{"Stop-Process", "Stop-Process -Name node", "process", "dangerous", true, false},
		
		// Network inspection
		{"curl", "curl -I https://example.com", "network", "safe", false, true},
		{"netstat", "netstat -ano", "network", "safe", false, true},
		
		// Package management
		{"npm install", "npm install react", "package", "moderate", true, true},
		{"apt install", "apt install htop", "package", "moderate", true, true},
		
		// System commands
		{"shutdown", "shutdown /s /t 0", "system", "dangerous", true, false},
		{"systemctl", "systemctl restart nginx", "system", "dangerous", true, true},
		{"chmod", "chmod 777 script.sh", "system", "dangerous", true, true},
		
		// Docker
		{"docker rm", "docker rm -f container", "docker", "dangerous", true, false},
		{"docker run", "docker run -d nginx", "docker", "moderate", true, true},
		{"docker ps", "docker ps", "docker", "safe", false, true},
		
		// Navigation
		{"cd", "cd ..", "filesystem", "safe", false, true},
		{"pwd", "pwd", "filesystem", "safe", false, true},
		
		// Unknown
		{"unknown", "some-unknown-cmd", "unknown", "moderate", true, true},
		
		// Chains
		{"chain max risk", "cd tmp && rm -rf *", "filesystem", "dangerous", true, false},
		{"pipe", "cat file.txt | grep foo | kill -9 1234", "process", "dangerous", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Predict(tt.command, "")
			if got.Category != tt.wantCat {
				t.Errorf("Predict(%q) category = %v, want %v", tt.command, got.Category, tt.wantCat)
			}
			if got.Risk != tt.wantRisk {
				t.Errorf("Predict(%q) risk = %v, want %v", tt.command, got.Risk, tt.wantRisk)
			}
			if got.WriteOp != tt.wantWrite {
				t.Errorf("Predict(%q) write = %v, want %v", tt.command, got.WriteOp, tt.wantWrite)
			}
			if got.Reversible != tt.wantRevers {
				t.Errorf("Predict(%q) revers = %v, want %v", tt.command, got.Reversible, tt.wantRevers)
			}
		})
	}
}

func TestFormatPreview(t *testing.T) {
	eff := Effect{Summary: "Deletes files recursively", Risk: "dangerous", WriteOp: true, Reversible: false}
	want := "[dangerous] Deletes files recursively (IRREVERSIBLE)"
	if got := eff.FormatPreview(); got != want {
		t.Errorf("FormatPreview() = %v, want %v", got, want)
	}

	eff2 := Effect{Summary: "Lists directory contents", Risk: "safe", WriteOp: false, Reversible: true}
	want2 := "[safe] Lists directory contents (read-only)"
	if got := eff2.FormatPreview(); got != want2 {
		t.Errorf("FormatPreview() = %v, want %v", got, want2)
	}
}
