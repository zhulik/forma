# Install

## Prerequisites

| Tool | Why |
|------|-----|
| **Go** | Application code and (later) the `forma` CLI. Match the version in [`.tool-versions`](../../.tool-versions) or `go.mod`. |
| **Node.js** | Host app, shadcn sources, and codegen (once P3 lands). |
| **GopherJS** | Compile Go to JavaScript for the browser loop (planned P3c). |

Optional: [asdf](https://asdf-vm.com/) with `.tool-versions` for pinned Go, Node, and Task.

## Install the forma CLI

> **Planned.** The command below is the target install path. The binary is not in this repository yet; see ROADMAP P3 and the `forma-cli` OpenSpec capability.

```bash
go install github.com/zhulik/forma/cmd/forma@latest
```

Verify:

```bash
forma version
```

## Create your first app

> **Planned** until `forma init` ships. This is the intended three-step flow after the CLI exists.

```bash
mkdir myapp && cd myapp
forma init github.com/you/myapp
forma dev
```

What **`forma init`** will do:

1. **`go mod init`** with your module path (no git remote setup).
2. Write a minimal scaffold: `main.go`, `forma.yaml`, `host/`, default imports of `github.com/zhulik/forma/ui` and `fn`.
3. Run **`forma sync`** so host deps, shadcn sources, generated `internal/ui` bindings, registry, and **`go mod tidy`** are applied.

Use your real module path (e.g. `github.com/you/myapp`), not a placeholder—you will import it from generated code under `internal/ui`.

## Learn the DSL today (no CLI)

While the CLI and browser loop are in progress:

1. Clone or browse [github.com/zhulik/forma](https://github.com/zhulik/forma).
2. Read [`main.go`](../../main.go) (todo demo).
3. Run `task test` and `task lint` from the repo root.

That exercises the **authoring model** only; `UIMount` in this repo is still a stub until P3.

## Next steps

- [CLI reference](cli.md)
- [forma.yaml](forma-yaml.md)
- [Writing UI](writing-ui.md)
