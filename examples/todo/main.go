package main

import (
	"context"
	"fmt"
	"strconv"

	. "github.com/zhulik/forma/fn" //nolint:revive,staticcheck // DSL
	. "github.com/zhulik/forma/ui" //nolint:revive,staticcheck // DSL

	todoUI "github.com/zhulik/forma/examples/todo/internal/ui"
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

// todoRoot re-builds the screen on every refresh so widgets read current state.
type todoRoot struct{}

func (todoRoot) Render(ctx context.Context) *Node {
	return TodoApp().Render(ctx)
}

func main() {
	UIMount("app-root", todoRoot{})
}

// --- Root Component ---

func TodoApp() Component {
	return Container{
		MaxWidth: "md",
		MarginX:  "auto",
		MarginY:  10,
		Class:    "px-4",
	}.With(
		todoUI.Card{
			Class: "shadow-xl border-slate-200",
			Title: Text{Text: "My Tasks"},
			Description: If(len(state.Todos) > 0, func() Component {
				return Text{
					Text:  fmt.Sprintf("%d task(s)", len(state.Todos)),
					Muted: true,
				}
			}),
			Content: Column{
				Gap: 6,
			}.With(
				// Composer: stacks on narrow viewports; Input+Add side-by-side from sm up.
				Container{
					Class: "flex w-full flex-col gap-2 sm:flex-row sm:items-center",
				}.With(
					todoUI.Input{
						BaseComponent: BaseComponent{Class: "flex-1"},
						Placeholder:   "What needs to be done?",
						Value:         state.InputValue,
						OnChange: func(ctx Context, val string) {
							ctx.Update(func() {
								state.InputValue = val
							})
						},
					},
					todoUI.Button{
						Variant: "default",
						OnClick: handleAddTodo,
					}.With(Text{Text: "Add"}),
				),

				Column{
					Gap: 3,
				}.With(renderTodoList()...).
					Or(
						Text{
							Class: "text-center py-6",
							Text:  "No tasks yet. Enjoy your day!",
							Muted: true,
						},
					),
			),
		},
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
				todoUI.Checkbox{
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

			todoUI.Button{
				Variant: "ghost",
				OnClick: func(ctx Context) {
					ctx.Update(func() {
						deleteTodo(todo.ID)
					})
				},
			}.With(Text{Text: "Delete"}),
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
