# Writing UI

## Mental model

- **State** lives in ordinary Go variables and structs—no React state for domain logic.
- **UI** is a function of state: build a tree of **`Component`** values.
- **Events** run Go handlers; after mutating state, call **`ctx.RefreshUI()`** (or **`ctx.Update`**).

The reference implementation is [`main.go`](../../main.go) in the forma repository.

## Mount

```go
func main() {
    UIMount("app-root", MyApp())
}
```

`app-root` must match `ui.mount` in **`forma.yaml`** and the host HTML mount element.

> **Today:** `UIMount` in the framework repo is a stub. Behavior is specified in **`ui-dsl`**; the browser bridge arrives in P3.

## Dot imports

Apps typically dot-import for fluency:

```go
import (
    . "github.com/zhulik/forma/fn"
    . "github.com/zhulik/forma/ui"
)
```

Use `//nolint:revive,staticcheck` on those lines if your linter flags dot imports.

Generated widgets import from **`your/module/internal/ui`** (or your configured `codegen.dir`), not from the framework root `ui` package.

## Composition

- **Children:** `.With(child...)` on layout and leaf widgets.
- **Empty states:** `.Or(fallback...)` when primary children are empty.
- **Text:** `Text{Text: "Hello"}` (not raw strings in `.With`).
- **Conditionals:** `If(cond, func() Component { ... })`.
- **Lists:** `fn.Map` over slices; set stable **`Key`** on repeated nodes (see `ui-dsl` spec).

## Handlers and refresh

Handlers receive **`ui.Context`** (embeds `context.Context`):

```go
OnClick: func(ctx Context) {
    ctx.Update(func() {
        state.Count++
    })
},
```

`Update` runs the mutation and calls **`RefreshUI()`** once on success.

After async work (goroutine), mutate state then **`RefreshUI()`** on the handler context, respecting cancellation.

## Compound widgets

shadcn-style compounds (`Dialog`, `Tabs`, …) use **field slots**, not separate `DialogTrigger`-style Go types:

```go
Dialog{
    Trigger: Button{Text: "Open"}.With(...),
    Content: Column{}.With(...),
}
```

See **`ui-dsl`** for slot rules.

## Escape hatches

`Class`, `Style`, and `Attrs` on **`BaseComponent`** for Tailwind or rare attributes—use sparingly.

## Related

- [`ui-dsl` specification](../../openspec/specs/ui-dsl/spec.md)
- [Project layout](project-layout.md)
