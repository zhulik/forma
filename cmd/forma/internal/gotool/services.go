package gotool

import (
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
)

// Provide registers the default Go toolchain adapter for pal injection.
func Provide() pal.ServiceDef {
	return pal.Provide[core.GoModRunner](&Tool{})
}
