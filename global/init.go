package global

import (
	"fmt"
	"log"

	"com.github/Fszta/gows/database"
	"com.github/Fszta/gows/pkg/dag"
)

var DagHandler = &dag.DagHandler{}

func RetrieveDag() {
	client, err := database.NewClient()

	if err != nil {
		log.Fatal("Fail to connect to database")
	}

	defer client.Close()

	fmt.Println("Database connection init")

	for _, dagConfig := range client.GetAllDags() {
		dag, err := dag.GetDagFromConfig(&dagConfig)

		if err != nil {
			fmt.Println(err)
		}

		DagHandler.AddDag(dag)
		dag.DagScheduler.RunScheduler()
	}
}
