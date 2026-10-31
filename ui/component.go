package ui

import "context"

type Component interface {
	Render(ctx context.Context) *Node
}

type BaseComponent struct {
	// Top-level escape hatches
	Class string            `js:"-"`
	Style map[string]string `js:"-"`
	ID    string            `js:"-"`
	Attrs map[string]any    `js:"-"`

	children []Component
}

func (c BaseComponent) With(children ...any) BaseComponent {
	c.children = append(c.children, toComponents(children)...)
	return c
}

func (c BaseComponent) Render(ctx context.Context) *Node {
	return &Node{
		Type:     "BaseComponent",
		Class:    c.Class,
		Props:    c,
		Children: parseChildren(ctx, c.children),
	}
}

func toComponents(items []any) []Component {
	var components []Component
	for _, item := range items {
		switch v := item.(type) {
		case Component:
			components = append(components, v)
		case string:
			components = append(components, Text{Text: v})
		}
	}
	return components
}
