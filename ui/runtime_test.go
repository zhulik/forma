package ui_test

import (
	"context"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/zhulik/forma/ui"
)

type recordingSink struct {
	trees []*Node
}

func (s *recordingSink) SetTree(tree *Node) {
	s.trees = append(s.trees, tree)
}

type staticRoot struct {
	label string
}

func (s staticRoot) Render(context.Context) *Node {
	return &Node{
		Type:  textType,
		Props: Text{Text: s.label},
	}
}

type clickRoot struct {
	onClick func(Context)
}

func (c clickRoot) Render(context.Context) *Node {
	return &Node{
		Type: "Button",
		Props: Button{
			OnClick: c.onClick,
		},
	}
}

type valueRoot struct {
	onChange  func(Context, string)
	onChecked func(Context, bool)
}

func (v valueRoot) Render(context.Context) *Node {
	return &Node{
		Type: "Column",
		Children: []*Node{
			{
				Type: "Input",
				Props: Input{
					OnChange: v.onChange,
				},
			},
			{
				Type: "Checkbox",
				Props: Checkbox{
					OnCheckedChange: v.onChecked,
				},
			},
		},
	}
}

const textType = "Text"

func TestUpdateSuccessfulRefresh(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	label := "a"
	root := &mutatingRoot{label: &label}
	ctx, unmount := MountWithSink(sink, root)
	t.Cleanup(unmount)

	require.Len(t, sink.trees, 1)
	require.Equal(t, "a", textOf(sink.trees[0]))

	ctx.Update(func() {
		label = "b"
	})
	require.Len(t, sink.trees, 2)
	require.Equal(t, "b", textOf(sink.trees[1]))
}

func TestUpdatePanicSkipsRefresh(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	label := "a"
	ctx, unmount := MountWithSink(sink, &mutatingRoot{label: &label})
	t.Cleanup(unmount)
	require.Len(t, sink.trees, 1)

	require.Panics(t, func() {
		ctx.Update(func() {
			label = "nope"
			panic("boom")
		})
	})

	require.Len(t, sink.trees, 1)
	require.Equal(t, "a", textOf(sink.trees[0]))
}

func TestManualRefreshUI(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	label := "a"
	ctx, unmount := MountWithSink(sink, &mutatingRoot{label: &label})
	t.Cleanup(unmount)

	label = "c"
	ctx.RefreshUI()
	require.Len(t, sink.trees, 2)
	require.Equal(t, "c", textOf(sink.trees[1]))
}

func TestRefreshCancelsPriorCallbackContext(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	var captured Context
	root := clickRoot{
		onClick: func(ctx Context) {
			captured = ctx
		},
	}
	ctx, unmount := MountWithSink(sink, root)
	t.Cleanup(unmount)
	require.Len(t, sink.trees, 1)

	invokeClick(t, sink.trees[0])
	require.NotNil(t, captured)
	require.NoError(t, captured.Err())

	ctx.RefreshUI()
	require.ErrorIs(t, captured.Err(), context.Canceled)
}

func TestUnmountCancelsCallbackContext(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	var captured Context
	_, unmount := MountWithSink(sink, clickRoot{
		onClick: func(ctx Context) {
			captured = ctx
		},
	})
	invokeClick(t, sink.trees[0])
	require.NoError(t, captured.Err())

	unmount()
	require.ErrorIs(t, captured.Err(), context.Canceled)
}

func TestBindClickStyleCallback(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	var called atomic.Bool
	_, unmount := MountWithSink(sink, clickRoot{
		onClick: func(ctx Context) {
			require.NotNil(t, ctx)
			called.Store(true)
			ctx.RefreshUI()
		},
	})
	t.Cleanup(unmount)

	invokeClick(t, sink.trees[0])
	require.True(t, called.Load())
	require.Len(t, sink.trees, 2)
}

func TestBindValueStyleCallbacks(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	var gotString string
	var gotBool bool
	_, unmount := MountWithSink(sink, valueRoot{
		onChange: func(ctx Context, val string) {
			gotString = val
			_ = ctx
		},
		onChecked: func(ctx Context, checked bool) {
			gotBool = checked
			_ = ctx
		},
	})
	t.Cleanup(unmount)

	props0, ok := sink.trees[0].Children[0].Props.(map[string]any)
	require.True(t, ok)
	onChange, ok := props0["OnChange"].(func(string))
	require.True(t, ok)
	onChange("hello")
	require.Equal(t, "hello", gotString)

	props1, ok := sink.trees[0].Children[1].Props.(map[string]any)
	require.True(t, ok)
	onChecked, ok := props1["OnCheckedChange"].(func(bool))
	require.True(t, ok)
	onChecked(true)
	require.True(t, gotBool)
}

func TestNativeUIMountNoop(t *testing.T) {
	t.Parallel()
	require.NotPanics(t, func() {
		UIMount("app-root", staticRoot{label: "x"})
	})
}

type mutatingRoot struct {
	label *string
}

func (m *mutatingRoot) Render(context.Context) *Node {
	return &Node{
		Type:  textType,
		Props: Text{Text: *m.label},
	}
}

func textOf(node *Node) string {
	props, ok := node.Props.(map[string]any)
	if !ok {
		return ""
	}
	text, _ := props["Text"].(string)

	return text
}

func invokeClick(t *testing.T, node *Node) {
	t.Helper()
	props, ok := node.Props.(map[string]any)
	require.True(t, ok)
	onClick, ok := props["OnClick"].(func())
	require.True(t, ok)
	onClick()
}
