# Tasks

P5 — `forma-todo-shadcn-slice` only. Do **not** add `examples/todo` or desktop work.

## 1. Codegen payload callbacks and overrides

- [x] 1.1 Extend TS→Go event mapping so change-style props emit `func(ctx ui.Context, value string)`, checked-change emit `func(ctx ui.Context, checked bool)`, and click-style stays context-only; emit those types from jennifer — verify with new/updated unit tests on `mapTSTypeToGo` (or analyze entrypoints)
- [x] 1.2 Ensure Value / Checked / Placeholder (and other scalar controlled props present on Input/Checkbox fixtures) are emitted as ordinary Go fields — verify via codegen golden or analyze snapshot
- [x] 1.3 Implement per-source override loading + merge (documented format); cover at least one override that forces a payload callback type — verify override unit test passes and maintainer docs exist (package comment or `doc/` note referenced from user guide advanced section)
- [x] 1.4 Add Input and Checkbox TSX fixtures + goldens (Go binding + registry as needed); regenerate goldens and verify `go test ./cmd/forma/internal/codegen/...` passes

## 2. Card → Panel and checkbox starter

- [x] 2.1 Rename framework layout `Card` to `Panel` (Go stub, type string, tests, root `main.go` smoke) — verify `go test ./ui/...` and smoke package build
- [x] 2.2 Update host builtin registry: remove layout `Card`, register `Panel` with the same padding/box behavior — verify host builtin unit tests
- [x] 2.3 Add `checkbox` to CLI starters (`button`, `input`, `card`, `checkbox`) and update init/sync tests that assert the starter list — verify `go test` for starters/sync packages
- [x] 2.4 Update author guide pages that list starters or layout `Card` (cli, install, adding-components, writing-ui as touched) for Panel + checkbox — verify docs mention the four starters and Panel vs Card

## 3. Tailwind static maps and Go Class content

- [x] 3.1 Replace dynamic `gap-` / `p-` / `max-w-` / `my-` class construction in scaffold `builtins.tsx` with static maps for the supported scale; omit unsupported values — verify `builtins` host tests assert literal class strings
- [x] 3.2 Extend scaffold `tailwind.config.js` content globs to include application Go sources that may carry `Class` utilities — verify config includes those paths and a short doc/troubleshooting note explains Class must be static strings under those paths
- [x] 3.3 Run scaffold host frontend tests (vitest or project equivalent for `web/src/host`) and fix regressions — verify host test suite green

## 4. Docs alignment and slice acceptance

- [x] 4.1 Finish `doc/user/` updates for typed OnChange/OnCheckedChange, Tailwind/Class scanning, and any remaining Panel/starter wording (troubleshooting / advanced as needed) — verify guide matches shipped behavior
- [x] 4.2 Run `task test` and `task lint` for the worktree after Go/host changes — verify both pass
- [x] 4.3 Smoke-check (manual or scripted): `forma init` in a temp dir materializes button/input/card/checkbox and generated Input/Checkbox Go types include payload callbacks — verify by inspecting generated `*.gen.go` field signatures (no need to ship `examples/todo`)

## Workflow follow-up

- Human reviews planning artifacts; builder applies via `/opsx-apply`.
- Archive after implementation review; then schedule `forma-todo-example`.
- Update ROADMAP tip row status when archiving (complete) if not already done by orch.
