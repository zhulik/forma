# forma.yaml

`forma.yaml` at the project root is the **desired state** for your forma app. The CLI reads it on **`forma sync`** and after **`forma init`**. Hand-edit mount and web path as needed.

## Example

```yaml
module: github.com/you/myapp

ui:
  mount: app-root

web:
  path: web
```

## Fields

### `module`

Go module path. Set by **`forma init`**; should match `go.mod`.

### `ui.mount`

DOM element id passed to **`UIMount`**. Must match the mount node in **`web/index.html`** (or equivalent).

### `web.path`

Root of the Vite + React host web tree (default **`web/`**). Look comes from **0build** CDN assets linked in `index.html`; host builtins live under `src/host/`.

There is **no** default **`codegen:`** section. Owned widgets are imported from **`github.com/zhulik/forma/ui`**.

## Future fields

Room for theme presets, desktop packaging, and third-party registry hints—added when features ship, with updates to this page.

## Related

- [Project layout](project-layout.md)
- [CLI reference](cli.md)
