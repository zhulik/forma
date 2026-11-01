package synccmd_test

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	synccmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/sync"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
	"github.com/zhulik/forma/cmd/forma/internal/testutil"
)

func TestRun_MissingFormaYAML(t *testing.T) {
	dir := t.TempDir()
	err := synccmd.Run(context.Background(), dir)
	require.Error(t, err)
}

func TestRun_StubSync(t *testing.T) {
	if testing.Short() {
		t.Skip("go mod and npm")
	}
	testutil.RequireNpm(t)

	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, formaconfig.FileName), []byte(`module: github.com/example/x
ui:
  components: []
web:
  path: web
`), 0o600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module github.com/example/x\n\ngo 1.27.3\n"), 0o600))
	webDir := filepath.Join(dir, "web")
	require.NoError(t, os.MkdirAll(webDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(webDir, "package.json"), []byte(`{"name":"test","private":true}`), 0o600))

	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok)
	repoRoot, err := filepath.Abs(filepath.Join(filepath.Dir(file), "..", "..", "..", "..", "..", ".."))
	require.NoError(t, err)

	err = synccmd.RunWithReplace(context.Background(), dir, repoRoot)
	require.NoError(t, err)
}
