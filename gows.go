package main

import (
	"gows/dag"
	"gows/operators"
	"gows/task"
)

func main() {

	dag1, _ := dag.CreateDag("my_dag1")
	operator1 := operators.CreateBashOperator()
	operator1.SetCmd("ps")
	task1, _ := task.CreateTask(operator1, "First Bash Task")
	dag1.AddTask(task1)

	operator2 := operators.CreateBashOperator()
	operator2.SetCmd("ls")
	task2, _ := task.CreateTask(operator2, "Second Bash Task")
	dag1.AddTask(task2)
	dag1.SetDependency(task2, task1)

	operator3 := operators.CreateBashOperator()
	operator3.SetCmd("echo 1")
	task3, _ := task.CreateTask(operator3, "Third Python Task")
	dag1.AddTask(task3)
	dag1.SetDependency(task3, task2)

	operator4 := operators.CreateBashOperator()
	operator4.SetCmd("echo 1")
	task4, _ := task.CreateTask(operator4, "Fourth Python Task")
	dag1.AddTask(task4)
	dag1.SetDependency(task4, task2)

	dag.ScheduleDag("*/3 * * * * *)", dag1)
}
