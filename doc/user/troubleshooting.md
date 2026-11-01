# Troubleshooting

## CLI or forma dev not found

The **`forma`** binary and **`forma dev`** are not shipped from this repository yet. See [Install](install.md) for what works today (`main.go`, `task test`).

## UI does not update after a click

- Confirm the handler mutates **Go state** and calls **`ctx.RefreshUI()`** or **`ctx.Update`**.
- Async: refresh must run after the goroutine updates state; check context cancellation.

## List items jump or remount

- Set stable **`Key`** on repeated nodes (`ui-dsl` requirement). Avoid using slice index alone when items reorder or delete.

## Input loses focus or cursor jumps

- Full-tree refresh on every keystroke can cause focus issues. The framework may use controlled/deferred input policies per widget (see DESIGN.md); until P3 documents per-component behavior, prefer blur/submit-driven updates for text fields if you hit this in a real host.

## forma sync / go mod tidy failures

- Ensure **`main.go`** (or scaffold) imports **`github.com/zhulik/forma/ui`** so tidy retains the framework module.
- Run sync from project root with valid **`forma.yaml`**.

## Host shows blank page

- **`ui.mount`** must match **`UIMount`** id and host HTML.
- Registry must include every **Type** string your Go tree emits (sync should register generated widgets).
- Check browser console for host errors (when host exists).

## Docs vs behavior mismatch

If the product behaves differently from this guide, that is a bug—either fix the code or update **`doc/user/`** in the same change. See [README](README.md#documentation-is-first-class).

## Related

- [Writing UI](writing-ui.md)
- [CLI reference](cli.md)
