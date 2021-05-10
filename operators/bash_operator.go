package operators

import (
	"encoding/json"
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

func NewBashOperator(operatorParameters []byte) *BashOperator {
	var bash = BashOperator{}

	err := json.Unmarshal(operatorParameters, &bash)
	if err != nil {
		panic(err)
	}
	return &bash
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
