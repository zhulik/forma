# Tasks

P1 — UI DSL v1. Verify Go changes with `go build ./...`, then `task test` and `task lint` when wired.

## 1. Core composition (`ui/component.go`, `ui/node.go`)

- [x] 1.1 Change `With` and `Or` to `...Component`, remove string coercion and `toComponents` panic path, and verify `go build ./...` succeeds
- [x] 1.2 Add `Key string` on `BaseComponent` and propagate to `Node` in `Render`, and verify a manual render of a keyed row includes the key on the node
- [x] 1.3 Implement skip/omit in `renderChildren` for conditional no-node results, and verify parent child list excludes them

## 2. Conditional helper (`ui/if.go` or equivalent)

- [x] 2.1 Add `If(cond bool, thunks ...func() Component) Component` (arity 1–2, panic on extras; lazy invoke), and verify `go build ./...`

## 3. Callback context (`ui/context.go`, stub)

- [x] 3.1 Add `Update(fn func())` on the stub `Context` implementation with post-success `RefreshUI`, and verify compile via `go build ./...`
- [x] 3.2 Document in code comments (brief) that `Context` is for handlers only; render uses stdlib context — matches design

## 4. Widget stubs (`ui/*.gen.go`)

- [x] 4.1 Update each wrapper’s `With` to `...Component` and add typed `Or` mirroring `With` on all hand-written gen files, and verify `go build ./...`
- [x] 4.2 Ensure `Input`/`Button`/`Checkbox` event fields still use handler context type, and verify compile

## 5. Functional helpers (`fn/fn.go`)

- [x] 5.1 Change list mapping in demo helpers to `[]Component` return type where applicable, and verify `main.go` splat compiles
- [x] 5.2 Add `fn.FilterMap` and use it in `renderChildren` to omit nil nodes, and verify `task lint` passes

## 6. Reference demo (`main.go`)

- [x] 6.1 Replace string `.With` with `Text{Text: "…"}` / field text; use `ctx.Update` for mutations; set `Key` on todo rows; use `If` where it clarifies conditional UI, and verify `go build -o /dev/null .`
- [x] 6.2 Keep dot imports and `.Or` empty-list fallback; confirm column gap preserved in rendered node type after apply (inspect `Render` output or debug print once)

## 7. Design doc alignment (`doc/DESIGN.md`)

- [x] 7.1 Update authoring sections: no string `.With`, `Text{Text: "…"}`, `ui.If`, `Update`, context split, keys, `fn` generic-only — and verify no contradiction with `openspec/changes/forma-ui-dsl/specs/ui-dsl/spec.md`

## 8. Verification

- [x] 8.1 Run `task test` and `task lint` from repo root and verify both exit 0 after all code changes

## Workflow follow-up

- Archive change after review; sync capability to `openspec/specs/ui-dsl/` and update `doc/ROADMAP.md` P1 status when done.
