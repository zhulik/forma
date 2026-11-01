package synccmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

// Service reconciles project module state (stub sync until host/codegen exist).
type Service struct {
	Logger      *slog.Logger
	GoModRunner core.GoModRunner
}

// Sync runs stub sync in dir with optional local replace path.
func (s *Service) Sync(ctx context.Context, dir, formaReplace string) error {
	target, err := resolveDir(dir)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(target, formaconfig.FileName)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s: %w", formaconfig.FileName, err)
		}
		return fmt.Errorf("stat %s: %w", formaconfig.FileName, err)
	}

	s.Logger.DebugContext(ctx, "sync stub: host and codegen steps skipped")

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
