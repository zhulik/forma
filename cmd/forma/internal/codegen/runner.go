package codegen

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Request holds paths for a codegen run derived from forma.yaml.
type Request struct {
	ComponentsDir string
	CodegenDir    string
	Package       string
	RegistryOut   string
	FormaUIImport string
}

// Runner runs in-process TSX → Go + registry codegen.
type Runner struct{}

// Generate analyzes components and writes Go bindings plus the host registry.
func (r *Runner) Generate(ctx context.Context, req Request) error {
	_ = ctx
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

	emitted := map[string]struct{}{}
	for _, w := range widgets {
		fileName := w.FileBase + ".gen.go"
		path := filepath.Join(req.CodegenDir, fileName)
		//nolint:gosec // generated project sources
		if err := os.WriteFile(path, []byte(emitGoFile(w, pkg, formaUI)), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
		emitted[fileName] = struct{}{}
	}

	//nolint:gosec // generated project sources
	if err := os.WriteFile(req.RegistryOut, []byte(emitRegistry(widgets)), 0o644); err != nil {
		return fmt.Errorf("write registry: %w", err)
	}

	entries, err := os.ReadDir(req.CodegenDir)
	if err != nil {
		return fmt.Errorf("read codegen dir: %w", err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !hasSuffix(name, ".gen.go") {
			continue
		}
		if _, ok := emitted[name]; ok {
			continue
		}
		if err := os.Remove(filepath.Join(req.CodegenDir, name)); err != nil {
			return fmt.Errorf("prune %s: %w", name, err)
		}
	}

	if err := gofmtDir(ctx, req.CodegenDir); err != nil {
		return fmt.Errorf("gofmt codegen output: %w", err)
	}
	return nil
}

func hasSuffix(name, suffix string) bool {
	return len(name) >= len(suffix) && name[len(name)-len(suffix):] == suffix
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
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if filepath.Ext(e.Name()) != ".go" {
			continue
		}
		files = append(files, filepath.Join(dir, e.Name()))
	}
	if len(files) == 0 {
		return nil
	}
	args := append([]string{"-w"}, files...)
	cmd := exec.CommandContext(ctx, "gofmt", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
