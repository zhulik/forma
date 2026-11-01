package core

import "context"

// GoModRunner runs go mod commands for CLI project setup.
type GoModRunner interface {
	ModInit(ctx context.Context, dir, module string) error
	RequireForma(ctx context.Context, dir string) error
	ReplaceForma(ctx context.Context, dir, localPath string) error
	ModTidy(ctx context.Context, dir string) error
}

// ProjectSyncer reconciles on-disk project state with forma.yaml (stub until web tree/codegen).
type ProjectSyncer interface {
	Sync(ctx context.Context, dir, formaReplace string) error
}
