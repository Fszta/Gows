package dag

import (
	"errors"
	"gows/task"
	"time"

	"github.com/google/uuid"
)

const (
	DefaultStatus = "not started"
	RunningStatus = "running"
	FailStatus    = "fail"
	SuccessStatus = "success"
	CancelStatus  = "cancel"
)

type DagTask struct {
	task         *task.Task
	dependencies []uuid.UUID
}

type Dag struct {
	uuid         uuid.UUID
	name         string
	status       string
	tasks        map[uuid.UUID]DagTask
	DagScheduler *DagScheduler
	lastRunTime  time.Time
}

func CreateDag(dagName string) (*Dag, error) {
	if dagName == "" {
		return nil, errors.New("ERROR: the Dag name provided was empty")
	}
	UUID := uuid.New()
	return &Dag{
		uuid:   UUID,
		name:   dagName,
		status: DefaultStatus,
		tasks:  map[uuid.UUID]DagTask{},
	}, nil
}

func (d *Dag) GetTaskLevel() map[uuid.UUID]int {
	dagLevel := make(map[uuid.UUID]int)

	for uuid, task := range d.tasks {
		if len(task.dependencies) == 0 {
			dagLevel[uuid] = 0
			continue
		}
		for _, dependency := range task.dependencies {
			parentLevel := dagLevel[dependency]
			dagLevel[uuid] = parentLevel + 1
		}
	}
	return dagLevel
}

func (d *Dag) AddTask(task *task.Task) {
	newTask := DagTask{task, []uuid.UUID{}}
	d.tasks[task.GetUuid()] = newTask
}

func (d *Dag) AddMultiplesTasks(tasks []*task.Task) {
	for _, task := range tasks {
		d.AddTask(task)
	}
}

func (d *Dag) GetTask(key uuid.UUID) *task.Task {
	return d.tasks[key].task
}

func (d *Dag) GetAllTasks() []*task.Task {
	allTasks := make([]*task.Task, 0)
	for _, taskItem := range d.tasks {
		allTasks = append(allTasks, taskItem.task)
	}
	return allTasks
}

func (d *Dag) SetScheduler(cronFormat string) {
	d.DagScheduler = NewScheduler(d, cronFormat)
}

func (d *Dag) SetDependency(task *task.Task, dependencyTask *task.Task) {
	dagTasksRef := d.tasks[task.GetUuid()]
	dagTasksRef.dependencies = append(dagTasksRef.dependencies, dependencyTask.GetUuid())
	d.tasks[task.GetUuid()] = dagTasksRef
}

func (d *Dag) SetMultiplesDependencies(task *task.Task, dependencyTasks []*task.Task) {
	for _, dependencyTask := range dependencyTasks {
		d.SetDependency(task, dependencyTask)
	}
}

func (d *Dag) GetTaskDependencies(task *task.Task) []uuid.UUID {
	return d.tasks[task.GetUuid()].dependencies
}

func (d *Dag) GetAllTaskStatus() map[uuid.UUID]string {
	tasksStatus := map[uuid.UUID]string{}
	for taskId, task := range d.tasks {
		status := task.task.GetStatus()
		tasksStatus[taskId] = status
	}
	return tasksStatus
}

func (d *Dag) GetTaskStatus(taskName string) (string, error) {
	for _, task := range d.tasks {
		if task.task.GetName() == taskName {
			status := task.task.GetStatus()
			return status, nil
		}
	}
	return "", errors.New("ERROR: Task %s doesn't exist")
}

func (d *Dag) GetUUID() uuid.UUID {
	return d.uuid
}

func (d *Dag) setTime() {
	d.lastRunTime = time.Now()
}
