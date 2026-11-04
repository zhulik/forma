package core

import "context"

// GoModWrapper runs go mod commands for CLI project setup.
type GoModWrapper interface {
	InitModule(ctx context.Context, dir, module string) error
	Require(ctx context.Context, dir string) error
	Replace(ctx context.Context, dir, localPath string) error
	Tidy(ctx context.Context, dir string) error
}

// SyncRequest is the input to ProjectSyncer.Sync.
type SyncRequest struct {
	Dir           string
	FormaReplace  string
	AddComponents []string // shadcn add these names; empty skips shadcn
}

// ProjectSyncer reconciles on-disk project state with forma.yaml.
type ProjectSyncer interface {
	Sync(ctx context.Context, req SyncRequest) error
}

// NpmWrapper installs npm dependencies under a web root.
type NpmWrapper interface {
	Install(ctx context.Context, dir string) error
}

// ShadcnWrapper materializes shadcn components under a web root.
type ShadcnWrapper interface {
	Add(ctx context.Context, dir string, components []string) error
}

// CodegenRequest is the input to CodeGenerator.Generate.
type CodegenRequest struct {
	ComponentsDir string
	CodegenDir    string
	Package       string
	RegistryOut   string
	FormaUIImport string
}

// CodeGenerator turns local shadcn TSX into Go bindings and a JS registry.
type CodeGenerator interface {
	Generate(ctx context.Context, req CodegenRequest) error
}
