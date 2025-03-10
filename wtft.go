package wtft

import (
	"fmt"

	"github.com/matsilva/wtft/internal/file"
	"github.com/matsilva/wtft/internal/helpers"
	"github.com/matsilva/wtft/internal/validators"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wtft",
	Short: "wtft is a CLI for determing what a file's real type is.",
	Long: `
	wtft is a simple command line tool that helps you figure out what the real file type is.
	It does this by looking at the file's magic number and comparing it to a list of known magic numbers.`,
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
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "The path to file to determine the real type of.")
	rootCmd.PersistentFlags().StringVar(&outputType, "output-type", "json", "The type of output to return. Valid values are json, text and yaml.")
}
