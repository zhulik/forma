package main

import (
	. "github.com/zhulik/forma/ui" //nolint:revive,staticcheck // DSL
)

func main() {
	UIMount("app-root", Smoke())
}

// Smoke is a thin DSL compile check. The runnable browser todo lives in examples/todo.
func Smoke() Component {
	return Container{
		MaxWidth: "md",
		MarginX:  "auto",
		MarginY:  10,
	}.With(
		Panel{
			Padding: 6,
		}.With(
			Text{Text: "forma DSL smoke — see examples/todo for the browser demo"},
		),
	)
}
