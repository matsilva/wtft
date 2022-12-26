package cmd

import (
	"fmt"

	"github.com/matsilva/wtft/lib/file"
	"github.com/matsilva/wtft/lib/helpers"
	"github.com/matsilva/wtft/lib/validators"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wtft",
	Short: "wtft is a CLI for determing what a file's real type is.",
	Long:  `wtft is a CLI for determing what a file's real type is.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := validators.CheckPrerequisites(filePath, outputType)
		if err != nil {
			helpers.ExitWithError(err)
		}

		header, err := file.GetFileHeader(filePath)
		if err != nil {
			helpers.ExitWithError(err)
		}

		sig, err := file.GetFileSignature(header)
		if err != nil {
			helpers.ExitWithError(err)
		}

		switch outputType {
		case "json":
			out, err := helpers.ToJSON(sig)
			if err != nil {
				helpers.ExitWithError(err)
			}
			fmt.Println(string(out))
		case "yaml":
			out, err := helpers.ToYAML(sig)
			if err != nil {
				helpers.ExitWithError(err)
			}
			fmt.Println(string(out))
		}
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var (
	filePath   string
	outputType string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&filePath, "file", "", "The path to file to determine the real type of.")
	rootCmd.PersistentFlags().StringVar(&outputType, "output-type", "json", "The type of output to return. Valid values are json, text and yaml.")
}
