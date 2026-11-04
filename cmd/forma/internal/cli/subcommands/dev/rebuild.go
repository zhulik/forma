package devcmd

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

// runRebuildLoop watches Go sources and rebuilds the browser artifact.
// Rebuild failures after the initial build are logged; the loop keeps running.
func runRebuildLoop(
	ctx context.Context,
	logger *slog.Logger,
	projectDir, artifact string,
	buildGo GoBuilder,
	watchGo func(ctx context.Context, roots []string, onChange func()) error,
	debounce time.Duration,
) error {
	triggers := make(chan struct{}, 1)
	notify := func() {
		select {
		case triggers <- struct{}{}:
		default:
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		RunSerializedRebuilds(ctx, logger, projectDir, artifact, buildGo, triggers, debounce)
	}()

	err := watchGo(ctx, []string{projectDir}, notify)
	wg.Wait()
	return err
}

// RunSerializedRebuilds debounces triggers and never overlaps GoBuilder calls.
// Exported for unit tests with a fake builder.
func RunSerializedRebuilds(
	ctx context.Context,
	logger *slog.Logger,
	projectDir, artifact string,
	buildGo GoBuilder,
	triggers <-chan struct{},
	debounce time.Duration,
) {
	var timer *time.Timer
	stopTimer := func() {
		if timer != nil {
			timer.Stop()
			timer = nil
		}
	}
	defer stopTimer()

	var timerC <-chan time.Time
	busy := false
	pending := false

	rebuild := func() {
		busy = true
		pending = false
		logger.InfoContext(ctx, "forma dev: rebuilding GoScript artifact", "out", artifact)
		err := buildGo(ctx, projectDir, artifact)
		busy = false
		if ctx.Err() != nil {
			return
		}
		if err != nil {
			logger.ErrorContext(ctx, "forma dev: GoScript rebuild failed (still watching)", "err", err)
		} else {
			logger.InfoContext(ctx, "forma dev: GoScript rebuild ok — reload the browser page")
		}
		if pending {
			pending = false
			stopTimer()
			timer = time.NewTimer(0)
			timerC = timer.C
		}
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-triggers:
			if busy {
				pending = true
				continue
			}
			stopTimer()
			timer = time.NewTimer(debounce)
			timerC = timer.C
		case <-timerC:
			timerC = nil
			timer = nil
			if busy {
				pending = true
				continue
			}
			rebuild()
		}
	}
}
