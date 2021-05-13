package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	task, _ := CreateTask("bash", "my_task")

	if task.name != "my_task" && task.logs != "" && task.status != "not started" {
		t.Errorf("Task was not properly created")
	}

	_, err := CreateTask("unknow", "my_task")

	if err == nil {
		t.Errorf("Task creation should failed when unknow operator name is passed")
	}
}

func TestSetStatus(t *testing.T) {
	task, _ := CreateTask("bash", "my_task")
	status := "fail"
	task.UpdateStatus(status)

	if task.status != status {
		t.Errorf("The status was not set properly to status field of Task")
	}
}
