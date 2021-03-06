package task

import (
	"testing"

	"com.github/Fszta/gows/pkg/operators"
)

func TestCreateTask(t *testing.T) {
	operator := operators.CreateBashOperator()
	operator.SetCmd("ls -all")
	task, _ := CreateTask(operator, "my_task")

	if task.name != "my_task" && task.logs != "" && task.status != "not started" {
		t.Errorf("Task was not properly created")
	}
}

func TestUpdateStatus(t *testing.T) {
	operator := operators.CreateBashOperator()
	operator.SetCmd("ls -all")
	task, _ := CreateTask(operator, "my_task")
	status := "fail"
	task.UpdateStatus(status)

	if task.status != status {
		t.Errorf("The status was not set properly to status field of Task")
	}
}
