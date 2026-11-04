package codegen

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
	"github.com/zhulik/forma/fn"
)

// CodeGenerator runs in-process TSX → Go + registry codegen.
type CodeGenerator struct{}

// Provide registers the code generator for pal injection.
func Provide() pal.ServiceDef {
	return pal.Provide[core.CodeGenerator](&CodeGenerator{})
}

// Generate analyzes components and writes Go bindings plus the host registry.
func (*CodeGenerator) Generate(ctx context.Context, req core.CodegenRequest) error {
	formaUI := req.FormaUIImport
	if formaUI == "" {
		formaUI = "github.com/zhulik/forma/ui"
	}
	pkg := req.Package
	if pkg == "" {
		pkg = "ui"
	}

	widgets, err := AnalyzeComponentsDir(req.ComponentsDir)
	if err != nil {
		return fmt.Errorf("analyze components: %w", err)
	}

	if err := os.MkdirAll(req.CodegenDir, 0o755); err != nil {
		return fmt.Errorf("mkdir codegen dir: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(req.RegistryOut), 0o755); err != nil {
		return fmt.Errorf("mkdir registry dir: %w", err)
	}

	emitted := fn.Associate(widgets, func(w Widget) (string, struct{}) {
		return w.FileBase + ".gen.go", struct{}{}
	})
	for _, w := range widgets {
		fileName := w.FileBase + ".gen.go"
		path := filepath.Join(req.CodegenDir, fileName)
		//nolint:gosec // generated project sources
		if err := os.WriteFile(path, []byte(emitGoFile(w, pkg, formaUI)), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}

	//nolint:gosec // generated project sources
	if err := os.WriteFile(req.RegistryOut, []byte(emitRegistry(widgets)), 0o644); err != nil {
		return fmt.Errorf("write registry: %w", err)
	}

	entries, err := os.ReadDir(req.CodegenDir)
	if err != nil {
		return fmt.Errorf("read codegen dir: %w", err)
	}
	stale := fn.FilterMap(entries, func(e os.DirEntry) (string, bool) {
		if e.IsDir() {
			return "", false
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".gen.go") {
			return "", false
		}
		if _, ok := emitted[name]; ok {
			return "", false
		}
		return name, true
	})
	for _, name := range stale {
		if err := os.Remove(filepath.Join(req.CodegenDir, name)); err != nil {
			return fmt.Errorf("prune %s: %w", name, err)
		}
	}

	if err := gofmtDir(ctx, req.CodegenDir); err != nil {
		return fmt.Errorf("gofmt codegen output: %w", err)
	}
	return nil
}

func gofmtDir(ctx context.Context, dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	files := fn.FilterMap(entries, func(e os.DirEntry) (string, bool) {
		if e.IsDir() || filepath.Ext(e.Name()) != ".go" {
			return "", false
		}
		return filepath.Join(dir, e.Name()), true
	})
	if len(files) == 0 {
		return nil
	}
	args := append([]string{"-w"}, files...)
	cmd := exec.CommandContext(ctx, "gofmt", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
