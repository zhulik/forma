package initcmd_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
	initcmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/init"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

const testModule = "github.com/example/app"

func TestRunOnlyConfig_ErrConfigExists(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, formaconfig.FileName), []byte("module: x\n"), 0o600))

	err := initcmd.Run(context.Background(), initcmd.Options{
		Module:     testModule,
		Directory:  dir,
		Template:   "minimal",
		OnlyConfig: true,
	})
	require.ErrorIs(t, err, clierrors.ErrConfigExists)
}

func TestRunFull_ErrNotEmpty(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "stray"), []byte("x"), 0o600))

	err := initcmd.Run(context.Background(), initcmd.Options{
		Module:    testModule,
		Directory: dir,
		Template:  "minimal",
	})
	require.ErrorIs(t, err, clierrors.ErrNotEmpty)
}

func TestRun_ErrModuleRequired(t *testing.T) {
	err := initcmd.Run(context.Background(), initcmd.Options{Directory: t.TempDir()})
	require.ErrorIs(t, err, clierrors.ErrModuleRequired)
}

func TestRunOnlyConfig_WritesYAML(t *testing.T) {
	dir := t.TempDir()
	err := initcmd.Run(context.Background(), initcmd.Options{
		Module:     testModule,
		Directory:  dir,
		Template:   "minimal",
		OnlyConfig: true,
	})
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(dir, formaconfig.FileName))
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	empty, err := initcmd.IsEmpty(dir)
	require.NoError(t, err)
	require.True(t, empty)

	require.NoError(t, os.WriteFile(filepath.Join(dir, "x"), []byte("a"), 0o600))
	empty, err = initcmd.IsEmpty(dir)
	require.NoError(t, err)
	require.False(t, empty)
}
