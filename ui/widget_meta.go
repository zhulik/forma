package ui

// WidgetMeta holds escape hatches shared by layout, leaf, and compound widgets.
// No syscall/js tags: Node fields are copied explicitly; tags would require
// embedding js.Value and break the js build for no benefit on this struct.
type WidgetMeta struct {
	Key   string
	Class string
	Style map[string]string
	ID    string
	Attrs map[string]any
}
