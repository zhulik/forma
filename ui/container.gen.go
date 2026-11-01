package ui

import "context"

type Container struct {
	BaseComponent
	MaxWidth string
	MarginX  string
	MarginY  int
}

func (c Container) With(children ...Component) Container {
	c.BaseComponent = c.BaseComponent.With(children...)
	return c
}

func (c Container) Or(fallback ...Component) Container {
	c.BaseComponent = c.BaseComponent.Or(fallback...)
	return c
}

func (b Container) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Container"
	node.Props = b
	return node
}
