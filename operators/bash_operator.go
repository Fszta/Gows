package operators

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type BashOperator struct {
	Path string `json:"path"`
}

func GetBashOperator(operatorParameters []byte) *BashOperator {
	var bash = BashOperator{}

	err := json.Unmarshal(operatorParameters, &bash)
	if err != nil {
		panic(err)
	}

	return &bash
}

func NewBashOperator() *BashOperator {
	return &BashOperator{
		Path: "test.sh",
	}
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

func (b *BashOperator) SetPath(path string) {
	b.Path = path
}

type OperatorConfig interface {
	SetConfig()
}

type BashConfig struct {
	name  string
	param string
}
