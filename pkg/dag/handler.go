package dag

import (
	"fmt"

	"github.com/google/uuid"
)

type DagHandler struct {
	dags map[string]*Dag
}

type DagInfo struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	LastRunTime string `json:"lastRunTime"`
	Status      string `json:"status"`
}

func NewHandler() *DagHandler {
	return &DagHandler{
		dags: make(map[string]*Dag),
	}
}

func (dh *DagHandler) AddDag(dag *Dag) {
	dh.dags[dag.uuid.String()] = dag
}

func (dh *DagHandler) RemoveDag(dagUUID string) error {
	if _, ok := dh.dags[dagUUID]; ok {
		delete(dh.dags, dagUUID)
		return nil
	}
	return fmt.Errorf("Dag %v not found", dagUUID)
}

func (dh *DagHandler) ScheduleDag(cronFormat string, dag *Dag) {
}

func (dh *DagHandler) StopDagScheduling(dagUUID uuid.UUID) {}

func (dh *DagHandler) ListDag() []DagInfo {
	var dags []DagInfo
	for _, dag := range dh.dags {
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
	dh.dags[dagUUID.String()].RunDag()
}
