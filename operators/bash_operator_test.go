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

func TestRunCmd(t *testing.T) {
	bash := NewBashOperator()
	cmd := "ls"
	bash.SetCmd(cmd)
	_, error := bash.RunTask()
	if error != nil {
		t.Errorf("An error occurs during the execution of the bash cmd")
	}
}

func TestRunCmdNoCodeFound(t *testing.T) {
	bash := NewBashOperator()
	_, error := bash.RunTask()
	// If no code is found, the function should return a error
	if error == nil {
		t.Errorf("The cmd should not have been run with no code")
	}
}
