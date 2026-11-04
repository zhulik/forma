# Design

## Context

See proposal.md — Why. Codegen analyzes TSX with tree-sitter and emits Go + `registry.gen.ts`. Golden fixtures under `testdata/goldens/` pin leaf (button) and compound (accordion) output. `github.com/zhulik/forma/fn` provides collection helpers; CLI already imports it. `cmd/forma/pkg/` previously held only `cmdexec`.

## Goals / Non-Goals

**Goals:**

- Readable, typed Go emission and a clear registry template
- Shared naming helpers with unit tests outside unexported analyze internals
- Prefer `fn` for slice transforms; leave AST index walks and I/O error loops as loops
- Preserve deterministic, golden-matching codegen behavior

**Non-Goals:**

- Changing TSX classification heuristics, prop mapping rules, or override hooks
- Splitting analyze into a `tsx` subpackage
- Author-facing CLI / docs / forma.yaml changes
- Removing the post-emit `gofmt` step

## Decisions

### 1. Go emit via jennifer; registry via `text/template`

- **Choice:** `github.com/dave/jennifer/jen` builds `*.gen.go` (`emit_go.go`). `text/template` renders `registry.gen.ts` (`emit_registry.go`).
- **Why:** Jennifer fits Go AST construction and import aliasing; templates fit TypeScript registry output.
- **Notes:** Jennifer may omit the blank line between stdlib and third-party imports; `ensureImportGroupBlank` restores goimports-style grouping so goldens stay byte-identical. Compound `Node` literals use jennifer `Custom` with `Multi: true` for field order/layout matching goldens.
- **Alternatives:** All templates; keep string builders; jennifer only (cannot emit TS).

### 2. Naming package `cmd/forma/pkg/ident`

- **Choice:** Export `ToPascal`, `ToExportedGoName`, and `ReceiverName` under `pkg/ident` with external-package unit tests.
- **Why:** Generic string/ident utilities; mirrors `pkg/cmdexec`; unlocks tests without violating external-test-package rules.

### 3. Dedup and collection helpers in `fn`

- **Choice:** Add `fn.UniqSorted` (replaces local `uniqSorted`). Also add `fn.FlatMap` and `fn.GroupBy` for registry import aggregation (`FlatMap` widget exports → `GroupBy` path → sorted keys via `slices.Sorted(maps.Keys(...))`).
- **Why:** Prefer-fn; registry import map-building was loop-heavy; helpers are generally reusable.

### 4. File split inside `internal/codegen`, no new analyze package

- **Choice:** Keep package `codegen`; emit split into `emit_go.go` and `emit_registry.go`; analyze stays one package (helpers in-file). Registry template data reuses `[]SlotField` directly — no parallel `registrySlot` DTO.
- **Why:** Scope stays a tidy-up; duplicate slot types added no value.

### 5. Loops that stay loops

Keep `for` when walking tree-sitter `NamedChildCount`, when collecting with early error return from I/O, and for short-circuit root-candidate search in `pickRoot`. Convert list filters/maps (tsx listing, prefixed exports → slots, emit field lists, gofmt/prune file lists, registry data prep) to `fn` where semantics match.

## Risks / Trade-offs

- **[Jennifer import / layout drift]** Mitigated by `ensureImportGroupBlank`, explicit multi-line `Node` literals, and golden tests.
- **[Whitespace / gofmt drift]** Runner still runs `gofmt` after write; goldens compared post-gofmt.
- **[Template determinism]** Registry sorts import paths and names before execute.

## Migration Plan

In-tree only: add dependency, refactor, run codegen golden tests + `task test` / `task lint`. No deploy or dual-write. Rollback = revert the change commit.
