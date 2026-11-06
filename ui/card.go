package ui

import "context"

type Card struct {
	BaseComponent

	Padding int
}

func (c Card) With(children ...Component) Card {
	c.BaseComponent = c.BaseComponent.With(children...)

	return c
}

func (c Card) Or(fallback ...Component) Card {
	c.BaseComponent = c.BaseComponent.Or(fallback...)

	return c
}

func (c Card) Render(ctx context.Context) *Node {
	node := c.BaseComponent.Render(ctx)
	node.Type = "Card"
	node.Props = c

	return node
}
