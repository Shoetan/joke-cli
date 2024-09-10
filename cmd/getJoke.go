package cmd

import (
	"github.com/Shoetan/joke-cli/jokeapi"
	"github.com/spf13/cobra"
)



var jokeCMD = &cobra.Command{
	Use: "Tell me a joke",
	Short: "Getting joke from jokeapi",
	Aliases: []string{"joke"},
	Long:  "Something to make you giggle",
	Run: func(cmd *cobra.Command, args []string) {
		jokeapi.GetJoke()
	},
}

func init(){
	rootCMD.AddCommand(jokeCMD)
}