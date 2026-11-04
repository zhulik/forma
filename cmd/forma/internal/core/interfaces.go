package core

import "context"

// GoModRunner runs go mod commands for CLI project setup.
type GoModRunner interface {
	ModInit(ctx context.Context, dir, module string) error
	RequireForma(ctx context.Context, dir string) error
	ReplaceForma(ctx context.Context, dir, localPath string) error
	ModTidy(ctx context.Context, dir string) error
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

// WebSyncRunner installs npm dependencies and materializes shadcn components under a web root.
type WebSyncRunner interface {
	NpmInstall(ctx context.Context, dir string) error
	ShadcnAdd(ctx context.Context, dir string, components []string) error
}

// CodegenRequest is the input to CodegenRunner.Generate.
type CodegenRequest struct {
	ComponentsDir string
	CodegenDir    string
	Package       string
	RegistryOut   string
	FormaUIImport string
}

// CodegenRunner turns local shadcn TSX into Go bindings and a JS registry.
type CodegenRunner interface {
	Generate(ctx context.Context, req CodegenRequest) error
}
