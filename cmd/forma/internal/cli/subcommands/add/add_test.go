package addcmd_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	formacli "github.com/zhulik/forma/cmd/forma/internal/cli"
	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
)

func TestAdd_RequiresComponentName(t *testing.T) {
	t.Parallel()

	cmd := formacli.NewRoot()
	err := cmd.Run(t.Context(), []string{"forma", "add"})
	require.ErrorIs(t, err, clierrors.ErrComponentNameRequired)
}
