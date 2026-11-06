package formaconfig

import (
	"path/filepath"
)

// UIConfig is the ui section of forma.yaml.
type UIConfig struct {
	Mount string `yaml:"mount"`
}

// WebConfig is the web section of forma.yaml.
type WebConfig struct {
	Path string `yaml:"path"`
}

// Project is the root forma.yaml shape used by the CLI.
type Project struct {
	Module string    `yaml:"module"`
	UI     UIConfig  `yaml:"ui"`
	Web    WebConfig `yaml:"web"`
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

// BrowserAppJSRel is the Vite-bundled GoScript browser artifact relative to the web root.
// Contract: Vite serves public/ at /, so the ES module is loaded as /app.js after
// main.tsx registers formaHost.setTree (host-before-Go order).
const BrowserAppJSRel = "public/app.js"

// BrowserAppJSPath returns the absolute browser Go bundle path under the web tree.
func (p Project) BrowserAppJSPath(projectDir string) string {
	return filepath.Join(p.WebRoot(projectDir), filepath.FromSlash(BrowserAppJSRel))
}
