package devcmd_test

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	devcmd "github.com/zhulik/forma/cmd/forma/internal/cli/subcommands/dev"
)

func TestRunSerializedRebuilds_DebounceAndSerialize(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		mu       sync.Mutex
		inflight int
		maxIn    int
		builds   atomic.Int32
	)

	buildStarted := make(chan struct{}, 8)
	buildGo := func(_ context.Context, _, _ string) error {
		mu.Lock()
		inflight++
		if inflight > maxIn {
			maxIn = inflight
		}
		mu.Unlock()
		defer func() {
			mu.Lock()
			inflight--
			mu.Unlock()
		}()

		builds.Add(1)
		select {
		case buildStarted <- struct{}{}:
		default:
		}
		time.Sleep(40 * time.Millisecond)
		return nil
	}

	triggers := make(chan struct{}, 8)
	done := make(chan struct{})
	go func() {
		devcmd.RunSerializedRebuilds(ctx, slog.Default(), "/proj", "/proj/web/public/app.js", buildGo, triggers, 30*time.Millisecond)
		close(done)
	}()

	// Burst of saves within debounce window → one rebuild.
	for range 5 {
		triggers <- struct{}{}
	}
	select {
	case <-buildStarted:
	case <-time.After(time.Second):
		t.Fatal("expected rebuild after debounce")
	}
	time.Sleep(80 * time.Millisecond)
	require.Equal(t, int32(1), builds.Load())

	// Trigger while busy → follow-up rebuild after first finishes.
	triggers <- struct{}{}
	<-buildStarted
	triggers <- struct{}{}
	triggers <- struct{}{}
	select {
	case <-buildStarted:
	case <-time.After(time.Second):
		t.Fatal("expected follow-up rebuild")
	}
	time.Sleep(80 * time.Millisecond)

	cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("rebuild loop did not stop")
	}

	mu.Lock()
	defer mu.Unlock()
	require.Equal(t, 1, maxIn, "overlapping rebuilds are not allowed")
	require.GreaterOrEqual(t, builds.Load(), int32(2))
}
