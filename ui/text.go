package ui

import "context"

type Text struct {
	BaseComponent

	Text          string
	Muted         bool
	Strikethrough bool
}

func (t Text) With(children ...Component) Text {
	t.BaseComponent = t.BaseComponent.With(children...)

	return t
}

func (t Text) Or(fallback ...Component) Text {
	t.BaseComponent = t.BaseComponent.Or(fallback...)

	return t
}

func (t Text) Render(ctx context.Context) *Node {
	node := t.BaseComponent.Render(ctx)
	node.Type = "Text"
	node.Props = t

	return node
}
