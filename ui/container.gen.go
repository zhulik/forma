package ui

import "context"

type Container struct {
	BaseComponent
	MaxWidth string
	MarginX  string
	MarginY  int
}

func (c Container) With(children ...any) Container {
	c.BaseComponent = c.BaseComponent.With(children...)
	return c
}

func (b Container) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Container"
	node.Props = b
	return node
}
