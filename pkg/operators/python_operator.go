package operators

import (
	"errors"
	"os/exec"
	"strings"
)

type PythonOperator struct {
	ArgumentsHandler
	src string
	cmd string
}

func CreatePythonOperator() *PythonOperator {
	return &PythonOperator{}
}

func (p *PythonOperator) SetSrc(srcPath string) {
	p.src = srcPath
}

func (p *PythonOperator) makeCmd() {
	var cmdBuilder strings.Builder
	cmdBuilder.WriteString("python ")
	cmdBuilder.WriteString(p.src)

	argsString := p.getArgsToString()
	if argsString != "" {
		cmdBuilder.WriteString(" " + argsString)
	}
	p.cmd = cmdBuilder.String()
}

func (p *PythonOperator) RunTask() (string, error) {

	p.makeCmd()

	if p.cmd == "" {
		return "", errors.New("ERROR no bash code was found")
	}

	cmd := exec.Command("bash", "-c", p.cmd)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.New("ERROR occured during the script execution")
	}
	return string(stdout), nil
}
