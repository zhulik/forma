# Proposal

## Why

forma’s reference demo and stub `ui` package encode a Go UI authoring model, but behavior is implicit and has known gaps (typed composition after `.Or`, string children, missing list keys). ROADMAP phase **P1** requires a locked **UI DSL v1** before user docs and runtime work so authors, agents, and the future React host share one contract.

## What Changes

- **BREAKING:** `With` and `Or` accept only `Component` values — no string children via `.With("…")`.
- **BREAKING:** Text leaf copy uses **`Text{Text: "…"}`** (exported struct; no bare string `.With`); other widgets may use dedicated text fields (e.g. `Heading{Text: "…"}`).
- Add typed **`Or`** on every component wrapper (same return type as `With`), mirroring codegen shape for future shadcn types.
- Add **`Key`** on `BaseComponent` (promoted on widgets); demo list rows set stable keys.
- Add **`If(cond, …func() Component)`** in the UI package (lazy branches; dot-import in apps); false with one branch renders nothing; optional second branch for else. Generic **`fn`** keeps only `Map` / `Filter`.
- Add **`Context.Update(fn func())`** — run mutation, then **`RefreshUI()` once** on success; document **`ui.Context` for callbacks only** and **`context.Context` at render/runtime**.
- Update **`main.go`** as the canonical example (dot imports, `Update`, keys, `Text`, typed `Or`).
- Update hand-written **`ui/*.gen.go`** stubs and **`ui/component.go`** / **`fn`** to match; no React host, GopherJS bridge, or input-sync policy in this change.

## Capabilities

### New Capabilities

- `ui-dsl`: Author-facing composition, control flow, keys, callback context, and render boundaries for forma apps.

### Modified Capabilities

- (none)

## Impact

- **ROADMAP phase:** P1 (UI DSL v1).
- **Code:** `main.go`, `ui/` (`component.go`, `if.go`, `context.go`, placeholders), `fn/fn.go` (`Map`, `Filter`, **`FilterMap`**); `doc/DESIGN.md` aligned with P1 rules.
- **Out of scope:** Controlled/deferred input policy (demo may keep per-keystroke refresh), `Fragment`, bridge/`Node`/`js` tags, codegen pipeline, React host, real `UIMount` behavior, P2 user docs.
