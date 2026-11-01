package devcmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
)

// New returns the dev subcommand (planned P3d).
func New() *cli.Command {
	return &cli.Command{
		Name:  "dev",
		Usage: "run the browser development loop (planned)",
		Action: func(_ context.Context, _ *cli.Command) error {
			return fmt.Errorf("forma dev: %w (see ROADMAP P3d)", clierrors.ErrNotImplemented)
		},
	}
}
