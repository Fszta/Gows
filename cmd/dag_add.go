package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	name string
	file string
)

var addDagCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new dag from json file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Add dag: %v based on file %v at %v \n", name, file, time.Now().String())
		_, err := http.Get("http://localhost:2128/add")
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	addDagCmd.Flags().StringVarP(&name, "name", "n", "", "name of the dag")
	addDagCmd.MarkFlagRequired("name")
	addDagCmd.Flags().StringVarP(&file, "file", "f", "", "path of the json file describing dag")
	addDagCmd.MarkFlagRequired("file")
}
