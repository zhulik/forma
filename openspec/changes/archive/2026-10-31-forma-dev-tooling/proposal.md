# Proposal

## Why

forma is greenfield with no linter, CI, Taskfile, or Cursor rules. Sister projects (hosehook, TubeGremlin) already enforce consistent Go quality, agent guidance, and CI gates; adopting that baseline now avoids drift before P1 DSL work and keeps humans and agents on the same standards.

## What Changes

- Add `.golangci.yaml` aligned with hosehook (module path `github.com/zhulik/forma`, path exclusions for `ui/*.gen.go` when present).
- Add root `Taskfile.yaml` with `default`, `lint`, `generate`, `fix`, `check-generate`, and `check-fix` (hosehook pattern); include a `test` target for later use but **do not** wire tests into `default` or CI yet.
- Add GitHub Actions CI (backend job: generate check, golangci-lint, fix check — **no test step** until the project has tests worth running).
- Copy and adapt Cursor rules: shared Go quality rules from hosehook/TubeGremlin, hosehook-only backend hygiene rules, and `early-stage-no-compat` tailored to forma; defer pal-specific and TubeGremlin-only rules until those structures exist.
- **No new test files or test dependencies** in this change (nothing meaningful to assert yet).
- Document tooling in `doc/ROADMAP.md` status or a short note in change design (no separate user guide yet).

## Capabilities

### New Capabilities

- `dev-tooling`: Local and CI engineering gates (Task, golangci-lint, generate/fix checks) and contributor Cursor rule coverage for Go code.

### Modified Capabilities

- (none)

## Impact

- **Code**: New config and workflow files only; no `_test.go` files.
- **Dependencies**: No Ginkgo/Gomega or other test-only modules added.
- **Agents**: `.cursor/rules/` becomes active for Go edits (including future-oriented testing rules, inactive until tests exist).
- **Out of scope**: Tests, React `host/` lint/CI, pal/DI rules, Ent/Connect/infra rules, Helm/Docker (later changes).
