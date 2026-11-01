package templates

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/zhulik/forma/fn"
)

//go:embed all:scaffolds
var scaffoldsFS embed.FS

// FS is the project template tree (one directory per template name).
var FS = mustSub(scaffoldsFS, "scaffolds")

func mustSub(fsys embed.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}
	return sub
}

// Data is available to template paths and file contents.
type Data struct {
	Package string
	Module  string
	Mount   string
}

// Exists reports whether name is a known template directory.
func Exists(name string) error {
	if name == "" {
		return fmt.Errorf("template name is required")
	}
	info, err := fs.Stat(FS, name)
	if err != nil {
		return fmt.Errorf("unknown template %q: %w", name, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("unknown template %q: not a template directory", name)
	}
	return nil
}

// Apply renders the named template into dst.
func Apply(dst, name string, data Data) error {
	if err := Exists(name); err != nil {
		return err
	}

	return fs.WalkDir(FS, name, func(embedPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(name, embedPath)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}

		renderedRel, err := execute(filepath.ToSlash(rel), data)
		if err != nil {
			return fmt.Errorf("render path %q: %w", rel, err)
		}
		if !d.IsDir() {
			renderedRel = strings.TrimSuffix(renderedRel, ".tmpl")
		}
		outPath := filepath.Join(dst, filepath.FromSlash(renderedRel))

		if d.IsDir() {
			return os.MkdirAll(outPath, 0o755)
		}

		raw, err := fs.ReadFile(FS, embedPath)
		if err != nil {
			return err
		}
		rendered, err := execute(string(raw), data)
		if err != nil {
			return fmt.Errorf("render file %q: %w", embedPath, err)
		}
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}
		//nolint:gosec // scaffold files are intentionally world-readable
		return os.WriteFile(outPath, []byte(rendered), 0o644)
	})
}

// RenderFile renders a single template file (path under template dir) to a string.
func RenderFile(name, relPath string, data Data) (string, error) {
	if err := Exists(name); err != nil {
		return "", err
	}
	embedPath := path.Join(name, relPath)
	raw, err := fs.ReadFile(FS, embedPath)
	if err != nil {
		return "", fmt.Errorf("read template %q: %w", embedPath, err)
	}
	return execute(string(raw), data)
}

// Names returns embedded template directory names.
func Names() ([]string, error) {
	entries, err := fs.ReadDir(FS, ".")
	if err != nil {
		return nil, err
	}
	return fn.FilterMap(entries, func(e fs.DirEntry) (string, bool) {
		return e.Name(), e.IsDir()
	}), nil
}

// PackageName derives the Go package name from a module path.
func PackageName(module string) string {
	return strings.ReplaceAll(path.Base(module), "-", "_")
}

func execute(text string, data Data) (string, error) {
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
