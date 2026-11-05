# Proposal

## Why

P5’s tip is an in-browser todo that proves the slice: default shadcn, responsive Class layout, `forma dev`. The framework-root `main.go` is DSL smoke without a `web/` tree, so authors still lack a runnable example project. Ship `examples/todo` now that the slice (Panel, checkbox starter, typed callbacks, Tailwind Go content) is assumed available.

## What Changes

- Add **`examples/todo`**: nested Go module + full `forma.yaml` / `web/` app tree, implementing the target responsive shadcn todo UI (CRUD, narrow/wide composer).
- Prefer **reuse of `forma init` + sync** to materialize the tree; replace scaffold hello with the todo; commit generated Go/TS artifacts like a real app.
- Slim framework-root **`main.go`** to a thin compile smoke; move the product narrative to `examples/todo`.
- Update **author docs**, README, DESIGN/ROADMAP pointers, and OpenSpec scenarios so the verification path is `examples/todo` + `forma dev`.
- Add a lightweight **CI/task gate** that builds (and host-tests) the committed example module.
- **Out of scope:** desktop (`forma-desktop`); new catalog widgets or breakpoint DSL; persistence/multi-list/filters; Playwright/visual e2e; re-opening archived slice design except as assumed base; removing framework `ui/*.gen.go` stubs wholesale.

## Capabilities

### New Capabilities

- (none)

### Modified Capabilities

- `ui-docs`: reference / verification path becomes `examples/todo`; document how to run it; entry points may cite the example beside (or instead of) root smoke `main.go`.
- `ui-dsl`: scenarios that treat the “canonical / reference todo” as the authoring exemplar point at `examples/todo` (requirements stay; location of the exemplar moves).
- `dev-tooling`: automation verifies the committed `examples/todo` module builds (and scaffold-equivalent web host checks as practical).

## Impact

- **ROADMAP phase: P5** — tip row `forma-todo-example` (after `forma-todo-shadcn-slice`).
- **Code:** new `examples/todo/` module (`go.mod` with replace to repo root, `forma.yaml`, `main.go`, `internal/ui/*.gen.go`, `web/`); slim root `main.go`; Taskfile (+ CI workflow if present) for example check.
- **Docs:** `doc/user/` (README, install, cli, writing-ui, project-layout, troubleshooting as touched), root README, `doc/DESIGN.md` north-star pointer, ROADMAP status when landing.
- **Assumed base (not re-designed here):** Panel (not layout Card), checkbox starter, typed `OnChange`/`OnCheckedChange`, static Tailwind maps + Go Class content scanning from `forma-todo-shadcn-slice`.
- **Breaking (narrative only):** docs that treated root `main.go` as the runnable browser demo redirect to `examples/todo`; root smoke remains compilable but is not the product demo.
