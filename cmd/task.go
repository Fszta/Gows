package cmd

import (
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.AddCommand(listTaskCmd)
	taskCmd.AddCommand(logsTaskCmd)
}
