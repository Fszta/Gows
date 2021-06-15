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
	operator1.SetCmd("sleep 1")
	task1, _ := task.CreateTask(operator1, "First Bash Task")
	dag1.AddTask(task1)

	operator2 := operators.CreateBashOperator()
	operator2.SetCmd("sleep 1")
	task2, _ := task.CreateTask(operator2, "Second Bash Task")
	dag1.AddTask(task2)

	dag1.SetScheduler("*/3 * * * * *")
	global.DagHandler.AddDag(dag1)
	global.DagHandler.Dags[dag1.GetUUID().String()].DagScheduler.RunScheduler()

	w.WriteHeader(http.StatusOK)
}

func ListDag(w http.ResponseWriter, req *http.Request) {
	dags := global.DagHandler.ListDag()
	dagsJson, err := json.Marshal(dags)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dagsJson)
}

func RemoveDag(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
	}

	err := global.DagHandler.RemoveDag(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func StopDag(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
	}
	err := global.DagHandler.StopDagScheduling(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func TriggerDag(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
	}
	err := global.DagHandler.TriggerDag(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

func RestartDag(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
	}
	err := global.DagHandler.StartDagScheduling(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetDagTasks(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
	}
	tasksInfo, err := global.DagHandler.GetDagTasks(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	tasksJson, err := json.Marshal(tasksInfo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tasksJson)
}
