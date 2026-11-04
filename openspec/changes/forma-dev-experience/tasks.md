# Tasks

Roadmap: **P3d** (`forma-dev-experience`). Stack base: `change/forma-gopherjs-runtime` (P3c). Verify Go changes with `task test` and `task lint` from repo root (and nested CLI checks as today). Do not re-implement P3b host or P3c bridge—consume their seams.

## 1. Align with P3c build seam

- [ ] 1.1 Rebase `change/forma-dev-experience` onto the latest `change/forma-gopherjs-runtime` tip; verify the worktree builds and P3c’s documented GopherJS output path / build command are readable from P3c design or code
- [ ] 1.2 Record the concrete build invocation + artifact path used by `forma dev` in a short comment or internal constant next to the orchestrator (matching P3c); verify a manual one-shot rebuild writes the expected file under the web tree

## 2. Process helpers and dependencies

- [ ] 2.1 Extend `cmdexec` (or add a thin sibling) to start long-running commands with cancelable context and forwarded stdio; verify unit test covers cancel kills the child
- [ ] 2.2 Add `fsnotify` (and `errgroup` if not already direct) to `cmd/forma` module; verify `go mod tidy` in `cmd/forma` succeeds and builds

## 3. Implement `forma dev`

- [ ] 3.1 Replace the stub Action with load-`forma.yaml`, resolve project/`web.path`, and fail-fast checks (missing yaml, missing `node_modules`, missing GopherJS/toolchain); verify unit/integration tests assert clear errors without starting Vite
- [ ] 3.2 Start `npm run dev` under `web.path` via the long-running helper; verify with a fake/stub executor or temp project that the npm args and working directory are correct
- [ ] 3.3 Implement debounced Go watch → GopherJS rebuild (serialized, no overlap), initial build on start, keep loop on rebuild failure after start; verify tests for debounce/serialize behavior with a fake builder
- [ ] 3.4 Wire SIGINT/SIGTERM cancel to stop Vite + watcher and exit; verify cancel path test or documented manual check
- [ ] 3.5 Register usage/help text as implemented (not “planned”); verify `forma dev --help` no longer says planned-only

## 4. Author docs and status

- [ ] 4.1 Update `doc/user/cli.md`, `install.md`, `troubleshooting.md`, and `doc/user/README.md` status so `forma dev` is available; document HMR vs Go reload and prerequisites; verify pages match shipped flags/behavior
- [ ] 4.2 Document the todo-demo browser verification path (adapt repo `main.go` patterns into an init’d app + `forma sync` + `forma dev`); verify the guide is followable without claiming framework-root `web/`
- [ ] 4.3 Update root `README.md` / any remaining “planned P3d” lines that this change owns; verify no doc still marks `forma dev` as planned

## 5. Integration verification

- [ ] 5.1 Manual smoke on stacked tip: init (or existing scaffold), sync, `forma dev`, open browser, confirm interactive UI via P3c bridge; save a Go file and confirm rebuild + reload path works; verify against P3d done criteria
- [ ] 5.2 Run `task test` and `task lint` (and CLI module tests) green on the branch; verify CI-equivalent checks pass locally

## Workflow follow-up

- Archive change and sync deltas to `openspec/specs/` after the shipping PR merges.
- Update `doc/ROADMAP.md` live fields (Status / PR) when the PR opens; mark P3d Done on archive.
