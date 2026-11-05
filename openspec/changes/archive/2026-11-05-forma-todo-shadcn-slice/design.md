## Context

See proposal.md — Why. Verified against worktree `forma-todo-shadcn-slice` @ `1422a3f`:

- Runtime (`ui/bind.go`) already injects `ui.Context` and forwards host payloads; host `sanitize.ts` already distinguishes click / string change / checked-change.
- Codegen `mapTSTypeToGo` collapses every `on*` prop to `func(ctx ui.Context)`; jennifer emit only handles that one func shape. Spec’d overrides are unimplemented.
- Starters: `button`, `input`, `card` only (`starters.go`). Framework `ui/checkbox.gen.go` is a hand stub; Checkbox is not materialized on init.
- Host `builtins.tsx` registers layout `Card` and builds `gap-${n}`, `p-${n}`, `max-w-${token}`, `my-${n}` dynamically. Scaffold `tailwind.config.js` content is only `./src/**/*.{js,ts,jsx,tsx}` — Go `Class:` strings are invisible to purge.
- Root `main.go` is a DSL-only todo smoke (no `web/`); shipping `examples/todo` is **`forma-todo-example`**, not this change.

### Target UI (invented once — gap baseline)

Minimal responsive default-shadcn todo the follow-on example must be able to express:

```text
Container (MaxWidth md, MarginX auto; Class for horizontal padding if needed)
  Card                    -- generated shadcn compound (Header / Title / Content)
    Header/Title: "My Tasks" (+ muted count Text optional)
    Content:
      composer: Class "flex w-full flex-col gap-2 sm:flex-row sm:items-center"
        Input (flex-1 via Class; Placeholder; Value; OnChange(ctx, string))
        Button "Add" (OnClick)
      list Column:
        each todo Row (Key; Class border rounded + padding; items-center justify-between)
          Row: Checkbox (Checked; OnCheckedChange(ctx, bool)) + Text (strikethrough when done)
          Button "Delete" (outline/ghost; OnClick)
      empty: Or(Text muted "No tasks yet…")
```

**Not** in the target: Select, Dialog, Dropdown, Toast, Tabs, icons catalog, filters beyond done checkbox, persistence, multi-list.

**Responsive contract:** narrow stacks composer vertically; `sm:` and up places Input and Add side-by-side. No new breakpoint DSL — static Tailwind utilities via `Class`.

## Goals / Non-Goals

**Goals**

- Make the target UI expressible after `forma init` + generate with correct Go types and host CSS.
- Resolve Card name collision so shadcn Card is usable as the todo shell.
- Keep catalog minimal (starters + codegen/host fixes only).

**Non-Goals**

- Implement `examples/todo` or rewrite the product narrative off root `main.go` (follow-on).
- Full shadcn catalog, third-party recipes, desktop, new layout breakpoint props API.
- Hybrid policy for Checkbox (boolean; existing sanitizer + controlled props suffice).

## Decisions

Builder may copy this list into the PR body.

1. **Typed payload callbacks in analyze + emit (reuse bind/sanitize)**  
   Extend `mapTSTypeToGo` / heuristics so `onChange` / `onValueChange` → `func(ctx ui.Context, value string)`, `onCheckedChange` → `func(ctx ui.Context, checked bool)`, click-style stays context-only. Extend jennifer `goTypeExpr` for those signatures. Prefer improving heuristics for common shadcn/Radix props over per-widget hand stubs. Runtime bind and host sanitize already match this shape — **reuse**, do not hand-roll a second callback path.

2. **Implement override files as the safety net**  
   Ship a small committed override format (YAML or JSON next to codegen testdata / embedded maintainer overrides for `input` and `checkbox` if TSX parsing still misses props). Document format for maintainers. Overrides merge on top of analysis (existing spec strengthened from MAY to SHALL when classification is insufficient). Prefer override over forever-growing special cases for edge TSX.

3. **Starter set = `button`, `input`, `card`, `checkbox`**  
   Add `checkbox` only. No Select/Dialog/etc. Card stays a starter so the todo shell is available without `forma add`.

4. **BREAKING: layout `Card` → `Panel`**  
   Rename framework layout type string and Go type to `Panel` (stub `ui/panel.gen.go`, builtin registry, tests, smoke `main.go`). Generated shadcn `Card` owns `"Card"`. Alternatives rejected: drop layout box entirely (loses Padding prop for non-shadcn surfaces); keep dual `Card` (registry override hides layout — current bug for authors expecting Padding Card after init). Early-stage: no dual-write / aliases.

5. **Static spacing maps in builtins; Tailwind content includes Go**  
   Replace dynamic template strings with explicit maps for supported Tailwind spacing scale (at least 0–12 used by demos) and max-width tokens (`sm`/`md`/`lg`/`xl`/`2xl`/`full`/…). Unsupported values omit the class. Extend scaffold `tailwind.config.js` `content` to include project Go trees that hold `Class` strings (e.g. app root `*.go` and `internal/**/*.go` relative to `web/`). Prefer content globs over a giant safelist (**reuse** Tailwind’s scanner). No new responsive Row API in this slice.

6. **Responsive via `Class`, not a breakpoint DSL**  
   Target composer uses static strings like `flex flex-col gap-2 sm:flex-row sm:items-center`. Decision (5) makes that viable. Defer first-class responsive props to later P5 if needed.

7. **Framework smoke `main.go`**  
   Update to `Panel` + keep Checkbox/Input signatures aligned with stubs so DSL smoke compiles. Do **not** add a `web/` tree or `examples/todo` here.

8. **Docs in the same change**  
   Update `doc/user/` (cli, install, adding-components, writing-ui, troubleshooting, project-layout as touched) per ui-docs obligation. No proposal opt-out.

9. **Reuse vs hand-roll**  
   - Reuse: existing codegen pipeline (tree-sitter analyze + jennifer), host sanitize/HybridInput, registry merge, `fn` helpers in analyze.  
   - Hand-roll only: static class maps; small override schema + loader; Panel rename; checkbox starter entry; Input/Checkbox goldens.

## Risks / Trade-offs

- **[Risk] Heuristics still miss exotic event types** → Mitigation: overrides + goldens for Input/Checkbox; click remains default for unknown `on*`.
- **[Risk] Go content globs pull noise / miss nested modules** → Mitigation: document expected paths; example change can extend `forma.yaml`/scaffold if monorepo layout differs; prefer broad `../**/*.go` under web-relative paths carefully to avoid `node_modules` (keep globs outside `web/node_modules`).
- **[Risk] Panel rename breaks WIP apps** → Mitigation: early-stage policy; docs + smoke update; no alias.
- **[Risk] Static maps omit a spacing the example wants** → Mitigation: include 0–12 and common max-w tokens; Class escape hatch covers the rest once Go is in content.
- **[Risk] Scope creep into examples/todo** → Mitigation: tasks explicitly exclude it; acceptance is “init project can generate and type-check the target surface,” not ship the example.

## Migration Plan

1. Land codegen + goldens + overrides.
2. Add checkbox starter; rename Card→Panel (framework + builtins + smoke + docs).
3. Static maps + Tailwind content; host tests.
4. Doc/user updates.
5. Follow-on `forma-todo-example` consumes the slice.

Rollback: revert the change branch; no production deploy.

## Open Questions

None material — deferred only: exact override file path/extension (YAML vs JSON) left to builder within “committed, documented, deterministic” constraints; prefer YAML beside codegen package if no existing format appears during apply.
