package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logsTaskCmd = &cobra.Command{
	Use:   "logs",
	Short: "Retrieve logs for a specific task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Logs")
	},
}
