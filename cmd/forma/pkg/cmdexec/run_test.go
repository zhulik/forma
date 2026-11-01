package cmdexec_test

import (
	"context"
	"testing"

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
