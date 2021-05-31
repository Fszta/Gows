// beginning of numbers.go
package cmd

import (
	"gows/api"
	"gows/dag"
	"gows/global"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start gows daemon",
	Run: func(cmd *cobra.Command, args []string) {
		global.DagHandler = dag.NewHandler()
		api.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
