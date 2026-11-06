# Tasks

ROADMAP **P5** — `forma-owned-widgets` (architecture migration). Do **not** ship `examples/todo` here.

## 1. Scaffold styling → 0build

- [x] 1.1 Replace minimal scaffold Tailwind/shadcn kit assets (`tailwind.config.js`, PostCSS Tailwind pipeline, `components.json`, `@tailwind` CSS, kit-only npm deps such as `tailwind-merge`/`cva`) with pinned **0build** CDN CSS + required runtime scripts in `index.html`/entry per upstream install docs; verify `npm install` in the template web tree succeeds without those kit packages
- [x] 1.2 Replace `cn`/`twMerge` helpers with a trivial class-join utility and update host imports; verify host unit tests that import the helper still compile/run
- [x] 1.3 Document the pinned CDN URLs/versions in a short scaffold comment or `doc/user/` styling note; verify the note names semantic `.z-*` first and utilities/`Class` as escape hatch

## 2. Hand-written Go widgets (`ui/`)

- [x] 2.1 Rename framework `ui/*.gen.go` leaf/layout types to ordinary `*.go` (one type per file) and keep APIs sufficient for `main.go` (Button, Input, Checkbox, Card, Text, Heading, Row, Column, Container); verify `go test ./ui/...` and that root `main.go` still typechecks against `ui`
- [x] 2.2 Adjust any `Class` strings in root `main.go` that assume Tailwind kit tokens to 0build-friendly escape hatches (or rely on defaults); verify `go build` / package tests for the module root still pass
- [x] 2.3 Confirm Fragment/slot helpers remain unchanged and unused compound-codegen assumptions are gone from library comments; verify existing `ui` slot tests pass

## 3. Host builtins for todo surface

- [x] 3.1 Add thin React host components for Button, Input, and Checkbox (and move/keep Card/layout/Text/Heading) under scaffold `web/src/host/` applying mapped `.z-*` classes; verify `builtinRegistry` resolves each todo-sized type without a generated registry
- [x] 3.2 Reimplement layout gap/padding/flex mapping with 0build utilities/semantic classes (no Tailwind class names); verify builtins host tests cover layout + interactive smoke renders
- [x] 3.3 Wire registry merge to **builtins + `registry.custom.ts` only** (drop `registry.gen` from `main.tsx` template); verify template tests / host merge tests expect no generated kit registry
- [x] 3.4 Keep hybrid text-input policy and event sanitization for Input/Checkbox/Button callbacks; verify existing sanitize/hybrid-input tests still pass (update fixtures if prop shapes change)

## 4. Retire shadcn / codegen product path (BREAKING)

- [x] 4.1 Remove init starter shadcn materialization (`starters`, `RunWithComponents` on init) and shadcn steps from sync; verify init/sync unit tests assert no shadcn invocation and sync still runs npm + `go mod tidy`
- [x] 4.2 Remove author-facing `forma add` and `forma generate` commands (and scaffold `generator.go` / check-generate task); verify CLI help/tests no longer expose them as product commands
- [x] 4.3 Delete `cmd/forma/internal/codegen`, shadcn wrapper wiring, and default `forma.yaml` `codegen:` fields from templates; verify `go test ./cmd/forma/...` passes and formaconfig no longer requires codegen dir for sync
- [x] 4.4 Update CI init smoke (`dev-tooling`) to expect owned-widgets init without shadcn/codegen/`generator.go`; verify the workflow/task definition matches the new smoke scenario

## 5. Docs and agent context

- [x] 5.1 Rewrite `doc/user/` pages that prescribe shadcn add/generate/codegen (`adding-components.md`, `cli.md`, `forma-yaml.md`, `project-layout.md`, `advanced.md`, `README.md` status table, etc.) for owned widgets + 0build; verify guide map links remain coherent and status labels match shipped CLI
- [x] 5.2 Update `doc/DESIGN.md` tip sections that still describe shadcn codegen as current SoT (architecture diagram codegen blurb, “How users work”, non-goals) to owned widgets + 0build; verify DESIGN and ROADMAP P5 tip do not contradict
- [x] 5.3 Update `openspec/config.yaml` context/rules/operations that still inject shadcn/Tailwind/codegen-as-SoT into agent planning; verify context no longer mandates kit codegen for new changes

## 6. Integration verification

- [x] 6.1 Run `task test` and `task lint` on the worktree after the above; verify both pass
- [x] 6.2 Manually or via automation: `forma init` in a temp dir, `go build` the app, and confirm scaffold host tests pass; verify no shadcn `components/ui` or `internal/ui/*.gen.go` are required for Hello/todo-shaped composition
- [x] 6.3 Confirm this change does **not** add `examples/todo` (leave for `forma-todo-example`); verify no new example package was introduced

## Workflow follow-up

- Human review of planning artifacts / implementation PR.
- Archive OpenSpec change after merge; sync delta specs into `openspec/specs/`.
- Queue **`forma-todo-example`** next (separate change) to ship in-browser todo on this stack.
