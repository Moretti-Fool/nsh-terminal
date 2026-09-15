package effects

import (
	"fmt"
	"regexp"
	"strings"
)

// Effect describes the predicted outcome of running a command.
type Effect struct {
	Summary    string // e.g. "Deletes 'node_modules' directory recursively"
	Category   string // filesystem, process, network, package, git, system, docker, unknown
	Risk       string // safe, moderate, dangerous
	Reversible bool
	WriteOp    bool // true if the command modifies state
}

// FormatPreview returns a human-readable one-line preview of the effect.
func (e Effect) FormatPreview() string {
	reversibleInfo := ""
	if e.WriteOp {
		if e.Reversible {
			reversibleInfo = " (reversible)"
		} else {
			reversibleInfo = " (IRREVERSIBLE)"
		}
	} else {
		reversibleInfo = " (read-only)"
	}
	return fmt.Sprintf("[%s] %s%s", e.Risk, e.Summary, reversibleInfo)
}

type rule struct {
	regex  *regexp.Regexp
	effect Effect
}

var rules = []rule{
	// Git operations
	{regexp.MustCompile(`^git\s+reset\s+--hard`), Effect{Summary: "Hard resets git repository", Category: "git", Risk: "dangerous", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^git\s+branch\s+-D`), Effect{Summary: "Force deletes a git branch", Category: "git", Risk: "dangerous", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^git\s+clean\s+-f`), Effect{Summary: "Removes untracked files from git", Category: "git", Risk: "dangerous", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^git\s+push\s+-f`), Effect{Summary: "Force pushes to a git remote", Category: "git", Risk: "dangerous", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^git\s+push`), Effect{Summary: "Pushes commits to a remote", Category: "git", Risk: "moderate", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^git\s+commit`), Effect{Summary: "Commits changes to git", Category: "git", Risk: "safe", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^git\s+add`), Effect{Summary: "Stages files in git", Category: "git", Risk: "safe", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^git\s+checkout`), Effect{Summary: "Switches git branches or restores files", Category: "git", Risk: "moderate", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^git\s`), Effect{Summary: "Reads git repository state", Category: "git", Risk: "safe", Reversible: true, WriteOp: false}},

	// Docker operations
	{regexp.MustCompile(`^docker\s+rm\s+-f`), Effect{Summary: "Force removes docker container", Category: "docker", Risk: "dangerous", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^docker\s+rm`), Effect{Summary: "Removes docker container", Category: "docker", Risk: "moderate", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^docker\s+rmi`), Effect{Summary: "Removes docker image", Category: "docker", Risk: "moderate", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^docker\s+run`), Effect{Summary: "Runs a docker container", Category: "docker", Risk: "moderate", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^docker\s+stop`), Effect{Summary: "Stops a docker container", Category: "docker", Risk: "safe", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^docker\s+build`), Effect{Summary: "Builds a docker image", Category: "docker", Risk: "moderate", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^docker\s`), Effect{Summary: "Inspects docker state", Category: "docker", Risk: "safe", Reversible: true, WriteOp: false}},

	// Package management
	{regexp.MustCompile(`^(npm|yarn|pnpm)\s+install`), Effect{Summary: "Installs Node packages", Category: "package", Risk: "moderate", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^(npm|yarn|pnpm)\s+uninstall`), Effect{Summary: "Uninstalls Node packages", Category: "package", Risk: "moderate", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^pip\s+install`), Effect{Summary: "Installs Python packages", Category: "package", Risk: "moderate", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^pip\s+uninstall`), Effect{Summary: "Uninstalls Python packages", Category: "package", Risk: "moderate", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^(apt|apt-get|yum|dnf|pacman|brew|choco)\s+(install|remove|purge)`), Effect{Summary: "Modifies system packages", Category: "package", Risk: "moderate", Reversible: true, WriteOp: true}},

	// System commands
	{regexp.MustCompile(`^(shutdown|reboot|halt|poweroff)`), Effect{Summary: "Restarts or stops the system", Category: "system", Risk: "dangerous", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^systemctl\s+(start|stop|restart|disable|enable)`), Effect{Summary: "Modifies system services", Category: "system", Risk: "dangerous", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^Set-ExecutionPolicy`), Effect{Summary: "Changes PowerShell execution policy", Category: "system", Risk: "dangerous", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^(chmod|chown|icacls)`), Effect{Summary: "Changes file permissions", Category: "system", Risk: "dangerous", Reversible: true, WriteOp: true}},

	// Process management
	{regexp.MustCompile(`^(kill|killall|Stop-Process|taskkill)`), Effect{Summary: "Terminates a process", Category: "process", Risk: "dangerous", Reversible: false, WriteOp: true}},

	// Filesystem destructive
	{regexp.MustCompile(`^(rm|del|Remove-Item|rmdir|erase)\b`), Effect{Summary: "Deletes files or directories", Category: "filesystem", Risk: "dangerous", Reversible: false, WriteOp: true}},

	// Filesystem move/rename
	{regexp.MustCompile(`^(mv|move|Rename-Item)\b`), Effect{Summary: "Moves or renames files", Category: "filesystem", Risk: "moderate", Reversible: true, WriteOp: true}},

	// Filesystem create/copy
	{regexp.MustCompile(`^(mkdir|md|New-Item|touch)\b`), Effect{Summary: "Creates files or directories", Category: "filesystem", Risk: "safe", Reversible: true, WriteOp: true}},
	{regexp.MustCompile(`^(cp|copy|Copy-Item|xcopy|robocopy)\b`), Effect{Summary: "Copies files or directories", Category: "filesystem", Risk: "safe", Reversible: true, WriteOp: true}},

	// File content modification
	{regexp.MustCompile(`^sed\s+-i`), Effect{Summary: "Modifies file contents inline", Category: "filesystem", Risk: "moderate", Reversible: false, WriteOp: true}},
	{regexp.MustCompile(`^(tee|Set-Content|Add-Content)\b`), Effect{Summary: "Writes to a file", Category: "filesystem", Risk: "moderate", Reversible: false, WriteOp: true}},

	// Network inspection
	{regexp.MustCompile(`^(netstat|ss|curl|wget|ping|Get-NetTCPConnection|Test-NetConnection|ipconfig|ifconfig|nslookup|dig)\b`), Effect{Summary: "Interacts with or inspects network", Category: "network", Risk: "safe", Reversible: true, WriteOp: false}},

	// Filesystem read
	{regexp.MustCompile(`^(ls|dir|cat|head|tail|find|grep|Get-ChildItem|Get-Content|Select-String|less|more|type|echo|printf|Write-Output)\b`), Effect{Summary: "Reads files or directories or prints text", Category: "filesystem", Risk: "safe", Reversible: true, WriteOp: false}},

	// Navigation
	{regexp.MustCompile(`^(cd|Set-Location|pwd|pushd|popd|Get-Location)\b`), Effect{Summary: "Navigates the filesystem", Category: "filesystem", Risk: "safe", Reversible: true, WriteOp: false}},
}

// Predict analyzes a command string and returns its predicted effect.
// This is a rules-based engine, not an LLM call — it's instant.
func Predict(command string, cwd string) Effect {
	command = strings.TrimSpace(command)
	if command == "" {
		return Effect{Summary: "Empty command", Category: "unknown", Risk: "safe", Reversible: true, WriteOp: false}
	}

	parts := splitCommand(command)

	maxRiskValue := 0
	var overallEffect Effect
	overallEffect.Category = "unknown"
	overallEffect.Risk = "safe"
	overallEffect.Reversible = true
	overallEffect.Summary = "Unknown command"

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		eff := evaluatePart(part)

		riskVal := riskToInt(eff.Risk)
		if riskVal > maxRiskValue {
			maxRiskValue = riskVal
			overallEffect = eff
		} else if riskVal == maxRiskValue && overallEffect.Category == "unknown" {
			overallEffect = eff
		}

		if eff.WriteOp {
			overallEffect.WriteOp = true
		}
		if !eff.Reversible {
			overallEffect.Reversible = false
		}
	}

	return overallEffect
}

func riskToInt(risk string) int {
	switch risk {
	case "safe":
		return 1
	case "moderate":
		return 2
	case "dangerous":
		return 3
	default:
		return 0
	}
}

func splitCommand(cmd string) []string {
	var parts []string
	replacer := strings.NewReplacer("&&", "\x00", "||", "\x00", ";", "\x00", "|", "\x00")
	s := replacer.Replace(cmd)

	for _, p := range strings.Split(s, "\x00") {
		parts = append(parts, p)
	}

	return parts
}

func evaluatePart(part string) Effect {
	hasOverwrite := strings.Contains(part, ">") && !strings.Contains(part, ">>")
	hasAppend := strings.Contains(part, ">>")

	cmdStr := part
	if idx := strings.Index(cmdStr, ">"); idx != -1 {
		cmdStr = strings.TrimSpace(cmdStr[:idx])
	}

	var eff Effect
	matched := false

	for _, r := range rules {
		if r.regex.MatchString(cmdStr) {
			eff = r.effect
			matched = true
			break
		}
	}

	if !matched {
		eff = Effect{
			Summary:    "Unknown command behavior",
			Category:   "unknown",
			Risk:       "moderate",
			Reversible: true,
			WriteOp:    true,
		}
	}

	if hasOverwrite {
		eff.WriteOp = true
		eff.Reversible = false
		if riskToInt(eff.Risk) < riskToInt("moderate") {
			eff.Risk = "moderate"
		}
		if matched {
			eff.Summary = eff.Summary + " and overwrites file"
		} else {
			eff.Summary = "Overwrites file"
			eff.Category = "filesystem"
		}
	} else if hasAppend {
		eff.WriteOp = true
		if riskToInt(eff.Risk) < riskToInt("safe") {
			eff.Risk = "safe"
		}
		if matched {
			eff.Summary = eff.Summary + " and appends to file"
		} else {
			eff.Summary = "Appends to file"
			eff.Category = "filesystem"
		}
	}

	return eff
}
