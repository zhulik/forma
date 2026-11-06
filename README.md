# forma

Go-native UI framework (in development): compose **owned widgets** in Go; a thin React host renders them with [0build](https://0build.dev) semantic CSS.

**Authors:** start with the [user guide](doc/user/README.md) (install, CLI, `forma.yaml`, writing UI). Architecture and phases: [doc/DESIGN.md](doc/DESIGN.md), [doc/ROADMAP.md](doc/ROADMAP.md).

**Canonical interactive demo:** [`examples/todo`](examples/todo/) — in-memory todo on owned widgets. From that directory: `forma sync` then `forma dev`. See [examples/todo/README.md](examples/todo/README.md).

Root [main.go](main.go) is a short Hello-style smoke for the framework module (not the browser todo SoT).

## Development

Tool versions are pinned in [.tool-versions](.tool-versions) when present (asdf: Go, golangci-lint, Task, Node); otherwise match `go.mod` / CI.

```bash
task          # static analysis (default)
task lint     # golangci-lint
task lint-fix # golangci-lint with autofix
task fix      # go fix modernizers
task generate # go generate ./...
task test     # go test -race ./...
```

CI runs lint and tests on every push. After `forma init`, the scaffold ships the React host (`web/src/host/`), GoScript bridge wiring (`formaHost` then `/app.js`), and **`forma dev`** (Vite + GoScript rebuild).
