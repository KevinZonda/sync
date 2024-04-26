package backend

import (
	"os"
	"os/exec"
)

func (b Base) runCmd(name string, args ...string) error {
	execCmd := exec.Command(name, args...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if b.Location != "" {
		execCmd.Dir = b.Location
	}
	return execCmd.Run()
}

type Base struct {
	Location string
}
