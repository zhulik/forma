# forma

Go-native UI framework (in development): compose semantic widgets in Go; a React + shadcn host renders them. See [doc/DESIGN.md](doc/DESIGN.md) and [doc/ROADMAP.md](doc/ROADMAP.md).

The reference DSL demo is [main.go](main.go).

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

CI runs `check-generate`, lint, and `check-fix` on every push. Tests are deferred until there is behavior worth asserting.
