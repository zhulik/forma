package ui_test

import (
	"context"
	"testing"

	"github.com/zhulik/forma/ui"
)

// compoundProbe mimics generated compound Render glue (not a catalog widget).
type compoundProbe struct {
	ui.WidgetMeta

	Trigger ui.Component
	Content ui.Component
}

func (p compoundProbe) Render(ctx context.Context) *ui.Node {
	slots := map[string][]*ui.Node{}
	ui.PutSlot(slots, ui.SlotKeyFromField("Trigger"), ui.RenderComponentSlot(ctx, p.Trigger))
	ui.PutSlot(slots, ui.SlotKeyFromField("Content"), ui.RenderComponentSlot(ctx, p.Content))

	if len(slots) == 0 {
		slots = nil
	}

	return &ui.Node{
		Type:  "CompoundProbe",
		Key:   p.Key,
		Class: p.Class,
		Props: p,
		Slots: slots,
	}
}

func TestCompoundProbeRenderSlots(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	node := compoundProbe{
		Trigger: ui.Button{Variant: "default"}.With(ui.Text{Text: "Open"}),
		Content: ui.Column{}.With(ui.Text{Text: "Body"}),
	}.Render(ctx)

	if len(node.Children) != 0 {
		t.Fatalf("Children: got %d want 0", len(node.Children))
	}

	trigger, ok := node.Slots["trigger"]
	if !ok || len(trigger) != 1 || trigger[0].Type != "Button" {
		t.Fatalf("trigger slot: %#v", node.Slots["trigger"])
	}

	content, ok := node.Slots["content"]
	if !ok || len(content) != 1 || content[0].Type != "Column" {
		t.Fatalf("content slot: %#v", node.Slots["content"])
	}
}

func TestCompoundProbeOmitsEmptySlot(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	node := compoundProbe{
		Content: ui.Text{Text: "Only body"},
	}.Render(ctx)

	if _, ok := node.Slots["trigger"]; ok {
		t.Fatal("expected trigger slot omitted")
	}

	if len(node.Slots["content"]) != 1 {
		t.Fatalf("content slot: %#v", node.Slots["content"])
	}
}

func TestRenderComponentSlotFragmentWrapper(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	nodes := ui.RenderComponentSlot(ctx, ui.Fragment{}.With(
		ui.Text{Text: "a"},
		ui.Text{Text: "b"},
	))

	if len(nodes) != 1 || nodes[0].Type != "Fragment" || len(nodes[0].Children) != 2 {
		t.Fatalf("slot nodes: %#v", nodes)
	}
}

func TestSlotKeyFromField(t *testing.T) {
	t.Parallel()

	if got, want := ui.SlotKeyFromField("Trigger"), "trigger"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestFragmentRender(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	node := ui.Fragment{}.With(
		ui.Text{Text: "a"},
		ui.Text{Text: "b"},
	).Render(ctx)

	if node.Type != "Fragment" {
		t.Fatalf("Type: got %q want Fragment", node.Type)
	}

	if len(node.Children) != 2 {
		t.Fatalf("Children: got %d want 2", len(node.Children))
	}
}

func TestFragmentOrFallback(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	node := ui.Fragment{}.Or(ui.Text{Text: "fallback"}).Render(ctx)

	if len(node.Children) != 1 || node.Children[0].Type != "Text" {
		t.Fatalf("Children: %#v", node.Children)
	}
}
