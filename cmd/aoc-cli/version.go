package main

import "fmt"

const (
	unset = "unset"
)

var ( // build info
	version   = unset
	date      = unset
	commit    = unset
	goversion = unset
)

func printVersion() {
	fmt.Printf("Go version: %s \n", goversion)
	fmt.Printf("Version info: %s \n", version)
	fmt.Printf("Build date: %s \n", date)
	fmt.Printf("commit: %s \n", commit)

	fmt.Println()
}
