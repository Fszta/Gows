package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"com.github/Fszta/gows/global"
	"com.github/Fszta/gows/pkg/dag"
	"com.github/Fszta/gows/pkg/operators"
	"com.github/Fszta/gows/pkg/task"
)

func AddDag(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Add dag:", time.Now().String())
	dag1, _ := dag.CreateDag("my_dag1")
	operator1 := operators.CreateBashOperator()
	operator1.SetCmd("ps")
	task1, _ := task.CreateTask(operator1, "First Bash Task")
	dag1.AddTask(task1)
	dag1.SetScheduler("*/3 * * * * *")

	global.DagHandler.AddDag(dag1)
}

func ListDag(w http.ResponseWriter, req *http.Request) {
	dags := global.DagHandler.ListDag()
	dagsJson, err := json.Marshal(dags)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dagsJson)
}
