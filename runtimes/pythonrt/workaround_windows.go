//go:build windows
// +build windows

package pythonrt

import (
	"errors"
	"os/exec"
)

const WNOHANG = 0

func setSysProcAttr(cmd *exec.Cmd) {
	// Windows doesn't support Setpgid, so leave it as nil.
	cmd.SysProcAttr = nil
}

func wait4Wrapper(pid int, status interface{}, options int, rusage interface{}) (int, error) {
	return 0, errors.New("Wait4 is not supported on Windows")
}
