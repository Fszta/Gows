package dag

import (
	"errors"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

type DagScheduler struct {
	dag         *Dag
	cron        *cron.Cron
	isScheduled bool
}

func NewScheduler(dag *Dag, cronFormat string) *DagScheduler {

	c := cron.New()

	c.AddFunc(cronFormat, func() {
		fmt.Printf("Start dag %v at %v \n", dag.name, time.Now())
		dag.Run()
	})

	return &DagScheduler{
		dag:         dag,
		cron:        c,
		isScheduled: false,
	}
}

func (s *DagScheduler) RunScheduler() error {
	fmt.Println("Start scheduler ", time.Now())
	if s == nil {
		return errors.New("Scheduler is not properly set")
	}
	s.isScheduled = true
	s.cron.Start()
	return nil
}

func (s *DagScheduler) stop() {
	fmt.Println("Stop dag at ", time.Now())
	s.isScheduled = false
	s.cron.Stop()
}
