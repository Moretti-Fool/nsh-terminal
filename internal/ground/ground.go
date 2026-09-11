package ground

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func ExecTool(name string, args map[string]any, cwd string, pathExists func(string) bool) string {
	cmdStr := ""
	if name == "run_command" {
		cmdStr = argString(args, "command")
	} else {
		// The 1.5B model often hallucinates the shell command directly into the 'name' field
		cmdStr = name
		if cmdArg := argString(args, "command"); cmdArg != "" && cmdArg != name {
			cmdStr = cmdStr + " " + cmdArg
		}
	}
	
	if cmdStr == "" {
		return "error: missing command argument"
	}
	
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", cmdStr)
	} else {
		cmd = exec.Command("bash", "-c", cmdStr)
	}
	cmd.Dir = cwd
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "error: " + err.Error() + "\noutput: " + string(out)
	}
	return string(out)
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
