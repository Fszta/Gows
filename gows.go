package main

import (
	"fmt"
	"gows/dag"
	"gows/operators"
	"gows/task"
)

func main() {

	dag, _ := dag.CreateDag("my_dag")

	operator1 := operators.CreateBashOperator()
	operator1.SetCmd("../sample-src/mockup.sh -n world")
	task1, _ := task.CreateTask(operator1, "First Bash Task")
	dag.AddTask(task1)

	operator2 := operators.CreateBashOperator()
	operator2.SetCmd("../sample-src/mockup.sh")
	operator2.AddArgument("-n", "world")
	task2, _ := task.CreateTask(operator2, "Second Bash Task")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	operator3 := operators.CreatePythonOperator()
	operator3.SetSrc("../sample-src/mockup.py")
	operator3.AddArgument("--name", "world")
	task3, _ := task.CreateTask(operator3, "Third Python Task")
	dag.AddTask(task3)
	dag.SetDependency(task3, task2)

	for _, value := range dag.GetAllTasks() {
		taskName := value.GetName()
		fmt.Println(taskName)
	}

	// dag.RunDag()

	// for uuid, status := range dag.GetAllTaskStatus() {
	// 	fmt.Printf("%s : %s \n", uuid.String(), status)
	// }
}
