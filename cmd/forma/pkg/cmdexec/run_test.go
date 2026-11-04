package cmdexec_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/cmd/forma/pkg/cmdexec"
)

func TestRun_GoVersion(t *testing.T) {
	if testing.Short() {
		t.Skip("subprocess")
	}
	t.Parallel()

	err := cmdexec.Run(context.Background(), "", "go", "version")
	require.NoError(t, err)
}

func TestStart_CancelKillsChild(t *testing.T) {
	if testing.Short() {
		t.Skip("subprocess")
	}
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	proc, err := cmdexec.Start(ctx, "", "sleep", "30")
	require.NoError(t, err)

	cancel()

	done := make(chan error, 1)
	go func() {
		done <- proc.Wait()
	}()

	select {
	case err := <-done:
		require.Error(t, err)
	case <-time.After(5 * time.Second):
		t.Fatal("child did not exit after context cancel")
	}
}
