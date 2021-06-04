package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var dagUUID string

var removeDagCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a dag",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:2128/remove?uuid=" + dagUUID)
		if err != nil {
			log.Fatalln(err)
		}
		if response.StatusCode == http.StatusAccepted {
			fmt.Printf("Dag %v has been successfully removed\n", dagUUID)
		}

		if response.StatusCode == http.StatusNotFound {
			fmt.Printf("Dag %v not found\n", dagUUID)
		}

	},
}

func init() {
	removeDagCmd.Flags().StringVarP(&dagUUID, "uuid", "u", "", "uuid of the dag")
	removeDagCmd.MarkFlagRequired("uuid")
}
