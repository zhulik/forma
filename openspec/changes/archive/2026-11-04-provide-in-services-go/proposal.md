# Proposal

## Why

Package-level `Provide` helpers that return `pal.ServiceDef` (or slices) are scattered—sometimes next to the type (`code_generator.go`), sometimes in `services.go`. Agents and humans need one durable home so DI registration stays discoverable and type files stay about behavior.

## What Changes

- Require every package-level `Provide` that returns `pal.ServiceDef` or `[]pal.ServiceDef` to live in that package’s `services.go`.
- Move any misplaced `Provide` (notably `codegen`) into `services.go`; leave composition roots assembling defs as they do today.
- Capture the convention in `.cursor/rules/pal-runners.mdc`.
- Author-facing CLI behavior, flags, and docs unchanged (explicit opt-out of `doc/user/` updates).

## Capabilities

### New Capabilities

<!-- none -->

### Modified Capabilities

<!-- none — no product/capability behavior change; see skip_specs in .openspec.yaml -->

## Impact

- `cmd/forma` packages that export `Provide` (`codegen`, `gotool`, `webtool`; future packages follow the same layout).
- `.cursor/rules/pal-runners.mdc` (Provide / `services.go` section).
- No change to published author CLI UX, `forma.yaml`, or OpenSpec product specs.
