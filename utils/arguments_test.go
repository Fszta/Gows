package operators

import (
	"testing"
)

func TestArgsToStringNoValue(t *testing.T) {
	arguments := []Argument{
		{
			Arg:   "-n",
			Value: "",
		},
		{
			Arg:   "--file",
			Value: "some_file.txt",
		},
	}
	expectedValue := "-n --file some_file.txt"
	actualValue := GetArgsToString(arguments)

	if actualValue != expectedValue {
		t.Errorf("The bash arguments were not properly stringified")
	}
}
