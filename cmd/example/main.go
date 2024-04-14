package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mfridman/buildversion"
)

func main() {
	versionPtr := flag.Bool("version", false, "")
	flag.Parse()
	if *versionPtr {
		fmt.Fprintln(os.Stdout, buildversion.New())
		return
	}
}
