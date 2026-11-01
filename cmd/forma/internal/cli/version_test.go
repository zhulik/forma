package cli_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	formacli "github.com/zhulik/forma/cmd/forma/internal/cli"
	"github.com/zhulik/forma/cmd/forma/internal/version"
)

func TestVersion_FlagAndSubcommandMatch(t *testing.T) {
	t.Parallel()

	want := "forma version " + version.String() + "\n"

	for _, args := range [][]string{
		{"forma", "--version"},
		{"forma", "-v"},
		{"forma", "version"},
	} {
		t.Run(strings.Join(args[1:], " "), func(t *testing.T) {
			t.Parallel()

			var out bytes.Buffer
			cmd := formacli.NewRoot()
			cmd.Writer = &out

			err := cmd.Run(context.Background(), args)
			require.NoError(t, err)
			require.Equal(t, want, out.String())
		})
	}
}
