//go:build linux
// +build linux

package pythonrt

import (
	"os/exec"
	"syscall"
)

const WNOHANG = syscall.WNOHAN

func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func wait4Wrapper(pid int, status *syscall.WaitStatus, options int, rusage *syscall.Rusage) (int, error) {
	return syscall.Wait4(pid, status, options, rusage)
}
