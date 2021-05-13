package task

import (
	"gows/operators"

	"github.com/google/uuid"
)

type Task struct {
	operator operators.IOperator
	uuid     uuid.UUID
	name     string
	status   string
	logs     string
}

func CreateTask(operatorName string, taskName string) (*Task, error) {
	taskUUID := uuid.New()
	operator, err := operators.NewOperator(operatorName)

	if err != nil {
		return nil, err
	}

	return &Task{
		operator: operator,
		uuid:     taskUUID,
		name:     taskName,
		status:   "not started",
	}, nil
}

func (t *Task) UpdateStatus(status string) {
	t.status = status
}

func (t *Task) GetLogs() string {
	return t.logs
}
