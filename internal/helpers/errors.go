package helpers

import (
	"fmt"
	"os"
)

// ExitWithError prints an error and exits the program
func ExitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
