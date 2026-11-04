package devcmd

import (
	"context"
	"fmt"

	"github.com/cortesi/moddwatch"
	"golang.org/x/sync/errgroup"
)

// Patterns for Go sources under the app project. web/ and other trees are
// excluded so Vite owns host assets and we never walk node_modules.
var (
	goWatchIncludes = []string{"**/*.go"}
	goWatchExcludes = []string{
		"**/.git/**",
		"**/node_modules/**",
		"**/vendor/**",
		"**/dist/**",
		"**/.idea/**",
		"**/.vscode/**",
		"**/web/**",
		"**/.goscript/**",
	}
)

func watchGoSources(ctx context.Context, roots []string, onChange func()) error {
	if len(roots) == 0 {
		return fmt.Errorf("watch: no roots")
	}

	g, gctx := errgroup.WithContext(ctx)
	for _, root := range roots {
		g.Go(func() error {
			return watchOneRoot(gctx, root, onChange)
		})
	}
	return g.Wait()
}

func watchOneRoot(ctx context.Context, root string, onChange func()) error {
	ch := make(chan *moddwatch.Mod, 8)
	watcher, err := moddwatch.Watch(root, goWatchIncludes, goWatchExcludes, rebuildDebounce, ch)
	if err != nil {
		return fmt.Errorf("watch %s: %w", root, err)
	}
	defer watcher.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case mod, ok := <-ch:
			if !ok {
				return nil
			}
			if mod == nil || mod.Empty() {
				continue
			}
			onChange()
		}
	}
}
