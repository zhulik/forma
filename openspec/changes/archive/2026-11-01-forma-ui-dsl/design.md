# Design

## Context

See [proposal.md](./proposal.md) for motivation. Today `main.go` and `ui/` demonstrate composition with `.With`, `.Or`, dot imports, and handlers, but string children, untyped `.Or` return types, and missing keys diverge from the intended P1 contract. Runtime (`UIMount`, React host) stays stubbed.

## Goals / Non-Goals

**Goals:**

- Encode the explored P1 decisions in types and the reference demo.
- Make `.Or` type-safe on every hand-written widget stub (template for future codegen).
- Replace `...any` children with `Component` only; add `ui.If` (lazy) and `Context.Update`.
- Add `Key` on `BaseComponent` and wire keys in the todo list.
- Align `doc/DESIGN.md` with text and context rules where it still mentions string `.With`.

**Non-Goals:**

- Text input controlled/deferred sync policy (demo may keep per-keystroke refresh).
- `Fragment`, bridge `js` tags, `Node` serialization details, codegen pipeline, React host.
- Tests (project has no test wiring in CI yet; follow repo when tests land).

## Decisions

### 1. `With` / `Or` signatures use `...Component`

**Choice:** `With(children ...Component)` and `Or(fallback ...Component)` on `BaseComponent`; wrappers delegate and return their own type.

**Alternatives:** Keep `...any` with string coercion — rejected (agents mix types; runtime panics).

**Implementation note:** `toComponents` becomes unnecessary or accepts only `Component`. List helpers return `[]Component`; splat into `With` unchanged in Go.

### 2. Typed `Or` on each wrapper

**Choice:** Mirror existing `With` boilerplate in each `*.gen.go`: copy embedded `BaseComponent`, call embedded `Or`, return `Row` / `Column` / etc.

**Alternatives:** Generic base struct — rejected for v1 (codegen and hand stubs stay explicit).

**Fixes:** `Column{Gap: 3}.With(...).Or(...)` keeps `Gap` because `Render` stays on `Column`.

### 3. Text: `Text{Text: "hello"}`

**Choice:** Exported **`Text`** struct with embedded **`BaseComponent`**; copy via **`Text{Text: "…"}`** (Go does not allow one-element **`Text{"…"}`** literals with an embedded base). Styling via struct fields (`Muted`, `Strikethrough`, `Class`, …). `Heading{Text: "…"}` stays field-first. Button labels: `.With(Text{Text: "Add Task"})`.

**Alternatives:** `Text()` constructor — tried and reverted; string `.With` — rejected.

### 3b. `fn.FilterMap` in child rendering

**Choice:** `renderChildren` uses **`fn.FilterMap`** to drop nil nodes after `Render` (conditional **`If`** / skip).

### 4. `If(cond bool, thunks ...func() Component)` in package `ui`

**Choice:** Package-level helper (dot-imported in apps as `If(...)`). Arity 1 or 2 thunks; more → panic. Lazy: only the taken branch runs. False with one thunk → sentinel **`skip`** (or equivalent) whose render produces **no node**; **`renderChildren` omits** skipped/nil contributions.

Example:

```go
Button{Variant: "default"}.With(
    If(showFoo,
        func() Component { return Text{Text: "foo"} },
        func() Component { return Text{Text: "bar"} },
    ),
)
```

**Alternatives:** Method `Button{}.If(...)` — rejected (author prefers free function, not receiver chain). **`fn.If`** — rejected; `fn` stays generic (`Map`, `Filter`), forma DSL lives in `ui`.

**File:** e.g. `ui/if.go` next to composition types.

### 5. `Context.Update(fn func())`

**Choice:** Run `fn()`, then `RefreshUI()` once if `fn` returns normally. `RefreshUI` inside `fn` discouraged but not forbidden (idempotent refresh acceptable in stub).

**Alternatives:** Update without refresh — rejected (defeats purpose).

**Stub:** `context.go` adds method on concrete stub type implementing `Context`; or embed in demo-only stub until runtime — implement minimal stub in `ui` for compile.

### 6. Keys

**Choice:** `Key string` on `BaseComponent`, copied to `Node` in `Render` (add `Key` field on `Node` if not present).

**Host (later):** Missing key on repeated siblings → dev warning; not specified in this change beyond carrying key on node.

### 7. Context split

**Choice:** Document in DESIGN: `Render(context.Context)`; handlers `func(ctx Context, ...)`. No rename in this change unless clarity needs `RenderContext` alias — defer rename.

### 8. Dot imports

**Choice:** Keep in `main.go` with existing nolint comments; spec captures convention.

## Risks / Trade-offs

- **[Risk] Hand-maintaining `With`/`Or` on every stub until codegen** → Mitigation: P3 codegen template includes both methods; this change sets the pattern.
- **[Risk] Full-tree refresh + demo input still hammers on each keystroke** → Mitigation: Explicitly out of scope; P2/P3 input spec later.
- **[Risk] `If` skip semantics vs host omit vs empty node** → Mitigation: Spec requires no visible node; runtime must match in P3.

## Migration Plan

Early-stage: no production users. Replace string `.With` usages in `main.go` only. Archive updates `openspec/specs/ui-dsl` on change completion; bump ROADMAP P1 status when archived.

## Open Questions

- Variant prop typing strictness (string vs const types) — defer to codegen change.
- Whether to add `Fragment` before first host — defer until layout pain appears.
