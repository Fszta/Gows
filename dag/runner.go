package dag

import (
	"github.com/google/uuid"
)

func (d *Dag) RunTaskWithoutDependencies(tasks map[uuid.UUID]DagTask, dagChannel chan uuid.UUID) {
	for _, task := range tasks {
		if len(task.dependencies) == 0 {
			go task.task.Run(dagChannel)
		}
	}
}

func (d *Dag) RunDependentTask(tasks map[uuid.UUID]DagTask, dagChannel chan uuid.UUID, failChannel chan uuid.UUID) {

	for _, task := range tasks {

		isReady := true

		if len(task.dependencies) != 0 {
			for _, dependencyUUID := range task.dependencies {
				if d.tasks[dependencyUUID].task.GetStatus() == FailStatus {
					failChannel <- task.task.GetUuid()
					failChannel <- dependencyUUID
					return
				}

				if d.tasks[dependencyUUID].task.GetStatus() != SuccessStatus {
					isReady = false
					break
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
	failChannel := make(chan uuid.UUID)

	remainingTasks := make(map[uuid.UUID]DagTask)

	for i, task := range d.tasks {
		remainingTasks[i] = task
	}

	d.RunTaskWithoutDependencies(remainingTasks, dagChannel)

	for {
		select {
		case taskUUID := <-dagChannel:
			delete(remainingTasks, taskUUID)
			if len(remainingTasks) == 0 {
				return
			}
			go d.RunDependentTask(remainingTasks, dagChannel, failChannel)

		case failUUID := <-failChannel:
			delete(remainingTasks, failUUID)
			if len(remainingTasks) == 0 {
				return
			}
			for _, task := range remainingTasks {
				for _, dependencyUUID := range task.dependencies {
					if failUUID == dependencyUUID {
						delete(remainingTasks, task.task.GetUuid())
						if len(remainingTasks) == 0 {
							break
						}
					}
				}
			}
		}

	}
}
