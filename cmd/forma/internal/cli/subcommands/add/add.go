package addcmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
	synccmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/sync"
)

const flagDirectory = "directory"

// New returns the add subcommand.
func New() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "materialize a shadcn component and sync (codegen + go mod)",
		ArgsUsage: "<component>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagDirectory,
				Aliases: []string{"d"},
				Usage:   "project directory (default: current working directory)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			name := cmd.Args().First()
			if name == "" {
				return clierrors.ErrComponentNameRequired
			}
			return synccmd.RunWithComponents(
				ctx,
				cmd.String(flagDirectory),
				os.Getenv(synccmd.EnvFormaReplace),
				[]string{name},
			)
		},
	}
}
