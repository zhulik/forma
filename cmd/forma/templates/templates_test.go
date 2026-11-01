package templates_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/cmd/forma/templates"
)

func TestApply_Minimal(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	data := templates.Data{
		Module:  "github.com/example/app",
		Package: "app",
		Mount:   "app-root",
	}
	require.NoError(t, templates.Apply(dir, "minimal", data))

	require.FileExists(t, filepath.Join(dir, "forma.yaml"))
	yaml, err := os.ReadFile(filepath.Join(dir, "forma.yaml"))
	require.NoError(t, err)
	require.Contains(t, string(yaml), "web:\n  path: web")
	require.FileExists(t, filepath.Join(dir, "main.go"))

	raw, err := os.ReadFile(filepath.Join(dir, "main.go"))
	require.NoError(t, err)
	require.Contains(t, string(raw), "UIMount(\"app-root\"")
}

func TestNames_IncludesMinimal(t *testing.T) {
	t.Parallel()

	names, err := templates.Names()
	require.NoError(t, err)
	require.Contains(t, names, "minimal")
}
