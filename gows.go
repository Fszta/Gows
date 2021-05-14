package main

import (
	"fmt"
	"gows/dag"
	"gows/operators"
	"gows/task"
)

func main() {
	operator, _ := operators.NewOperator("bash")
	operator.SetCmd("ls")
	output, _ := operator.RunTask()
	fmt.Println(output)

	task1, _ := task.CreateTask("bash", "task1")
	task1.Operator.SetCmd("ls -lah")

	dag, _ := dag.CreateDag("my_dag")
	dag.AddTask(task1)

	task2, _ := task.CreateTask("bash", "task2")
	task2.Operator.SetCmd("sleep 3 && tree")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	task3, _ := task.CreateTask("bash", "task3")
	task3.Operator.SetCmd("docker ps -aq")
	dag.AddTask(task3)
	dag.SetDependency(task3, task2)

	dag.RunDag()
}
