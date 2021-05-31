package api

import (
	"encoding/json"
	"fmt"
	"gows/dag"
	"gows/global"
	"gows/operators"
	"gows/task"
	"log"
	"net/http"
	"time"
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

func Start() {
	http.HandleFunc("/add", AddDag)
	http.HandleFunc("/list", ListDag)
	fmt.Println("Server started at port 2128")
	log.Fatal(http.ListenAndServe(":2128", nil))
}
