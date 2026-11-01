package main

import (
	. "github.com/zhulik/forma/fn" //nolint:revive,staticcheck // DSL
	. "github.com/zhulik/forma/ui" //nolint:revive,staticcheck // DSL
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
	UIMount("app-root", TodoApp())
}

// --- Root Component ---

func TodoApp() Component {
	return Container{
		MaxWidth: "md",
		MarginX:  "auto",
		MarginY:  10,
	}.With(
		Card{
			Class:   "shadow-xl border-slate-200", // Escape hatch alongside props!
			Padding: 6,
		}.With(
			Column{
				Gap: 6,
			}.With(
				Heading{
					Level: 2,
					Text:  "My Tasks",
				},

				// Input Row
				Row{
					Gap: 2,
				}.With(
					Input{
						Placeholder: "What needs to be done?",
						Value:       state.InputValue,
						OnChange: func(ctx Context, val string) {
							state.InputValue = val
							ctx.RefreshUI()
						},
					},
					Button{
						Variant: "default", // Typed prop
						OnClick: handleAddTodo,
					}.With("Add Task"),
				),

				// Dynamic Task List
				Column{
					Gap: 3,
				}.With(renderTodoList()...).
					Or(
						Text{
							Class: "text-center py-6",
							Muted: true,
						}.With("No tasks yet. Enjoy your day!"),
					),
			),
		),
	)
}

// --- Dynamic List Renderer ---

func renderTodoList() []any {
	return Map(state.Todos, func(todo *Todo) any {
		return Row{
			Align:   Center,
			Justify: Between,
			Padding: 3,
			Border:  true,
		}.With(
			Row{
				Align: Center,
				Gap:   3,
			}.With(
				Checkbox{
					Checked: todo.Done,
					OnCheckedChange: func(ctx Context, checked bool) {
						todo.Done = checked
						ctx.RefreshUI()
					},
				},
				Text{
					Strikethrough: todo.Done,
					Muted:         todo.Done,
				}.With(todo.Text),
			),

			Button{
				Variant: "ghost",
				Size:    "icon",
				OnClick: func(ctx Context) {
					deleteTodo(todo.ID)
					ctx.RefreshUI()
				},
			}.With("✕"),
		)
	})
}

// --- Business Logic ---

func handleAddTodo(ctx Context) {
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

	ctx.RefreshUI()
}

func deleteTodo(id int) {
	state.Todos = Filter(state.Todos, func(t *Todo) bool {
		return t.ID != id
	})
}
