package ui

import "context"

type Column struct {
	BaseComponent
	Gap int
}

func (c Column) With(children ...any) Column {
	c.BaseComponent = c.BaseComponent.With(children...)
	return c
}

func (b Column) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Column"
	node.Props = b
	return node
}
