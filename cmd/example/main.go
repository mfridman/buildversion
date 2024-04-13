package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mfridman/buildversion"
)

var (
	versionPtr = flag.Bool("version", false, "")

	// These may be set at build time with -ldflags "-X 'main.version=1.2.3'"
	version string
)

func main() {
	flag.Parse()
	if *versionPtr {
		fmt.Fprintf(os.Stdout, "%s\n", buildversion.New(version))
		return
	}
}
