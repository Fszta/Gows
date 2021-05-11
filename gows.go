package main

import (
	"gows/operators"
)

func main() {
	operator, err := operators.NewOperator("bash")

	if err != nil {
		panic(err)
	}
	operator.RunTask()
}
