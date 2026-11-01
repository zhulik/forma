package version

import (
	"fmt"
	"runtime/debug"
	"time"
)

const (
	devel   = "(devel)"
	unknown = "unknown"
	shaLen  = 7
)

// String returns the CLI version for --version and the version subcommand.
func String() string {
	info, ok := debug.ReadBuildInfo()
	return FromBuildInfo(info, ok)
}

// FromBuildInfo formats a version string from Go build metadata.
func FromBuildInfo(info *debug.BuildInfo, ok bool) string {
	if !ok || info == nil {
		return unknown
	}

	if v := info.Main.Version; v != "" && v != devel {
		return v
	}

	var revision, vcsTime string
	modified := false
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			revision = s.Value
		case "vcs.time":
			vcsTime = s.Value
		case "vcs.modified":
			modified = s.Value == "true"
		}
	}

	if revision == "" {
		return "devel"
	}
	if len(revision) > shaLen {
		revision = revision[:shaLen]
	}
	if modified {
		revision += "-dirty"
	}
	if vcsTime == "" {
		return revision
	}
	if t, err := time.Parse(time.RFC3339, vcsTime); err == nil {
		vcsTime = t.UTC().Format("2006-01-02")
	}
	return fmt.Sprintf("%s %s", revision, vcsTime)
}
