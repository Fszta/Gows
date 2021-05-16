package dag

import (
	"gows/task"
	"testing"
)

func TestRunDagSequentialSuccess(t *testing.T) {
	dag, _ := CreateDag("my_dag")

	task1, _ := task.CreateTask("bash", "task1")
	task1.Operator.SetCmd("ls -lah")
	dag.AddTask(task1)

	task2, _ := task.CreateTask("bash", "task2")
	task2.Operator.SetCmd("sleep 1 && tree")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	task3, _ := task.CreateTask("bash", "task3")
	task3.Operator.SetCmd("ps")
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
	dag, _ := CreateDag("my_dag")

	task1, _ := task.CreateTask("bash", "task1")
	task1.Operator.SetCmd("ls -lah")
	dag.AddTask(task1)

	task2, _ := task.CreateTask("bash", "task2")
	task2.Operator.SetCmd("sleep 1 && tree")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	task3, _ := task.CreateTask("bash", "task3")
	task3.Operator.SetCmd("ps")
	dag.AddTask(task3)
	dag.SetDependency(task3, task1)

	task4, _ := task.CreateTask("bash", "task4")
	task4.Operator.SetCmd("ps")
	dag.AddTask(task4)
	dag.SetDependency(task4, task2)

	task5, _ := task.CreateTask("bash", "task5")
	task5.Operator.SetCmd("ps")
	dag.AddTask(task5)
	dag.SetDependency(task5, task2)

	task6, _ := task.CreateTask("bash", "task6")
	task6.Operator.SetCmd("ps")
	dag.AddTask(task6)
	dag.SetDependency(task6, task3)

	task7, _ := task.CreateTask("bash", "task7")
	task7.Operator.SetCmd("ps")
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
	dag, _ := CreateDag("my_dag")
	task1, _ := task.CreateTask("bash", "task1")
	task1.Operator.SetCmd("ls -lah")
	dag.AddTask(task1)

	task2, _ := task.CreateTask("bash", "task2")
	task2.Operator.SetCmd("a bad bash command")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	task3, _ := task.CreateTask("bash", "task3")
	task3.Operator.SetCmd("ps")
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
	dag, _ := CreateDag("my_dag")

	task1, _ := task.CreateTask("bash", "task1")
	task1.Operator.SetCmd("ls -lah")
	dag.AddTask(task1)

	task2, _ := task.CreateTask("bash", "task2")
	task2.Operator.SetCmd("a failling command")
	dag.AddTask(task2)
	dag.SetDependency(task2, task1)

	task3, _ := task.CreateTask("bash", "task3")
	task3.Operator.SetCmd("tree")
	dag.AddTask(task3)
	dag.SetDependency(task3, task1)

	task4, _ := task.CreateTask("bash", "task4")
	task4.Operator.SetCmd("ls")
	dag.AddTask(task4)
	dag.SetDependency(task4, task2)

	task5, _ := task.CreateTask("bash", "task5")
	task5.Operator.SetCmd("ps")
	dag.AddTask(task5)
	dag.SetDependency(task5, task2)

	task6, _ := task.CreateTask("bash", "task6")
	task6.Operator.SetCmd("echo toto")
	dag.AddTask(task6)
	dag.SetDependency(task6, task3)

	task7, _ := task.CreateTask("bash", "task7")
	task7.Operator.SetCmd("echo toto && ls")
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
			"task1, task3, task6 should be succesfull, have : %s, %s, %s, task2 should be failed, have: %s, task4 and task5 should be not started, have: %s and %s ",
			task1Status, task3Status, task6Status, task2Status, task4Status, task5Status)
	}
}
