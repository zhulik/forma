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

// CodegenRelDir returns the configured codegen directory, defaulting to internal/ui.
func (p Project) CodegenRelDir() string {
	if p.Codegen.Dir != "" {
		return p.Codegen.Dir
	}
	return "internal/ui"
}

// CodegenPackageName returns the Go package name for generated files, defaulting to ui.
func (p Project) CodegenPackageName() string {
	if p.Codegen.Package != "" {
		return p.Codegen.Package
	}
	return "ui"
}

// CodegenAbsDir returns the absolute codegen output directory.
func (p Project) CodegenAbsDir(projectDir string) string {
	return filepath.Join(projectDir, p.CodegenRelDir())
}

// ComponentsUIDir returns the absolute shadcn components/ui directory under the web root.
// Scaffold aliases map @/ to src/, so widgets live at web/src/components/ui.
func (p Project) ComponentsUIDir(projectDir string) string {
	return filepath.Join(p.WebRoot(projectDir), "src", "components", "ui")
}

// RegistryOutPath returns the absolute path for registry.gen.ts under the web src tree.
func (p Project) RegistryOutPath(projectDir string) string {
	return filepath.Join(p.WebRoot(projectDir), "src", "registry.gen.ts")
}

// CodegenRequest builds a core-compatible path set for generation from this project.
func (p Project) ToCodegenPaths(projectDir string) (componentsDir, codegenDir, pkg, registryOut string) {
	return p.ComponentsUIDir(projectDir),
		p.CodegenAbsDir(projectDir),
		p.CodegenPackageName(),
		p.RegistryOutPath(projectDir)
}
