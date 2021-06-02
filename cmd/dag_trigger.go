package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var triggerDagCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Trigger dag execution",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Trigger")
	},
}
