package ui

// WidgetMeta holds escape hatches shared by layout, leaf, and compound widgets.
type WidgetMeta struct {
	Key   string            `js:"-"`
	Class string            `js:"-"`
	Style map[string]string `js:"-"`
	ID    string            `js:"-"`
	Attrs map[string]any    `js:"-"`
}
