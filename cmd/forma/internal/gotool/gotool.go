package gotool

import (
	"context"

	"github.com/zhulik/forma/cmd/forma/pkg/cmdexec"
)

const FormaModule = "github.com/zhulik/forma"

// Tool runs go commands in a project directory.
type Tool struct{}

func (Tool) ModInit(ctx context.Context, dir, module string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "init", module)
}

func (Tool) RequireForma(ctx context.Context, dir string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "edit", "-require="+FormaModule+"@v0.0.0")
}

func (Tool) ReplaceForma(ctx context.Context, dir, localPath string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "edit", "-replace="+FormaModule+"="+localPath)
}

func (Tool) ModTidy(ctx context.Context, dir string) error {
	return cmdexec.Run(ctx, dir, "go", "mod", "tidy")
}
