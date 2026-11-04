package fn_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/fn"
)

var errBoom = errors.New("boom")

func TestMapErr_Success(t *testing.T) {
	t.Parallel()

	got, err := fn.MapErr([]int{1, 2, 3}, func(n int) (int, error) {
		return n * 2, nil
	})
	require.NoError(t, err)
	require.Equal(t, []int{2, 4, 6}, got)
}

func TestMapErr_FirstError(t *testing.T) {
	t.Parallel()

	got, err := fn.MapErr([]int{1, 2, 3}, func(n int) (int, error) {
		if n == 2 {
			return 0, errBoom
		}

		return n, nil
	})
	require.ErrorIs(t, err, errBoom)
	require.Nil(t, got)
}

func TestFilterMapErr_Success(t *testing.T) {
	t.Parallel()

	got, err := fn.FilterMapErr([]int{1, 2, 3, 4}, func(n int) (int, bool, error) {
		if n%2 == 0 {
			return n * 10, true, nil
		}

		return 0, false, nil
	})
	require.NoError(t, err)
	require.Equal(t, []int{20, 40}, got)
}

func TestFilterMapErr_FirstError(t *testing.T) {
	t.Parallel()

	got, err := fn.FilterMapErr([]string{"a", "b", "c"}, func(s string) (string, bool, error) {
		if s == "b" {
			return "", false, errBoom
		}

		return s, true, nil
	})
	require.ErrorIs(t, err, errBoom)
	require.Nil(t, got)
}
