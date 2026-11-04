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

func TestUniqSorted(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"a", "b", "c"}, fn.UniqSorted([]string{"c", "a", "b", "a", "c"}))
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

func TestFlatMap(t *testing.T) {
	t.Parallel()

	got := fn.FlatMap([]string{"ab", "c"}, func(s string) []rune {
		return []rune(s)
	})
	require.Equal(t, []rune{'a', 'b', 'c'}, got)
}

func TestGroupBy(t *testing.T) {
	t.Parallel()

	got := fn.GroupBy([]string{"aa", "b", "cc", "dd"}, func(s string) int {
		return len(s)
	})
	require.Equal(t, []string{"aa", "cc", "dd"}, got[2])
	require.Equal(t, []string{"b"}, got[1])
}
