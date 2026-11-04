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

type recordingWebRunner struct {
	npmDir     string
	shadcnDir  string
	shadcnArgs []string
}

func (r *recordingWebRunner) NpmInstall(_ context.Context, dir string) error {
	r.npmDir = dir
	return nil
}

func (r *recordingWebRunner) ShadcnAdd(_ context.Context, dir string, components []string) error {
	r.shadcnDir = dir
	r.shadcnArgs = append([]string(nil), components...)
	return nil
}

type recordingCodegen struct {
	reqs []core.CodegenRequest
}

func (r *recordingCodegen) Generate(_ context.Context, req core.CodegenRequest) error {
	r.reqs = append(r.reqs, req)
	return nil
}

type noopGoModRunner struct{}

func (noopGoModRunner) ModInit(context.Context, string, string) error { return nil }
func (noopGoModRunner) RequireForma(context.Context, string) error    { return nil }
func (noopGoModRunner) ReplaceForma(context.Context, string, string) error {
	return nil
}
func (noopGoModRunner) ModTidy(context.Context, string) error { return nil }

func TestService_Sync_WebCodegenAndGo(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, formaconfig.FileName), []byte(`module: github.com/example/x
ui:
  mount: app-root
web:
  path: web
codegen:
  dir: internal/ui
  package: ui
`), 0o600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module github.com/example/x\n\ngo 1.27.3\n"), 0o600))
	webDir := filepath.Join(dir, "web")
	require.NoError(t, os.MkdirAll(webDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(webDir, "package.json"), []byte("{}"), 0o600))

	web := &recordingWebRunner{}
	cg := &recordingCodegen{}
	svc := &synccmd.Service{
		Logger:        slog.Default(),
		GoModRunner:   noopGoModRunner{},
		WebSyncRunner: web,
		CodegenRunner: cg,
	}

	err := svc.Sync(context.Background(), core.SyncRequest{
		Dir:           dir,
		AddComponents: []string{"button", "button", "card"},
	})
	require.NoError(t, err)
	require.Equal(t, webDir, web.npmDir)
	require.Equal(t, webDir, web.shadcnDir)
	require.Equal(t, []string{"button", "card"}, web.shadcnArgs)
	require.Len(t, cg.reqs, 1)
	require.Equal(t, filepath.Join(webDir, "src", "components", "ui"), cg.reqs[0].ComponentsDir)
	require.Equal(t, filepath.Join(dir, "internal", "ui"), cg.reqs[0].CodegenDir)
	require.Equal(t, "ui", cg.reqs[0].Package)
	require.Equal(t, filepath.Join(webDir, "src", "registry.gen.ts"), cg.reqs[0].RegistryOut)
}

func TestService_Sync_NoShadcnWhenEmpty(t *testing.T) {
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

	web := &recordingWebRunner{}
	cg := &recordingCodegen{}
	svc := &synccmd.Service{
		Logger:        slog.Default(),
		GoModRunner:   noopGoModRunner{},
		WebSyncRunner: web,
		CodegenRunner: cg,
	}

	err := svc.Sync(context.Background(), core.SyncRequest{Dir: dir})
	require.NoError(t, err)
	require.Empty(t, web.shadcnArgs)
	require.Len(t, cg.reqs, 1)
}
