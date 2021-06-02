package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopDagCmd = &cobra.Command{
	Use:   "stop",
	Short: "Suspend dag scheduling",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stop dag scheduling")
	},
}
