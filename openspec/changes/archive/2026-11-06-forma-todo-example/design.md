# Design — forma-todo-example

## Context

See `proposal.md` (Why / What). **P5** tip after owned widgets: stack is React thin host + **0build** + hand-written Go widgets from **`forma-owned-widgets` @ `f853abb` (PR #8)**. Root `main.go` already implements an in-memory todo on that surface, but the framework repo has **no** app `web/` / `forma.yaml`, so `forma dev` cannot run it in-place (`doc/user/install.md`). `forma dev` compiles with `goscript compile --package .` from the project directory that owns `forma.yaml`.

Locked: **not shadcn**; do not reopen #5/#6. No new widget catalog. Keep `fn/`.

## Goals / Non-Goals

**Goals**

- Ship a **runnable** `examples/todo` mini-app on owned widgets + 0build via `forma sync` + `forma dev`.
- Prove CRUD + minimal responsive layout without new APIs.
- Make docs/README/ROADMAP point at that run path as the P5 tip demo.

**Non-Goals (design-level)**

- Persistence, auth, backend, desktop (P4).
- New widgets, breakpoint DSL, or kit codegen.
- Publishing the example as a separate product module outside the repo.
- Perfect visual polish beyond “usable + 0build semantic defaults.”
- Requiring full `forma dev` in CI (manual / documented verification is enough; optional `go build` of the example module).

## Decisions

Copy-ready for PR body:

1. **Location** — Canonical runnable demo lives at **`examples/todo/`** (nested app), not framework root. Rationale: `forma dev` needs a project root with `forma.yaml` + `web/` + `package main` at `.`; framework root stays library + CLI.

2. **Nested module** — `examples/todo` has its own **`go.mod`** (`github.com/zhulik/forma/examples/todo`) with **`replace github.com/zhulik/forma => ../..`**. Rationale: matches how apps consume the library; `goscript --package .` works from that directory; sync/`FORMA_REPLACE` patterns stay familiar.

3. **Reuse scaffold host (no hand-roll)** — Materialize **`web/`** from **`cmd/forma/templates/scaffolds/minimal`** (owned builtins + 0build CDN). Do **not** invent a parallel host or CSS stack. Commit the tree under `examples/todo/web/` so clone → sync → dev works without re-running init into a non-empty dir.

4. **App UI** — Port the current root **`main.go` todo** into `examples/todo/main.go` (in-memory `AppState`, `fn.Map`/`Filter`, `ctx.Update`, owned widgets only). Empty-list `.Or(...)` kept. Mount a tiny root `Component` whose `Render` rebuilds `TodoApp()` each refresh so Values/list stay live (static one-shot `TodoApp()` freezes the tree). No shared extracted package unless duplication becomes painful later.

5. **Root `main.go`** — **Slim to Hello-style** owned-widget smoke (scaffold-like). Canonical **interactive** demo is `examples/todo`. Update `openspec/config.yaml` apply guidance / context that still say “preserve main.go as canonical DSL example” to point at `examples/todo`.

6. **Responsive = layout props + escape hatch** — Use `Container` max-width + margins, card padding, row/column gaps; add small horizontal padding via `Class` utilities if needed. **No** new media-query widget. “Minimal responsive” = phone-width readable, no page-chrome horizontal scroll for default content.

7. **Docs in the same slice** — Update `doc/user/` (install, project-layout, README index as needed), root `README.md`, and add **`examples/todo/README.md`** with sync + `forma dev` steps. Remove planned-only wording for dedicated `examples/todo`.

8. **Ignore generated/dev artifacts** — Ensure `examples/todo` ignores `web/node_modules/`, `web/.goscript/`, and other sync/dev outputs consistent with scaffold `.gitignore`.

9. **Verification** — Builder manually (or scripted once) runs sync + `forma dev` and confirms CRUD. Optionally add `go build` / `go test` for the example module to root `Taskfile`—**not** mandatory Vite e2e in CI this change.

10. **Reuse vs hand-roll** — **Reuse** owned `ui/` widgets, `fn/`, scaffold host/0build, existing `forma sync`/`forma dev`. **Hand-write** only the example app Go + thin README. **Do not** replace `fn/` or add a third-party UI kit.

11. **ROADMAP** — On ship/archive, mark `forma-todo-example` done with branch/PR; keep #5/#6 superseded. Impact phase remains **P5**.

### Alternatives considered

| Topic | Rejected | Why |
|-------|----------|-----|
| Put `web/` at framework root | Pollutes library module; confuses init vs framework | Nested example is clearer |
| Same-module `examples/todo` without nested go.mod | `goscript --package .` / module path awkward for `@goscript/<module>/main` | Nested module matches scaffold |
| Keep dual full todos (root + examples) | Drift | Single interactive SoT |
| Extract shared `internal/tododemo` | Extra indirection for tip demo | Port once; slim root |
| Add responsive breakpoint widgets | Catalog expansion | Out of scope for tip |
| Wire Playwright/CI browser | Heavy for tip | Docs + manual verify |

## Risks / Trade-offs

- **[Risk] Scaffold host drifts from committed `examples/todo/web`** → Document “refresh from minimal scaffold when host changes”; optional follow-up task to sync. Prefer copy-from-scaffold at implement time over long-lived forks of host logic.
- **[Risk] Nested module CI surprise (lint/test paths)** → Scope lint to framework + `cmd/forma` as today; add explicit example `go build` only if Taskfile grows; don’t break root `go test ./...`.
- **[Risk] GoScript module path / replace edge cases** → Follow scaffold `goscript-entry` pattern with example module path; verify during apply.
- **[Trade-off] Slimming root `main.go`** → Readers lose one-file todo at repo root; docs must link `examples/todo` prominently (accepted).

## Migration Plan

1. Land planning artifacts (this change).
2. Builder implements `examples/todo`, slims root `main.go`, updates docs/config/ROADMAP pointers.
3. Manual `forma sync` + `forma dev` verification of CRUD.
4. `task test` / `task lint` on framework + CLI.
5. No runtime migration—docs mental-model break only; early-stage OK.

## Open Questions

None blocking. Deferrable: whether to add a root Taskfile target that `cd`s into the example for `go build`—Builder may add if cheap.
