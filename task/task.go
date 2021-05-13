package task

import (
	"fmt"
	"gows/operators"

	"github.com/google/uuid"
)

type Task struct {
	operator operators.IOperator
	uuid     uuid.UUID
	name     string
	status   string
}

func CreateTask(operatorName string, taskName string) *Task {
	taskUUID := uuid.New()
	operator, err := operators.NewOperator(operatorName)

	if err != nil {
		fmt.Println(err)
	}

	return &Task{
		operator: operator,
		uuid:     taskUUID,
		name:     taskName,
	}
}

func (t *Task) SetStatus(status string) {
	t.status = status
}
