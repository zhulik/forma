package fn_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/fn"
)

func TestMapValues(t *testing.T) {
	t.Parallel()

	got := fn.MapValues(map[string]int{"a": 1, "b": 2}, func(n int) int { return n * 10 })
	require.Equal(t, map[string]int{"a": 10, "b": 20}, got)
}

func TestMapEntries(t *testing.T) {
	t.Parallel()

	got := fn.MapEntries(map[string]int{"a": 1}, func(k string, v int) (string, string) {
		return k + k, k + strconv.Itoa(v)
	})
	require.Equal(t, map[string]string{"aa": "a1"}, got)
}

func TestFilterEntries(t *testing.T) {
	t.Parallel()

	got := fn.FilterEntries(map[string]int{"a": 1, "b": 2, "c": 3}, func(_ string, v int) bool {
		return v%2 == 0
	})
	require.Equal(t, map[string]int{"b": 2}, got)
}

func TestFilterMapEntries(t *testing.T) {
	t.Parallel()

	got := fn.FilterMapEntries(map[string]int{"a": 1, "b": 2}, func(k string, v int) (string, int, bool) {
		if v%2 != 0 {
			return "", 0, false
		}

		return k, v * 10, true
	})
	require.Equal(t, map[string]int{"b": 20}, got)
}

func TestMapValuesErr_FirstError(t *testing.T) {
	t.Parallel()

	got, err := fn.MapValuesErr(map[string]int{"a": 1, "b": 2}, func(n int) (int, error) {
		if n == 2 {
			return 0, errBoom
		}

		return n, nil
	})
	require.ErrorIs(t, err, errBoom)
	require.Nil(t, got)
}

func TestMapEntriesErr_Success(t *testing.T) {
	t.Parallel()

	got, err := fn.MapEntriesErr(map[string]int{"x": 3}, func(k string, v int) (string, int, error) {
		return k, v + 1, nil
	})
	require.NoError(t, err)
	require.Equal(t, map[string]int{"x": 4}, got)
}

func TestFilterMapEntriesErr_FirstError(t *testing.T) {
	t.Parallel()

	got, err := fn.FilterMapEntriesErr(map[string]int{"a": 1}, func(_ string, _ int) (string, int, bool, error) {
		return "", 0, false, errBoom
	})
	require.ErrorIs(t, err, errBoom)
	require.Nil(t, got)
}
