package ui

import "context"

type Input struct {
	BaseComponent
	Placeholder string
	Value       string
	OnChange    func(ctx Context, val string)
}

func (i Input) With(children ...any) Input {
	i.BaseComponent = i.BaseComponent.With(children...)
	return i
}

func (i Input) Render(ctx context.Context) *Node {
	node := i.BaseComponent.Render(ctx)
	node.Type = "Input"
	node.Props = i
	return node
}
