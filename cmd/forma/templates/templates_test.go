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
	require.NotContains(t, string(yaml), "components:")
	require.FileExists(t, filepath.Join(dir, "main.go"))
	require.FileExists(t, filepath.Join(dir, "generator.go"))
	require.FileExists(t, filepath.Join(dir, "Taskfile.yaml"))
	require.FileExists(t, filepath.Join(dir, "web", "package.json"))
	require.FileExists(t, filepath.Join(dir, "web", "index.html"))
	require.FileExists(t, filepath.Join(dir, "web", "src", "main.tsx"))

	html, err := os.ReadFile(filepath.Join(dir, "web", "index.html"))
	require.NoError(t, err)
	require.Contains(t, string(html), `id="app-root"`)

	raw, err := os.ReadFile(filepath.Join(dir, "main.go"))
	require.NoError(t, err)
	require.Contains(t, string(raw), "UIMount(\"app-root\"")

	gen, err := os.ReadFile(filepath.Join(dir, "generator.go"))
	require.NoError(t, err)
	require.Contains(t, string(gen), "//go:generate forma generate")

	taskfile, err := os.ReadFile(filepath.Join(dir, "Taskfile.yaml"))
	require.NoError(t, err)
	require.Contains(t, string(taskfile), "check-generate:")
	require.Contains(t, string(taskfile), "git diff --exit-code")

	mainTSX, err := os.ReadFile(filepath.Join(dir, "web", "src", "main.tsx"))
	require.NoError(t, err)
	require.Contains(t, string(mainTSX), `from "./registry.gen"`)
	require.Contains(t, string(mainTSX), `FormaHost`)
	require.Contains(t, string(mainTSX), `mergeRegistry`)

	require.FileExists(t, filepath.Join(dir, "web", "src", "host", "index.ts"))
	require.FileExists(t, filepath.Join(dir, "web", "src", "registry.custom.ts"))

	pkgJSON, err := os.ReadFile(filepath.Join(dir, "web", "package.json"))
	require.NoError(t, err)
	require.Contains(t, string(pkgJSON), `"test"`)
	require.Contains(t, string(pkgJSON), "vitest")
}

func TestNames_IncludesMinimal(t *testing.T) {
	t.Parallel()

	names, err := templates.Names()
	require.NoError(t, err)
	require.Contains(t, names, "minimal")
}
