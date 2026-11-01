package ui

type Node struct {
	Type     string
	Key      string
	Class    string
	Props    any
	Children []*Node
}
