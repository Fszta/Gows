package operators

import (
	"os"
	"strings"
	"testing"
)

func TestMakePythonCmdWithArgs(t *testing.T) {
	python := CreatePythonOperator()
	python.SetSrc("main.py")
	python.AddArgument("--name", "some_name")
	python.AddArgument("--schedule", "some_schedule")
	python.makeCmd()

	if python.cmd != "python main.py --name some_name --schedule some_schedule" {
		t.Errorf("The cmd was not properly built for the Python Operator")
	}

}

func TestMakePythonCmdNoArgs(t *testing.T) {
	python := CreatePythonOperator()
	python.SetSrc("main.py")
	python.makeCmd()

	if python.cmd != "python main.py" {
		t.Errorf("The cmd was not properly built for the Python Operator")
	}

}

func TestRunPythonTask(t *testing.T) {
	pythonCode := "import argparse\r\n" +
		"parser = argparse.ArgumentParser(description = 'my description')\r\n" +
		"parser.add_argument('-n', '--name')\r\n" +
		"args = parser.parse_args()\r\n" +
		"print('hello ' + args.name)"

	pythonCodePath := "mockup.py"
	f, _ := os.Create(pythonCodePath)
	f.WriteString(pythonCode)
	expectedOutput := "hello world"

	python := CreatePythonOperator()
	python.SetSrc(pythonCodePath)
	python.AddArgument("--name", "world")
	output, error := python.RunTask()
	output = strings.TrimSpace(output)
	if error != nil {
		t.Errorf("The python script return an error")
	}
	if output != expectedOutput {
		t.Errorf("The python script did not generated the expected output")
	}

	os.Remove(pythonCodePath)
}

func TestRunPythonTaskNoCode(t *testing.T) {
	pythonCodePath := "mockup.python"
	python := CreatePythonOperator()
	python.SetSrc(pythonCodePath)
	_, error := python.RunTask()

	// If no code is found, the function should return a error
	if error == nil {
		t.Errorf("The cmd should not have been run with no code")
	}
	os.Remove(pythonCodePath)
}
