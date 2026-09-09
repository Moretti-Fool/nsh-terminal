package ollama

import (
	"fmt"
	"regexp"
	"strings"
)

func NormalizeShell(name string) string {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "pwsh", "powershell":
		return "powershell"
	case "cmd", "cmd.exe":
		return "cmd"
	case "bash", "sh", "zsh":
		return "bash"
	default:
		return strings.ToLower(strings.TrimSpace(name))
	}
}

var cmdDirSwitch = regexp.MustCompile(`(?i)\bdir\s+/[a-z]`)
var unixFind = regexp.MustCompile(`(?i)\bfind\s+\.\s+-name\b`)

// ValidatePlan rejects unusable output without deciding which user task it is.
func ValidatePlan(plan Plan, expectedShell string) error {
	if plan.Join() == "" {
		return fmt.Errorf("no commands in model output")
	}
	joined := plan.Join()
	if LooksLikeBinDump(joined) {
		return fmt.Errorf("output looks like a binary list, not a command")
	}
	want := NormalizeShell(expectedShell)
	if plan.Dialect != "" && NormalizeShell(plan.Dialect) != want {
		return fmt.Errorf("model used dialect %s; this session uses %s", plan.Dialect, want)
	}
	for _, cmd := range plan.Commands {
		if reason := DialectError(cmd, want); reason != "" {
			return fmt.Errorf("%s", reason)
		}
	}
	return nil
}

// DialectError returns a repair hint if cmd is clearly the wrong shell family.
func DialectError(cmd, shell string) string {
	lower := strings.ToLower(cmd)
	switch NormalizeShell(shell) {
	case "powershell":
		if cmdDirSwitch.MatchString(cmd) {
			return "cmd.exe dir switches are not valid in PowerShell: " + cmd
		}
		if unixFind.MatchString(cmd) {
			return "unix find is not PowerShell; use Get-ChildItem -Recurse -Filter: " + cmd
		}
	case "cmd":
		if strings.Contains(lower, "get-childitem") || strings.Contains(lower, "sort-object") ||
			strings.Contains(lower, "select-object") || strings.Contains(lower, "select-string") {
			return "PowerShell cmdlet is not valid in cmd.exe: " + cmd
		}
	case "bash":
		if strings.Contains(lower, "get-childitem") {
			return "PowerShell cmdlet is not valid in bash: " + cmd
		}
	}
	return ""
}
