package devcmd

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func watchGoSources(ctx context.Context, roots []string, onChange func()) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	for _, root := range roots {
		if err := addWatchTree(watcher, root); err != nil {
			return err
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-watcher.Errors:
			if err != nil {
				return err
			}
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if shouldIgnoreWatchPath(event.Name) {
				continue
			}
			if event.Has(fsnotify.Create) {
				if info, err := os.Stat(event.Name); err == nil && info.IsDir() {
					_ = addWatchTree(watcher, event.Name)
				}
			}
			if !isGoSourceEvent(event) {
				continue
			}
			onChange()
		}
	}
}

func addWatchTree(watcher *fsnotify.Watcher, root string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		if shouldIgnoreWatchPath(path) {
			return filepath.SkipDir
		}
		return watcher.Add(path)
	})
}

func shouldIgnoreWatchPath(path string) bool {
	base := filepath.Base(path)
	switch base {
	case ".git", "node_modules", "vendor", "dist", ".idea", ".vscode":
		return true
	}
	// Ignore the web tree except we still skip walking into it from project root.
	if base == "web" {
		return true
	}
	return false
}

func isGoSourceEvent(event fsnotify.Event) bool {
	if event.Has(fsnotify.Chmod) && !event.Has(fsnotify.Write) && !event.Has(fsnotify.Create) && !event.Has(fsnotify.Rename) && !event.Has(fsnotify.Remove) {
		return false
	}
	name := event.Name
	if strings.HasSuffix(name, "~") || strings.HasSuffix(name, ".swp") {
		return false
	}
	return strings.HasSuffix(name, ".go")
}
