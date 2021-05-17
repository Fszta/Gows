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
	dag, _ := CreateDag("empty_dag")
	if dag.name != "my_dag" && dag.status != DefaultStatus || len(dag.tasks) != 0 {
		t.Errorf("The Dag was not created properly")
	}

	_, error := CreateDag("")
	if error == nil {
		t.Errorf("The Dag creation should have failed when the Dag name is empty")
	}
}

func TestAddTask(t *testing.T) {
	dag, _ := CreateDag("one_task_dag")
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
	dag, _ := CreateDag("multiple_tasks_dag")

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
	dag, _ := CreateDag("a_dag")
	task, _ := task.CreateTask(getTestingBashOperator(), "my_task")
	taskUuid := task.GetUuid()
	dag.AddTask(task)

	retrievedTask := dag.GetTask(taskUuid)
	if taskUuid != retrievedTask.GetUuid() {
		t.Errorf("The task was not properly retrieved from the Dag")
	}
}

func TestGetAllTask(t *testing.T) {
	dag, _ := CreateDag("dag_test")

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
	dag, _ := CreateDag("dag")

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
	dag, _ := CreateDag("dag_task_multiple_dependencies")

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
	dag, _ := CreateDag("any_dag")
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
	dag, _ := CreateDag("another_dag")
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
