package webtool

import (
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
)

// Provide registers npm and shadcn wrappers for pal injection.
func Provide() []pal.ServiceDef {
	return []pal.ServiceDef{
		pal.Provide[core.NpmWrapper](&NpmWrapper{}),
		pal.Provide[core.ShadcnWrapper](&ShadcnWrapper{}),
	}
}
