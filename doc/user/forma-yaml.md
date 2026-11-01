# forma.yaml

`forma.yaml` at the project root is the **desired state** for your forma app. The CLI reads it on **`forma sync`** (and after **`forma init`** / **`forma add`**). Hand-edit advanced fields; use **`forma add`** for the component list in normal workflow.

## Example

```yaml
module: github.com/you/myapp

ui:
  mount: app-root
  components:
    - button
    - input
    - card

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

### `ui.components`

shadcn component names (CLI registry ids, typically lowercase) to include in the project. **`forma add`** appends here then syncs.

### `codegen.dir`

Directory for generated `*.gen.go` widget bindings.

- **Default:** `internal/ui`
- Override for monorepos or custom layout (see [Advanced](advanced.md)).

### `codegen.package`

Go package name for generated files (usually `ui`). Import path is `<module>/internal/ui`, not to be confused with **`github.com/zhulik/forma/ui`** (framework primitives: `Component`, `Context`, layout helpers, `UIMount`).

### `web.path`

Root of the Vite + React + shadcn web tree (default **`web/`**). The CLI manages shadcn config under this directory; authors rarely run `npx shadcn` directly.

## Future fields

Room for theme presets, desktop packaging, codegen overrides per component, and third-party registry entries—added when features ship, with updates to this page.

## Related

- [Project layout](project-layout.md)
- [CLI reference](cli.md)
