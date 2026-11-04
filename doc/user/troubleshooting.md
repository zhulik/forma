# Troubleshooting

## CLI or forma dev not found

Install the **`forma`** CLI per [Install](install.md). **`forma dev`** needs **Node.js**, **`npm`**, **`goscript`** on PATH (pinned **v0.6.3**), a readable **`forma.yaml`**, and **`web.path/node_modules`** (run **`forma sync`** once). Missing deps fail fast with a clear error—`forma dev` does not auto-sync.

## UI does not update after a click

- Confirm the handler mutates **Go state** and calls **`ctx.RefreshUI()`** or **`ctx.Update`**.
- Async: refresh must run after the goroutine updates state; check context cancellation.
- Confirm **`globalThis.formaHost.setTree`** exists before the Go bundle runs (scaffold `main.tsx` registers it, then imports `/app.js`). Missing seam → developer-visible panic/failure from `UIMount`, not a silent blank tree.
- After a Go rebuild, **reload the page**—Go state is not preserved across GoScript rebuilds (not true Go HMR).

## List items jump or remount

- Set stable **`Key`** on repeated nodes (`ui-dsl` requirement). Avoid using slice index alone when items reorder or delete.

## Input loses focus or cursor jumps

- The host applies a **hybrid text-input policy** for `Input`: local React state while focused, sanitized string `OnChange`, value-prop sync when unfocused. Domain state stays in Go.
- If caret still jumps, check that refreshes are not overwriting the field with a *different* value while focused, and that custom registry entries did not bypass the host `Input` wrapper path.

## forma sync / go mod tidy failures

- Ensure **`main.go`** (or scaffold) imports **`github.com/zhulik/forma/ui`** so tidy retains the framework module.
- Run sync from project root with valid **`forma.yaml`**.

## Browser shows blank page

- **`ui.mount`** must match **`UIMount`** id and the mount node in **`web/index.html`** (or your `web.path` HTML). Missing mount fails closed (no alternate root); check the console for `[forma] mount element … not found`.
- Registry must include every **Type** string your tree emits (builtins + generated + `registry.custom.ts`). Unknown types are skipped with a console warning.
- Tree stays null until Go calls **`UIMount`** / refresh. Confirm **`web/public/app.js`** exists (initial `forma dev` build) and the console has no script load errors for `/app.js`.
- Check the browser console for JavaScript errors (when the web tree is built and served).

## Docs vs behavior mismatch

If the product behaves differently from this guide, that is a bug—either fix the code or update **`doc/user/`** in the same change. See [README](README.md#documentation-is-first-class).

## Related

- [Writing UI](writing-ui.md)
- [CLI reference](cli.md)
