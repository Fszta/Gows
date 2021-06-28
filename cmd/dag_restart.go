package cmd

import (
	"fmt"
	"log"
	"net/http"

	"com.github/Fszta/gows/api"
	"github.com/spf13/cobra"
)

var dagToRestartUUID string

var restartDagCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart dag scheduling",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:2128" + api.RestartDagRoute + "?uuid=" + dagToRestartUUID)

		if err != nil {
			log.Fatalln(err)
		}

		if response.StatusCode == http.StatusOK {
			fmt.Printf("Dag %v has been successfully restarted\n", dagToRestartUUID)
		}

		if response.StatusCode == http.StatusNotFound {
			fmt.Printf("Dag %v not found\n", dagToRestartUUID)
		}
	},
}

func init() {
	restartDagCmd.Flags().StringVarP(&dagToRestartUUID, "uuid", "u", "", "uuid of the dag")
	restartDagCmd.MarkFlagRequired("uuid")
}
