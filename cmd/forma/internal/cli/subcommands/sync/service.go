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
	CodegenRunner core.CodegenRunner
}

// Sync runs web, codegen, and Go reconciliation.
func (s *Service) Sync(ctx context.Context, req core.SyncRequest) error {
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

	if err := s.WebSyncRunner.NpmInstall(ctx, webRoot); err != nil {
		return err
	}

	components := dedupeComponents(req.AddComponents)
	if err := s.WebSyncRunner.ShadcnAdd(ctx, webRoot, components); err != nil {
		return err
	}

	if err := runCodegen(ctx, s.CodegenRunner, target, project); err != nil {
		return err
	}

	if err := s.GoModRunner.RequireForma(ctx, target); err != nil {
		return fmt.Errorf("require forma library: %w", err)
	}
	if req.FormaReplace != "" {
		if err := s.GoModRunner.ReplaceForma(ctx, target, req.FormaReplace); err != nil {
			return fmt.Errorf("replace forma library: %w", err)
		}
	}
	if err := s.GoModRunner.ModTidy(ctx, target); err != nil {
		return fmt.Errorf("go mod tidy: %w", err)
	}

	return nil
}

func runCodegen(ctx context.Context, runner core.CodegenRunner, projectDir string, project formaconfig.Project) error {
	componentsDir, codegenDir, pkg, registryOut := project.ToCodegenPaths(projectDir)
	if err := runner.Generate(ctx, core.CodegenRequest{
		ComponentsDir: componentsDir,
		CodegenDir:    codegenDir,
		Package:       pkg,
		RegistryOut:   registryOut,
	}); err != nil {
		return fmt.Errorf("codegen: %w", err)
	}
	return nil
}

func dedupeComponents(names []string) []string {
	return fn.Uniq(fn.Filter(names, func(name string) bool { return name != "" }))
}
