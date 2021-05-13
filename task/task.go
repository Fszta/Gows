package task

import (
	"gows/operators"

	"github.com/google/uuid"
)

var (
	defaultStatus = "not started"
	runningStatus = "running"
	failStatus    = "fail"
	successStatus = "success"
)

type Task struct {
	operator operators.Operator
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
		status:   defaultStatus,
	}, nil
}

func (t *Task) Run() {
	t.updateStatus(runningStatus)

	logs, err := t.operator.RunTask()

	if err != nil {
		t.setLogs(err.Error())
		t.updateStatus(failStatus)
	}

	t.setLogs(logs)
	t.updateStatus(successStatus)
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
