# buildversion

Very simple package to generate an opinionated release version string for Go applications, such as
CLI tools.

Works with Go modules.

```
go install github.com/mfridman/buildversion/cmd/example@latest
example --version
v1.2.3
```

It is intended to be used in CLI tools where you want to display the version string with something
like `mytool version` or `mytool --version`.

## But why?

I've ended up copying this code across a few projects, so I decided to make it a package.

The `New` function returns the version string from the
[BuildInfo](https://pkg.go.dev/runtime/debug#BuildInfo), if available. **It will always return a
non-empty string.**

- If the version is explicitly set, it returns the string as is. Useful for setting the version at
  build time. For example, `-ldflags "-X 'main.version=1.2.3'"` and just passing the version string
  to this function.

- If the build info is not available, it returns `devel`. This can happen if the binary was built
  without module support, if the Go version is too old or `-buildvcs=false` was set.

- If building from source, it returns `devel` followed by the first 12 characters of the VCS
  revision, followed by `, dirty` if the working directory was dirty. For example,

  - `devel (abcdef012345, dirty)`
  - `devel (abcdef012345)`
  - `devel (unknown revision)`

Note, vcs info not stamped when built listing .go files directly. For example,

```
go build main.go
go build .
```

For more information, see https://github.com/golang/go/issues/51279
