package main

import (
	"fmt"
	"gows/operators"
)

func main() {
	operator, err := operators.NewOperator("bash")
	operator.SetCmd("ls")
	operator.SetArguments([]operators.Argument{
		{Arg: "-all", Value: ""},
	})
	output, _ := operator.RunTask()
	fmt.Println(output)

	if err != nil {
		panic(err)
	}
	operator.RunTask()
}
