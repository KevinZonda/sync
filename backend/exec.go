package backend

import (
	"fmt"
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
	err := execCmd.Run()
	fmt.Println("COV", execCmd.String())

	return err
}

type Base struct {
	Location string
}
