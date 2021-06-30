package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"com.github/Fszta/gows/api"
	"github.com/spf13/cobra"
)

var file string

var addDagCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new dag from json file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Try to add dag from file %v at %v \n", file, time.Now().String())

		response, err := http.Get("http://localhost:2128" + api.AddDagRoute + "?file=" + file)
		if err != nil {
			log.Fatalln(err)
		}

		responseMessage, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		if response.StatusCode == http.StatusOK {
			fmt.Printf("Dag successfully created from %v \n", file)
		}

		if response.StatusCode == http.StatusNotFound {
			fmt.Printf("Fail to find file %v \n", file)
		}

		if response.StatusCode == http.StatusNotAcceptable {
			fmt.Printf("File %v is not properly structured, %v \n", file, string(responseMessage))
		}

		if response.StatusCode == http.StatusBadRequest {
			fmt.Printf("Dag name already exists in database")
		}
	},
}

func init() {
	addDagCmd.Flags().StringVarP(&file, "file", "f", "", "path of the json file describing dag")
	addDagCmd.MarkFlagRequired("file")
}
