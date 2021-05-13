package dag

import (
	"errors"
	"gows/task"

	"github.com/google/uuid"
)

const (
	DefaultStatus = "not started"
	RunningStatus = "running"
	FailStatus    = "fail"
	SuccessStatus = "success"
)

type DagTask struct {
	task         task.Task
	dependencies []uuid.UUID
}

type Dag struct {
	uuid   uuid.UUID
	name   string
	status string
	tasks  map[uuid.UUID]DagTask
}

func CreateDag(dagName string) (*Dag, error) {
	if dagName == "" {
		return nil, errors.New("Error the Dag name provided was empty")
	}
	taskUUID := uuid.New()
	return &Dag{
		taskUUID,
		dagName,
		DefaultStatus,
		map[uuid.UUID]DagTask{},
	}, nil
}

func (d *Dag) AddTask(task *task.Task) {
	newTask := DagTask{*task, []uuid.UUID{}}
	d.tasks[task.GetUuid()] = newTask
}

func (d *Dag) AddMultipleTasks(tasks []*task.Task) {
	for _, taskToAdd := range tasks {
		newTask := DagTask{*taskToAdd, []uuid.UUID{}}
		d.tasks[taskToAdd.GetUuid()] = newTask
	}
}

func (d *Dag) GetTask(key uuid.UUID) *task.Task {
	var requestedTask task.Task = d.tasks[key].task
	return &requestedTask
}
