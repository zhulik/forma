package ui

import "context"

type Heading struct {
	BaseComponent
	Level int
	Text  string
}

func (h Heading) With(children ...Component) Heading {
	h.BaseComponent = h.BaseComponent.With(children...)
	return h
}

func (h Heading) Or(fallback ...Component) Heading {
	h.BaseComponent = h.BaseComponent.Or(fallback...)
	return h
}

func (h Heading) Render(ctx context.Context) *Node {
	node := h.BaseComponent.Render(ctx)
	node.Type = "Heading"
	node.Props = h
	return node
}
