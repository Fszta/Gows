package task

import (
	"gows/operators"

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

func CreateTask(operator operators.Operator, taskName string) (*Task, error) {
	taskUUID := uuid.New()

	return &Task{
		Operator: operator,
		uuid:     taskUUID,
		name:     taskName,
		status:   defaultStatus,
	}, nil
}

func (t *Task) Run(dagChannel chan uuid.UUID) error {
	t.updateStatus(runningStatus)
	logs, err := t.Operator.RunTask()
	dagChannel <- t.uuid

	if err != nil {
		t.setLogs(err.Error())
		t.updateStatus(failStatus)
		return err
	}

	t.setLogs(logs)
	t.updateStatus(successStatus)

	return nil
}

func (t *Task) updateStatus(status string) {
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
