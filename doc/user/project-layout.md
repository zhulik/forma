# Project layout

## Your app (after forma init)

Target layout once the CLI exists:

```
myapp/
  forma.yaml           # desired state (components, paths, mount)
  go.mod
  go.sum               # after sync / go mod tidy
  main.go              # UIMount, app root component
  internal/
    ui/                # generated *.gen.go (default codegen.dir)
      button.gen.go
      ...
  host/                # React + shadcn + Vite (host.path)
    package.json
    components/ui/     # shadcn TSX sources (committed)
    ...
```

### Import split

| Import | Role |
|--------|------|
| `github.com/zhulik/forma/ui` | Framework: `Component`, `Context`, `UIMount`, layout widgets, `Fragment`, … |
| `github.com/zhulik/forma/fn` | `Map`, `Filter`, … for lists and transforms |
| `your/module/internal/ui` | **Generated** shadcn bindings (`Button`, `Input`, …) |

Application code often dot-imports framework `ui` and `fn` for a dense DSL (see [Writing UI](writing-ui.md)).

## This repository (github.com/zhulik/forma)

The forma repo is **not** a typical user app today:

- **`ui/`** at repo root = framework DSL + **hand-written** `*.gen.go` stubs for development until P3a codegen.
- **`main.go`** = reference todo demo at repo root.
- No **`host/`** or **`forma.yaml`** yet.

When you build a real app, use **`forma init`** into its own module. This repo remains the library and CLI source; examples may move under `examples/` later.

## Version control

Commit:

- `forma.yaml`, `go.mod`, `go.sum`
- `main.go` and your app packages
- `host/` sources including shadcn **`components/ui/*.tsx`**
- Generated **`internal/ui/*.gen.go`**

Do not commit `host/node_modules/`. After **`forma sync`**, review diffs before commit—sync may add or update generated files and lockfiles.

## Related

- [forma.yaml](forma-yaml.md)
- [Adding components](adding-components.md)
