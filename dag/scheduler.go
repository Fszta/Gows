package dag

import (
	"fmt"

	"github.com/robfig/cron"
)

type DagScheduler struct {
	dag  *Dag
	cron *cron.Cron
}

func NewScheduler(dag *Dag, cronFormat string) *DagScheduler {

	c := cron.New()

	c.AddFunc(cronFormat, func() {
		func() {
			dag.RunDag()
			fmt.Println("finish", dag.name)
		}()
	})

	return &DagScheduler{
		dag:  dag,
		cron: c,
	}
}

func (s *DagScheduler) RunScheduler() {
	s.cron.Start()
	select {}
}

func (s *DagScheduler) Stop() {
	s.cron.Stop()
}
