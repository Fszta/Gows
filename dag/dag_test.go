package dag

import (
	"gows/operators"
	"gows/task"
	"testing"
)

func getTestingBashOperator() operators.Operator {
	operator := operators.CreateBashOperator()
	operator.SetCmd("ls -all")
	return operator
}

func TestCreateDag(t *testing.T) {
	dag, _ := CreateDag("my_dag")
	if dag.name != "my_dag" && dag.status != DefaultStatus || len(dag.tasks) != 0 {
		t.Errorf("The Dag was not created properly")
	}

	_, error := CreateDag("")
	if error == nil {
		t.Errorf("The Dag creation should have failed when the Dag name is empty")
	}
}

func TestAddTask(t *testing.T) {
	dag, _ := CreateDag("my_dag")
	task, _ := task.CreateTask(getTestingBashOperator(), "my_task")
	taskUuid := task.GetUuid()

	initialDagSize := len(dag.tasks)

	dag.AddTask(task)

	if len(dag.tasks) != initialDagSize+1 {
		t.Errorf("The task was not added to the Dag")
	}

	addedTask := dag.GetTask(taskUuid)
	if addedTask.GetUuid() != taskUuid {
		t.Errorf("The task was not added properly at the proper key in the Dag")
	}
}

func TestAddMultipleTasks(t *testing.T) {
	dag, _ := CreateDag("my_dag")

	/* Creation of tasks */
	task1, _ := task.CreateTask(getTestingBashOperator(), "first_task")
	task1Uuid := task1.GetUuid()
	task2, _ := task.CreateTask(getTestingBashOperator(), "second_task")
	task2Uuid := task2.GetUuid()

	initialDagSize := len(dag.tasks)

	dag.AddMultiplesTasks([]*task.Task{task1, task2})

	if len(dag.tasks) != initialDagSize+2 {
		t.Errorf("The task was not added to the Dag")
	}

	addedTask := dag.GetTask(task1Uuid)
	if addedTask.GetUuid() != task1Uuid {
		t.Errorf("The task was not added properly at the proper key in the Dag")
	}

	addedTask = dag.GetTask(task2Uuid)
	if addedTask.GetUuid() != task2Uuid {
		t.Errorf("The task was not added properly at the proper key in the Dag")
	}

}

func TestGetTask(t *testing.T) {
	dag, _ := CreateDag("my_dag")
	task, _ := task.CreateTask(getTestingBashOperator(), "my_task")
	taskUuid := task.GetUuid()
	dag.AddTask(task)

	retrievedTask := dag.GetTask(taskUuid)
	if taskUuid != retrievedTask.GetUuid() {
		t.Errorf("The task was not properly retrieved from the Dag")
	}
}

func TestGetAllTask(t *testing.T) {
	dag, _ := CreateDag("my_dag")

	/* Creation of tasks */
	task1, _ := task.CreateTask(getTestingBashOperator(), "first_task")
	task1Uuid := task1.GetUuid()
	task2, _ := task.CreateTask(getTestingBashOperator(), "second_task")
	task2Uuid := task2.GetUuid()

	dag.AddMultiplesTasks([]*task.Task{task1, task2})

	allTasks := dag.GetAllTasks()

	if allTasks[0].GetUuid() != task1Uuid || allTasks[1].GetUuid() != task2Uuid {
		t.Errorf("All the tasks were not properly retreived")
	}
}

func TestSetDependency(t *testing.T) {
	dag, _ := CreateDag("my_dag")

	/* Creation of tasks */
	task1, _ := task.CreateTask(getTestingBashOperator(), "first_task")
	task2, _ := task.CreateTask(getTestingBashOperator(), "second_task")

	dag.AddMultiplesTasks([]*task.Task{task1, task2})
	dag.SetDependency(task1, task2)

	task1Dependencies := dag.GetTaskDependencies(task1)
	if task1Dependencies[0] != task2.GetUuid() {
		t.Errorf("The task2 was not properly added as a dependency of task1")
	}
}

func TestSetMultiplesDependencies(t *testing.T) {
	dag, _ := CreateDag("my_dag")

	/* Creation of tasks */
	task1, _ := task.CreateTask(getTestingBashOperator(), "first_task")
	task2, _ := task.CreateTask(getTestingBashOperator(), "second_task")
	task3, _ := task.CreateTask(getTestingBashOperator(), "third_task")

	dag.AddMultiplesTasks([]*task.Task{task1, task2, task3})
	dag.SetMultiplesDependencies(task1, []*task.Task{task2, task3})

	task1Dependencies := dag.GetTaskDependencies(task1)
	if task1Dependencies[0] != task2.GetUuid() && task1Dependencies[2] != task3.GetUuid() {
		t.Errorf("The task2 & task3 were not properly added as a dependency of task1")
	}
}

func TestGetAllTaskStatus(t *testing.T) {
	dag, _ := CreateDag("my_dag")
	task, _ := task.CreateTask(getTestingBashOperator(), "first_task")
	dag.AddTask(task)
	tasksStatus := dag.GetAllTaskStatus()

	if len(tasksStatus) != 1 {
		t.Errorf("The taskStatus map was not properly created")
	}

	if tasksStatus[task.GetUuid()] != DefaultStatus {
		t.Errorf("The wrong status code has been returned")
	}
}

func TestGetTaskStatus(t *testing.T) {
	dag, _ := CreateDag("my_dag")
	taskName := "my_task"

	task, _ := task.CreateTask(getTestingBashOperator(), taskName)
	dag.AddTask(task)

	taskStatus, _ := dag.GetTaskStatus(taskName)

	if taskStatus != DefaultStatus {
		t.Errorf("The wrong status code has been returned")
	}

	_, err := dag.GetTaskStatus("unknow_task")
	if err == nil {
		t.Errorf("An error should be returned when task doesn't exists")
	}
}

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

	if task3Status != DefaultStatus {
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
		task4Status != DefaultStatus || task5Status != DefaultStatus ||
		task6Status != SuccessStatus || task7Status != DefaultStatus {
		t.Errorf(
			"task1, task3, task6 should be succesfull, have : %s, %s, %s, task2 should be failed, have: %s, task4 and task5 should be not started, have: %s and %s ",
			task1Status, task3Status, task6Status, task2Status, task4Status, task5Status)
	}
}
