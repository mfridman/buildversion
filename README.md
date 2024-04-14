# buildversion

A simple package to generate a release version for Go applications. Compatible with Go modules.

Ideal in CLI tools when you want to display the version using commands such as `mytool --version`.

## Usage

Here's a simple example of how you can use this package in your application. Now when `go build` or
`go install` is run, the version will be stamped into the binary in a consistent way.

```go
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
```

### Linking the version at build time

The reason `New()` takes a variadic `string` argument is to allow you to pass in the version at
build time. This can be useful when building a release binary with tools like
[goreleaser](https://goreleaser.com/).

Define a version variable in your main package:

```go
package main

var version string

func main() {
	buildversion.New(version)
}
```

Then, when building the binary, pass in the version using the `-ldflags` flag:

```
go build -ldflags "-X main.version=v1.2.3" -o bin/example ./cmd/example
```

## Example

### No tags (pseudo-version)

```
$ go install github.com/mfridman/buildversion/cmd/example@latest

example --version
v0.0.0-20240413170022-fe4dc7cb6b9d
```

### Tagged release

```
$ go install github.com/mfridman/buildversion/cmd/example@latests

example --version
v0.1.0
```

### Building from source

```
$ go build -o bin/example ./cmd/example

./bin/example --version
devel (fe4dc7cb6b9d, dirty)
```

## But why?

I've ended up copying this simple function across a few projects, so I decided to make it a package.

The `New()` function returns the version string from the
[BuildInfo](https://pkg.go.dev/runtime/debug#BuildInfo), if available.

**`New()` will always return a non-empty string.**

- If the build info is not available, it returns `devel`. This can happen if the binary was built
  without module support, if the Go version is too old or `-buildvcs=false` was set.

- If building from source, it returns `devel` followed by the first 12 characters of the VCS
  revision, followed by `, dirty` if the working directory was dirty. For example,

  - `devel (abcdef012345, dirty)`
  - `devel (abcdef012345)`
  - `devel (unknown revision)`

Note, VCS info not stamped when built listing .go files directly. For example,

```
go build main.go
go build .
```

For more information, see https://github.com/golang/go/issues/51279
