package dag

import (
	"gows/operators"
	"gows/task"
	"testing"
)

func TestRunDagSequentialSuccess(t *testing.T) {

	dag, _ := CreateDag("sequential_success_dag")

	operator1 := operators.CreateBashOperator()
	operator1.SetCmd("ls -lah")
	task1, _ := task.CreateTask(operator1, "task1")
	dag.AddTask(task1)

	operator2 := operators.CreateBashOperator()
	operator2.SetCmd("sleep 1 && tree")
	task2, _ := task.CreateTask(operator2, "task2")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	operator3 := operators.CreateBashOperator()
	task3, _ := task.CreateTask(operator3, "task3")
	operator3.SetCmd("ps")
	dag.AddTask(task3)
	dag.SetDependency(task3, task2)

	dag.RunDag()

	task1Status, _ := dag.GetTaskStatus("task1")
	task2Status, _ := dag.GetTaskStatus("task2")
	task3Status, _ := dag.GetTaskStatus("task3")

	if task1Status != SuccessStatus || task2Status != SuccessStatus || task3Status != SuccessStatus {
		t.Errorf("All tasks should be successful %s %s %s", task1Status, task2Status, task3Status)
	}
}

func TestRunDagParallelSuccess(t *testing.T) {
	dag, _ := CreateDag("parallel_success_dag")

	operator1 := operators.CreateBashOperator()
	operator1.SetCmd("ls -lah")
	task1, _ := task.CreateTask(operator1, "task1")
	dag.AddTask(task1)

	operator2 := operators.CreateBashOperator()
	operator2.SetCmd("sleep 1 && echo toto")
	task2, _ := task.CreateTask(operator2, "task2")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	operator3 := operators.CreateBashOperator()
	operator3.SetCmd("tree")
	task3, _ := task.CreateTask(operator3, "task3")
	dag.AddTask(task3)
	dag.SetDependency(task3, task1)

	operator4 := operators.CreateBashOperator()
	operator4.SetCmd("ls")
	task4, _ := task.CreateTask(operator4, "task4")
	dag.AddTask(task4)
	dag.SetDependency(task4, task2)

	operator5 := operators.CreateBashOperator()
	operator5.SetCmd("ps")
	task5, _ := task.CreateTask(operator5, "task5")
	dag.AddTask(task5)
	dag.SetDependency(task5, task2)

	operator6 := operators.CreateBashOperator()
	operator6.SetCmd("echo toto")
	task6, _ := task.CreateTask(operator6, "task6")
	dag.AddTask(task6)
	dag.SetDependency(task6, task3)

	operator7 := operators.CreateBashOperator()
	operator7.SetCmd("echo toto && ls")
	task7, _ := task.CreateTask(operator7, "task7")
	dag.AddTask(task7)
	dag.SetMultiplesDependencies(task7, []*task.Task{task4, task5, task6})

	dag.RunDag()

	task1Status, _ := dag.GetTaskStatus("task1")
	task2Status, _ := dag.GetTaskStatus("task2")
	task3Status, _ := dag.GetTaskStatus("task3")
	task4Status, _ := dag.GetTaskStatus("task1")
	task5Status, _ := dag.GetTaskStatus("task2")
	task6Status, _ := dag.GetTaskStatus("task3")
	task7Status, _ := dag.GetTaskStatus("task3")

	if task1Status != SuccessStatus || task2Status != SuccessStatus || task3Status != SuccessStatus ||
		task4Status != SuccessStatus || task5Status != SuccessStatus ||
		task6Status != SuccessStatus || task7Status != SuccessStatus {
		t.Errorf("All tasks should be successful %s %s %s", task1Status, task2Status, task3Status)
	}
}

func TestRunDagSequentialFail(t *testing.T) {
	dag, _ := CreateDag("sequantial_fail_dag")

	operator1 := operators.CreateBashOperator()
	operator1.SetCmd("ls -lah")
	task1, _ := task.CreateTask(operator1, "task1")
	dag.AddTask(task1)

	operator2 := operators.CreateBashOperator()
	operator2.SetCmd("a bad bash command")
	task2, _ := task.CreateTask(operator2, "task2")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	operator3 := operators.CreateBashOperator()
	operator3.SetCmd("ps")
	task3, _ := task.CreateTask(operator3, "task3")
	dag.AddTask(task3)
	dag.SetDependency(task3, task2)

	dag.RunDag()

	task1Status, _ := dag.GetTaskStatus("task1")
	task2Status, _ := dag.GetTaskStatus("task2")
	task3Status, _ := dag.GetTaskStatus("task3")

	if task1Status != SuccessStatus {
		t.Errorf("Task1 should be successful")
	}

	if task2Status != FailStatus {
		t.Errorf("Task2 should be failed %s", task2Status)
	}

	if task3Status != CancelStatus {
		t.Errorf("Task3 should be not started %s", task3Status)
	}
}

func TestRunDagParallelFail(t *testing.T) {
	dag, _ := CreateDag("parallel_dag_fail")

	operator1 := operators.CreateBashOperator()
	operator1.SetCmd("ls -lah")
	task1, _ := task.CreateTask(operator1, "task1")
	dag.AddTask(task1)

	operator2 := operators.CreateBashOperator()
	operator2.SetCmd("a failling command")
	task2, _ := task.CreateTask(operator2, "task2")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	operator3 := operators.CreateBashOperator()
	operator3.SetCmd("tree")
	task3, _ := task.CreateTask(operator3, "task3")
	dag.AddTask(task3)
	dag.SetDependency(task3, task1)

	operator4 := operators.CreateBashOperator()
	operator4.SetCmd("ls")
	task4, _ := task.CreateTask(operator4, "task4")
	dag.AddTask(task4)
	dag.SetDependency(task4, task2)

	operator5 := operators.CreateBashOperator()
	operator5.SetCmd("ps")
	task5, _ := task.CreateTask(operator5, "task5")
	dag.AddTask(task5)
	dag.SetDependency(task5, task2)

	operator6 := operators.CreateBashOperator()
	operator6.SetCmd("echo toto")
	task6, _ := task.CreateTask(operator6, "task6")
	dag.AddTask(task6)
	dag.SetDependency(task6, task3)

	operator7 := operators.CreateBashOperator()
	operator7.SetCmd("echo toto && ls")
	task7, _ := task.CreateTask(operator7, "task7")
	dag.AddTask(task7)
	dag.SetMultiplesDependencies(task7, []*task.Task{task4, task5, task6})

	dag.RunDag()

	task1Status, _ := dag.GetTaskStatus("task1")
	task2Status, _ := dag.GetTaskStatus("task2")
	task3Status, _ := dag.GetTaskStatus("task3")
	task4Status, _ := dag.GetTaskStatus("task4")
	task5Status, _ := dag.GetTaskStatus("task5")
	task6Status, _ := dag.GetTaskStatus("task6")
	task7Status, _ := dag.GetTaskStatus("task7")

	if task1Status != SuccessStatus || task2Status != FailStatus || task3Status != SuccessStatus ||
		task4Status != CancelStatus || task5Status != CancelStatus ||
		task6Status != SuccessStatus || task7Status != CancelStatus {
		t.Errorf(
			"task1, task3, task6 should be succesfull, have : %s, %s, %s, task2 should be failed, have: %s, task4 and task5 should be canceled, have: %s and %s, task7 : %s ",
			task1Status, task3Status, task6Status, task2Status, task4Status, task5Status, task7Status)
	}
}
