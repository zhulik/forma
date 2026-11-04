package devcmd_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	formacli "github.com/zhulik/forma/cmd/forma/internal/cli"
)

func TestDevHelp_NotPlannedOnly(t *testing.T) {
	t.Parallel()

	var out bytes.Buffer
	cmd := formacli.NewRoot()
	cmd.Writer = &out
	err := cmd.Run(context.Background(), []string{"forma", "dev", "--help"})
	require.NoError(t, err)
	require.Contains(t, out.String(), "Vite host + GopherJS rebuild")
	require.NotContains(t, out.String(), "planned")
}
