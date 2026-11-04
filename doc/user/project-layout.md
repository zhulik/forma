# Project layout

## Your app (after forma init)

Layout after **`forma init`** (sync materializes starter shadcn components and runs codegen):

```
myapp/
  forma.yaml           # mount, web path, codegen paths
  generator.go         # //go:generate forma generate
  Taskfile.yaml        # check-generate: go generate + git diff --exit-code
  go.mod
  go.sum               # after sync / go mod tidy
  main.go              # UIMount, app root component
  internal/
    ui/                # generated *.gen.go (default codegen.dir)
      button.gen.go
      ...
  web/                 # React + shadcn + Vite (web.path)
    package.json
    src/
      components/ui/   # shadcn TSX sources (committed)
      main.tsx
      registry.gen.ts  # generated host registry
    ...
```

### Import split

| Import | Role |
|--------|------|
| `github.com/zhulik/forma/ui` | Framework: `Component`, `Context`, `UIMount`, layout widgets, `Fragment`, … |
| `github.com/zhulik/forma/fn` | `Map`, `Filter`, `MapErr`, `FilterMapErr`, map helpers (`MapValues`, `MapEntries`, …) for lists and transforms |
| `your/module/internal/ui` | **Generated** shadcn bindings (`Button`, `Input`, …) |

Application code often dot-imports framework `ui` and `fn` for a dense DSL (see [Writing UI](writing-ui.md)).

## This repository (github.com/zhulik/forma)

The forma repo is **not** a typical user app today:

- **`ui/`** at repo root = framework DSL + **hand-written** `*.gen.go` stubs until a follow-up moves demos to **`examples/`**.
- **`main.go`** = reference todo demo at repo root.
- No app-level **`web/`** or **`forma.yaml`** at the framework root.

When you build a real app, use **`forma init`** into its own module. This repo remains the library and CLI source.

## Version control

Commit:

- `forma.yaml`, `generator.go`, `Taskfile.yaml`, `go.mod`, `go.sum`
- `main.go` and your app packages
- `web/` sources including shadcn **`components/ui/*.tsx`**
- Generated **`internal/ui/*.gen.go`** and **`web/src/registry.gen.ts`**

Do not commit `web/node_modules/`. After **`forma sync`** or **`forma generate`**, review diffs before commit.

## Related

- [forma.yaml](forma-yaml.md)
- [Adding components](adding-components.md)
