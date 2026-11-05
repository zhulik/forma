package ui

import "context"

type Panel struct {
	BaseComponent
	Padding int
}

func (p Panel) With(children ...Component) Panel {
	p.BaseComponent = p.BaseComponent.With(children...)
	return p
}

func (p Panel) Or(fallback ...Component) Panel {
	p.BaseComponent = p.BaseComponent.Or(fallback...)
	return p
}

func (p Panel) Render(ctx context.Context) *Node {
	node := p.BaseComponent.Render(ctx)
	node.Type = "Panel"
	node.Props = p
	return node
}
