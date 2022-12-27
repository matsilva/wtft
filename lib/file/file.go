package file

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var HeaderLength = 24

// GetFileHeader returns the first 24 bytes of a file
func GetFileHeader(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buff, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return buff[:HeaderLength], nil
}

// GetFileSignature returns the signature of a file for a given file header
func GetFileSignature(header []byte) (*Signature, error) {
	for _, sig := range Signatures {
		match := true
		for i, b := range sig.HeaderSeq {
			fmt.Println(b, header[i], b == header[i])
			if b != header[i] {
				match = false
				break
			}
		}
		if match {
			return &sig, nil
		}
	}
	return nil, errors.New("No file type could be found for file. File may be corrupt or signature does not exist in wtft codebase yet...")
}
