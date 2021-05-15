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
	task         *task.Task
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
	return "", errors.New("ERROR TASK %s DOESN'T EXISTS")
}

func (d *Dag) RunTaskWithoutDependencies(tasks map[uuid.UUID]DagTask, dagChannel chan uuid.UUID) {
	for _, task := range tasks {
		if len(task.dependencies) == 0 {
			go task.task.Run(dagChannel)
		}
	}
}

func (d *Dag) RunDependentTask(tasks map[uuid.UUID]DagTask, dagChannel chan uuid.UUID) {

	for _, task := range tasks {

		isReady := true

		if len(task.dependencies) != 0 {
			for _, dependencyUUID := range task.dependencies {
				if d.tasks[dependencyUUID].task.GetStatus() != SuccessStatus || d.tasks[dependencyUUID].task.GetStatus() != FailStatus && !isReady {
					isReady = false
					continue
				}
				isReady = true
			}
			if isReady {
				go task.task.Run(dagChannel)
				delete(tasks, task.task.GetUuid())
			}
		}
	}
}

func (d *Dag) RunDag() {
	dagChannel := make(chan uuid.UUID)

	remainingTasks := make(map[uuid.UUID]DagTask)

	for i, task := range d.tasks {
		remainingTasks[i] = task
	}

	d.RunTaskWithoutDependencies(remainingTasks, dagChannel)

	for {
		taskUUID := <-dagChannel
		delete(remainingTasks, taskUUID)

		if len(remainingTasks) == 0 {
			break
		}
		d.RunDependentTask(remainingTasks, dagChannel)
	}
}
