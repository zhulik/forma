package devcmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v3"
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
)

const flagDirectory = "directory"

// New returns the forma dev subcommand (Vite host + GoScript rebuild loop).
func New() *cli.Command {
	return &cli.Command{
		Name:  "dev",
		Usage: "run the browser development loop (Vite host + GoScript rebuild)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagDirectory,
				Aliases: []string{"d"},
				Usage:   "project directory (default: current working directory)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
			defer stop()

			return app.Run(ctx, pal.Provide(&Runner{
				Dir: cmd.String(flagDirectory),
			}))
		},
	}
}
