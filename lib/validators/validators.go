package validators

import (
	"errors"
	"os"
)

// CLI validations

// CheckPrerequisites checks if the CLI arguments are valid
func CheckPrerequisites(file, outputType string) error {
	if file == "" {
		return errors.New("File is required")
	}
	if !IsValidFile(file) {
		return errors.New("File does not exist")
	}
	if !IsValidOutputType(outputType) {
		return errors.New("Invalid output type")
	}
	return nil
}

// IsValidOutputType checks if the output type is valid
func IsValidOutputType(outputType string) bool {
	return outputType == "json" || outputType == "yaml"
}

//File validations

// IsValidFile checks if a file exists
func IsValidFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsNotDirectory checks if a path is not a directory
func IsNotDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}
