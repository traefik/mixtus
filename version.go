package main

import (
	"fmt"
	"runtime"
)

var (
	version = "dev"
	date    = "I don't remember exactly"
)

// displayVersion displays the version.
func displayVersion() {
	fmt.Printf(`Lasius Mixtus:
 version     : %s
 build date  : %s
 go version  : %s
 go compiler : %s
 platform    : %s/%s
`, version, date, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
}
