package dag

import (
	"fmt"

	"github.com/google/uuid"
)

type DagHandler struct {
	dags map[uuid.UUID]*Dag
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

func (dh *DagHandler) ListDag() map[uuid.UUID]*Dag {
	for dag := range dh.dags {
		fmt.Println(dag)
	}
	return dh.dags
}

func (dh *DagHandler) TriggerDag(dagUUID uuid.UUID) {
	dh.dags[dagUUID].RunDag()
}
