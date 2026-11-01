package versioncmd

import (
	"context"

	"github.com/urfave/cli/v3"
)

// New returns the version subcommand.
func New() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "print the forma CLI version",
		Action: func(_ context.Context, cmd *cli.Command) error {
			cli.ShowVersion(cmd.Root())
			return nil
		},
	}
}
