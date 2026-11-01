package version

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/zhulik/forma/fn"
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

	settings := fn.Associate(info.Settings, func(s debug.BuildSetting) (string, string) {
		return s.Key, s.Value
	})
	revision := settings["vcs.revision"]
	vcsTime := settings["vcs.time"]
	modified := settings["vcs.modified"] == "true"

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
