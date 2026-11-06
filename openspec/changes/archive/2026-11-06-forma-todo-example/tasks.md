# Tasks

## 1. Scaffold examples/todo mini-app

- [x] 1.1 Create `examples/todo/` with nested `go.mod` (`github.com/zhulik/forma/examples/todo`), `replace` to `../..`, and `forma.yaml` (module + `ui.mount` + `web.path`); verify `go mod tidy` succeeds from that directory
- [x] 1.2 Materialize `examples/todo/web/` from `cmd/forma/templates/scaffolds/minimal` (owned host builtins + 0build CDN), with module/mount substituted for the example; verify `web/index.html` links 0build CSS and mount id matches `forma.yaml`
- [x] 1.3 Add ignore rules for example sync/dev outputs (`node_modules`, `.goscript`, etc.) consistent with the scaffold; verify those paths are not tracked as intentional sources

## 2. Todo UI on owned widgets

- [x] 2.1 Port the in-memory todo app (add / list / toggle / delete, empty-state `.Or`, `fn.Map`/`Filter`, `ctx.Update`) into `examples/todo/main.go` using only owned widgets; verify `go build .` succeeds in `examples/todo` with the replace
- [x] 2.2 Apply minimal responsive layout via existing Container/Card/Row/Column props (and Class escape hatch if needed); verify no new widget types were added under framework `ui/`
- [x] 2.3 Slim root `main.go` to a short Hello-style owned-widget smoke; verify `go build .` at framework root still succeeds

## 3. Author docs and pointers (same slice)

- [x] 3.1 Add `examples/todo/README.md` with sync + `forma dev` run steps; verify the documented commands match `forma.yaml` / module layout
- [x] 3.2 Update `doc/user/` (at least install + project-layout; README index if needed) so `examples/todo` is the runnable demo and is not labeled planned; verify wording matches `ui-docs` delta scenarios
- [x] 3.3 Update root `README.md` and `openspec/config.yaml` context/apply guidance to point the canonical interactive demo at `examples/todo`; verify no remaining “root main.go is the only north-star” claim for the browser demo

## 4. Integration verification (P5)

- [x] 4.1 From `examples/todo`, run `forma sync` then `forma dev` and confirm add, toggle-done, and delete in the browser; verify empty-state appears with zero tasks
- [x] 4.2 Run `task test` and `task lint` from the framework root; verify both pass
- [x] 4.3 Update `doc/ROADMAP.md` tip-row status for `forma-todo-example` when shipping (branch/PR as known); verify Impact still cites phase **P5** and stack on owned-widgets/#8
## Workflow follow-up

- Archive the change after review (`/opsx-archive` / finalize flow) so delta specs merge into main specs.
- Do not reopen superseded shadcn tip PRs #5/#6.
