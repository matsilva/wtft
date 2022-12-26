package file

import (
	"errors"
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

func GetFileSignature(header []byte) (*Signature, error) {
	for _, sig := range Signatures {
		for i, b := range sig.HeaderSeq {
			if b != header[i] {
				break
			}
			return &sig, nil
		}
	}
	return nil, errors.New("No file type could be found for file. File may be corrupt.")
}
