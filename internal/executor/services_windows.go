//go:build windows

package executor

import (
	"fmt"
	"os/exec"
	"syscall"
)

// setServiceProcessGroup puts the service process into its own process group.
// This is critical: without it, console signals (Ctrl+C, CTRL_CLOSE_EVENT,
// and system-generated events like sleep/wake, RDP disconnect, Windows Update
// interruptions) propagate from the nsh console to every child process in the
// same group, silently killing all running services.
func setServiceProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// CREATE_NEW_PROCESS_GROUP (0x00000200) detaches the child from the
		// parent's console group so it does NOT receive the parent's Ctrl
		// events.  This is the single most important fix for services dying
		// unexpectedly.
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
}

// gracefulStopService sends CTRL_BREAK_EVENT to the service's process group,
// giving it a chance to shut down cleanly (e.g. close DB connections, flush
// logs).  CTRL_BREAK cannot be ignored by the child, but well-behaved programs
// treat it as a shutdown signal.
func gracefulStopService(cmd *exec.Cmd) {
	if cmd.Process == nil {
		return
	}
	// Send CTRL_BREAK_EVENT to the process group.
	// This is the correct way to signal a process that was started with
	// CREATE_NEW_PROCESS_GROUP.
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return
	}
	proc, err := dll.FindProc("GenerateConsoleCtrlEvent")
	if err != nil {
		return
	}
	// CTRL_BREAK_EVENT = 1, target = the process group ID (same as PID when
	// the process was created with CREATE_NEW_PROCESS_GROUP).
	proc.Call(1, uintptr(cmd.Process.Pid))
}

// forceKillService uses taskkill /F /T to kill the process tree.
func forceKillService(cmd *exec.Cmd) {
	if cmd.Process == nil {
		return
	}
	exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprintf("%d", cmd.Process.Pid)).Run()
}
