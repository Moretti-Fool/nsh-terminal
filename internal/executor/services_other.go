//go:build !windows

package executor

import (
	"os/exec"
	"syscall"
)

// setServiceProcessGroup puts the service process into its own process group
// on Unix so that signals sent to nsh (e.g. SIGINT) do not propagate to
// service child processes.
func setServiceProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}

// gracefulStopService sends SIGTERM to the process group, giving the service
// a chance to shut down cleanly.
func gracefulStopService(cmd *exec.Cmd) {
	if cmd.Process == nil {
		return
	}
	// Send SIGTERM to the entire process group (negative PID).
	syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM)
}

// forceKillService sends SIGKILL to the process group.
func forceKillService(cmd *exec.Cmd) {
	if cmd.Process == nil {
		return
	}
	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
}
