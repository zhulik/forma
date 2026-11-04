package ident_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/cmd/forma/pkg/ident"
)

func TestToPascal(t *testing.T) {
	t.Parallel()

	require.Equal(t, "Button", ident.ToPascal("button"))
	require.Equal(t, "Accordion", ident.ToPascal("accordion"))
	require.Equal(t, "AlertDialog", ident.ToPascal("alert-dialog"))
	require.Equal(t, "FooBar", ident.ToPascal("foo_bar"))
	require.Equal(t, "", ident.ToPascal(""))
}

func TestToExportedGoName(t *testing.T) {
	t.Parallel()

	require.Equal(t, "OnClick", ident.ToExportedGoName("onClick"))
	require.Equal(t, "Variant", ident.ToExportedGoName("variant"))
	require.Equal(t, "Size", ident.ToExportedGoName("size"))
	require.Equal(t, "", ident.ToExportedGoName(""))
}

func TestReceiverName(t *testing.T) {
	t.Parallel()

	require.Equal(t, "b", ident.ReceiverName("Button"))
	require.Equal(t, "a", ident.ReceiverName("Accordion"))
	require.Equal(t, "v", ident.ReceiverName(""))
	require.Equal(t, "v", ident.ReceiverName("button"))
}
