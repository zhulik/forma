# Design

## Context

See [proposal.md](./proposal.md). Shipped the first **`cmd/forma`** binary: nested module, pal, embedded **`minimal`** template, init (+ `--only-config`), stub sync, CI smoke, and author doc updates. Reference: pal CLI (`cmd/pal`) for urfave/cli v3 + embed templates; forma differs on in-place init (no staging/promote).

## Goals / Non-Goals

**Goals:** (delivered)

- **`go install github.com/zhulik/forma/cmd/forma@latest`**
- Full init: empty dir, in-place, **`minimal`**, stub sync
- Config-only init; multi-template embed pipeline (v1: **`minimal` only**)
- Pal DI via **`app.Run`** + explicit `pal.ServiceDef` lists; **`InjectSlog()`**
- **`internal/core/interfaces.go`**: `GoModRunner`, `ProjectSyncer`
- CI: **`task check-init-scaffold`** with in-repo temp dir + **`FORMA_REPLACE`**

**Non-Goals:** Full sync (npm/codegen/host), **`add`**, **`dev`**, wizard, **`git init`**, second template in v1.

## Decisions

### 1. urfave/cli v3 + pal

Thin `Action` → `app.Run(ctx, …)` with `pal.Provide(&runner{…})`. Only CLI-bound fields on the runner struct; **`GoModRunner`** and **`ProjectSyncer`** injected by pal.

### 2. No global service registration

`app.Run` accepts explicit service defs. Init uses **`deps.Services(pal.Provide(&runner{…}))…`**; synccmd inlines **`gotool.Provide(), synccmd.Provide(), pal.Provide(…)`** to avoid an import cycle with **`deps`**.

### 3. Nested module `cmd/forma/go.mod`

CLI deps isolated; root **`task test`** / **`task lint`** recurse into `cmd/forma`.

### 4. Embedded templates

Package **`cmd/forma/templates`**: `//go:embed all:scaffolds`; paths under **`scaffolds/<name>/`**; **`.tmpl`** suffix stripped on write; **`text/template`**.

### 5. Stub sync

`RequireForma` via `go mod edit -require`; optional **`-replace`** from **`FORMA_REPLACE`**; `go mod tidy`. Debug log notes host/codegen skipped.

### 6. Layout (as built)

```
cmd/forma/
  main.go
  go.mod
  templates/              # embed + Apply
  internal/cli/           # root, deps/, version tests
  internal/cli/app/       # Run + InjectSlog
  internal/cli/subcommands/{init,sync,add,dev,version}/
  internal/core/          # interfaces.go
  internal/gotool/
  internal/formaconfig/
  pkg/cmdexec/
```

### 7. CI smoke

Temp dir under repo root (`.tmp-scaffold.*`) so asdf Go resolves; **`FORMA_REPLACE=$root`**; no separate `go mod edit -replace` in Task (init sync applies replace).

## Risks / Trade-offs

- Partial init on failure → documented in user guide / troubleshooting.
- Stub sync vs full sync → cli.md status table.
- Unpublished module → **`FORMA_REPLACE`** until tagged releases.

## Migration Plan

Archive syncs **`forma-cli`** purpose and requirements. ROADMAP marks scaffold in progress / delivered.

## Open Questions

_(none — resolved in implementation: empty **`ui.components`**, root Task includes CLI module.)_
