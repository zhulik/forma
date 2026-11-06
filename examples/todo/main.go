package main

import (
	"context"
	"fmt"
	"strconv"

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
	// Root re-renders from state on every RefreshUI (static TodoApp() once would freeze Values/list).
	UIMount("app-root", todoRoot{})
}

type todoRoot struct{}

func (todoRoot) Render(ctx context.Context) *Node {
	return TodoApp().Render(ctx)
}

// --- Root Component ---

func TodoApp() Component {
	return Container{
		MaxWidth: "md",
		MarginX:  "auto",
		MarginY:  10,
		// Horizontal padding so phone-width viewports stay readable without page chrome scroll.
		Class: "px=4",
	}.With(
		Card{
			Class:   "shadow=xl", // 0build utility escape hatch alongside semantic .z-card
			Padding: 6,
		}.With(
			Column{
				Gap: 6,
			}.With(
				Heading{
					Level: 2,
					Text:  "My Tasks",
				},

				If(len(state.Todos) > 0, func() Component {
					return Text{
						Text:  fmt.Sprintf("%d task(s)", len(state.Todos)),
						Muted: true,
					}
				}),

				// Input Row
				Row{
					Gap: 2,
				}.With(
					Input{
						Placeholder: "What needs to be done?",
						Value:       state.InputValue,
						OnChange: func(ctx Context, val string) {
							ctx.Update(func() {
								state.InputValue = val
							})
						},
					},
					Button{
						Variant: "default",
						OnClick: handleAddTodo,
					}.With(Text{Text: "Add Task"}),
				),

				// Dynamic Task List
				Column{
					Gap: 3,
				}.With(renderTodoList()...).
					Or(
						Text{
							Class: "text-align=center py=6",
							Text:  "No tasks yet. Enjoy your day!",
							Muted: true,
						},
					),
			),
		),
	)
}

// --- Dynamic List Renderer ---

func renderTodoList() []Component {
	return Map(state.Todos, func(todo *Todo) Component {
		return Row{
			Key:     strconv.Itoa(todo.ID),
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
						ctx.Update(func() {
							todo.Done = checked
						})
					},
				},
				Text{
					Text:          todo.Text,
					Muted:         todo.Done,
					Strikethrough: todo.Done,
				},
			),

			Button{
				Variant: "ghost",
				Size:    "icon",
				OnClick: func(ctx Context) {
					ctx.Update(func() {
						deleteTodo(todo.ID)
					})
				},
			}.With(Text{Text: "✕"}),
		)
	})
}

// --- Business Logic ---

func handleAddTodo(ctx Context) {
	if state.InputValue == "" {
		return
	}

	ctx.Update(func() {
		state.Todos = append(state.Todos, &Todo{
			ID:   state.nextID,
			Text: state.InputValue,
			Done: false,
		})
		state.nextID++
		state.InputValue = ""
	})
}

func deleteTodo(id int) {
	state.Todos = Filter(state.Todos, func(t *Todo) bool {
		return t.ID != id
	})
}
