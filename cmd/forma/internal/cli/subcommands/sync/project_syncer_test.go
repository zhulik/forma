package synccmd_test

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	synccmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/sync"
	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

type recordingNpm struct {
	dir   string
	calls int
}

func (r *recordingNpm) Install(_ context.Context, dir string) error {
	r.dir = dir
	r.calls++
	return nil
}

type noopGoMod struct{}

func (noopGoMod) InitModule(context.Context, string, string) error { return nil }
func (noopGoMod) Require(context.Context, string) error            { return nil }
func (noopGoMod) Replace(context.Context, string, string) error    { return nil }
func (noopGoMod) Tidy(context.Context, string) error               { return nil }

func TestProjectSyncer_Sync_NpmAndGoOnly(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, formaconfig.FileName), []byte(`module: github.com/example/x
ui:
  mount: app-root
web:
  path: web
`), 0o600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module github.com/example/x\n\ngo 1.27.3\n"), 0o600))
	webDir := filepath.Join(dir, "web")
	require.NoError(t, os.MkdirAll(webDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(webDir, "package.json"), []byte("{}"), 0o600))

	npm := &recordingNpm{}
	svc := &synccmd.ProjectSyncer{
		Logger:       slog.Default(),
		GoModWrapper: noopGoMod{},
		NpmWrapper:   npm,
	}

	err := svc.Sync(context.Background(), core.SyncRequest{Dir: dir})
	require.NoError(t, err)
	require.Equal(t, webDir, npm.dir)
	require.Equal(t, 1, npm.calls)
}
