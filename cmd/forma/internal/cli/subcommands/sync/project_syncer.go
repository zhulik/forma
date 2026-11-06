package synccmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
	"github.com/zhulik/forma/cmd/forma/internal/gotool"
	"github.com/zhulik/forma/cmd/forma/internal/webtool"
)

// EnvFormaReplace is the environment variable for a local replace of the forma module (cmd only).
const EnvFormaReplace = "FORMA_REPLACE"

// ProjectSyncer reconciles on-disk project state with forma.yaml.
type ProjectSyncer struct {
	Dir          string
	FormaReplace string

	Logger       *slog.Logger
	GoModWrapper core.GoModWrapper
	NpmWrapper   core.NpmWrapper
}

// Run reconciles project state using CLI-bound fields (pal.Runner entry for sync).
func (s *ProjectSyncer) Run(ctx context.Context) error {
	return s.Sync(ctx, core.SyncRequest{
		Dir:          s.Dir,
		FormaReplace: s.FormaReplace,
	})
}

// Sync runs web install and Go reconciliation (no kit materialization or codegen).
func (s *ProjectSyncer) Sync(ctx context.Context, req core.SyncRequest) error {
	target, err := resolveDir(req.Dir)
	if err != nil {
		return err
	}

	configPath := filepath.Join(target, formaconfig.FileName)
	if _, err := os.Stat(configPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s: %w", formaconfig.FileName, err)
		}
		return fmt.Errorf("stat %s: %w", formaconfig.FileName, err)
	}

	project, err := formaconfig.Load(target)
	if err != nil {
		return fmt.Errorf("load %s: %w", formaconfig.FileName, err)
	}

	webRoot := project.WebRoot(target)
	packageJSON := filepath.Join(webRoot, "package.json")
	if _, err := os.Stat(packageJSON); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%w at %s", clierrors.ErrWebPackageMissing, packageJSON)
		}
		return fmt.Errorf("stat web package.json: %w", err)
	}

	if err := s.NpmWrapper.Install(ctx, webRoot); err != nil {
		return err
	}

	if err := s.GoModWrapper.Require(ctx, target); err != nil {
		return fmt.Errorf("require forma library: %w", err)
	}
	if req.FormaReplace != "" {
		if err := s.GoModWrapper.Replace(ctx, target, req.FormaReplace); err != nil {
			return fmt.Errorf("replace forma library: %w", err)
		}
	}
	if err := s.GoModWrapper.Tidy(ctx, target); err != nil {
		return fmt.Errorf("go mod tidy: %w", err)
	}

	return nil
}

// Run reconciles project state from forma.yaml (web tree + Go module).
func Run(ctx context.Context, dir string) error {
	return runViaPal(ctx, &ProjectSyncer{Dir: dir})
}

// RunWithReplace is like Run but applies formaReplace when non-empty.
func RunWithReplace(ctx context.Context, dir, formaReplace string) error {
	return runViaPal(ctx, &ProjectSyncer{
		Dir:          dir,
		FormaReplace: formaReplace,
	})
}

func runViaPal(ctx context.Context, s *ProjectSyncer) error {
	services := []pal.ServiceDef{gotool.Provide()}
	services = append(services, webtool.Provide()...)
	services = append(services, pal.Provide(s))
	return app.Run(ctx, services...)
}

func resolveDir(directory string) (string, error) {
	if directory == "" {
		return os.Getwd()
	}

	return filepath.Abs(directory)
}
