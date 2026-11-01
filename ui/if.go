package ui

import "context"

type skip struct{}

func (skip) Render(context.Context) *Node {
	return nil
}

const (
	ifBranchThen = 1
	ifBranchElse = 2
)

// If selects a lazy branch. One thunk: rendered when cond is true; false yields no node.
// Two thunks: first when true, second when false. More than two thunks panics.
func If(cond bool, thunks ...func() Component) Component {
	switch len(thunks) {
	case ifBranchThen:
		if cond {
			return thunks[0]()
		}

		return skip{}
	case ifBranchElse:
		if cond {
			return thunks[0]()
		}

		return thunks[1]()
	default:
		panic("ui.If: want 1 or 2 branch functions")
	}
}
