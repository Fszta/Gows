package dag

import (
	"github.com/google/uuid"
)

type DagHandler struct {
	dags map[uuid.UUID]*Dag
}

type DagInfo struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	LastRunTime string `json:"lastRunTime"`
	Status      string `json:"status"`
}

func NewHandler() *DagHandler {
	return &DagHandler{
		dags: make(map[uuid.UUID]*Dag),
	}
}

func (dh *DagHandler) AddDag(dag *Dag) {
	dh.dags[dag.uuid] = dag
}

func (dh *DagHandler) RemoveDag(dagUUID uuid.UUID) {
	delete(dh.dags, dagUUID)
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
	dh.dags[dagUUID].RunDag()
}
