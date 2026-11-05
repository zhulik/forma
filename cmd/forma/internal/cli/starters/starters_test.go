package starters_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/cmd/forma/internal/cli/starters"
)

func TestComponents_IncludesCheckbox(t *testing.T) {
	t.Parallel()
	require.Equal(t, []string{"button", "input", "card", "checkbox"}, starters.Components)
}
