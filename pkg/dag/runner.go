package dag

import (
	"fmt"
	"time"

	"com.github/Fszta/gows/pkg/task"
	"github.com/google/uuid"
)

func (d *Dag) Run() {
	statusChannel := make(chan task.TaskStatus)
	remainingTasks := copyTasksMap(d.tasks)
	aTaskFinish := false

	d.setTime()
	d.resetDagStatus()
	d.runTaskWithoutDependencies(remainingTasks, statusChannel)

	for {
		if len(remainingTasks) == 0 && d.allTaskCompleted() {
			d.setStatus()
			fmt.Println("Dag ended at ", time.Now())
			break
		}

		newTaskStatus, ok := <-statusChannel

		if ok {
			aTaskFinish = true
			d.tasks[newTaskStatus.UUID].task.UpdateStatus(newTaskStatus.Status)

			if newTaskStatus.Status == SuccessStatus {
				delete(remainingTasks, newTaskStatus.UUID)
			}

			if newTaskStatus.Status == CancelStatus || newTaskStatus.Status == FailStatus {
				delete(remainingTasks, newTaskStatus.UUID)
				d.cancelDependenciesTask(remainingTasks, newTaskStatus.UUID, statusChannel)
			}
		}

		if aTaskFinish {
			d.tryRunTaskWithDependencies(remainingTasks, statusChannel)
			aTaskFinish = false
		}
	}
}

// runTaskWithoutDependencies run dag's top level task, which have no dependencies
func (d *Dag) runTaskWithoutDependencies(tasks map[uuid.UUID]DagTask, statusChannel chan task.TaskStatus) {
	for _, task := range tasks {
		if len(task.dependencies) == 0 {
			go task.task.Run(statusChannel)
		}
	}
}

// tryRunTaskWithDependencies run dag's tasks which have any dependencies only
// if the status of their dependencies is SuccessStatus
// removes the task if any dependency is failed or canceled
// otherwise continue without doing anything
func (d *Dag) tryRunTaskWithDependencies(remainingTasks map[uuid.UUID]DagTask, statusChannel chan task.TaskStatus) {
	for _, remainingTask := range remainingTasks {
		isReady := false
		for _, dependency := range remainingTask.dependencies {
			if d.tasks[dependency].task.GetStatus() == SuccessStatus {
				isReady = true
				continue
			}

			if d.tasks[dependency].task.GetStatus() == RunningStatus || d.tasks[dependency].task.GetStatus() == DefaultStatus {
				isReady = false
				continue
			}

			if d.tasks[dependency].task.GetStatus() == CancelStatus {
				isReady = false
				statusChannel <- task.TaskStatus{UUID: remainingTask.task.GetUuid(), Status: CancelStatus}
				break
			}
		}

		if isReady && d.tasks[remainingTask.task.GetUuid()].task.GetStatus() == DefaultStatus {
			go remainingTask.task.Run(statusChannel)
			delete(remainingTasks, remainingTask.task.GetUuid())
		}

	}
}

func (d *Dag) cancelDependenciesTask(remainingTasks map[uuid.UUID]DagTask, canceledTaskUUID uuid.UUID, statusChannel chan task.TaskStatus) {
	for _, remaining := range remainingTasks {
		for _, dependencyUUID := range remaining.dependencies {
			if canceledTaskUUID == dependencyUUID {
				fmt.Printf("remove task %s because have a canceled or failed dependency : %s\n", remaining.task.GetName(), d.tasks[canceledTaskUUID].task.GetName())
				d.tasks[remaining.task.GetUuid()].task.UpdateStatus(CancelStatus)
				delete(remainingTasks, remaining.task.GetUuid())
				d.cancelDependenciesTask(remainingTasks, remaining.task.GetUuid(), statusChannel)
			}
		}
	}
}

func (d *Dag) allTaskCompleted() bool {
	for _, dagTask := range d.tasks {
		if dagTask.task.GetStatus() == RunningStatus || dagTask.task.GetStatus() == DefaultStatus {
			return false
		}
	}
	return true
}

func (d *Dag) setStatus() {
	for _, dagTask := range d.tasks {
		if dagTask.task.GetStatus() == FailStatus {
			d.status = FailStatus
			return
		}
	}
	d.status = SuccessStatus
}

func (d *Dag) resetDagStatus() {
	for _, task := range d.tasks {
		task.task.UpdateStatus(DefaultStatus)
	}
}

func copyTasksMap(tasksToCopy map[uuid.UUID]DagTask) map[uuid.UUID]DagTask {
	copiedTasks := make(map[uuid.UUID]DagTask)

	for i, task := range tasksToCopy {
		copiedTasks[i] = task
	}

	return copiedTasks
}
