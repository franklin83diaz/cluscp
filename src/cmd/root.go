package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cluscp",
	Short: "cluscp is a tool copy files to cluster efficiently and fast",
	Long: `cluscp is a tool for efficiently and quickly copying files to a cluster
	 using cascade copying. It is capable of providing similar copy speeds 
	for copying files to 1 node as to 1000 nodes, without the need to use 
	multicast or broadcast.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			// execute help
			cmd.Help()
			return
		}

		//Copy files to cluster
		if len(args) == 2 {
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
