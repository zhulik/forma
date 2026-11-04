# Tasks

## 1. Durable rule

- [x] 1.1 Add a “Provide lives in services.go” section to `.cursor/rules/pal-runners.mdc` (filename plural; allow `ServiceDef` or `[]ServiceDef`; forbid Provide beside type files) and verify the section is present and clear.

## 2. Sweep Provide placement

- [x] 2.1 Move any package-level `Provide` that is not in `services.go` into that package’s `services.go` (expected: `codegen`), leave composition roots unchanged, and verify `rg 'func Provide\\('` under `cmd/forma` only hits `**/services.go`.
- [x] 2.2 Confirm `gotool` and `webtool` already expose `Provide` from `services.go` and verify those files still compile with their callers.

## 3. Verification

- [x] 3.1 Run `task test` and `task lint` from the repo root and verify both succeed.

## Workflow follow-up

- Archive the change after review (`/opsx-archive` or `/opsx-finalize`). No main-spec sync expected (`skip_specs: true`).
