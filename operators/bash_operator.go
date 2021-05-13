package operators

import (
	"errors"
	"os/exec"
)

type BashOperator struct {
	Cmd string
}

func NewBashOperator() *BashOperator {
	return &BashOperator{}
}

func (b *BashOperator) SetCmd(cmd string) {
	b.Cmd = cmd
}

func (b *BashOperator) RunTask() (string, error) {
	if b.Cmd == "" {
		return "", errors.New("Error some bash code was found")
	}

	cmd := exec.Command("bash", "-c", b.Cmd)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.New("Error occured during the script execution")
	}
	return string(stdout), nil
}
