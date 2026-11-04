# forma

Go-native UI framework (in development): compose semantic widgets in Go; a React + shadcn host renders them.

**Authors:** start with the [user guide](doc/user/README.md) (install, CLI, `forma.yaml`, writing UI). Architecture and phases: [doc/DESIGN.md](doc/DESIGN.md), [doc/ROADMAP.md](doc/ROADMAP.md).

The reference DSL demo is [main.go](main.go) (dot-imported [`ui`](ui/) and [`fn`](fn/), `.With` / `.Or`, `ui.Context`).

## Development

Tool versions are pinned in [.tool-versions](.tool-versions) (asdf: Go, golangci-lint, Task, Node).

```bash
task          # static analysis (default)
task lint     # golangci-lint
task lint-fix # golangci-lint with autofix
task fix      # go fix modernizers
task generate # go generate ./...
task test     # go test -race ./... (no tests yet; wired for later)
```

CI runs `check-generate`, lint, and `check-test` on every push. After `forma init`, the scaffold ships the React host (`web/src/host/`), GoScript bridge wiring (`formaHost` then `/app.js`), and **`forma dev`** (Vite + GoScript rebuild).
