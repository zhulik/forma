package webtool

import (
	"context"
	"fmt"

	"github.com/zhulik/forma/cmd/forma/pkg/cmdexec"
)

// NpmWrapper runs npm install in a web project directory.
type NpmWrapper struct{}

func (NpmWrapper) Install(ctx context.Context, dir string) error {
	if err := requireNpm(); err != nil {
		return err
	}

	if err := cmdexec.Run(ctx, dir, "npm", "install"); err != nil {
		return fmt.Errorf("npm install: %w", err)
	}

	return nil
}
