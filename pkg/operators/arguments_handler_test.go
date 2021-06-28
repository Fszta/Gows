package operators

import (
	"testing"
)

func TestSetArgument(t *testing.T) {
	argumentsHandler := ArgumentsHandler{}
	argumentsHandler.AddArgument("--file", "some_file.txt")

	if (len(argumentsHandler.arguments)) != 1 {
		t.Errorf("The argument was not added to the argument list")
	}

	if argumentsHandler.arguments[0].arg != "--file" || argumentsHandler.arguments[0].value != "some_file.txt" {
		t.Errorf("The wrong value habe been added to the arguments list")
	}
}

func TestArgsToString(t *testing.T) {
	arguments := []Argument{
		{
			arg:   "-n",
			value: "",
		},
		{
			arg:   "--file",
			value: "some_file.txt",
		},
	}
	expectedValue := "-n --file some_file.txt"
	argumentsHandler := ArgumentsHandler{}
	argumentsHandler.arguments = arguments
	actualValue := argumentsHandler.getArgsToString()

	if actualValue != expectedValue {
		t.Errorf("The bash arguments were not properly stringified")
	}
}

func TestArgsToStringNoArgs(t *testing.T) {
	arguments := []Argument{}
	expectedValue := ""
	argumentsHandler := ArgumentsHandler{}
	argumentsHandler.arguments = arguments
	actualValue := argumentsHandler.getArgsToString()

	if actualValue != expectedValue {
		t.Errorf("The bash arguments were not properly stringified")
	}
}
