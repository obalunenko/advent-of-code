// Package version provides utilities for retrieving application build information
// such as version number, build date, commit hash, application name, and the Go
// language version used during the build process.
//
// This package leverages the `runtime/debug` package to read build metadata
// embedded in the binary during the build process with Go modules.
//
// Build information can be accessed via the exposed methods:
//
//   - GetGoVersion: Returns the Go version used to build the app.
//   - GetVersion: Returns the application's version.
//   - GetBuildDate: Returns the date the application was built.
//   - GetCommit: Returns the full commit hash the application was built from.
//   - GetShortCommit: Returns a shortened commit hash (7 characters).
//   - GetAppName: Returns the application's module path.
//
// When the application is built from sources with uncommitted changes, the
// commit information will be suffixed with "+CHANGES".
package version

import (
	"runtime/debug"
)

const unset = "unset"

var ( // build info
	version     = unset
	builddate   = unset
	commit      = unset
	shortcommit = unset
	appname     = unset
	goversion   = unset
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	bi := parseBuildInfo(info)

	version = bi.version
	builddate = bi.builddate
	commit = bi.commit
	shortcommit = bi.shortcommit
	appname = bi.appname
	goversion = bi.goversion
}

type buildInfo struct {
	version     string
	builddate   string
	commit      string
	shortcommit string
	appname     string
	goversion   string
}

// parseBuildInfo extracts build information from info.
// Values missing in info stay unset.
func parseBuildInfo(info *debug.BuildInfo) buildInfo {
	bi := buildInfo{
		version:     info.Main.Version,
		builddate:   unset,
		commit:      unset,
		shortcommit: unset,
		appname:     info.Path,
		goversion:   info.GoVersion,
	}

	var modified bool

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			bi.commit = setting.Value
		case "vcs.time":
			bi.builddate = setting.Value
		case "vcs.modified":
			// Go writes this key for every VCS build: "true" or "false".
			modified = setting.Value == "true"
		}
	}

	if len(bi.commit) < 7 {
		bi.shortcommit = bi.commit
	} else {
		bi.shortcommit = bi.commit[:7]
	}

	if modified {
		bi.commit += "+CHANGES"
		bi.shortcommit += "+CHANGES"
	}

	return bi
}

// GetGoVersion returns the go version
func GetGoVersion() string {
	return goversion
}

// GetVersion returns the app version
func GetVersion() string {
	return version
}

// GetBuildDate returns the build date
func GetBuildDate() string {
	return builddate
}

// GetCommit returns the git commit
func GetCommit() string {
	return commit
}

// GetAppName returns the app name
func GetAppName() string {
	return appname
}

// GetShortCommit returns the short git commit
func GetShortCommit() string {
	return shortcommit
}
