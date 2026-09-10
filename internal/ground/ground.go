package ground

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const maxListing = 40

// Snapshot is cheap local evidence for the NL translator. It is not a task catalog.
type Snapshot struct {
	CWD     string
	Listing string
	Present []string
}

// Capture lists the current directory and a tiny allowlist of tools that exist on PATH.
func Capture(cwd string, pathExists func(string) bool) Snapshot {
	if cwd == "" {
		cwd, _ = os.Getwd()
	}
	s := Snapshot{CWD: cwd, Listing: ListCWD(cwd, maxListing)}
	for _, name := range []string{"git", "docker", "rg", "python", "python3"} {
		if pathExists != nil && pathExists(name) {
			s.Present = append(s.Present, name)
		}
	}
	return s
}

// ListCWD returns a compact directory listing. Directories end with /.
func ListCWD(cwd string, limit int) string {
	if limit <= 0 {
		limit = maxListing
	}
	ents, err := os.ReadDir(cwd)
	if err != nil {
		return "(could not read directory: " + err.Error() + ")"
	}
	var b strings.Builder
	n := 0
	for _, e := range ents {
		if n >= limit {
			fmt.Fprintf(&b, "... %d more\n", len(ents)-n)
			break
		}
		name := e.Name()
		if e.IsDir() {
			b.WriteString(name)
			b.WriteString("/\n")
			n++
			continue
		}
		size := int64(0)
		if info, err := e.Info(); err == nil {
			size = info.Size()
		}
		fmt.Fprintf(&b, "%s %d\n", name, size)
		n++
	}
	return strings.TrimSpace(b.String())
}

// ExecTool runs one translator tool. Unknown names return an error string, not a panic.
func ExecTool(name string, args map[string]any, cwd string, pathExists func(string) bool) string {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "list_cwd":
		return ListCWD(cwd, maxListing)
	case "which":
		prog := argString(args, "name")
		if prog == "" {
			return "missing argument: name"
		}
		if pathExists != nil && pathExists(prog) {
			if p, err := exec.LookPath(prog); err == nil {
				return p
			}
			return prog + " is on PATH"
		}
		return prog + " is not on PATH"
	case "git_status":
		if pathExists != nil && !pathExists("git") {
			return "git is not on PATH"
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "git", "-C", cwd, "status", "--short", "-b")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			return strings.TrimSpace(out.String() + "\n" + err.Error())
		}
		s := strings.TrimSpace(out.String())
		if s == "" {
			return "(clean working tree)"
		}
		return trimRunes(s, 1500)
	case "netstat":
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "powershell", "-NoProfile", "-Command", "Get-NetTCPConnection -State Listen | Select-Object -Property LocalAddress,LocalPort,OwningProcess | ConvertTo-Json")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			return trimRunes(strings.TrimSpace(out.String()+"\n"+err.Error()), 2000)
		}
		return trimRunes(strings.TrimSpace(out.String()), 2000)
	case "env_var":
		name := argString(args, "name")
		if name == "" {
			return "missing argument: name"
		}
		if val, ok := os.LookupEnv(name); ok {
			if val == "" {
				// Just return empty if actually set to empty
				return val
			}
			return val
		}
		return "(not set)"
	case "read_file_head":
		path := argString(args, "path")
		if path == "" {
			return "missing argument: path"
		}
		fullPath := path
		if !filepath.IsAbs(path) {
			fullPath = filepath.Join(cwd, path)
		}
		f, err := os.Open(fullPath)
		if err != nil {
			return err.Error()
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		var b strings.Builder
		for i := 0; i < 20 && scanner.Scan(); i++ {
			b.WriteString(scanner.Text())
			b.WriteString("\n")
		}
		if err := scanner.Err(); err != nil {
			b.WriteString(err.Error())
		}
		return trimRunes(strings.TrimSpace(b.String()), 2000)
	default:
		return "unknown tool: " + name
	}
}

func argString(args map[string]any, key string) string {
	if args == nil {
		return ""
	}
	if v, ok := args[key]; ok {
		return strings.TrimSpace(fmt.Sprint(v))
	}
	for k, v := range args {
		if strings.EqualFold(k, key) {
			return strings.TrimSpace(fmt.Sprint(v))
		}
	}
	return ""
}

func trimRunes(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
