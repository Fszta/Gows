package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"

	"com.github/Fszta/gows/api"
	"com.github/Fszta/gows/pkg/dag"
	"github.com/spf13/cobra"
)

var (
	minWidth int  = 0
	tabWidth int  = 8
	padding  int  = 3
	padchar  byte = '\t'
	quiet    bool
)

var listDagsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List dags",
	Run: func(cmd *cobra.Command, args []string) {
		var dagsInfo []dag.DagInfo
		r, err := http.Get("http://localhost:2128" + api.ListDagsRoute)

		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &dagsInfo)

		if err != nil {
			fmt.Println(err)
		}

		writer := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padchar, tabwriter.AlignRight)

		if quiet {
			for _, info := range dagsInfo {
				fmt.Fprintln(writer, "ID")
				fmt.Fprintf(writer, "%v\n", info.UUID)
			}
			return
		}

		fmt.Fprintln(writer, "NAME\tSTATUS\tLAST-RUNTIME\tID")
		for _, info := range dagsInfo {
			fmt.Fprintf(writer, "%v\t%v\t%v\t%v\n", info.Name, info.Status, info.LastRunTime, info.UUID)
			writer.Flush()
		}
	},
}

func init() {
	listDagsCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "return only dag id when activate")
}
