package initcmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
	"github.com/zhulik/forma/cmd/forma/internal/cli/deps"
	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
	synccmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/sync"
	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
	"github.com/zhulik/forma/cmd/forma/templates"
)

type runner struct {
	Opts          Options
	FormaReplace  string
	Logger        *slog.Logger
	GoModRunner   core.GoModRunner
	ProjectSyncer core.ProjectSyncer
}

// Run initializes a project (exported for tests).
func Run(ctx context.Context, opts Options) error {
	return app.Run(ctx, deps.Services(pal.Provide(&runner{
		Opts: opts,
	}))...)
}

func (r *runner) Run(ctx context.Context) error {
	target, err := resolveTarget(r.Opts.Directory)
	if err != nil {
		return err
	}

	if r.Opts.OnlyConfig {
		return r.runOnlyConfig(ctx, target)
	}

	return r.runFull(ctx, target)
}

func (r *runner) runOnlyConfig(ctx context.Context, target string) error {
	_ = ctx
	configPath := filepath.Join(target, formaconfig.FileName)
	if _, err := os.Stat(configPath); err == nil {
		return clierrors.ErrConfigExists
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat %s: %w", formaconfig.FileName, err)
	}

	if r.Opts.Module == "" {
		return clierrors.ErrModuleRequired
	}
	if err := templates.Exists(r.Opts.Template); err != nil {
		return fmt.Errorf("%w: %w", clierrors.ErrUnknownTemplate, err)
	}

	data := templateData(r.Opts.Module)
	content, err := templates.RenderFile(r.Opts.Template, "forma.yaml.tmpl", data)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(target, 0o755); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}
	//nolint:gosec // project config file
	if err := os.WriteFile(configPath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", formaconfig.FileName, err)
	}

	return nil
}

func (r *runner) runFull(ctx context.Context, target string) error {
	if r.Opts.Module == "" {
		return clierrors.ErrModuleRequired
	}

	empty, err := IsEmpty(target)
	if err != nil {
		return err
	}
	if !empty {
		return clierrors.ErrNotEmpty
	}

	if err := templates.Exists(r.Opts.Template); err != nil {
		return fmt.Errorf("%w: %w", clierrors.ErrUnknownTemplate, err)
	}

	if err := r.GoModRunner.ModInit(ctx, target, r.Opts.Module); err != nil {
		return fmt.Errorf("go mod init: %w", err)
	}

	data := templateData(r.Opts.Module)
	if err := templates.Apply(target, r.Opts.Template, data); err != nil {
		return fmt.Errorf("apply template: %w", err)
	}

	replace := r.FormaReplace
	if replace == "" {
		replace = os.Getenv(synccmd.EnvFormaReplace)
	}
	if err := r.ProjectSyncer.Sync(ctx, target, replace); err != nil {
		return fmt.Errorf("sync: %w", err)
	}

	r.Logger.InfoContext(ctx, "project initialized", "dir", target, "module", r.Opts.Module)

	return nil
}

func templateData(module string) templates.Data {
	return templates.Data{
		Module:  module,
		Package: templates.PackageName(module),
		Mount:   defaultMount,
	}
}
