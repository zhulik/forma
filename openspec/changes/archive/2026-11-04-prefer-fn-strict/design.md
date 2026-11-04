# Design

## Context

See proposal.md — Why. Today `fn` lives in a single `fn.go` with slice-only helpers; codegen already uses them heavily. A handful of raw loops remain: fallible collect in analyze, side-effect each in runner write/prune, tree-sitter index walks, short-circuit `pickRoot`, string builder in `ident.ToPascal`, and a table-driven test. Prefer-fn is soft guidance only.

## Goals / Non-Goals

**Goals:**

- Split `fn` into `slices.go` / `maps.go` with clear naming (no Go overload of `Map`/`Filter` for maps).
- Add fail-fast `MapErr` / `FilterMapErr` and map-entry analogues; convert collecting fallible call sites.
- Strict prefer-fn allowlist + no-op proxy ban in the Cursor rule.
- Keep author docs discoverable for the expanded catalog.

**Non-Goals:**

- `EachErr` or forcing side-effect walks through `fn`.
- Custom golangci analyzer banning `for`.
- Error aggregation (`[]error` / `errors.Join` inside helpers).
- Third-party collection libraries.

## Decisions

1. **Each stays raw `for`** — Side-effect walks (write files, delete stale, pure I/O without building a returned collection as the purpose) remain allowlisted. No `Each` / `EachErr` in `fn`.
   - Alternative considered: `EachErr` — rejected; adds API surface for something Go already expresses clearly.

2. **Fallible collect uses `MapErr` / `FilterMapErr`** — AnalyzeComponentsDir-style loops that build `[]T` and can fail convert even when they also do I/O inside the callback (parser closed over is fine).
   - Fail-fast: first error returns; no partial slice.

3. **Map helpers naming** — Same package, distinct names: `MapValues`, `MapEntries`, `FilterEntries`, `FilterMapEntries`, plus `*Err` variants as needed. Stdlib `maps.Keys` / `maps.Values` stay for key/value extraction.
   - Alternative: subpackage `fn/maps` — rejected; authors already dot-import `fn`.

4. **File split** — `slices.go` (existing + Err), `maps.go` (map analogues), tests split or kept as `fn_test.go` / `slices_test.go` + `maps_test.go` as convenient.

5. **No-op proxy ban** — Rule examples: `fn.Map(items, toJSON)` not `fn.Map(items, func(item Item) JSON { return toJSON(item) })`. Applies when signatures match exactly.

6. **Allowlist** (prefer-fn.mdc):
   - Inside `fn/`
   - Generated `*.gen.go` (and similar)
   - Table-driven tests
   - Side-effect-only each (slice or map)
   - Channels / `select` / infinite loops; string rune walks when needed
   - Indexed count APIs (no slice to range)
   - Short-circuit search
   - Index-heavy algorithms

7. **Docs** — Brief updates in `doc/user/writing-ui.md` and/or `project-layout.md`; no new dedicated page.

## Risks / Trade-offs

- [Hybrid I/O+collect looks like "each"] → Mitigation: purpose test — if the function returns a built collection, use `*Err` helpers.
- [Map helpers unused initially] → Acceptable; ships for rule completeness and future call sites.
- [Strict rule without lint] → Agents follow Cursor rules; humans rely on review. Analyzer deferred.

## Migration Plan

In-repo only: rewrite call sites in the same change. No deployment. Early-stage — additive `fn` API, no dual paths.
