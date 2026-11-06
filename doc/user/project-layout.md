# Project layout

## Your app (after forma init)

```
myapp/
  forma.yaml           # mount + web path
  Taskfile.yaml        # optional local tasks
  go.mod
  go.sum               # after sync / go mod tidy
  main.go              # UIMount, app root component
  web/                 # React host + Vite (web.path)
    package.json
    index.html         # 0build CDN CSS/JS + mount node
    src/
      host/            # write-once React host (builtins, renderer, sanitize)
        widgets/       # thin Button/Input/… mapped to .z-* classes
      main.tsx         # mounts FormaHost; merges builtins + registry.custom
      registry.custom.ts  # optional third-party / override registry
    ...
```

### Import split

| Import | Role |
|--------|------|
| `github.com/zhulik/forma/ui` | Framework: `Component`, `Context`, `UIMount`, owned widgets (Button, Input, layout, …) |
| `github.com/zhulik/forma/fn` | `Map`, `Filter`, … for lists and transforms |

Application code often dot-imports framework `ui` and `fn` for a dense DSL (see [Writing UI](writing-ui.md)).

## This repository (github.com/zhulik/forma)

The forma repo is **not** a typical user app today:

- **`ui/`** at repo root = framework DSL + **hand-written** owned widgets.
- **`main.go`** = reference todo demo at repo root.
- No app-level **`web/`** or **`forma.yaml`** at the framework root (scaffold lives under `cmd/forma/templates/`).

When you build a real app, use **`forma init`** into its own module. This repo remains the library and CLI source.

## Version control

Commit:

- `forma.yaml`, `Taskfile.yaml`, `go.mod`, `go.sum`
- `main.go` and your app packages
- `web/` sources (host, `registry.custom.ts`, lockfile)

Do not commit `web/node_modules/`.

## Related

- [forma.yaml](forma-yaml.md)
- [Owned widgets](adding-components.md)
