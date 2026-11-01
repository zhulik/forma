package addcmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
)

// New returns the add subcommand (not implemented until full sync ships).
func New() *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "add a shadcn component to forma.yaml and sync",
		Action: func(_ context.Context, _ *cli.Command) error {
			return fmt.Errorf("forma add: %w (see ROADMAP P3)", clierrors.ErrNotImplemented)
		},
	}
}
