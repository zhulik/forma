package fn_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/fn"
)

func TestUniq_PreservesOrder(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"a", "b", "c"}, fn.Uniq([]string{"a", "b", "a", "c", "b"}))
}

func TestAssociate_LastKeyWins(t *testing.T) {
	t.Parallel()

	byKey := fn.Associate([]struct{ k, v string }{
		{"a", "1"},
		{"a", "2"},
	}, func(p struct{ k, v string }) (string, string) {
		return p.k, p.v
	})
	require.Equal(t, "2", byKey["a"])
}
