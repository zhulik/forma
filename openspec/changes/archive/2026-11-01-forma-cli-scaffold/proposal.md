# Proposal

## Why

The `forma-cli` capability described target commands but the binary did not exist. Authors and CI could not verify install → init → build. This change ships a minimal CLI before P3 host/codegen: scaffold, stub sync, and CI smoke so runtime work plugs into a real module layout.

## What Changes

- Nested-module **`forma` CLI** under `cmd/forma/` using **urfave/cli/v3** and **pal** (InjectSlog, explicit service lists—no global registration).
- **`forma init`**: empty-directory full scaffold from embedded **`minimal`** template; in-place writes; stub **`forma sync`** after success.
- **`forma init --only-config`**: write **`forma.yaml`** when missing; fail if present (non-empty dirs allowed).
- **`forma version`**, root **`--log-level`** / **`FORMA_LOG_LEVEL`**; stub **`add`** / **`dev`** with clear not-implemented errors.
- **`FORMA_REPLACE`** for local `go mod edit -replace` during init/sync (development and CI).
- Cross-package DI interfaces in **`internal/core/interfaces.go`** (`GoModRunner`, `ProjectSyncer`).
- **CI / Task**: `check-init-scaffold`; root **`task test`** / **`task lint`** include `cmd/forma`.
- **`doc/user/`** (install, CLI, manual setup), **ROADMAP**, **`.cursor/rules/forma-cli.mdc`**.

## Capabilities

### New Capabilities

_(none — behavior extends existing capabilities)_

### Modified Capabilities

- **`forma-cli`**: Init rules, `--only-config`, templates, stub sync, version, nested CLI module.
- **`dev-tooling`**: Init scaffold smoke in CI.
- **`ui-docs`**: Manual setup; CLI availability vs planned commands.

## Impact

- **ROADMAP**: CLI scaffold delivered ahead of P3a; `add` / full sync / `dev` remain P3.
- **Dependencies**: `cmd/forma/go.mod` (`urfave/cli/v3`, `pal`, …); root library module unchanged.
- **Author docs**: Partial CLI availability documented; browser loop still planned.
