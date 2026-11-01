package ui

import "context"

type Button struct {
	BaseComponent
	Variant string
	Size    string
	OnClick func(ctx Context)
}

func (b Button) With(children ...any) Button {
	b.BaseComponent = b.BaseComponent.With(children...)
	return b
}

func (b Button) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Button"
	node.Props = b
	return node
}
