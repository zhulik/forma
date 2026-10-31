package ui

import "context"

type Text struct {
	BaseComponent
	Text          string
	Muted         bool
	Strikethrough bool
}

func (t Text) With(children ...any) Text {
	t.BaseComponent = t.BaseComponent.With(children...)
	return t
}

func (b Text) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Text"
	node.Props = b
	return node
}
