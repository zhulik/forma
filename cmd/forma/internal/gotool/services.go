package gotool

import (
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
)

// Provide registers the default Go module wrapper for pal injection.
func Provide() pal.ServiceDef {
	return pal.Provide[core.GoModWrapper](&GoModWrapper{})
}
