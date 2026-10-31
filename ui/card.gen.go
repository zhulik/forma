package ui

import "context"

type Card struct {
	BaseComponent
	Padding int
}

func (c Card) With(children ...any) Card {
	c.BaseComponent = c.BaseComponent.With(children...)
	return c
}

func (b Card) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Card"
	node.Props = b
	return node
}
