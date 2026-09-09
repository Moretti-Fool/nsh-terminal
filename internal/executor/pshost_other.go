//go:build !windows

package executor

import "os/exec"

func hideWindow(cmd *exec.Cmd) {}
