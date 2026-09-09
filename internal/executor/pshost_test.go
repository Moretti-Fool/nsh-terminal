package executor

import (
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestLooksLikePowerShell(t *testing.T) {
	yes := []string{
		"Get-ChildItem",
		"Get-ChildItem | Sort-Object Length -Descending | Select-Object -First 1",
		"Get-Process | Where-Object { $_.CPU -gt 1 }",
		"Write-Output 'hi'",
		"Restart-Service Spooler",
	}
	no := []string{
		"git status",
		"echo hello",
		"ls -la",
		"dir /b",
		"ipconfig /all",
	}
	for _, c := range yes {
		if !looksLikePowerShell(c) {
			t.Errorf("expected powershell: %s", c)
		}
	}
	for _, c := range no {
		if looksLikePowerShell(c) {
			t.Errorf("did not expect powershell: %s", c)
		}
	}
}

func TestLooksLikeUnknownCommand(t *testing.T) {
	if !LooksLikeUnknownCommand("'Get-ChildItem' is not recognized as an internal or external command") {
		t.Fatal("expected unknown-command match")
	}
	if LooksLikeUnknownCommand("Get-ChildItem: Cannot find path") {
		t.Fatal("path errors should not look like unknown command")
	}
}

func TestPSHostWarmReuse(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-only")
	}
	e := New("auto", "")
	defer e.Close()

	first, err := e.RunGenerated(`Write-Output 'nsh-host-ok'`)
	if err != nil {
		t.Fatalf("first: %v", err)
	}
	if first.ExitCode != 0 {
		t.Fatalf("first exit %d output %s", first.ExitCode, first.Output)
	}
	if !strings.Contains(first.Output, "nsh-host-ok") {
		t.Fatalf("first output: %s", first.Output)
	}

	start := time.Now()
	second, err := e.RunGenerated(`Write-Output 'nsh-host-ok-2'`)
	elapsed := time.Since(start)
	if err != nil {
		t.Fatalf("second: %v", err)
	}
	if second.ExitCode != 0 {
		t.Fatalf("second exit %d output %s", second.ExitCode, second.Output)
	}
	if !strings.Contains(second.Output, "nsh-host-ok-2") {
		t.Fatalf("second output: %s", second.Output)
	}
	if elapsed > 1500*time.Millisecond {
		t.Fatalf("warm command took %s, expected well under 1.5s", elapsed)
	}
	if second.DurationMs > 1500 {
		t.Fatalf("reported duration %dms too high for warm host", second.DurationMs)
	}
}

func TestPSHostGetChildItem(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-only")
	}
	e := New("auto", "")
	defer e.Close()

	result, err := e.RunGenerated(`Get-ChildItem -Name | Select-Object -First 5`)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.ExitCode != 0 {
		t.Fatalf("exit %d output %s", result.ExitCode, result.Output)
	}
	if strings.TrimSpace(result.Output) == "" {
		t.Fatal("expected file names")
	}
}

func TestPSHostLargestFile(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-only")
	}
	e := New("auto", "")
	defer e.Close()

	cmd := `Get-ChildItem -File | Sort-Object Length -Descending | Select-Object -First 1 -ExpandProperty Name`
	result, err := e.RunGenerated(cmd)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.ExitCode != 0 {
		t.Fatalf("exit %d output %s", result.ExitCode, result.Output)
	}
}

func TestRunTypedCmdletUsesHost(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-only")
	}
	e := New("auto", "")
	defer e.Close()
	result, err := e.Run(`Write-Output 'via-run'`)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.ExitCode != 0 || !strings.Contains(result.Output, "via-run") {
		t.Fatalf("exit %d output %s", result.ExitCode, result.Output)
	}
}

func TestPSHostFindstrDoesNotHang(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-only")
	}
	e := New("auto", "")
	defer e.Close()
	start := time.Now()
	result, err := e.RunGenerated(`findstr /n "package" pshost.go`)
	if time.Since(start) > 8*time.Second {
		t.Fatalf("findstr hung: %s", time.Since(start))
	}
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.ExitCode != 0 && !strings.Contains(result.Output, "package") {
		t.Fatalf("exit %d output %s", result.ExitCode, result.Output)
	}
}

func TestNLShellName(t *testing.T) {
	e := New("auto", "")
	if runtime.GOOS == "windows" && e.NLShellName() != "powershell" {
		t.Fatalf("got %s", e.NLShellName())
	}
	if runtime.GOOS == "windows" {
		shell, _ := e.DetectShell()
		if !strings.Contains(strings.ToLower(shell), "cmd") {
			t.Fatalf("typed auto shell should remain cmd, got %s", shell)
		}
	}
}
