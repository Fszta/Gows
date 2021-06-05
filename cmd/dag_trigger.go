package cmd

import (
	"fmt"
	"log"
	"net/http"

	"com.github/Fszta/gows/api"
	"github.com/spf13/cobra"
)

var dagToTriggerUUID string

var triggerDagCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Trigger dag execution",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:2128" + api.TriggerDagRoute + "?uuid=" + dagToTriggerUUID)

		if err != nil {
			log.Fatalln(err)
		}

		if response.StatusCode == http.StatusOK {
			fmt.Printf("Dag %v has been successfully triggered\n", dagToTriggerUUID)
		}

		if response.StatusCode == http.StatusNotFound {
			fmt.Printf("Dag %v not found\n", dagToTriggerUUID)
		}
	},
}

func init() {
	triggerDagCmd.Flags().StringVarP(&dagToTriggerUUID, "uuid", "u", "", "uuid of the dag")
	triggerDagCmd.MarkFlagRequired("uuid")
}
