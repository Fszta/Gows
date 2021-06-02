package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeDagCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a dag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove dag")
	},
}
