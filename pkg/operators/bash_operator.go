package operators

import (
	"errors"
	"os/exec"
	"strings"
)

type BashOperator struct {
	ArgumentsHandler
	cmd string
}

func CreateBashOperator() *BashOperator {
	return &BashOperator{}
}

func (b *BashOperator) SetCmd(cmd string) {
	b.cmd = cmd
}

func (b *BashOperator) makeCmd() {
	var cmdBuilder strings.Builder
	cmdBuilder.WriteString(b.cmd)

	argsString := b.getArgsToString()
	if argsString != "" {
		cmdBuilder.WriteString(" " + argsString)
	}
	b.cmd = cmdBuilder.String()
}

func (b *BashOperator) RunTask() (string, error) {

	b.makeCmd()

	if b.cmd == "" {
		return "", errors.New("ERROR no bash code was found")
	}

	cmd := exec.Command("bash", "-c", b.cmd)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.New("ERROR occured during the script execution")
	}
	return string(stdout), nil
}
