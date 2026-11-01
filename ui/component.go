package ui

import (
	"context"

	"github.com/zhulik/forma/fn"
)

type Component interface {
	Render(ctx context.Context) *Node
}

type BaseComponent struct {
	WidgetMeta

	children []Component
	fallback []Component
}

func (c BaseComponent) With(children ...Component) BaseComponent {
	c.children = append(c.children, children...)

	return c
}

func (c BaseComponent) Or(fallback ...Component) BaseComponent {
	c.fallback = append(c.fallback, fallback...)

	return c
}

func (c BaseComponent) Render(ctx context.Context) *Node {
	return &Node{
		Type:     "BaseComponent",
		Key:      c.Key,
		Class:    c.Class,
		Props:    c,
		Children: renderChildren(ctx, c.children, c.fallback),
	}
}

func renderChildren(ctx context.Context, children []Component, fallback []Component) []*Node {
	if len(children) == 0 {
		children = fallback
	}

	return fn.FilterMap(children, func(child Component) (*Node, bool) {
		node := child.Render(ctx)

		return node, node != nil
	})
}
