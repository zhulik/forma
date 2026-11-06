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
	require.NotContains(t, string(yaml), "codegen:")
	require.NotContains(t, string(yaml), "components:")
	require.FileExists(t, filepath.Join(dir, "main.go"))
	require.NoFileExists(t, filepath.Join(dir, "generator.go"))
	require.FileExists(t, filepath.Join(dir, "Taskfile.yaml"))
	require.FileExists(t, filepath.Join(dir, "web", "package.json"))
	require.FileExists(t, filepath.Join(dir, "web", "index.html"))
	require.FileExists(t, filepath.Join(dir, "web", "src", "main.tsx"))
	require.NoFileExists(t, filepath.Join(dir, "web", "components.json"))
	require.NoFileExists(t, filepath.Join(dir, "web", "tailwind.config.js"))

	html, err := os.ReadFile(filepath.Join(dir, "web", "index.html"))
	require.NoError(t, err)
	require.Contains(t, string(html), `id="app-root"`)
	require.Contains(t, string(html), "cdn.0build.dev/0build.css")
	require.Contains(t, string(html), "cdn.0build.dev/0build.js")
	require.Contains(t, string(html), "z-not-ready")

	raw, err := os.ReadFile(filepath.Join(dir, "main.go"))
	require.NoError(t, err)
	require.Contains(t, string(raw), "UIMount(\"app-root\"")

	taskfile, err := os.ReadFile(filepath.Join(dir, "Taskfile.yaml"))
	require.NoError(t, err)
	require.NotContains(t, string(taskfile), "check-generate:")
	require.NotContains(t, string(taskfile), "forma generate")

	mainTSX, err := os.ReadFile(filepath.Join(dir, "web", "src", "main.tsx"))
	require.NoError(t, err)
	require.NotContains(t, string(mainTSX), `registry.gen`)
	require.Contains(t, string(mainTSX), `registry.custom`)
	require.Contains(t, string(mainTSX), `FormaBridgeHost`)
	require.Contains(t, string(mainTSX), `registerFormaHost`)
	require.Contains(t, string(mainTSX), `mergeRegistry`)
	require.Contains(t, string(mainTSX), `/app.js`)
	require.Contains(t, string(mainTSX), `goBundle.type = "module"`)

	require.FileExists(t, filepath.Join(dir, "web", "src", "host", "index.ts"))
	require.FileExists(t, filepath.Join(dir, "web", "src", "host", "widgets", "button.tsx"))
	require.FileExists(t, filepath.Join(dir, "web", "src", "registry.custom.ts"))
	require.FileExists(t, filepath.Join(dir, "web", "goscript-entry.ts"))
	require.FileExists(t, filepath.Join(dir, "web", "vite.goscript.config.ts"))
	require.DirExists(t, filepath.Join(dir, "web", "public"))

	entry, err := os.ReadFile(filepath.Join(dir, "web", "goscript-entry.ts"))
	require.NoError(t, err)
	require.Contains(t, string(entry), `@goscript/github.com/example/app/main.gs.ts`)

	pkgJSON, err := os.ReadFile(filepath.Join(dir, "web", "package.json"))
	require.NoError(t, err)
	require.Contains(t, string(pkgJSON), `"test"`)
	require.Contains(t, string(pkgJSON), "vitest")
	require.Contains(t, string(pkgJSON), `"build:go"`)
	require.NotContains(t, string(pkgJSON), "tailwind")
	require.NotContains(t, string(pkgJSON), "tailwind-merge")
	require.NotContains(t, string(pkgJSON), "class-variance-authority")
}

func TestNames_IncludesMinimal(t *testing.T) {
	t.Parallel()

	names, err := templates.Names()
	require.NoError(t, err)
	require.Contains(t, names, "minimal")
}
