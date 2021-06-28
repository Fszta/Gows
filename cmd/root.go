package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "gows",
	Version: "0.1",
	Short:   "A golang workflow scheduler",
	Long:    `A golang workflow scheduler`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	
}
