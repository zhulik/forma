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
	Dir          string
	FormaReplace string
}

// ProjectSyncer reconciles on-disk project state with forma.yaml.
type ProjectSyncer interface {
	Sync(ctx context.Context, req SyncRequest) error
}

// NpmWrapper installs npm dependencies under a web root.
type NpmWrapper interface {
	Install(ctx context.Context, dir string) error
}
