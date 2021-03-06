package dag

import (
	"errors"
	"fmt"
)

var timeFormat = "2006-01-02 15:04:05"

type DagHandler struct {
	Dags map[string]*Dag
}

type DagInfo struct {
	Name            string `json:"name"`
	UUID            string `json:"uuid"`
	LastRunTime     string `json:"lastRunTime"`
	Status          string `json:"status"`
	SchedulerFormat string `json:"scheduler"`
}

type TaskInfo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	UUID   string `json:"uuid"`
}

func NewHandler() *DagHandler {
	return &DagHandler{
		Dags: make(map[string]*Dag),
	}
}

func (dh *DagHandler) AddDag(dag *Dag) {
	dh.Dags[dag.uuid.String()] = dag
}

func (dh *DagHandler) RemoveDag(dagUUID string) error {
	fmt.Println("Try to remove dag", dagUUID)
	if _, ok := dh.Dags[dagUUID]; ok {
		dh.StopDagScheduling(dagUUID)
		delete(dh.Dags, dagUUID)
		fmt.Println("Successfully remove dag", dagUUID)
		return nil
	}
	return fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) StopDagScheduling(dagUUID string) error {
	fmt.Println("Try to stop scheduling for dag", dagUUID)
	if _, ok := dh.Dags[dagUUID]; ok {
		dh.Dags[dagUUID].DagScheduler.stop()
		return nil
	}
	return fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) StartDagScheduling(dagUUID string) error {
	fmt.Println("Try to start scheduling for dag", dagUUID)
	if _, ok := dh.Dags[dagUUID]; ok {
		dh.Dags[dagUUID].DagScheduler.RunScheduler()
		return nil
	}
	return fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) ListDag() []DagInfo {
	var dags []DagInfo
	for _, dag := range dh.Dags {
		dagInfo := DagInfo{
			Name:            dag.name,
			UUID:            dag.uuid.String(),
			LastRunTime:     dag.lastRunTime.Format(timeFormat),
			Status:          dag.status,
			SchedulerFormat: dag.GetSchedulerFormat(),
		}
		dags = append(dags, dagInfo)
	}
	return dags
}

func (dh *DagHandler) TriggerDag(dagUUID string) error {
	if _, ok := dh.Dags[dagUUID]; ok {
		fmt.Println("Trigger dag", dagUUID)

		if dh.Dags[dagUUID].DagScheduler.isScheduled {
			fmt.Printf("Dag %v has been triggered while already scheduled \n", dagUUID)
		}

		dh.Dags[dagUUID].Run()
		return nil
	}
	return fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) GetDagTasks(dagUUID string) ([]TaskInfo, error) {
	if _, ok := dh.Dags[dagUUID]; ok {
		fmt.Println("Retrieve tasks list from dag", dagUUID)

		tasks := dh.Dags[dagUUID].GetAllTasks()

		var tasksInfo []TaskInfo

		for _, task := range tasks {
			taskInfo := TaskInfo{
				Name:   task.GetName(),
				Status: task.GetStatus(),
				UUID:   task.GetUuid().String(),
			}
			tasksInfo = append(tasksInfo, taskInfo)
		}

		return tasksInfo, nil
	}
	return nil, fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) GetTaskLogs(dagUUID string, taskName string) (string, error) {
	if dag, ok := dh.Dags[dagUUID]; ok {
		task, err := dag.GetTaskByName(taskName)
		if err != nil {
			return "", nil
		}

		return task.GetLogs(), nil
	}
	return "", errors.New(fmt.Sprintf("Dag %s not found", dagUUID))
}

func (d *Dag) GetTaskLogs(name string) (string, error) {
	for _, dagTask := range d.tasks {
		task := dagTask.task
		if task.GetName() == name {
			return task.GetLogs(), nil
		}

	}
	return "", errors.New("Fail to find task")
}
