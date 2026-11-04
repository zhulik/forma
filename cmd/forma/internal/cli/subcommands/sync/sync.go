package synccmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
	"github.com/zhulik/forma/cmd/forma/internal/codegen"
	"github.com/zhulik/forma/cmd/forma/internal/gotool"
	"github.com/zhulik/forma/cmd/forma/internal/webtool"
)

const flagDirectory = "directory"

// New returns the sync subcommand.
func New() *cli.Command {
	return &cli.Command{
		Name:  "sync",
		Usage: "reconcile project with forma.yaml (npm, codegen, go mod tidy)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagDirectory,
				Aliases: []string{"d"},
				Usage:   "project directory (default: current working directory)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return app.Run(ctx, gotool.Provide(), webtool.Provide(), codegen.Provide(), Provide(), pal.Provide(&runner{
				Dir:          cmd.String(flagDirectory),
				FormaReplace: os.Getenv(EnvFormaReplace),
			}))
		},
	}
}
