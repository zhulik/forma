package initcmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/cli/app"
	"github.com/zhulik/forma/cmd/forma/internal/cli/deps"
	synccmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/sync"
)

// New returns the init subcommand.
func New() *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "initialize a new forma app",
		Flags:     Flags(),
		Arguments: Arguments(),
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return app.Run(ctx, deps.Services(pal.Provide(&runner{
				Opts:         OptionsFromCommand(cmd),
				FormaReplace: os.Getenv(synccmd.EnvFormaReplace),
			}))...)
		},
	}
}
