package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"com.github/Fszta/gows/api"
	"github.com/spf13/cobra"
)

var (
	parentDagUUID string
	taskToLogName string
)

var logsTaskCmd = &cobra.Command{
	Use:   "logs",
	Short: "Retrieve logs for a specific task",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:2128" + api.GetTaskLogsRoute + "?uuid=" + parentDagUUID + "&name=" + taskToLogName)

		if err != nil {
			fmt.Println(err)
		}

		if response.StatusCode == http.StatusNotFound {
			fmt.Println("Dag or task not found")
		}

		logs, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(logs))
	},
}

func init() {
	logsTaskCmd.Flags().StringVarP(&parentDagUUID, "uuid", "u", "", "uuid of the dag")
	logsTaskCmd.MarkFlagRequired("uuid")
	logsTaskCmd.Flags().StringVarP(&taskToLogName, "task", "t", "", "name of the task")
	logsTaskCmd.MarkFlagRequired("task")
}
