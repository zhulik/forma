package core

import "context"

// GoModRunner runs go mod commands for CLI project setup.
type GoModRunner interface {
	ModInit(ctx context.Context, dir, module string) error
	RequireForma(ctx context.Context, dir string) error
	ReplaceForma(ctx context.Context, dir, localPath string) error
	ModTidy(ctx context.Context, dir string) error
}

// ProjectSyncer reconciles on-disk project state with forma.yaml.
type ProjectSyncer interface {
	Sync(ctx context.Context, dir, formaReplace string) error
}

// WebSyncRunner installs npm dependencies and materializes shadcn components under a web root.
type WebSyncRunner interface {
	NpmInstall(ctx context.Context, dir string) error
	ShadcnAdd(ctx context.Context, dir string, components []string) error
}
