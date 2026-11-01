package ui

import "context"

type Align string
type Justify string

const (
	Center  Align   = "center"
	Between Justify = "between"
)

type Row struct {
	BaseComponent
	Gap     int
	Align   Align
	Justify Justify
	Padding int
	Border  bool
}

func (r Row) With(children ...Component) Row {
	r.BaseComponent = r.BaseComponent.With(children...)
	return r
}

func (r Row) Or(fallback ...Component) Row {
	r.BaseComponent = r.BaseComponent.Or(fallback...)
	return r
}

func (b Row) Render(ctx context.Context) *Node {
	node := b.BaseComponent.Render(ctx)
	node.Type = "Row"
	node.Props = b
	return node
}
