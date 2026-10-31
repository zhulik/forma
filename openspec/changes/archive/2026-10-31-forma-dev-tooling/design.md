# Design

## Context

forma has Go 1.27.3 and `.tool-versions` pins matching hosehook/TubeGremlin but no `.golangci.yaml`, `Taskfile.yaml`, GitHub CI, or `.cursor/rules/`. See [proposal.md](./proposal.md) for motivation.

## Goals / Non-Goals

**Goals:**

- Reuse proven hosehook backend tooling shape (golangci config, Task targets, CI job order minus tests).
- Copy Cursor rules that apply to a small Go library + root demo, adapted for `github.com/zhulik/forma` and `doc/ROADMAP.md`.
- Keep CI green without adding test files (lint + generate/fix checks only).

**Non-Goals:**

- Adding tests, Ginkgo/Gomega, or CI test steps in this change.
- TubeGremlin depguard matrix (Ent, core-only packages, Connect protos).
- pal/DI rules until `cmd/` uses pal.
- Frontend/host lint rules until `host/` exists (separate change; adapt TubeGremlin `frontend-*` globs later).
- Docker, Helm, release counters, issue workflows.

## Decisions

### 1. golangci-lint baseline: hosehook, not TubeGremlin

**Choice:** Copy hosehook `.golangci.yaml` structure (`default: all`, same disabled linters, strict issue limits, goimports local prefix `github.com/zhulik/forma`).

**Rationale:** forma has no `internal/core`, Ent, or HTTP edge yet; TubeGremlin rules would be dead weight.

**Additions:**

- Path exclusion for `ui/**/*.gen.go` under `exclusions.paths` (and formatter exclusions if needed) so generated shadcn bindings do not fight exhaustruct/revive unnecessarily.
- `depguard` rules only for `cmd/` import ban and future `pkg/` → no `internal/` when those directories appear (copy hosehook pattern with forma module path).

**Alternatives:** Minimal golangci enable-list — rejected; sister repos use `default: all` with explicit disables.

### 2. Taskfile: hosehook backend subset without test in default

**Choice:** `default` → `lint` only; `generate`, `fix`, `_check-tree`, `check-generate`, `check-fix`, `lint`, `lint-fix`; optional `test` → `go test -race ./...` for future use (document in README, not in default or CI).

**Rationale:** User directive: no tests until there is behavior worth testing. Empty `go test` passing is not useful signal.

**Note:** `generate` runs `go generate ./...` even if empty until codegen lands; CI `check-generate` still passes.

### 3. CI: backend job without tests

**Choice:** `.github/workflows/ci.yaml` like hosehook but **omit** the Test step: setup-go (go.mod), mod cache, task, `check-generate`, golangci-lint-action (5m timeout), `check-fix`.

**Alternatives:** Run `go test` with no test files — rejected; add test job when first real specs land (likely P1 DSL or P3 runtime).

### 4. Cursor rules inventory

Copy from **both** repos where identical or hosehook is source of truth; edit product names in `early-stage-no-compat.mdc`.

| Rule file | Source | Notes |
|-----------|--------|--------|
| `early-stage-no-compat.mdc` | hosehook | Replace product name; point to `doc/ROADMAP.md` |
| `modern-go.mdc` | either | |
| `err-variable.mdc` | either | |
| `error-wrapping.mdc` | either | |
| `one-type-one-file.mdc` | either | Important for `ui/*.gen.go` + hand types |
| `copy-dont-mutate.mdc` | either | Fits immutable UI trees |
| `naming.mdc` | either | |
| `no-interface-assertions.mdc` | either | |
| `no-env-in-libraries.mdc` | either | |
| `verify-before-done.mdc` | either | |
| `keep-files.mdc` | either | |
| `testpackage.mdc` | either | Applies when tests exist |
| `ginkgo-testing.mdc` | hosehook | Applies when `*_test.go` added later |
| `test-slog.mdc` | either | |
| `no-speculative-guards.mdc` | hosehook | |
| `evidence-based-guards.mdc` | TubeGremlin | Complements above |
| `error-path-at-caller.mdc` | hosehook | |
| `no-context-in-structs.mdc` | hosehook | |
| `no-delegating-methods.mdc` | hosehook | |
| `no-init-wiring-guards.mdc` | hosehook | |
| `distinct-types.mdc` | TubeGremlin | Useful for domain IDs later |

**Deferred (copy in later change):**

- `pal-components`, `cross-package-di`, `mock-core-ports`, `cli-config-validation`, `no-manual-init` — when pal + `cmd/` exist.
- `frontend-*`, `docs-diagrams` (forma globs), Ent/infra/protobuf rules.

**hooks.json:** Copy empty `{ "version": 1, "hooks": {} }` for parity.

### 5. Tests deferred

**Choice:** Do not add `_test.go` files or test dependencies in this change. Wire CI and `task default` without a test step.

**Follow-up:** When P1 DSL or runtime behavior is specified, a separate change adds Ginkgo (per sister repos) and re-enables `test` in default + CI.

### 6. ROADMAP note

Add a short row or footnote under phased delivery: foundation tooling (this change) precedes P1 DSL; not a replacement for P1.

## Risks / Trade-offs

- **[Risk] golangci strict `default: all` fails on stub ui code** → Run lint during apply; fix or narrow exclusions for placeholder APIs.
- **[Risk] No CI test gate until later** → Acceptable until meaningful behavior exists; document in ROADMAP follow-up.
- **[Risk] Rules reference testing before tests exist** → Ginkgo/test rules use `*_test.go` globs; harmless until files appear.

## Migration Plan

Single PR: add configs, rules, CI (no tests). No runtime or DSL behavior change. Contributors run `task` locally after pulling.

## Open Questions

- Whether to add `README.md` section "Development" in this change or wait for P2 user docs (tasks include minimal README pointer).
