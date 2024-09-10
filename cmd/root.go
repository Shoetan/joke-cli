package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use: "Joke",
	Short: "A simple CLI to tell you a joke when you are bored",
	Long:  "Not super fancy at all",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome! Get jokes delivered to your terminal whenever you are bored")
	},
}


func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing your command %s", err)
		os.Exit(1)
	}
}