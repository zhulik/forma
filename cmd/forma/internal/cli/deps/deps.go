package deps

import (
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/gotool"
)

// Services returns shared pal service definitions followed by extras (e.g. command runners).
// Sync runs in its own pal graph via synccmd.Run* helpers (ProjectSyncer is the sync runner).
func Services(extras ...pal.ServiceDef) []pal.ServiceDef {
	out := make([]pal.ServiceDef, 0, 1+len(extras))
	out = append(out, gotool.Provide())
	out = append(out, extras...)

	return out
}
