# Design

## Context

See proposal.md — Why. Worktree tip is stacked on `forma-todo-shadcn-slice` @ `4809718` (PR #5). Assumed available:

- Framework layout boxed surface is **Panel**; generated shadcn **Card** owns `"Card"` after init.
- Init starters: `button`, `input`, `card`, `checkbox`.
- Typed `OnChange(ctx, string)` / `OnCheckedChange(ctx, bool)`; host sanitize already matches.
- Scaffold Tailwind `content` includes `../*.go` and `../internal/**/*.go`; builtins use static spacing maps.

Framework root today: `main.go` is a full DSL todo smoke with **no** `web/` / `forma.yaml`. Docs still treat root `main.go` as the browser verification narrative. No `examples/` tree yet. `FORMA_REPLACE` + `go mod edit -replace` is the established local-library path (init/sync/CI).

Target UI (from slice — implement, do not re-invent):

```text
Container (MaxWidth md, MarginX auto; Class padding as needed)
  Card  -- generated shadcn compound (Header / Title / Content slots as emitted)
    Header/Title: "My Tasks" (+ muted count optional)
    Content:
      composer: Class "flex w-full flex-col gap-2 sm:flex-row sm:items-center"
        Input (Placeholder; Value; OnChange(ctx, string); flex-1 via Class)
        Button "Add" (OnClick)
      list Column:
        each todo Row (Key; Class; items)
          Checkbox (Checked; OnCheckedChange(ctx, bool)) + Text
          Button "Delete" (OnClick)
      empty: Or(Text muted…)
```

## Goals / Non-Goals

**Goals**

- Ship a committed, runnable `examples/todo` that matches the target UI and proves CRUD + responsive composer via `forma dev`.
- Reuse init/sync/codegen/host rather than hand-rolling a parallel web stack.
- Move author-facing “run the todo” docs to that example; keep root smoke thin and compiling.

**Non-Goals**

- New widgets, breakpoint DSL, desktop, persistence, multi-list, filters, icons catalog.
- Re-opening slice codegen/host decisions.
- Full e2e browser automation (Playwright/Cypress).
- Deleting all framework `ui/*.gen.go` stubs (keep whatever thin smoke still needs; no wholesale stub purge).

## Decisions

Builder may copy this list into the PR body.

1. **Nested module at `examples/todo`**  
   Module path `github.com/zhulik/forma/examples/todo`. Own `go.mod`, `forma.yaml`, `web/`, `internal/ui/`, `main.go`. Committed `replace github.com/zhulik/forma => ../..`.  
   **Why:** ROADMAP says own module; mirrors a real `forma init` app; keeps framework root free of app `web/`.  
   **Rejected:** same-module `examples/todo` package under root `go.mod` (forces root `forma.yaml`/`web/` or a non-standard layout); `go.work` as primary wiring (extra mental model; replace already matches `FORMA_REPLACE`).

2. **Materialize via `forma init` + replace (reuse), then commit**  
   One-time: `FORMA_REPLACE=<repo-root> forma init github.com/zhulik/forma/examples/todo -d examples/todo` (empty dir), then replace scaffold hello with the todo UI; run `forma sync` / `generate` as needed; commit TSX, `*.gen.go`, `registry.gen.ts`, lockfiles.  
   **Why:** Prefer reuse over hand-roll (`prefer-reuse-over-handroll`); guarantees starters/host/Tailwind match production scaffold.  
   **Rejected:** hand-copying scaffold files piecemeal; regenerating the whole example from init on every CI run (non-deterministic churn; example is a product artifact).

3. **UI implements the slice target literally**  
   Generated **Card** compound (slot fields as emitted — typically Header / Title / Content); composer **Class** responsive stack; list via **`fn.Map`** with stable **Key**; empty via **`.Or`**; in-memory state + **`ctx.Update`**. Import split: dot-import framework `ui` + `fn`; normal import of `…/examples/todo/internal/ui` for Button/Input/Card/Checkbox (or a short local alias).  
   **Why:** Slice invented the UI once; this change only proves it. Exact Card field names come from generated output after init — do not invent parallel compound APIs.  
   **Rejected:** sticking with framework **Panel** as the shell (would not exercise generated Card); new catalog widgets.

4. **Root `main.go` becomes thin smoke**  
   Reduce to a minimal compilable DSL check (e.g. Container + Panel + Text). Product narrative and DESIGN/README “north-star runnable” pointers move to `examples/todo`.  
   **Why:** ROADMAP explicitly allows thin smoke once the example owns the narrative; avoids two competing full todos.  
   **Rejected:** deleting root `main.go`; keeping a second full Card-based todo at root without `web/` (misleading).

5. **Docs migrate in the same change**  
   Update `doc/user/` (README, install, cli browser check, writing-ui reference, project-layout “this repo”), root README, DESIGN north-star sentence, ROADMAP tip status on land. Add a short `examples/todo/README.md` with sync + `forma dev` (+ note that `FORMA_REPLACE`/`replace` is already in `go.mod` for the monorepo checkout).  
   **Why:** ui-docs obligation; no proposal opt-out.

6. **CI/Task: `check-example-todo` (reuse scaffold check spirit)**  
   Task builds `examples/todo` with its replace; installs/runs `web/` `npm test` + build (or the scaffold’s equivalent scripts). Wire into CI beside `check-init-scaffold`. Do **not** require headed browser CRUD automation.  
   **Why:** Keeps the example from rotting; matches existing Node+CGO CI capacity.  
   **Rejected:** Playwright viewport suite in this change; relying only on manual verification.

7. **Responsive proof = Class + manual acceptance**  
   Composer Class string from the target UI; acceptance checklist in example README / tasks (narrow stacks, `sm:+` side-by-side). No new layout API.  
   **Why:** Slice already chose Class over breakpoint DSL.

8. **Reuse vs hand-roll**  
   - **Reuse:** `forma init`/sync/generate, starters, generated Card/Input/Button/Checkbox, host, Tailwind content globs, `fn.Map`/`Filter`, `ctx.Update`.  
   - **Hand-roll only:** example app state + screen composition; thin root smoke; docs; Task/CI glue for the committed module.

## Risks / Trade-offs

- **[Risk] Nested module ignored by root `go test ./...`** → Mitigation: dedicated Task/CI check; document that the example is its own module.
- **[Risk] Card slot field names differ slightly from docs (e.g. Description present)** → Mitigation: builder binds to whatever codegen emits; use Header/Title/Content as available; keep Title copy “My Tasks”.
- **[Risk] Committed `web/node_modules` or `.goscript` artifacts** → Mitigation: respect scaffold `.gitignore`; never commit `node_modules` / compile outputs.
- **[Risk] `replace` path breaks when example is copied outside the repo** → Mitigation: example README states it is an in-repo demo; external apps use published forma or `FORMA_REPLACE`.
- **[Risk] Doc drift leaving root main.go as “the” demo** → Mitigation: explicit doc tasks + ui-docs delta scenarios.

## Migration Plan

1. Init + commit `examples/todo` skeleton (replace in go.mod).
2. Implement todo UI; verify `forma dev` locally (CRUD + responsive Class).
3. Slim root `main.go`; update docs/README/DESIGN.
4. Add Task + CI smoke for the example module.
5. Update ROADMAP tip row status when merging.

Rollback: revert the change branch; no production deploy.

## Open Questions

None material. Minor: whether Delete uses `Variant: "ghost"` vs `"outline"` and icon vs text “Delete” — prefer visible **text “Delete”** (matches target sketch; clearer for demos) unless generated Button sizing fights layout, then ghost + text is fine.
