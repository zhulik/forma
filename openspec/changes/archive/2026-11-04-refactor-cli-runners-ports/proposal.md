# Proposal

## Why

CLI pal wiring uses thin proxy `runner` types that only forward to `Service`/`Tool` collaborators, and several types named `Runner` do not implement `pal.Runner`. That naming and layering confuse DI ports with process lifecycle and makes durable conventions hard to follow.

## What Changes

- **BREAKING** (internal `cmd/forma` only): rename DI ports and concrete types to role names that match across `core` and implementing packages; drop `Service` / `Tool` / `adapter` aliases for those ports.
- Split `WebSyncRunner` into `NpmWrapper` + `ShadcnWrapper` (same `webtool` package).
- Rename `GoModRunner` → `GoModWrapper`, `CodegenRunner`/`codegen.Runner` → `CodeGenerator`; keep `ProjectSyncer` name but implement as `synccmd.ProjectSyncer` (not `Service`).
- Shorten port methods to verbs (`InitModule`, `Require`, `Replace`, `Tidy`, `Install`, `Add`, `Generate`, `Sync`) without clashing with pal lifecycle (`Init`, `Run`, `Shutdown`, `HealthCheck`).
- Remove generate/sync proxy `runner` structs: put CLI-bound fields on the real component and make that type the `pal.Runner` (`ProjectSyncer`, `generatecmd.Runner`).
- Init calls nested `synccmd.RunWithComponents` after scaffold (avoids registering `ProjectSyncer` with `Run` as a passive collaborator).
- Add Cursor durable rules for pal runners, no proxy runners, and port/impl naming.
- Author-facing CLI behavior, flags, and docs unchanged (explicit opt-out of `doc/user/` updates).

## Capabilities

### New Capabilities

<!-- none -->

### Modified Capabilities

<!-- none — no product/capability behavior change; see skip_specs in .openspec.yaml -->

## Impact

- `cmd/forma/internal/core/interfaces.go` and all injectors/callers (`gotool`, `webtool`, `codegen`, `synccmd`, `initcmd`, `generatecmd`, tests).
- Cursor rules under `.cursor/rules/` (new or update `forma-cli.mdc`).
- No change to published author CLI UX, `forma.yaml`, or OpenSpec product specs.
