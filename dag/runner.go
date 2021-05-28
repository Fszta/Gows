package dag

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// RunTaksWithoutDependencies run dag's top level task, which have no dependencies
func (d *Dag) RunTaskWithoutDependencies(tasks map[uuid.UUID]DagTask, dagChannel, failChannel chan uuid.UUID) {
	for _, task := range tasks {
		if len(task.dependencies) == 0 {
			go task.task.Run(dagChannel, failChannel)
		}
	}
}

// RunDependentTask run dag's tasks which have any dependencies only
// if the status of their dependencies is SuccessStatus
// removes the task if any dependency is failed
// otherwise continue without doing anything
func (d *Dag) RunDependentTask(tasks map[uuid.UUID]DagTask, dagChannel, failChannel, cancelChannel chan uuid.UUID) {

	for _, task := range tasks {
		isReady := true

		if len(task.dependencies) != 0 {
			for _, dependencyUUID := range task.dependencies {
				// Check if a task depends on a failed or canceled task
				if d.tasks[dependencyUUID].task.GetStatus() == FailStatus || d.tasks[dependencyUUID].task.GetStatus() == CancelStatus {
					cancelChannel <- dependencyUUID
					continue
				}

				if d.tasks[dependencyUUID].task.GetStatus() != SuccessStatus {
					isReady = false
					break
				}

				isReady = true
			}
			if isReady && task.task.GetStatus() == DefaultStatus {
				go task.task.Run(dagChannel, failChannel)
			}
		}
	}
}

func (d *Dag) RunDag() {
	d.resetDagStatus()
	d.setTime()

	successChannel := make(chan uuid.UUID)
	cancelChannel := make(chan uuid.UUID)
	failChannel := make(chan uuid.UUID)
	remainingTasks := copyTasksMap(d.tasks)

	// Run top level tasks
	d.RunTaskWithoutDependencies(remainingTasks, successChannel, failChannel)

	for {
		select {
		// UUID sent by a task which ended successfully
		case successfulTaskUUID := <-successChannel:
			delete(remainingTasks, successfulTaskUUID)
			if len(remainingTasks) == 0 {
				fmt.Println("INFO: Dag ended at ", time.Now())
				return
			}
			go d.RunDependentTask(remainingTasks, successChannel, failChannel, cancelChannel)

		case canceledTaskUUID := <-cancelChannel:
			// Cancel tasks which depends on canceledTaskUUID
			d.cancelDependenciesTask(remainingTasks, canceledTaskUUID)

		// UUID sent by a task which failed
		case failedTaskUUID := <-failChannel:
			// Remove failing task from remaining tasks
			delete(remainingTasks, failedTaskUUID)

			// Cancel tasks which depends on failedTaskUUID
			d.cancelDependenciesTask(remainingTasks, failedTaskUUID)
			if len(remainingTasks) == 0 {
				fmt.Println("INFO: Dag ended at ", time.Now())
				return
			}
		}

	}
}

func (d *Dag) cancelDependenciesTask(remainingTasks map[uuid.UUID]DagTask, canceledTaskUUID uuid.UUID) {
	for _, task := range remainingTasks {
		for _, dependencyUUID := range task.dependencies {
			if canceledTaskUUID == dependencyUUID {
				fmt.Printf("remove task %s because have a canceled or failed dependency : %s\n", task.task.GetName(), d.tasks[canceledTaskUUID].task.GetName())
				task.task.UpdateStatus(CancelStatus)
				delete(remainingTasks, task.task.GetUuid())
			}
		}
	}
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
