package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mfridman/buildversion"
)

var version string

func main() {
	versionPtr := flag.Bool("version", false, "")
	flag.Parse()
	if *versionPtr {
		fmt.Fprintln(os.Stdout, buildversion.New(version))
		return
	}
}
