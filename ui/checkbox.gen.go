package ui

import "context"

type Checkbox struct {
	BaseComponent
	Checked         bool
	OnCheckedChange func(ctx Context, checked bool)
}

func (c Checkbox) With(children ...any) Checkbox {
	c.BaseComponent = c.BaseComponent.With(children...)
	return c
}

func (c Checkbox) Render(ctx context.Context) *Node {
	node := c.BaseComponent.Render(ctx)
	node.Type = "Checkbox"
	node.Props = c
	return node
}
