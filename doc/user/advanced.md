# Advanced

## Custom Go components

Return `ui.Component` from functions or structs by composing owned widgets:

```go
func UserRow(name string) Component {
	return Row{Align: Center, Gap: 2}.With(
		Text{Text: name},
	)
}
```

## Custom host islands

1. Implement a thin React wrapper that sanitizes events.
2. Export it from **`web/src/registry.custom.ts`** (merged after builtins; last wins).
3. Add a matching Go struct that renders a `Node` with that type name and props.

Do not put business logic in the host.

## Escape hatches

- **`Class` / `Style` / `Attrs`** on widgets (via `WidgetMeta` / `BaseComponent`).
- Prefer semantic `.z-*`; use 0build utility shorthand (for example `p=4`, `gap=2`) when a one-off is clearer than a new semantic class.

## Slots / future compounds

Fragment and `Node.Slots` remain for field-only compound regions. New compounds should be hand-maintained Go + host pairs—not kit-generated catalogs.

## Related

- [Owned widgets](adding-components.md)
- [Project layout](project-layout.md)
