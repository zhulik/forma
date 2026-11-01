package ui

import (
	"context"

	"github.com/zhulik/forma/fn"
)

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
	fallback []Component
}

func (c BaseComponent) With(children ...any) BaseComponent {
	c.children = append(c.children, toComponents(children)...)

	return c
}

func (c BaseComponent) Or(fallback ...any) BaseComponent {
	c.fallback = append(c.fallback, toComponents(fallback)...)

	return c
}

func (c BaseComponent) Render(ctx context.Context) *Node {
	return &Node{
		Type:     "BaseComponent",
		Class:    c.Class,
		Props:    c,
		Children: renderChildren(ctx, c.children, c.fallback),
	}
}

func toComponents(items []any) []Component {
	return fn.Map(items, func(item any) Component {
		switch v := item.(type) {
		case Component:
			return v
		case string:
			return Text{Text: v}
		default:
			panic("unknown type")
		}
	})
}

func renderChildren(ctx context.Context, children []Component, fallback []Component) []*Node {
	if len(children) == 0 {
		children = fallback
	}

	return fn.Map(children, func(child Component) *Node {
		return child.Render(ctx)
	})
}
