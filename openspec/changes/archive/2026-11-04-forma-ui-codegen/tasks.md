# Tasks

## 1. Codegen library (pure Go)

- [x] 1.1 Add in-process Go TSX analysis + emit under `cmd/forma/internal/codegen` (official tree-sitter + tree-sitter-typescript TSX; no Node); verify generate on fixture TSX with exit 0
- [x] 1.2 Implement leaf widget emission (props, With/Or, Render, Type string) and verify golden output matches checked-in fixture for button
- [x] 1.3 Implement compound classification (root + prefixed parts → slot fields + registry slot map) and verify golden output for accordion or dialog fixture
- [x] 1.4 Implement stale output pruning when TSX removed and verify second run deletes orphan `*.gen.go`
- [x] 1.5 Emit single `registry.gen.ts` and verify it imports the same React symbols as the TSX sources for fixtures
- [x] 1.6 Remove `cmd/forma/tools/forma-codegen` Node package and drop Node build steps from `task test-cli` for codegen

## 2. CLI integration

- [x] 2.1 Add `core.CodegenRunner` port + implementation invoking the codegen library with paths from `formaconfig`; unit test with fake runner records invoked paths
- [x] 2.2 Implement `forma generate` subcommand that loads `forma.yaml` and runs codegen/registry only (no npm, shadcn, or `go mod tidy`); verify generate test
- [x] 2.3 Wire sync: after npm install, call the same generate path (remove skip logs); pass init flag or starter shadcn list only on init-driven sync; verify `sync_test.go` expects codegen call
- [x] 2.4 Define CLI constant for init starter components (`button`, `input`, `card`); init path runs shadcn for starters then codegen via sync; verify init integration test sees three shadcn adds once
- [x] 2.5 Implement `forma add <component>` (shadcn add + sync) replacing `ErrNotImplemented`; verify add test
- [x] 2.6 Remove `ui.components` from `forma.yaml.tmpl` and `formaconfig` if unused; update `project_test.go`; verify load still requires mount, web, codegen

## 3. Scaffold and host wiring

- [x] 3.1 Import generated `registry.gen.ts` from scaffold `main.tsx` (or thin wrapper); verify TypeScript build of web tree succeeds after mock codegen output in test fixture
- [x] 3.2 Add root `generator.go` to init template with `//go:generate forma generate`; verify template includes the file
- [x] 3.3 Add minimal `Taskfile.yaml` to init template with `check-generate` running `go generate` then `git diff --exit-code` on the whole repo; verify template test or smoke includes task file

## 4. Author docs and roadmap

- [x] 4.1 Update `doc/user/forma-yaml.md`, `cli.md`, `adding-components.md`, `project-layout.md`, `install.md`, README status for `forma generate`, full sync calling generate, no component list, add workflow, committed gen output, scaffold `generator.go` / `check-generate`; verify no doc claims partial sync skip for codegen or documents `--check` flags
- [x] 4.2 Clarify docs that generate needs no Node; Node remains for npm/shadcn on sync/init/add; update ROADMAP P3a when complete

## 5. CI and quality gates

- [x] 5.1 Extend init scaffold smoke (`task check-init-scaffold` or CI job) to assert `internal/ui/*.gen.go` and `generator.go` exist and `go build` succeeds after init; verify locally
- [x] 5.2 Run `task test` and `task lint` in `cmd/forma` and framework root after Go codegen rewrite; verify both pass

## Workflow follow-up

- Follow-up change: `examples/` demos and remove framework `ui/*.gen.go` stubs.
