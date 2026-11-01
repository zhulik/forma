package synccmd

import (
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
)

// Provide registers the sync service for pal injection.
func Provide() pal.ServiceDef {
	return pal.Provide[core.ProjectSyncer](&Service{})
}
