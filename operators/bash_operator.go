package operators

import (
	"errors"
	"os/exec"
	"strings"
)

type Argument struct {
	Arg   string
	Value string
}

type BashOperator struct {
	Cmd       string
	Arguments []Argument
}

func NewBashOperator() *BashOperator {
	return &BashOperator{}
}

func (b *BashOperator) SetCmd(cmd string) {
	b.Cmd = cmd
}

func (b *BashOperator) SetArguments(args []Argument) {
	b.Arguments = args
}

func (b *BashOperator) GetArgsToString() string {
	var argsString strings.Builder
	for _, arg := range b.Arguments {
		if arg.Arg != "" {
			argsString.WriteString(" " + arg.Arg)
		}
		if arg.Value != "" {
			argsString.WriteString(" " + arg.Value)
		}
	}
	// Return the trimed stringified arguments
	return strings.TrimSpace(argsString.String())
}

func (b *BashOperator) RunTask() (string, error) {
	if b.Cmd == "" {
		return "", errors.New("Error some bash code was found")
	}
	cmdString := b.Cmd + " " + b.GetArgsToString()
	cmd := exec.Command("bash", "-c", cmdString)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.New("Error occured during the script execution")
	}
	return string(stdout), nil
}
