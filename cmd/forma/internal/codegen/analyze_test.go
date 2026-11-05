package codegen_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhulik/forma/cmd/forma/internal/codegen"
)

func TestAnalyzeComponentsDir_InputCheckboxPayloads(t *testing.T) {
	t.Parallel()

	dir := filepath.Join(testdataDir(t), "fixtures", "input")
	widgets, err := codegen.AnalyzeComponentsDir(dir)
	require.NoError(t, err)
	require.Len(t, widgets, 1)
	props := propMap(widgets[0].Props)
	require.Equal(t, "func(ctx ui.Context, value string)", props["OnChange"])
	require.Equal(t, "string", props["Value"])
	require.Equal(t, "string", props["Placeholder"])

	dir = filepath.Join(testdataDir(t), "fixtures", "checkbox")
	widgets, err = codegen.AnalyzeComponentsDir(dir)
	require.NoError(t, err)
	require.Len(t, widgets, 1)
	props = propMap(widgets[0].Props)
	require.Equal(t, "func(ctx ui.Context, checked bool)", props["OnCheckedChange"])
	require.Equal(t, "bool", props["Checked"])
}

func TestAnalyzeComponentsDir_ClickStaysContextOnly(t *testing.T) {
	t.Parallel()

	dir := filepath.Join(testdataDir(t), "fixtures", "button")
	widgets, err := codegen.AnalyzeComponentsDir(dir)
	require.NoError(t, err)
	require.Len(t, widgets, 1)
	props := propMap(widgets[0].Props)
	require.Equal(t, "func(ctx ui.Context)", props["OnClick"])
}

func TestAnalyzeComponentsDir_OverrideForcesPayloadCallback(t *testing.T) {
	t.Parallel()

	dir := filepath.Join(testdataDir(t), "fixtures", "override_force")
	widgets, err := codegen.AnalyzeComponentsDir(dir)
	require.NoError(t, err)
	require.Len(t, widgets, 1)
	props := propMap(widgets[0].Props)
	require.Equal(t, "func(ctx ui.Context, value string)", props["OnSelect"])
}

func TestAnalyzeComponentsDir_EmbeddedOverridesForShadcnForwardRef(t *testing.T) {
	t.Parallel()

	dir := filepath.Join(testdataDir(t), "fixtures", "shadcn_input")
	widgets, err := codegen.AnalyzeComponentsDir(dir)
	require.NoError(t, err)
	require.Len(t, widgets, 1)
	props := propMap(widgets[0].Props)
	require.Equal(t, "func(ctx ui.Context, value string)", props["OnChange"])
	require.Equal(t, "string", props["Value"])

	dir = filepath.Join(testdataDir(t), "fixtures", "shadcn_checkbox")
	widgets, err = codegen.AnalyzeComponentsDir(dir)
	require.NoError(t, err)
	require.Len(t, widgets, 1)
	props = propMap(widgets[0].Props)
	require.Equal(t, "func(ctx ui.Context, checked bool)", props["OnCheckedChange"])
	require.Equal(t, "bool", props["Checked"])
}

func TestMergeOverride_ReplacesPropType(t *testing.T) {
	t.Parallel()

	w := codegen.MergeOverride(codegen.Widget{
		Kind: "leaf",
		Name: "X",
		Props: []codegen.PropField{
			{Name: "OnSelect", GoType: "func(ctx ui.Context)"},
		},
	}, codegen.SourceOverride{
		Props: []codegen.PropOverride{
			{Name: "OnSelect", GoType: "func(ctx ui.Context, checked bool)"},
		},
	})
	require.Equal(t, "func(ctx ui.Context, checked bool)", w.Props[0].GoType)
}

func propMap(props []codegen.PropField) map[string]string {
	out := make(map[string]string, len(props))
	for _, p := range props {
		out[p.Name] = p.GoType
	}
	return out
}
