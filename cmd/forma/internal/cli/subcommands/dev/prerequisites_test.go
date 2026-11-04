package devcmd_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
	devcmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/dev"
)

func TestCheckPrerequisites_MissingYAML(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	_, _, err := devcmd.CheckPrerequisites(dir, lookPathOK)
	require.Error(t, err)
	require.ErrorIs(t, err, os.ErrNotExist)
	require.Contains(t, err.Error(), "forma.yaml")
}

func TestCheckPrerequisites_MissingNodeModules(t *testing.T) {
	t.Parallel()

	dir := writeMinimalProject(t, false)
	_, _, err := devcmd.CheckPrerequisites(dir, lookPathOK)
	require.ErrorIs(t, err, devcmd.ErrNodeModulesMissing)
}

func TestCheckPrerequisites_MissingGopherJS(t *testing.T) {
	t.Parallel()

	dir := writeMinimalProject(t, true)
	_, _, err := devcmd.CheckPrerequisites(dir, func(file string) (string, error) {
		if file == "gopherjs" {
			return "", errors.New("not found")
		}
		return "/bin/" + file, nil
	})
	require.ErrorIs(t, err, devcmd.ErrGopherJSMissing)
}

func TestCheckPrerequisites_MissingPackageJSON(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "forma.yaml"), []byte("module: m\nweb:\n  path: web\n"), 0o600))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "web"), 0o755))

	_, _, err := devcmd.CheckPrerequisites(dir, lookPathOK)
	require.ErrorIs(t, err, clierrors.ErrWebPackageMissing)
}

func TestCheckPrerequisites_OK(t *testing.T) {
	t.Parallel()

	dir := writeMinimalProject(t, true)
	project, webRoot, err := devcmd.CheckPrerequisites(dir, lookPathOK)
	require.NoError(t, err)
	require.Equal(t, "web", project.WebPath())
	require.Equal(t, filepath.Join(dir, "web"), webRoot)
}

func lookPathOK(file string) (string, error) {
	return "/usr/bin/" + file, nil
}

func writeMinimalProject(t *testing.T, withNodeModules bool) string {
	t.Helper()
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "forma.yaml"), []byte("module: example.com/app\nweb:\n  path: web\n"), 0o600))
	web := filepath.Join(dir, "web")
	require.NoError(t, os.MkdirAll(web, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(web, "package.json"), []byte(`{"name":"t","private":true}`), 0o600))
	if withNodeModules {
		require.NoError(t, os.MkdirAll(filepath.Join(web, "node_modules"), 0o755))
	}
	return dir
}
