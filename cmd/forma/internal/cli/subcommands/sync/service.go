package synccmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/zhulik/forma/fn"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

// Service reconciles on-disk project state with forma.yaml.
type Service struct {
	Logger        *slog.Logger
	GoModRunner   core.GoModRunner
	WebSyncRunner core.WebSyncRunner
}

// Sync runs web and Go reconciliation in dir with optional local replace path.
func (s *Service) Sync(ctx context.Context, dir, formaReplace string) error {
	target, err := resolveDir(dir)
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

	if err := s.WebSyncRunner.NpmInstall(ctx, webRoot); err != nil {
		return err
	}

	components := dedupeComponents(project.UI.Components)
	if err := s.WebSyncRunner.ShadcnAdd(ctx, webRoot, components); err != nil {
		return err
	}

	s.Logger.InfoContext(ctx, "sync: Go widget codegen skipped (not implemented)")
	s.Logger.InfoContext(ctx, "sync: JS component registry update skipped (not implemented)")

	if err := s.GoModRunner.RequireForma(ctx, target); err != nil {
		return fmt.Errorf("require forma library: %w", err)
	}
	if formaReplace != "" {
		if err := s.GoModRunner.ReplaceForma(ctx, target, formaReplace); err != nil {
			return fmt.Errorf("replace forma library: %w", err)
		}
	}
	if err := s.GoModRunner.ModTidy(ctx, target); err != nil {
		return fmt.Errorf("go mod tidy: %w", err)
	}

	return nil
}

func dedupeComponents(names []string) []string {
	return fn.Uniq(fn.Filter(names, func(name string) bool { return name != "" }))
}
