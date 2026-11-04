package devcmd_test

import (
	"context"
	"log/slog"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	devcmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/dev"
)

func TestRunner_StartsViteWithCorrectArgs(t *testing.T) {
	t.Parallel()

	dir := writeMinimalProject(t, true)
	webRoot := filepath.Join(dir, "web")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		viteDir  string
		viteName string
		viteArgs []string
		builds   atomic.Int32
	)

	viteStarted := make(chan struct{})
	runner := &devcmd.Runner{
		Dir:    dir,
		Logger: slog.Default(),
		LookPath: func(file string) (string, error) {
			return "/usr/bin/" + file, nil
		},
		BuildGo: func(context.Context, string, string) error {
			builds.Add(1)
			return nil
		},
		RunLong: func(ctx context.Context, d, name string, args ...string) error {
			viteDir = d
			viteName = name
			viteArgs = append([]string{}, args...)
			close(viteStarted)
			<-ctx.Done()
			return ctx.Err()
		},
		WatchGo: func(ctx context.Context, _ []string, _ func()) error {
			<-ctx.Done()
			return nil
		},
	}

	done := make(chan error, 1)
	go func() {
		done <- runner.Run(ctx)
	}()

	select {
	case <-viteStarted:
	case <-time.After(2 * time.Second):
		t.Fatal("vite did not start")
	}

	require.Equal(t, webRoot, viteDir)
	require.Equal(t, "npm", viteName)
	require.Equal(t, []string{"run", "dev"}, viteArgs)
	require.Equal(t, int32(1), builds.Load(), "initial build must run once")

	cancel()
	select {
	case err := <-done:
		require.NoError(t, err)
	case <-time.After(2 * time.Second):
		t.Fatal("runner did not stop on cancel")
	}
}

func TestRunner_FailFastDoesNotStartVite(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	var viteCalls atomic.Int32
	runner := &devcmd.Runner{
		Dir:    dir,
		Logger: slog.Default(),
		LookPath: func(string) (string, error) {
			return "/usr/bin/x", nil
		},
		RunLong: func(context.Context, string, string, ...string) error {
			viteCalls.Add(1)
			return nil
		},
		BuildGo: func(context.Context, string, string) error {
			t.Fatal("build must not run without prerequisites")
			return nil
		},
	}

	err := runner.Run(context.Background())
	require.Error(t, err)
	require.Contains(t, err.Error(), "forma.yaml")
	require.Equal(t, int32(0), viteCalls.Load())
}
