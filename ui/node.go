package ui

import "context"

type Node struct {
	Type     string
	Class    string
	Props    any
	Children []*Node
}

func parseChildren(ctx context.Context, children []Component) []*Node {
	var nodes []*Node
	for _, child := range children {
		nodes = append(nodes, child.Render(ctx))
	}
	return nodes
}
