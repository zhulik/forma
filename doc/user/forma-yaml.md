# forma.yaml

`forma.yaml` at the project root is the **desired state** for your forma app. The CLI reads it on **`forma sync`**, **`forma generate`**, and after **`forma init`** / **`forma add`**. Hand-edit paths and mount; use **`forma add`** to materialize new shadcn sources.

## Example

```yaml
module: github.com/you/myapp

ui:
  mount: app-root

codegen:
  dir: internal/ui
  package: ui

web:
  path: web
```

## Fields

### `module`

Go module path. Set by **`forma init`**; should match `go.mod`.

### `ui.mount`

DOM element id passed to **`UIMount`**. Must match the mount node in **`web/index.html`** (or equivalent).

### `codegen.dir`

Directory for generated `*.gen.go` widget bindings.

- **Default:** `internal/ui`
- Override for monorepos or custom layout (see [Advanced](advanced.md)).

### `codegen.package`

Go package name for generated files (usually `ui`). Import path is `<module>/internal/ui`, not to be confused with **`github.com/zhulik/forma/ui`** (framework primitives: `Component`, `Context`, layout helpers, `UIMount`).

### `web.path`

Root of the Vite + React + shadcn web tree (default **`web/`**). The CLI manages shadcn config under this directory; authors rarely run `npx shadcn` directly.

Widget inventory is **committed TSX** under **`{web.path}/components/ui/`**, not a list in this file.

## Future fields

Room for theme presets, desktop packaging, codegen overrides per component, and third-party registry entries—added when features ship, with updates to this page.

## Related

- [Project layout](project-layout.md)
- [CLI reference](cli.md)
