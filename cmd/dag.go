package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var dagCmd = &cobra.Command{
	Use:   "dag",
	Short: "Manage dags",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var addDagCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new dag from json file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add dag:", time.Now().String())
		_, err := http.Get("http://0.0.0.0:2128/add")
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dagCmd)
	dagCmd.AddCommand(addDagCmd)
	dagCmd.AddCommand(removeDagCmd)
	dagCmd.AddCommand(stopDagCmd)
	dagCmd.AddCommand(triggerDagCmd)
	dagCmd.AddCommand(listDagsCmd)
}
