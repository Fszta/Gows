package dag

import (
	"fmt"

	"github.com/google/uuid"
)

type DagHandler struct {
	Dags map[string]*Dag
}

type DagInfo struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	LastRunTime string `json:"lastRunTime"`
	Status      string `json:"status"`
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
	fmt.Println("INFO: Try to remove dag", dagUUID)
	if _, ok := dh.Dags[dagUUID]; ok {
		dh.StopDagScheduling(dagUUID)
		delete(dh.Dags, dagUUID)
		fmt.Println("INFO: Successfully remove dag", dagUUID)
		return nil
	}
	return fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) StopDagScheduling(dagUUID string) error {
	fmt.Println("INFO: Try to stop scheduling for dag", dagUUID)
	if _, ok := dh.Dags[dagUUID]; ok {
		dh.Dags[dagUUID].DagScheduler.stop()
		return nil
	}
	return fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) ListDag() []DagInfo {
	var dags []DagInfo
	for _, dag := range dh.Dags {
		dagInfo := DagInfo{
			Name:        dag.name,
			UUID:        dag.uuid.String(),
			LastRunTime: dag.lastRunTime.String(),
			Status:      dag.status,
		}
		dags = append(dags, dagInfo)
	}
	return dags
}

func (dh *DagHandler) TriggerDag(dagUUID uuid.UUID) {
	dh.Dags[dagUUID.String()].RunDag()
}
