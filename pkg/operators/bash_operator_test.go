package operators

import (
	"testing"
)

func TestSetCmd(t *testing.T) {
	bash := CreateBashOperator()
	cmdString := "some_cmd"
	bash.SetCmd(cmdString)

	if bash.cmd != cmdString {
		t.Errorf("The cmd was not properly set to the Cmd field of BashOperator")
	}
}

func TestMakeCmdWithArgs(t *testing.T) {
	bash := CreateBashOperator()
	bash.SetCmd("ls")
	bash.AddArgument("-n", "")
	bash.AddArgument("--file", "some_file.txt")
	bash.makeCmd()

	expectedValue := "ls -n --file some_file.txt"

	if bash.cmd != expectedValue {
		t.Errorf("The cmd string was not properly built")
	}
}

func TestMakeCmdNoArgs(t *testing.T) {
	bash := CreateBashOperator()
	bash.SetCmd("ls")
	bash.makeCmd()

	expectedValue := "ls"

	if bash.cmd != expectedValue {
		t.Errorf("The cmd string was not properly built")
	}
}

func TestRunCmd(t *testing.T) {
	bash := CreateBashOperator()
	cmd := "ls"
	bash.SetCmd(cmd)
	_, error := bash.RunTask()
	if error != nil {
		t.Errorf("An error occurs during the execution of the bash cmd")
	}
}

func TestRunCmdNoCodeFound(t *testing.T) {
	bash := CreateBashOperator()
	_, error := bash.RunTask()
	// If no code is found, the function should return a error
	if error == nil {
		t.Errorf("The cmd should not have been run with no code")
	}
}
