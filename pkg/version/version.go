package version

import (
	"runtime/debug"
)

var Version = "dev"

func String() string {
	// version injected with go releaser (-ldflags -X spike/pkg/version.Version={{.Version}})
	if Version != "dev" {
		return Version
	}

	// built using go toolchain (go install @x.x.x)
	if info, ok := debug.ReadBuildInfo(); ok {
		// (devel) is a magic string which indicates binary was built from a local dir
		// not documented in go currently: https://github.com/golang/go/issues/29228
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}
	// local build
	return Version
}
