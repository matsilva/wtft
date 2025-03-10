package file_test

import (
	"fmt"
	"testing"

	"github.com/matsilva/wtft/lib/file"
)

// TestGetFileHeader tests the GetFileHeader function.
func TestGetFileHeader(t *testing.T) {
	//test case 1: given a valid jpg file, expect the first 3 bytes to be 0xff, 0xd8, 0xff
	jpgHeader := []byte{0xff, 0xd8, 0xff}
	header, err := file.GetFileHeader("../../testdata/chippy.jpg")
	if err != nil {
		t.Errorf("Error getting file header: %v", err)
	}
	for i, b := range header[:3] {
		if b != jpgHeader[i] {
			t.Error("Expected", jpgHeader[i], "got", b)
		}
	}
	//test case 2: given an invalid file, expect an error
	_, err = file.GetFileHeader("../../testdata/invalid.txt")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestGetFileSignature(t *testing.T) {
	//test case 1: given a valid jpg file header, expect the file signature to be jpg
	jpgHeader := []byte{0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10, 0x4a, 0x46, 0x49, 0x46, 0x00, 0x01, 0x01, 0x01, 0x00, 0x48, 0x00, 0x48, 0x00, 0x00, 0xff, 0xdb, 0x00, 0x43}
	sig, err := file.GetFileSignature(jpgHeader)
	if err != nil {
		t.Error(err)
	}
	if sig.Ext != "jpg" {
		t.Error("Expected jpg, got", sig.Ext)
	}
	//test case 2: given an invalid file header, expect an error
	invalidHeader := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	sig, err = file.GetFileSignature(invalidHeader)
	fmt.Println(sig)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
