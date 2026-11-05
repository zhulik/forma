# Tasks

P5 — `forma-todo-example` only. Assume slice landings (Panel, checkbox starter, typed callbacks, Tailwind Go content). Do **not** add desktop work or new catalog widgets.

## 1. Scaffold examples/todo module

- [x] 1.1 Create empty `examples/todo` and run `FORMA_REPLACE=<repo-root> forma init github.com/zhulik/forma/examples/todo -d examples/todo`; verify `forma.yaml`, `web/`, `internal/ui/` with button/input/card/checkbox `*.gen.go`, and `go.mod` with `replace github.com/zhulik/forma => ../..`
- [x] 1.2 Ensure scaffold ignores stay correct (`web/node_modules`, GoScript outputs not committed); verify `git status` would not stage those paths
- [x] 1.3 Add `examples/todo/README.md` with sync + `forma dev` run steps (and that replace points at the monorepo root); verify the README commands match the tree layout

## 2. Implement responsive todo UI

- [x] 2.1 Replace scaffold hello `main.go` with the target UI (Container → generated Card slots → responsive Class composer → Mapped rows with Checkbox/Text/Delete → `.Or` empty); verify `go build` in `examples/todo` succeeds
- [x] 2.2 Wire in-memory CRUD (add / toggle done / delete) via `ctx.Update` and `fn.Map`/`Filter` with stable Keys; verify native `go test`/`go build` still pass and handlers compile against generated Input/Checkbox signatures
- [x] 2.3 Manually run `forma sync` (if needed) then `forma dev` from `examples/todo`; verify browser CRUD works and composer stacks narrow / sits Input+Add side-by-side at `sm:+` (note checklist result in PR)

## 3. Thin framework smoke + docs

- [x] 3.1 Slim root `main.go` to thin DSL smoke (layout-only or equivalent); verify root `go build` / `task test` still pass
- [x] 3.2 Update `doc/user/` (README status, install, cli browser check, writing-ui reference, project-layout “this repo”) so the runnable demo is `examples/todo`; verify links resolve
- [x] 3.3 Update root README and `doc/DESIGN.md` north-star pointer to `examples/todo`; verify they no longer claim root `main.go` is the browser product demo
- [x] 3.4 Update `doc/ROADMAP.md` tip row / status for `forma-todo-example` as appropriate for the PR landing; verify P5 phase reference remains accurate

## 4. Automation

- [x] 4.1 Add Task `check-example-todo` that builds `examples/todo` and runs its `web/` npm test + build (install deps if required); verify `task check-example-todo` passes locally
- [x] 4.2 Wire that task into CI alongside existing scaffold checks; verify the workflow file references it
- [x] 4.3 Run `task test` and `task lint` at repo root (and CLI module as usual); verify both pass after the change

## Workflow follow-up

- Archive the change after review/merge requirements are satisfied (`/opsx-archive` or project process).
- Confirm ROADMAP tip row shows `examples/todo` shipped once archived.
