package codegen

import (
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
)

// Provide registers the code generator for pal injection.
func Provide() pal.ServiceDef {
	return pal.Provide[core.CodeGenerator](&CodeGenerator{})
}
