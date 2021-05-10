package main

import (
	"gows/operators"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("operator.json")

	if err != nil {
		panic(err)
	}

	operator, _ := operators.NewOperator(data)

	operator.RunTask()
}
