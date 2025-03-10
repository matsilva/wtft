package main

import (
	"fmt"
	"os"

	"github.com/matsilva/wtft"
)

func main() {
	if err := wtft.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
