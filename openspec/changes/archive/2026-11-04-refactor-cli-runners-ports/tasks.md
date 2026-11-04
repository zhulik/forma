# Tasks

## 1. Durable rules

- [x] 1.1 Add or update `.cursor/rules/` (and sync `forma-cli.mdc` Pal section) for: runner structs must implement `pal.Runner`; no proxy runners; ports + main impls share role names (no `Service`/`Tool`/`adapter` for ports); short verb methods that do not clash with pal lifecycle (`Init`/`Run`/`Shutdown`/`HealthCheck`). Verify the rule files exist and state those constraints clearly.

## 2. Core ports and leaf wrappers

- [x] 2.1 Rewrite `internal/core/interfaces.go` to the new port names/methods (`GoModWrapper`, `NpmWrapper`, `ShadcnWrapper`, `CodeGenerator`, `ProjectSyncer`) and verify `go build` under `cmd/forma` fails only at expected call sites until callers are updated.
- [x] 2.2 Rename `gotool.Tool` → `GoModWrapper` with `InitModule`/`Require`/`Replace`/`Tidy`; update `Provide`; verify package tests or compile of `gotool`.
- [x] 2.3 Split `webtool.Tool` into `NpmWrapper` (`Install`) and `ShadcnWrapper` (`Add`); update `Provide` to register both; verify `webtool` compiles.
- [x] 2.4 Rename `codegen.Runner` → `CodeGenerator`, remove `adapter`, implement `core.CodeGenerator` on that type, rename source file to match; update tests; verify `go test ./internal/codegen/...` under `cmd/forma`.

## 3. Sync / generate / init wiring

- [x] 3.1 Rename `synccmd.Service` → `ProjectSyncer`, inject new ports, drop proxy `runner`, put CLI fields on `ProjectSyncer` with `Run` + `Sync`; update helpers/`Provide`/tests; verify sync package tests pass.
- [x] 3.2 Fold `generatecmd.Service` + proxy `runner` into one pal-runner type (not named `Service`); update tests; verify generate package tests pass.
- [x] 3.3 Update `initcmd.runner` and `deps` for new port field names/methods; verify init-related tests still pass.
- [x] 3.4 Sweep remaining `cmd/forma` references (`*Runner`, `Tool`, `Service` aliases, old method names) and verify `rg` finds no stale port names.

## 4. Verification

- [x] 4.1 Run `task test` and `task lint` from the repo root and verify both succeed.

## Workflow follow-up

- Archive the change after review (`/opsx-archive` or `/opsx-finalize`). No main-spec sync expected (`skip_specs: true`).
