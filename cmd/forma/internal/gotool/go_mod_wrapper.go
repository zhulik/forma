package gotool

import (
	"context"

	"github.com/zhulik/forma/cmd/forma/pkg/cmdexec"
)

const FormaModule = "github.com/zhulik/forma"

// GoModWrapper runs go mod commands in a project directory.
type GoModWrapper struct{}

func (GoModWrapper) InitModule(ctx context.Context, dir, module string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "init", module)
}

func (GoModWrapper) Require(ctx context.Context, dir string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "edit", "-require="+FormaModule+"@v0.0.0")
}

func (GoModWrapper) Replace(ctx context.Context, dir, localPath string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "edit", "-replace="+FormaModule+"="+localPath)
}

func (GoModWrapper) Tidy(ctx context.Context, dir string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "tidy")
}
