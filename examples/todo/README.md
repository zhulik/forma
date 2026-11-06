# Todo example

In-memory todo demo on forma **owned widgets** + [0build](https://0build.dev). Nested module with `replace github.com/zhulik/forma => ../..`. No backend or persistence.

## Prerequisites

- Go (match framework `go.mod`)
- Node.js / npm
- `forma` CLI (`cd ../../cmd/forma && go build -o forma .`, or `go install`)
- GoScript on PATH (`go install github.com/s4wave/goscript/cmd/goscript@v0.6.3`)

## Run

From this directory (`examples/todo`):

```bash
forma sync
forma dev
```

Open the URL Vite prints. Confirm:

- empty state (“No tasks yet…”) with zero tasks
- add a task
- toggle done
- delete a task

`forma.yaml` sets `ui.mount: app-root` and `web.path: web` (matches `web/index.html` and `web/src/main.tsx`).

## Layout

- Module: `github.com/zhulik/forma/examples/todo`
- `web/` is the minimal scaffold host (owned builtins + 0build CDN)
- Sync/dev outputs (`web/node_modules/`, `web/.goscript/`, `web/public/app.js`, …) are gitignored via `web/.gitignore`
