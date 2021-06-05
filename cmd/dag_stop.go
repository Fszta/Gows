package cmd

import (
	"fmt"
	"log"
	"net/http"

	"com.github/Fszta/gows/api"
	"github.com/spf13/cobra"
)

var dagToStopUUID string

var stopDagCmd = &cobra.Command{
	Use:   "stop",
	Short: "Suspend dag scheduling",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:2128" + api.StopDagRoute + "?uuid=" + dagToStopUUID)

		if err != nil {
			log.Fatalln(err)
		}

		if response.StatusCode == http.StatusOK {
			fmt.Printf("Dag %v has been successfully stopped\n", dagToStopUUID)
		}

		if response.StatusCode == http.StatusNotFound {
			fmt.Printf("Dag %v not found\n", dagToStopUUID)
		}
	},
}

func init() {
	stopDagCmd.Flags().StringVarP(&dagToStopUUID, "uuid", "u", "", "uuid of the dag")
	stopDagCmd.MarkFlagRequired("uuid")
}
