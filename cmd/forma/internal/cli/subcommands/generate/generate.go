package generatecmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
	"github.com/zhulik/forma/cmd/forma/internal/codegen"
	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

const flagDirectory = "directory"

// New returns the generate subcommand.
func New() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "regenerate Go widget bindings and JS registry from shadcn sources",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagDirectory,
				Aliases: []string{"d"},
				Usage:   "project directory (default: current working directory)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return app.Run(ctx, codegen.Provide(), Provide(), pal.Provide(&runner{
				Dir: cmd.String(flagDirectory),
			}))
		},
	}
}

// Service loads forma.yaml and runs codegen only.
type Service struct {
	Logger        *slog.Logger
	CodegenRunner core.CodegenRunner
}

// Provide registers the generate service (also used as Codegen orchestration for the command runner).
func Provide() pal.ServiceDef {
	return pal.Provide(&Service{})
}

// Generate reads forma.yaml under dir and regenerates bindings and registry.
func (s *Service) Generate(ctx context.Context, dir string) error {
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

	componentsDir, codegenDir, pkg, registryOut := project.ToCodegenPaths(target)
	if err := s.CodegenRunner.Generate(ctx, core.CodegenRequest{
		ComponentsDir: componentsDir,
		CodegenDir:    codegenDir,
		Package:       pkg,
		RegistryOut:   registryOut,
	}); err != nil {
		return err
	}

	s.Logger.InfoContext(ctx, "generate complete", "codegen", codegenDir, "registry", registryOut)
	return nil
}

type runner struct {
	Dir     string
	Service *Service
}

func (r *runner) Run(ctx context.Context) error {
	return r.Service.Generate(ctx, r.Dir)
}

func resolveDir(directory string) (string, error) {
	if directory == "" {
		return os.Getwd()
	}
	return filepath.Abs(directory)
}
