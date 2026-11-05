## Why

P5 needs a **minimal responsive default-shadcn todo** as the next example, but forma still cannot author that UI from `forma init` alone: codegen collapses value/checked callbacks to click-shaped signatures, Checkbox is missing from starters, layout `Card` clashes with generated shadcn `Card`, and Tailwind purges the dynamic spacing classes builtins emit (and would purge Go `Class` utilities). This change gap-fills only those blockers so a later `forma-todo-example` can ship `examples/todo`.

## What Changes

- **Codegen:** emit typed payload callbacks for change/checked-change style props (`OnChange(ctx, string)`, `OnCheckedChange(ctx, bool)`) plus ordinary value props (`Value`, `Checked`, `Placeholder`, …); implement the already-spec’d per-source override hook where heuristics are insufficient; add Input/Checkbox (and Card compound) goldens.
- **Starters:** add **`checkbox`** to the CLI init starter set alongside button, input, and card.
- **BREAKING — layout Card → Panel:** rename the framework layout/leaf boxed surface from type `Card` to `Panel` so generated shadcn `Card` owns the `Card` registry type after init; update framework stubs, host builtins, smoke `main.go`, and docs.
- **Tailwind / layout spacing:** replace dynamic `gap-${n}` / `p-${n}` / `max-w-*` / `my-*` construction in host builtins with **static class maps**; extend scaffold Tailwind `content` globs so utilities written in Go `Class` fields (responsive stacks) are not purged.
- **Author docs:** document starter set, typed callbacks, Panel, and Class/Tailwind scanning expectations.
- **Out of scope:** shipping `examples/todo` (`forma-todo-example`); desktop (`forma-desktop`); full shadcn catalog; new breakpoint DSL; Select/Dialog/Dropdown/Toast and other non-todo widgets.

## Capabilities

### New Capabilities

- (none)

### Modified Capabilities

- `ui-codegen`: typed value/checked callback signatures in generated Go; overrides consulted and documented as a real input.
- `forma-cli`: init starter set includes checkbox.
- `ui-host`: builtin layout spacing/max-width classes are statically enumerated; scaffold Tailwind content covers Go Class sources; builtin boxed layout type is Panel.
- `ui-dsl`: framework layout boxed surface is Panel; `Card` type string reserved for generated shadcn Card after starters land.
- `ui-docs`: author guide reflects starters, Panel, typed callbacks, and Tailwind/Class scanning.

## Impact

- **ROADMAP phase: P5** — tip row `forma-todo-shadcn-slice` (gap-fill before `forma-todo-example`).
- **Code:** `cmd/forma/internal/codegen/` (analyze, emit, overrides, goldens); `cmd/forma/internal/cli/starters/`; scaffold host `builtins.tsx`, `tailwind.config.js`; framework `ui/` layout stubs; root smoke `main.go`; host/builtin tests.
- **Specs / docs:** deltas under the capabilities above; `doc/user/` CLI, adding-components, writing-ui, troubleshooting as needed.
- **Follow-on:** `forma-todo-example` implements `examples/todo` on this slice; does not belong in this change.
- **Breaking:** any author/code using framework layout `Card` must switch to `Panel` (early-stage, no compat shims).
