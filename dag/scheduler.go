package dag

import (
	"fmt"

	"github.com/robfig/cron"
)

func ScheduleDag(cronFormat string, dag *Dag) *cron.Cron {
	c := cron.New() // Create a new timed task object

	c.AddFunc("*/3 * * * * *", func() {
		func() {
			dag.RunDag()
			fmt.Println("finish", dag.name)
		}()

	})
	c.Start()
	select {}
}
