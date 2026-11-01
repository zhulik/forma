package webtool

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/pkg/cmdexec"
)

// Tool runs npm and shadcn CLI in a web project directory.
type Tool struct{}

// Provide registers the web sync runner for pal injection.
func Provide() pal.ServiceDef {
	return pal.Provide[core.WebSyncRunner](&Tool{})
}

func (Tool) NpmInstall(ctx context.Context, dir string) error {
	if err := requireNpm(); err != nil {
		return err
	}

	if err := cmdexec.Run(ctx, dir, "npm", "install"); err != nil {
		return fmt.Errorf("npm install: %w", err)
	}

	return nil
}

func (Tool) ShadcnAdd(ctx context.Context, dir string, components []string) error {
	if len(components) == 0 {
		return nil
	}
	if err := requireNpm(); err != nil {
		return err
	}

	args := []string{"--yes", "shadcn@latest", "add", "--yes"}
	args = append(args, components...)
	if err := cmdexec.Run(ctx, dir, "npx", args...); err != nil {
		return fmt.Errorf("shadcn add: %w", err)
	}

	return nil
}

func requireNpm() error {
	if _, err := exec.LookPath("npm"); err != nil {
		return errors.New("npm not found on PATH; install Node.js to sync the web tree")
	}

	return nil
}
