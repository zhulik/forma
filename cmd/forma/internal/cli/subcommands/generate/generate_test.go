package generatecmd_test

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	generatecmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/generate"
	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

type recordingCodegen struct {
	reqs []core.CodegenRequest
}

func (r *recordingCodegen) Generate(_ context.Context, req core.CodegenRequest) error {
	r.reqs = append(r.reqs, req)
	return nil
}

func TestRunner_Generate_LoadsConfigAndRunsCodegen(t *testing.T) {
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

	cg := &recordingCodegen{}
	r := &generatecmd.Runner{
		Logger:        slog.Default(),
		CodeGenerator: cg,
	}
	err := r.Generate(context.Background(), dir)
	require.NoError(t, err)
	require.Len(t, cg.reqs, 1)
	require.Equal(t, filepath.Join(dir, "web", "src", "components", "ui"), cg.reqs[0].ComponentsDir)
	require.Equal(t, filepath.Join(dir, "internal", "ui"), cg.reqs[0].CodegenDir)
	require.Equal(t, "ui", cg.reqs[0].Package)
	require.Equal(t, filepath.Join(dir, "web", "src", "registry.gen.ts"), cg.reqs[0].RegistryOut)
}

func TestRunner_Generate_MissingConfig(t *testing.T) {
	t.Parallel()

	r := &generatecmd.Runner{
		Logger:        slog.Default(),
		CodeGenerator: &recordingCodegen{},
	}
	err := r.Generate(context.Background(), t.TempDir())
	require.Error(t, err)
	require.Contains(t, err.Error(), formaconfig.FileName)
}
