package main

import (
	"fmt"
	"gows/operators"
)

func main() {
	operator, _ := operators.NewOperator("bash")
	operator.SetCmd("ls")
	output, _ := operator.RunTask()
	fmt.Println(output)
}
