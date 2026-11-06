package formaconfig_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

func TestLoad_UIWeb(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, formaconfig.FileName), []byte(`module: github.com/example/app
ui:
  mount: app-root
web:
  path: frontend
`), 0o600))

	project, err := formaconfig.Load(dir)
	require.NoError(t, err)
	require.Equal(t, "github.com/example/app", project.Module)
	require.Equal(t, "app-root", project.UI.Mount)
	require.Equal(t, "frontend", project.Web.Path)
	require.Equal(t, filepath.Join(dir, "frontend"), project.WebRoot(dir))
	require.Equal(t, filepath.Join(dir, "frontend", "public", "app.js"), project.BrowserAppJSPath(dir))
}
