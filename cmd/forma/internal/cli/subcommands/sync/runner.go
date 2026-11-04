package synccmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
	"github.com/zhulik/forma/cmd/forma/internal/codegen"
	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/gotool"
	"github.com/zhulik/forma/cmd/forma/internal/webtool"
)

// EnvFormaReplace is the environment variable for a local replace of the forma module (cmd only).
const EnvFormaReplace = "FORMA_REPLACE"

type runner struct {
	Dir           string
	FormaReplace  string
	AddComponents []string
	ProjectSyncer core.ProjectSyncer
}

// Run reconciles project state from forma.yaml (web tree + Go module).
func Run(ctx context.Context, dir string) error {
	return runViaPal(ctx, &runner{
		Dir: dir,
	})
}

// RunWithReplace is like Run but applies formaReplace when non-empty.
func RunWithReplace(ctx context.Context, dir, formaReplace string) error {
	return runViaPal(ctx, &runner{
		Dir:          dir,
		FormaReplace: formaReplace,
	})
}

// RunWithComponents runs sync and materializes the given shadcn components first.
func RunWithComponents(ctx context.Context, dir, formaReplace string, components []string) error {
	return runViaPal(ctx, &runner{
		Dir:           dir,
		FormaReplace:  formaReplace,
		AddComponents: components,
	})
}

func runViaPal(ctx context.Context, r *runner) error {
	return app.Run(ctx, gotool.Provide(), webtool.Provide(), codegen.Provide(), Provide(), pal.Provide(r))
}

func (r *runner) Run(ctx context.Context) error {
	return r.ProjectSyncer.Sync(ctx, core.SyncRequest{
		Dir:           r.Dir,
		FormaReplace:  r.FormaReplace,
		AddComponents: r.AddComponents,
	})
}

func resolveDir(directory string) (string, error) {
	if directory == "" {
		return os.Getwd()
	}

	return filepath.Abs(directory)
}
