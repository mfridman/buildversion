package buildversion

import (
	"runtime/debug"
	"strings"
	"time"
)

// New returns a version string from the build info, if available. It will always return a non-empty
// string, see below for details.
//
// A version string can be provided as an argument. If provided, only the first argument is used and
// it is returned as is. This is useful for setting the version string at build time with the
// -ldflags flag. For example:
//
//	go build -ldflags "-X main.version=1.2.3" ./cmd/example
//
// The version string is constructed as follows:
//
//   - If the build info is not available, it returns "devel". This can happen if the binary was
//     built without module support, if the Go version is too old or -buildvcs=false was set.
//
//   - If building from source, it returns "devel" followed by the first 12 characters of the VCS
//     revision, followed by ", dirty" if the working directory was dirty. For example,
//
//     "devel (abcdef012345, dirty)"
//     "devel (abcdef012345)"
//     "devel (unknown revision)"
//
// Note, vcs info not stamped when built listing .go files directly. For example,
//   - `go build main.go`
//   - `go build .`
//
// For more information, see https://github.com/golang/go/issues/51279
func New(version ...string) string {
	if len(version) > 0 && version[0] != "" {
		return strings.TrimSpace(version[0])
	}
	const defaultVersion = "devel"

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		// Should only happen if -buildvcs=false is set or using a really old version of Go.
		return defaultVersion
	}
	// The (devel) string is not documented, but it is the value used by the Go toolchain. See
	// https://github.com/golang/go/issues/29228
	if s := buildInfo.Main.Version; s != "" && s != "(devel)" {
		return buildInfo.Main.Version
	}
	var vcs struct {
		revision string
		time     time.Time
		modified bool
	}
	for _, setting := range buildInfo.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcs.revision = setting.Value
		case "vcs.time":
			vcs.time, _ = time.Parse(time.RFC3339, setting.Value)
		case "vcs.modified":
			vcs.modified = (setting.Value == "true")
		}
	}

	var b strings.Builder
	b.WriteString(defaultVersion)
	b.WriteString(" (")
	if vcs.revision == "" || len(vcs.revision) < 12 {
		b.WriteString("unknown revision")
	} else {
		b.WriteString(vcs.revision[:12])
	}
	if vcs.modified {
		b.WriteString(", dirty")
	}
	b.WriteString(")")
	return b.String()
}
