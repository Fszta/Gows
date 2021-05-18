package dag

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/robfig/cron"
)

type DagHandler struct {
	dags          map[uuid.UUID]*Dag
	scheduledDags []uuid.UUID
	numberOfDag   int
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

func (dh *DagHandler) ScheduleDag(cronFormat string, dag *Dag) *cron.Cron {
	dh.scheduledDags = append(dh.scheduledDags, dag.uuid)

	c := cron.New() // Create a new timed task object

	c.AddFunc(cronFormat, func() {
		func() {
			dag.RunDag()
			fmt.Println("finish", dag.name)
		}()

	})
	c.Start()
	select {}
}

func (dh *DagHandler) StopDagScheduling(dagUUID uuid.UUID) {}

func (dh *DagHandler) ListDag() map[uuid.UUID]*Dag {
	return dh.dags
}

func (dh *DagHandler) TriggerDag(dagUUID uuid.UUID) {
	dh.dags[dagUUID].RunDag()
}
