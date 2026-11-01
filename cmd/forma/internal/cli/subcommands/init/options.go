package initcmd

import "github.com/urfave/cli/v3"

const (
	flagDirectory   = "directory"
	flagTemplate    = "template"
	flagOnlyConfig  = "only-config"
	argModule       = "module"
	defaultTemplate = "minimal"
	defaultMount    = "app-root"
)

// Options holds init inputs.
type Options struct {
	Module     string
	Directory  string
	Template   string
	OnlyConfig bool
}

// OptionsFromCommand reads init options from the urfave command.
func OptionsFromCommand(cmd *cli.Command) Options {
	template := cmd.String(flagTemplate)
	if template == "" {
		template = defaultTemplate
	}
	return Options{
		Module:     cmd.StringArg(argModule),
		Directory:  cmd.String(flagDirectory),
		Template:   template,
		OnlyConfig: cmd.Bool(flagOnlyConfig),
	}
}
