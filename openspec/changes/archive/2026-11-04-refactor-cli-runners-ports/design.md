# Design

## Context

See proposal.md — Why. Pre-change: thin proxy `runner` types forwarded to `Service` / ports; `codegen.Runner` and `*Runner` ports were not `pal.Runner`; `webtool.Tool` bundled npm + shadcn; init injected `ProjectSyncer` in the same pal graph as the init runner.

## Goals / Non-Goals

**Goals:**

- One naming scheme: port interface and main concrete type share the role name; no `Service` / `Tool` / `adapter` for those ports.
- Structs named `Runner`/`runner` implement `pal.Runner`; library types are not named `Runner`.
- No proxy pal runners whose `Run` only delegates to another component.
- Split npm vs shadcn into two ports in `webtool`.
- Capture conventions as Cursor durable rules; sweep the whole `cmd/forma` tree to match.

**Non-Goals:**

- Changing CLI flags, exit codes, or author docs.
- Splitting `GoModWrapper` into multiple ports.
- New product capabilities or OpenSpec main-spec deltas.

## Decisions

1. **Port + impl name map**

   | Port (`core`) | Package impl | Methods |
   |---------------|--------------|---------|
   | `GoModWrapper` | `gotool.GoModWrapper` | `InitModule`, `Require`, `Replace`, `Tidy` |
   | `NpmWrapper` | `webtool.NpmWrapper` | `Install` |
   | `ShadcnWrapper` | `webtool.ShadcnWrapper` | `Add` |
   | `CodeGenerator` | `codegen.CodeGenerator` | `Generate` |
   | `ProjectSyncer` | `synccmd.ProjectSyncer` | `Sync` (+ `Run` as pal entry for sync CLI) |

   **Alternatives:** keep `*Runner` / `Tool` — rejected (clashes with pal lifecycle naming). Split packages for npm/shadcn — rejected (keep one `webtool` package).

2. **Avoid pal lifecycle method names on ports**

   Do not name domain methods `Init`, `Run`, `Shutdown`, or `HealthCheck` on non-runner collaborators. `go mod init` → `InitModule`. `Run` remains only on intentional `pal.Runner` types (init/sync/generate entrypoints).

3. **Kill proxy runners**

   - **sync:** CLI-bound fields (`Dir`, `FormaReplace`, `AddComponents`) live on `ProjectSyncer`. `Run` → `Sync` with those fields. Drop `synccmd.runner`.
   - **generate:** `generatecmd.Runner` owns load-config + codegen (`Run` / `Generate`); drop the old `Service` + proxy `runner`.
   - **init:** keep `initcmd.runner` (real work); inject `GoModWrapper` only; after scaffold call `synccmd.RunWithComponents` (nested pal) so `ProjectSyncer` is never a passive collaborator with `Run`.

4. **Codegen adapter**

   Implement `core.CodeGenerator` directly on `codegen.CodeGenerator` (map `core.CodegenRequest` ↔ package request inline or unify request types). Remove `adapter` and rename `runner.go` → `code_generator.go` (one-type-one-file).

5. **Provide wiring**

   - `gotool.Provide()` → `pal.Provide[core.GoModWrapper](&GoModWrapper{})`
   - `webtool.Provide()` → `[]pal.ServiceDef` for `NpmWrapper` + `ShadcnWrapper`
   - `codegen.Provide()` → `pal.Provide[core.CodeGenerator](&CodeGenerator{})`
   - Sync CLI / `Run*` helpers: `pal.Provide(&ProjectSyncer{Dir:…})` as the sole sync runner (has `Run` + `Sync`). Do **not** also register an empty `ProjectSyncer` in init’s graph (pal would start `Run` on it). Init calls `synccmd.RunWithComponents` (nested pal) after scaffold.

6. **Durable rules**

   Add/update `.cursor/rules/` (prefer focused rule(s) + touch `forma-cli.mdc`):
   - Runner structs must implement `pal.Runner`.
   - No proxy runners (promote the component).
   - Ports and main impls share role names; no `Service`/`Tool`/`adapter` for ports; method verbs; no pal lifecycle name clashes on non-runners.

## Risks / Trade-offs

- **[Nested pal for init→sync]** Init starts a second pal graph via `RunWithComponents` → Mitigation: acceptable early-stage cost; avoids registering a `ProjectSyncer` with `Run` as a passive collaborator.
- **[Generate type name]** `generatecmd.Runner` implements `pal.Runner` and owns generate work (not a proxy).
- **[Test churn]** Recording fakes rename with ports → Mitigation: update tests in the same change; run `task test` / `task lint`.

## Migration Plan

In-place rewrite under `cmd/forma` (early-stage, no compat shims). No deploy/rollback.

## Open Questions

None — naming and method verbs locked in exploration.
