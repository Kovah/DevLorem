package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "devlorem",
	Short: "DevLorem is your Lorem Ipsum generator made from the best movie quotes.",
	Long: `Real quotes ready to copy and paste. No more "Lorem ipsum dolor". Get some quotes from President Obama,
Samuel L Jackson, Daisy Ridley or Morgan Freeman!`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func CmdExecute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
