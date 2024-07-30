package utils

import (
	"os/exec"
)

func OSExecuteCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.Output()
	return string(stdout), err
}
