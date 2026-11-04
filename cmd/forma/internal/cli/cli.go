package cli

import (
	"github.com/urfave/cli/v3"

	addcmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/add"
	devcmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/dev"
	generatecmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/generate"
	initcmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/init"
	synccmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/sync"
	versioncmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/version"
	"github.com/zhulik/forma/cmd/forma/internal/version"
)

// NewRoot returns the root forma CLI with subcommands registered.
func NewRoot() *cli.Command {
	return &cli.Command{
		Name:    "forma",
		Usage:   "forma UI framework CLI",
		Version: version.String(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  FlagLogLevel,
				Usage: "log level (debug, info, warn, error)",
			},
		},
		Commands: []*cli.Command{
			initcmd.New(),
			synccmd.New(),
			generatecmd.New(),
			addcmd.New(),
			devcmd.New(),
			versioncmd.New(),
		},
	}
}
