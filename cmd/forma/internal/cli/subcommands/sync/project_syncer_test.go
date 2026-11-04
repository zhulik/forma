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
	dir string
}

func (r *recordingNpm) Install(_ context.Context, dir string) error {
	r.dir = dir
	return nil
}

type recordingShadcn struct {
	dir  string
	args []string
}

func (r *recordingShadcn) Add(_ context.Context, dir string, components []string) error {
	r.dir = dir
	r.args = append([]string(nil), components...)
	return nil
}

type recordingCodegen struct {
	reqs []core.CodegenRequest
}

func (r *recordingCodegen) Generate(_ context.Context, req core.CodegenRequest) error {
	r.reqs = append(r.reqs, req)
	return nil
}

type noopGoMod struct{}

func (noopGoMod) InitModule(context.Context, string, string) error { return nil }
func (noopGoMod) Require(context.Context, string) error            { return nil }
func (noopGoMod) Replace(context.Context, string, string) error    { return nil }
func (noopGoMod) Tidy(context.Context, string) error               { return nil }

func TestProjectSyncer_Sync_WebCodegenAndGo(t *testing.T) {
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

	npm := &recordingNpm{}
	shadcn := &recordingShadcn{}
	cg := &recordingCodegen{}
	svc := &synccmd.ProjectSyncer{
		Logger:        slog.Default(),
		GoModWrapper:  noopGoMod{},
		NpmWrapper:    npm,
		ShadcnWrapper: shadcn,
		CodeGenerator: cg,
	}

	err := svc.Sync(context.Background(), core.SyncRequest{
		Dir:           dir,
		AddComponents: []string{"button", "button", "card"},
	})
	require.NoError(t, err)
	require.Equal(t, webDir, npm.dir)
	require.Equal(t, webDir, shadcn.dir)
	require.Equal(t, []string{"button", "card"}, shadcn.args)
	require.Len(t, cg.reqs, 1)
	require.Equal(t, filepath.Join(webDir, "src", "components", "ui"), cg.reqs[0].ComponentsDir)
	require.Equal(t, filepath.Join(dir, "internal", "ui"), cg.reqs[0].CodegenDir)
	require.Equal(t, "ui", cg.reqs[0].Package)
	require.Equal(t, filepath.Join(webDir, "src", "registry.gen.ts"), cg.reqs[0].RegistryOut)
}

func TestProjectSyncer_Sync_NoShadcnWhenEmpty(t *testing.T) {
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
	shadcn := &recordingShadcn{}
	cg := &recordingCodegen{}
	svc := &synccmd.ProjectSyncer{
		Logger:        slog.Default(),
		GoModWrapper:  noopGoMod{},
		NpmWrapper:    npm,
		ShadcnWrapper: shadcn,
		CodeGenerator: cg,
	}

	err := svc.Sync(context.Background(), core.SyncRequest{Dir: dir})
	require.NoError(t, err)
	require.Empty(t, shadcn.args)
	require.Len(t, cg.reqs, 1)
}
