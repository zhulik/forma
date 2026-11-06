package main

import (
	. "github.com/zhulik/forma/ui" //nolint:revive,staticcheck // DSL
)

// Root package main is a short Hello-style smoke for the framework module.
// The canonical interactive todo demo is examples/todo (forma sync + forma dev).

func main() {
	UIMount("app-root", App())
}

func App() Component {
	return Container{
		MaxWidth: "md",
		MarginX:  "auto",
		MarginY:  10,
	}.With(
		Heading{
			Level: 1,
			Text:  "Hello from forma",
		},
		Text{
			Text:  "Run the interactive todo at examples/todo with forma sync && forma dev.",
			Muted: true,
		},
	)
}
