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

`app-root` must match `ui.mount` in **`forma.yaml`** and the mount element in **`web/index.html`** (or your configured `web.path`).

On a **browser (GoScript) build**, `UIMount` performs an initial refresh into the host via `globalThis.formaHost.setTree`. The scaffold `main.tsx` registers that seam before importing `/app.js` (the Vite-bundled GoScript artifact). On **native** `go test` / `go build`, `UIMount` is a no-op stub.

Run the browser loop with **`forma dev`** in an initialized app (after **`forma sync`**, with GoScript on PATH). Host TS/CSS use Vite HMR; Go rebuilds rewrite `web/public/app.js` and need a **page reload**—Go heap is not preserved. See [CLI — forma dev](cli.md#forma-dev).

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
- **Lists:** `fn.Map` / `fn.Filter` (and related helpers) over slices; set stable **`Key`** on repeated nodes (see `ui-dsl` spec).
- **Fallible collect:** `fn.MapErr` / `fn.FilterMapErr` when building a slice can fail (stop on first error).
- **Maps:** `fn.MapValues`, `fn.MapEntries`, `fn.FilterEntries`, `fn.FilterMapEntries` (and `*Err` variants) for map transforms—not hand-written accumulate loops.

## Handlers and refresh

Handlers receive **`ui.Context`** (embeds `context.Context`) as the first argument. Click-style handlers are context-only; value and checked handlers take a sanitized payload:

```go
OnClick: func(ctx Context) {
    ctx.Update(func() {
        state.Count++
    })
},
OnChange: func(ctx Context, value string) {
    ctx.Update(func() {
        state.Draft = value
    })
},
OnCheckedChange: func(ctx Context, checked bool) {
    ctx.Update(func() {
        todo.Done = checked
    })
},
```

`Update` runs the mutation and calls **`RefreshUI()`** once on success.

After async work (goroutine), mutate state then **`RefreshUI()`** on the handler context, respecting cancellation.

## Layout vs generated Card

Framework layout uses **`Panel`** for a padded bordered box (`Padding` prop). After **`forma init`**, **`Card`** refers to the generated shadcn Card binding (Header / Title / Content slots)—do not expect a framework layout type named `Card`.

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

`Class` values must be **static** Tailwind utility strings (including responsive variants such as `sm:flex-row`). The scaffold `tailwind.config.js` scans `../*.go` and `../internal/**/*.go` relative to `web/` so those utilities are not purged. Dynamic string concatenation in Go will not appear in the built CSS.

## Related

- [`ui-dsl` specification](../../openspec/specs/ui-dsl/spec.md)
- [Project layout](project-layout.md)
