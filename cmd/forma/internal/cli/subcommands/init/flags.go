package initcmd

import "github.com/urfave/cli/v3"

// Flags returns init subcommand flags.
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    flagDirectory,
			Aliases: []string{"d"},
			Usage:   "project directory (must be empty for full init)",
		},
		&cli.StringFlag{
			Name:  flagTemplate,
			Value: defaultTemplate,
			Usage: "embedded scaffold template name",
		},
		&cli.BoolFlag{
			Name:  flagOnlyConfig,
			Usage: "write forma.yaml only when it does not exist",
		},
	}
}

// Arguments returns the module path positional argument.
func Arguments() []cli.Argument {
	return []cli.Argument{
		&cli.StringArg{
			Name: argModule,
		},
	}
}
