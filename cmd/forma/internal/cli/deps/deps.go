package deps

import (
	"github.com/zhulik/pal"

	synccmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/sync"
	"github.com/zhulik/forma/cmd/forma/internal/codegen"
	"github.com/zhulik/forma/cmd/forma/internal/gotool"
	"github.com/zhulik/forma/cmd/forma/internal/webtool"
)

// Services returns shared pal service definitions followed by extras (e.g. command runners).
// synccmd cannot import this package (cycle); it lists gotool.Provide() and synccmd.Provide() inline instead.
func Services(extras ...pal.ServiceDef) []pal.ServiceDef {
	out := make([]pal.ServiceDef, 0, 4+len(extras))
	out = append(out, gotool.Provide(), webtool.Provide(), codegen.Provide(), synccmd.Provide())
	out = append(out, extras...)

	return out
}
