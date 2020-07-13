package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var format []string
var numParagraphs int
var addParagraphs bool

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringArrayVarP(&format, "format", "f", []string{"json", "text"}, "Number of paragraphs returned")
	cliCmd.Flags().IntVarP(&numParagraphs, "number", "n", 5, "Number of paragraphs returned")
	cliCmd.Flags().BoolVarP(&addParagraphs, "paragraphs", "p", false, "Show paragraphs in the generated paragraphs")
}

var cliCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate quotes directly in the CLI.",
	Long:  `Generate quotes directly in the CLI. You can specify the number of paragraphs, if paragraph tags should be included, as well as the format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleCliGenerator(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func handleCliGenerator() error {
	source := GetNumLines(numParagraphs, addParagraphs)

	if format[0] == "text" {
		fmt.Println(source.Source)
		for _, paragraph := range source.Paragraphs {
			fmt.Println(paragraph)
		}

		return nil
	}

	writer := json.NewEncoder(log.Writer())
	writer.SetEscapeHTML(false)
	err := writer.Encode(source)

	return err
}
