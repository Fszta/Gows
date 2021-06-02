package cmd

import (
	"github.com/spf13/cobra"
)

var dagCmd = &cobra.Command{
	Use:   "dag",
	Short: "Manage dags",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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
