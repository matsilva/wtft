package validators

import (
	"errors"
	"os"
)

// CLI validations
func CheckPrerequisites(file, outputType string) error {
	if !IsValidFile(file) {
		return errors.New("File does not exist")
	}
	if !IsValidOutputType(outputType) {
		return errors.New("Invalid output type")
	}
	return nil
}

// CLI argument validations
func IsValidOutputType(outputType string) bool {
	return outputType == "json" || outputType == "yaml"
}

//File validations

// IsValidFile checks if a file exists
func IsValidFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsNotDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}
