package synccmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
	"github.com/zhulik/forma/cmd/forma/internal/gotool"
)

const flagDirectory = "directory"

// New returns the sync subcommand.
func New() *cli.Command {
	return &cli.Command{
		Name:  "sync",
		Usage: "reconcile project with forma.yaml (stub: Go module only until host ships)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagDirectory,
				Aliases: []string{"d"},
				Usage:   "project directory (default: current working directory)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return app.Run(ctx, gotool.Provide(), Provide(), pal.Provide(&runner{
				Dir:          cmd.String(flagDirectory),
				FormaReplace: os.Getenv(EnvFormaReplace),
			}))
		},
	}
}
