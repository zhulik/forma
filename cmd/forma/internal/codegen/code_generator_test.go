package codegen_test

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/cmd/forma/internal/codegen"
	"github.com/zhulik/forma/cmd/forma/internal/core"
)

func TestCodeGenerator_Generate_ButtonFixture(t *testing.T) {
	t.Parallel()

	fixture := filepath.Join(testdataDir(t), "fixtures", "button")
	out := t.TempDir()
	codegenDir := filepath.Join(out, "ui")
	registry := filepath.Join(out, "registry.gen.ts")

	err := (&codegen.CodeGenerator{}).Generate(context.Background(), core.CodegenRequest{
		ComponentsDir: fixture,
		CodegenDir:    codegenDir,
		Package:       "ui",
		RegistryOut:   registry,
	})
	require.NoError(t, err)

	got, err := os.ReadFile(filepath.Join(codegenDir, "button.gen.go"))
	require.NoError(t, err)
	want, err := os.ReadFile(filepath.Join(testdataDir(t), "goldens", "button", "button.gen.go"))
	require.NoError(t, err)
	require.Equal(t, string(want), string(got))

	reg, err := os.ReadFile(registry)
	require.NoError(t, err)
	wantReg, err := os.ReadFile(filepath.Join(testdataDir(t), "goldens", "button", "registry.gen.ts"))
	require.NoError(t, err)
	require.Equal(t, string(wantReg), string(reg))
}

func TestCodeGenerator_Generate_AccordionFixture(t *testing.T) {
	t.Parallel()

	fixture := filepath.Join(testdataDir(t), "fixtures", "accordion")
	out := t.TempDir()
	codegenDir := filepath.Join(out, "ui")

	err := (&codegen.CodeGenerator{}).Generate(context.Background(), core.CodegenRequest{
		ComponentsDir: fixture,
		CodegenDir:    codegenDir,
		Package:       "ui",
		RegistryOut:   filepath.Join(out, "registry.gen.ts"),
	})
	require.NoError(t, err)

	got, err := os.ReadFile(filepath.Join(codegenDir, "accordion.gen.go"))
	require.NoError(t, err)
	want, err := os.ReadFile(filepath.Join(testdataDir(t), "goldens", "accordion", "accordion.gen.go"))
	require.NoError(t, err)
	require.Equal(t, string(want), string(got))
}

func TestCodeGenerator_Generate_PruneOrphan(t *testing.T) {
	t.Parallel()

	fixture := filepath.Join(testdataDir(t), "fixtures", "button")
	out := t.TempDir()
	codegenDir := filepath.Join(out, "ui")
	require.NoError(t, os.MkdirAll(codegenDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(codegenDir, "orphan.gen.go"), []byte("package ui\n"), 0o600))

	err := (&codegen.CodeGenerator{}).Generate(context.Background(), core.CodegenRequest{
		ComponentsDir: fixture,
		CodegenDir:    codegenDir,
		Package:       "ui",
		RegistryOut:   filepath.Join(out, "registry.gen.ts"),
	})
	require.NoError(t, err)
	_, err = os.Stat(filepath.Join(codegenDir, "orphan.gen.go"))
	require.True(t, os.IsNotExist(err))
	_, err = os.Stat(filepath.Join(codegenDir, "button.gen.go"))
	require.NoError(t, err)
}

func testdataDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok)
	return filepath.Join(filepath.Dir(file), "testdata")
}
