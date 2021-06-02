package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listTaskCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tasks in a dag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List task in dag")
	},
}
