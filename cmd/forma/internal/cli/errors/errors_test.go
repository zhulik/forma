package clierrors_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
)

func TestSentinelsDistinct(t *testing.T) {
	t.Parallel()

	require.NotErrorIs(t, clierrors.ErrNotEmpty, clierrors.ErrConfigExists)
	require.ErrorIs(t, clierrors.ErrNotEmpty, clierrors.ErrNotEmpty)
}
