package formaconfig

import (
	"path/filepath"
)

// UIConfig is the ui section of forma.yaml.
type UIConfig struct {
	Mount      string   `yaml:"mount"`
	Components []string `yaml:"components"`
}

// WebConfig is the web section of forma.yaml.
type WebConfig struct {
	Path string `yaml:"path"`
}

// CodegenConfig is the codegen section of forma.yaml.
type CodegenConfig struct {
	Dir     string `yaml:"dir"`
	Package string `yaml:"package"`
}

// Project is the root forma.yaml shape used by the CLI.
type Project struct {
	Module  string        `yaml:"module"`
	UI      UIConfig      `yaml:"ui"`
	Web     WebConfig     `yaml:"web"`
	Codegen CodegenConfig `yaml:"codegen"`
}

// WebPath returns the configured web root relative path, defaulting to web.
func (p Project) WebPath() string {
	if p.Web.Path != "" {
		return p.Web.Path
	}

	return "web"
}

// WebRoot returns the absolute web tree directory under the project dir.
func (p Project) WebRoot(projectDir string) string {
	return filepath.Join(projectDir, p.WebPath())
}
