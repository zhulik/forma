package ui

type Node struct {
	Type     string
	Class    string
	Props    any
	Children []*Node
}
