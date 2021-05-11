package operators

import (
	"fmt"
	"os/exec"
)

type BashOperator struct {
	Path      string `json:"path"`
	Arguments []struct {
		Arg   string `json:"arg"`
		Value string `json:"value"`
	} `json:"arguments"`
}

func NewBashOperator() *BashOperator {
	return &BashOperator{}
}

func (b *BashOperator) RunTask() {
	cmd := exec.Command("bash", b.Path)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}
