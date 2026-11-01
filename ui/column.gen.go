package ui

import "context"

type Column struct {
	BaseComponent
	Gap int
}

func (c Column) With(children ...Component) Column {
	c.BaseComponent = c.BaseComponent.With(children...)
	return c
}

func (c Column) Or(fallback ...Component) Column {
	c.BaseComponent = c.BaseComponent.Or(fallback...)
	return c
}

func (b Column) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Column"
	node.Props = b
	return node
}
