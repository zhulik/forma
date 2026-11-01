package cli_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitSmoke_Subprocess(t *testing.T) { //nolint:paralleltest // subprocess + go mod; serial with other mod tests
	if testing.Short() {
		t.Skip("subprocess")
	}

	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok)
	repoRoot, err := filepath.Abs(filepath.Join(filepath.Dir(file), "..", "..", "..", ".."))
	require.NoError(t, err)
	cliDir := filepath.Join(repoRoot, "cmd", "forma")

	bin := filepath.Join(t.TempDir(), "forma")
	build := exec.CommandContext(context.Background(), "go", "build", "-o", bin, ".")
	build.Dir = cliDir
	build.Env = append(os.Environ(), "FORMA_REPLACE="+repoRoot)
	require.NoError(t, build.Run())

	target := t.TempDir()
	cmd := exec.CommandContext(context.Background(), bin, "init", "github.com/zhulik/forma/smoke", "-d", target)
	cmd.Env = append(os.Environ(), "FORMA_REPLACE="+repoRoot)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, string(out))

	buildApp := exec.CommandContext(context.Background(), "go", "build", "-o", filepath.Join(t.TempDir(), "app"), ".")
	buildApp.Dir = target
	buildApp.Env = append(os.Environ(), "FORMA_REPLACE="+repoRoot)
	require.NoError(t, buildApp.Run())
}
