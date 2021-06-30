package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"com.github/Fszta/gows/database"
	"com.github/Fszta/gows/global"
	"com.github/Fszta/gows/pkg/dag"
)

func AddDag(w http.ResponseWriter, req *http.Request) {
	dagConfigFile := req.FormValue("file")

	fmt.Println(dagConfigFile)
	if dagConfigFile == "" {
		http.Error(w, "Missing file parameter", http.StatusBadRequest)
	}

	// Read dag configuration from json file
	dagConfigJson, err := dag.ReadDagConfig(dagConfigFile)

	if err != nil {
		http.Error(w, fmt.Sprintf("File %v not found", dagConfigFile), http.StatusNotFound)
		return
	}

	config, err := dag.UnmarshalDagConfig(dagConfigJson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create Dag from configuration
	dag, err := dag.GetDagFromConfig(config)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	// Try to add dag to database
	client, err := database.NewClient()
	if err != nil {
		http.Error(w, "Fail to open connect to database", http.StatusInternalServerError)
	}
	defer client.Close()

	dagUUID := dag.GetUUID().String()

	// Store Dag to database
	err = client.SaveDagConfig(config)
	if err != nil {
		fmt.Println(err)
	}

	global.DagHandler.AddDag(dag)
	global.DagHandler.Dags[dagUUID].DagScheduler.RunScheduler()

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

	// Init connection to database
	client, err := database.NewClient()
	if err != nil {
		http.Error(w, "Fail to open connect to database", http.StatusInternalServerError)
	}
	defer client.Close()

	// Remove dag from database
	client.RemoveDag(uuid)

	err = global.DagHandler.RemoveDag(uuid)
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

func GetTaskLogs(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
	}

	taskName := req.FormValue("name")
	if taskName == "" {
		http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
	}

	logs, err := global.DagHandler.GetTaskLogs(uuid, taskName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte(logs))
}
