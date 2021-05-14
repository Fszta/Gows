package main

import (
	"fmt"
	"gows/dag"
	"gows/task"
)

func main() {

	task1, _ := task.CreateTask("bash", "task1")
	task1.Operator.SetCmd("ls -lah")

	dag, _ := dag.CreateDag("my_dag")
	dag.AddTask(task1)

	task2, _ := task.CreateTask("bash", "task2")
	task2.Operator.SetCmd("sleep 3 && tree")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	task3, _ := task.CreateTask("bash", "task3")
	task3.Operator.SetCmd("docker ps -aqp")
	dag.AddTask(task3)
	dag.SetDependency(task3, task2)

	dag.RunDag()

	for uuid, status := range dag.GetAllTaskStatus() {
		fmt.Printf("%s : %s \n", uuid.String(), status)
	}
}
