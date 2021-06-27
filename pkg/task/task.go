package task

import (
	"fmt"

	"com.github/Fszta/gows/pkg/operators"

	"github.com/google/uuid"
)

const (
	defaultStatus = "not started"
	runningStatus = "running"
	failStatus    = "fail"
	successStatus = "success"
)

type Task struct {
	Operator operators.Operator
	uuid     uuid.UUID
	name     string
	status   string
	logs     string
}

type TaskStatus struct {
	UUID   uuid.UUID
	Status string
}

func CreateTask(operator operators.Operator, taskName string) (*Task, error) {
	taskUUID := uuid.New()

	return &Task{
		Operator: operator,
		uuid:     taskUUID,
		name:     taskName,
		status:   defaultStatus,
	}, nil
}

func (t *Task) Run(statusChannel chan TaskStatus) error {
	fmt.Printf("Start running %v\n", t.name)
	statusChannel <- TaskStatus{UUID: t.uuid, Status: runningStatus}

	logs, err := t.Operator.RunTask()
	if err != nil {
		t.setLogs(err.Error())
		fmt.Printf("Task %v failed \n", t.name)
		statusChannel <- TaskStatus{UUID: t.uuid, Status: failStatus}
		return err
	}

	t.setLogs(logs)
	fmt.Printf("Task %v successfully ended \n", t.name)
	statusChannel <- TaskStatus{UUID: t.uuid, Status: successStatus}

	return nil
}

func (t *Task) UpdateStatus(status string) {
	t.status = status
}

func (t *Task) setLogs(logs string) {
	t.logs = logs
}

func (t *Task) GetLogs() string {
	return t.logs
}

func (t *Task) GetUuid() uuid.UUID {
	return t.uuid
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetStatus() string {
	return t.status
}
