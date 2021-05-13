package dag

import (
	"gows/task"
	"testing"
)

func TestCreateDag(t *testing.T) {
	dag, _ := CreateDag("my_dag")
	if dag.name != "my_dag" && dag.status != DefaultStatus && len(dag.tasks) != 0 {
		t.Errorf("The Dag was not created properly")
	}

	_, error := CreateDag("")
	if error == nil {
		t.Errorf("The Dag creation should have failed when the Dag name is empty")
	}
}

func TestAddTask(t *testing.T) {
	dag, _ := CreateDag("my_dag")
	task, _ := task.CreateTask("bash", "my_task")
	taskUuid := task.GetUuid()

	initialDagSize := len(dag.tasks)

	dag.AddTask(task)

	if len(dag.tasks) != initialDagSize+1 {
		t.Errorf("The task was not added to the Dag")
	}

	addedTask := dag.GetTask(taskUuid)
	if addedTask.GetUuid() != task.GetUuid() {
		t.Errorf("The task was not added properly at the proper key in the Dag")
	}
}
