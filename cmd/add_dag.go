// beginning of numbers.go
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var addDagCmd = &cobra.Command{
	Use:   "add-dag",
	Short: "Add a new dag from json file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add dag:", time.Now().String())
		_, err := http.Get("http://0.0.0.0:2128/add")
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addDagCmd)
}
