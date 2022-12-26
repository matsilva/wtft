package main

import (
	"fmt"
	"os"

	"github.com/matsilva/wtft/cmd"
)

func main() {
	// defer os.Exit(0)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
