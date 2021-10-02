package helpers

import (
	"fmt"
	"os"
)

func PrintError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func PrintErrorAndExit(err error) {
	PrintError(err)
	os.Exit(1)
}
