package webtool

import (
	"context"
	"fmt"

	"github.com/zhulik/forma/cmd/forma/pkg/cmdexec"
)

// ShadcnWrapper runs the shadcn CLI in a web project directory.
type ShadcnWrapper struct{}

func (ShadcnWrapper) Add(ctx context.Context, dir string, components []string) error {
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
