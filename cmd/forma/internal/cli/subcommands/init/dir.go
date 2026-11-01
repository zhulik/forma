package initcmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// IsEmpty reports whether dir contains no entries.
func IsEmpty(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, fmt.Errorf("read directory: %w", err)
	}
	return len(entries) == 0, nil
}

func resolveTarget(directory string) (target string, err error) {
	if directory == "" {
		return os.Getwd()
	}
	abs, err := filepath.Abs(directory)
	if err != nil {
		return "", fmt.Errorf("resolve directory: %w", err)
	}
	info, err := os.Stat(abs)
	if err == nil {
		if !info.IsDir() {
			return "", fmt.Errorf("%s: not a directory", abs)
		}
		return abs, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("stat directory: %w", err)
	}
	if err := os.MkdirAll(abs, 0o755); err != nil {
		return "", fmt.Errorf("create directory %s: %w", abs, err)
	}
	return abs, nil
}
