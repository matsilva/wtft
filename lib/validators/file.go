package validators

import "os"

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
