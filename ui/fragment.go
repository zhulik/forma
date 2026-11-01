package ui

import "context"

type Fragment struct {
	BaseComponent
}

func (f Fragment) With(children ...Component) Fragment {
	f.BaseComponent = f.BaseComponent.With(children...)

	return f
}

func (f Fragment) Or(fallback ...Component) Fragment {
	f.BaseComponent = f.BaseComponent.Or(fallback...)

	return f
}

func (f Fragment) Render(ctx context.Context) *Node {
	node := f.BaseComponent.Render(ctx)
	node.Type = "Fragment"

	return node
}
