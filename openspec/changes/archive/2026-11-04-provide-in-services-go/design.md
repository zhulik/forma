# Design

## Context

See proposal.md — Why. Packages under `cmd/forma` register pal defs via package-level `Provide`. `gotool` and `webtool` already use `services.go`; `codegen` had `Provide` in `code_generator.go`. Cursor rule `pal-runners.mdc` is the durable home for related pal conventions.

## Goals / Non-Goals

**Goals:**

- One file per package for package-level `Provide`: always `services.go`.
- Allow `Provide() pal.ServiceDef` or `Provide() []pal.ServiceDef` (multi-def packages like `webtool`).
- Durable Cursor rule so new packages follow without rediscovery.
- Sweep existing `Provide` sites into compliance.

**Non-Goals:**

- Changing what is registered, injection keys, or composition-root assembly (`app.Run`, `deps.Services`).
- Renaming `Provide` or introducing a global registry.
- Spec deltas or author-doc updates.

## Decisions

1. **Filename: `services.go` (plural)**

   Matches existing `gotool`/`webtool` and the example in `no-delegating-methods.mdc`. Singular `service.go` rejected.

2. **Return type: single or slice**

   Packages with one def return `pal.ServiceDef`; packages registering several return `[]pal.ServiceDef`. Do not force one-def-per-function or split files.

3. **Rule placement**

   Extend `.cursor/rules/pal-runners.mdc` (same `cmd/forma/**` globs) with a “Provide lives in services.go” section rather than a near-duplicate rule file.

4. **Composition roots stay thin**

   Subcommand Actions and helpers call `pkg.Provide()`; they do not define other packages’ `Provide` bodies.

## Risks / Trade-offs

- **[Already partially applied]** Rule + `codegen/services.go` may already exist before apply → Mitigation: tasks verify presence and sweep leftovers; no double-move.
- **[Name collision]** Unrelated `services.go` content → Mitigation: file is only for package `Provide` registration in this convention; keep it small.

## Migration Plan

In-place under `cmd/forma` (early-stage, no compat shims). No deploy/rollback.

## Open Questions

None — filename and return-type flexibility locked in exploration.
