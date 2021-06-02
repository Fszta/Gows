package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"

	"com.github/Fszta/gows/global"
	"com.github/Fszta/gows/pkg/dag"
	"github.com/spf13/cobra"
)

var listDagsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List dags",
	Run: func(cmd *cobra.Command, args []string) {
		global.DagHandler.ListDag()
		var dagsInfo []dag.DagInfo
		r, err := http.Get("http://0.0.0.0:2128/list")

		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &dagsInfo)

		if err != nil {
			fmt.Println(err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(writer, "NAME\tSTATUS\tLAST-RUNTIME\tID")
		for _, info := range dagsInfo {
			fmt.Fprintf(writer, "%v\t%v\t%v\t%v\n", info.Name, info.Status, info.LastRunTime, info.UUID)
			writer.Flush()
		}
	},
}
