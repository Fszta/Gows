package dag

import (
	"errors"
	"fmt"
	"time"

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
			fmt.Println("INFO: Start dag", dag.name, "at", time.Now())
			dag.RunDag()
		}()
	})

	return &DagScheduler{
		dag:  dag,
		cron: c,
	}
}

func (s *DagScheduler) RunScheduler() error {
	fmt.Println("Start scheduler ", time.Now())
	if s == nil {
		return errors.New("ERROR: scheduler is not properly set")
	}
	s.cron.Start()
	select {}
}

func (s *DagScheduler) Stop() {
	s.cron.Stop()
}
