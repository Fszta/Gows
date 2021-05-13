package operators

import (
	"testing"
)

func TestSetCmd(t *testing.T) {
	bash := NewBashOperator()
	cmdString := "some_cmd"
	bash.SetCmd(cmdString)

	if bash.Cmd != cmdString {
		t.Errorf("The cmd was not properly set to the Cmd field of BashOperator")
	}
}

func TestSetArguments(t *testing.T) {
	bash := NewBashOperator()
	args := []Argument{
		{
			Arg:   "-n",
			Value: "some-value",
		},
		{
			Arg:   "-f",
			Value: "some-file",
		},
	}
	bash.SetArguments(args)
	if bash.Arguments[0] != args[0] && bash.Arguments[1] != args[1] {
		t.Errorf("The arguments were not properly set to the Argument filed of BashOperator")
	}
}

func TestArgsToString(t *testing.T) {
	bash := NewBashOperator()
	bash.SetArguments([]Argument{
		{
			Arg:   "-n",
			Value: "some_name",
		},
		{
			Arg:   "--file",
			Value: "some_file",
		},
	})
	expectedArgsToString := "-n some_name --file some_file"
	actualArgsToString := bash.GetArgsToString()

	if actualArgsToString != expectedArgsToString {
		t.Errorf("The bash arguments were not properly stringified")
	}
}

func TestArgsToStringNoValue(t *testing.T) {
	bash := NewBashOperator()
	bash.SetArguments([]Argument{
		{
			Arg:   "-n",
			Value: "",
		},
		{
			Arg:   "-f",
			Value: "",
		},
	})
	expectedValue := "-n -f"
	actualValue := bash.GetArgsToString()

	if actualValue != expectedValue {
		t.Errorf("The bash arguments were not properly stringified")
	}
}

func TestRunCmd(t *testing.T) {
	bash := NewBashOperator()
	cmd := "ls"
	args := []Argument{
		{
			Arg:   "-all",
			Value: "",
		},
	}
	bash.SetCmd(cmd)
	bash.SetArguments(args)
	_, error := bash.RunTask()
	if error != nil {
		t.Errorf("An error occurs during the execution of the bash cmd")
	}
}
