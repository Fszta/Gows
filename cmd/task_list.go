package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"com.github/Fszta/gows/api"
	"com.github/Fszta/gows/pkg/dag"
	"github.com/spf13/cobra"
)

var dagToListUUID string

var listTaskCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tasks in a dag",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:2128" + api.GetDagTasksRoute + "?uuid=" + dagToListUUID)

		if err != nil {
			log.Fatalln(err)
		}

		if response.StatusCode == http.StatusNotFound {
			fmt.Printf("Dag %v not found\n", dagToListUUID)
		}

		if response.StatusCode == http.StatusOK {
			fmt.Printf("Dag's %v tasks successfully retrieved \n", dagToListUUID)
			var tasksInfo []dag.TaskInfo

			body, err := ioutil.ReadAll(response.Body)
			json.Unmarshal(body, &tasksInfo)

			if err != nil {
				fmt.Println(err)
			}
			writer := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padchar, tabwriter.AlignRight)

			fmt.Fprintln(writer, "NAME\tSTATUS\tUUID")
			for _, info := range tasksInfo {
				fmt.Fprintf(writer, "%v\t%v\t%v\n", info.Name, info.Status, info.UUID)
				writer.Flush()
			}
		}
	},
}

func init() {
	listTaskCmd.Flags().StringVarP(&dagToListUUID, "uuid", "u", "", "uuid of the dag")
	listTaskCmd.MarkFlagRequired("uuid")
}
