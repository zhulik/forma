package main

import (
	"context"

	"github.com/zhulik/forma/ui"
)

// --- Application State ---

type Todo struct {
	ID   int
	Text string
	Done bool
}

type AppState struct {
	Todos      []*Todo
	InputValue string
	nextID     int
}

var state = &AppState{ //nolint:gochecknoglobals // demo app root state
	Todos:  []*Todo{},
	nextID: 1,
}

func main() {
	ui.Mount("app-root", TodoApp())
}

// --- Root Component ---

func TodoApp() ui.Component {
	return ui.Container{
		MaxWidth: "md",
		MarginX:  "auto",
		MarginY:  10,
	}.With(
		ui.Card{
			Class:   "shadow-xl border-slate-200", // Escape hatch alongside props!
			Padding: 6,
		}.With(
			ui.Column{
				Gap: 6,
			}.With(
				ui.Heading{
					Level: 2,
					Text:  "My Tasks",
				},

				// Input Row
				ui.Row{
					Gap: 2,
				}.With(
					ui.Input{
						Placeholder: "What needs to be done?",
						Value:       state.InputValue,
						OnChange: func(ctx context.Context, val string) {
							_ = ctx
							state.InputValue = val
							ui.Refresh()
						},
					},
					ui.Button{
						Variant: "default", // Typed prop
						OnClick: handleAddTodo,
					}.With("Add Task"),
				),

				// Dynamic Task List
				ui.Column{
					Gap: 3,
				}.With(renderTodoList()...),
			),
		),
	)
}

// --- Dynamic List Renderer ---

func renderTodoList() []any {
	if len(state.Todos) == 0 {
		return []any{
			ui.Text{
				Class: "text-center py-6",
				Muted: true,
			}.With("No tasks yet. Enjoy your day!"),
		}
	}

	var list []any
	for _, t := range state.Todos {
		todo := t // Capture loop variable

		row := ui.Row{
			Align:   ui.Center,
			Justify: ui.Between,
			Padding: 3,
			Border:  true,
		}.With(
			ui.Row{
				Align: ui.Center,
				Gap:   3,
			}.With(
				ui.Checkbox{
					Checked: todo.Done,
					OnCheckedChange: func(ctx context.Context, checked bool) {
						_ = ctx
						todo.Done = checked
						ui.Refresh()
					},
				},
				ui.Text{
					Strikethrough: todo.Done,
					Muted:         todo.Done,
				}.With(todo.Text),
			),

			ui.Button{
				Variant: "ghost",
				Size:    "icon",
				OnClick: func(ctx context.Context) {
					_ = ctx
					deleteTodo(todo.ID)
					ui.Refresh()
				},
			}.With("✕"),
		)

		list = append(list, row)
	}

	return list
}

// --- Business Logic ---

func handleAddTodo(ctx context.Context) {
	_ = ctx
	if state.InputValue == "" {
		return
	}

	state.Todos = append(state.Todos, &Todo{
		ID:   state.nextID,
		Text: state.InputValue,
		Done: false,
	})
	state.nextID++
	state.InputValue = ""

	ui.Refresh()
}

func deleteTodo(id int) {
	var updated []*Todo
	for _, t := range state.Todos {
		if t.ID != id {
			updated = append(updated, t)
		}
	}
	state.Todos = updated
}
