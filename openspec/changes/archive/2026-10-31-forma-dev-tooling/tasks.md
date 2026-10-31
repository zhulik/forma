# Tasks

## 1. Static analysis and tasks

- [x] 1.1 Add `.golangci.yaml` from hosehook with module `github.com/zhulik/forma` and `ui/**/*.gen.go` path exclusions; verify `golangci-lint run --timeout=5m` exits 0 (fix any findings in existing Go code in the same task)
- [x] 1.2 Add root `Taskfile.yaml` (default → lint only; generate, fix, check-generate, check-fix, lint-fix; optional `test` target not in default); verify `task default` succeeds

## 2. Cursor rules

- [x] 2.1 Copy rule files listed in [design.md](./design.md) into `.cursor/rules/`, adapting `early-stage-no-compat.mdc` for forma; verify each file has valid frontmatter and expected globs
- [x] 2.2 Add `.cursor/hooks.json` (empty hooks) matching sister repos; verify JSON parses

## 3. Continuous integration

- [x] 3.1 Add `.github/workflows/ci.yaml` (hosehook backend job without test step); verify locally `task check-generate`, `task lint`, and `task check-fix` match CI steps

## 4. Project docs

- [x] 4.1 Add a short **Development** section to `README.md` (`task`, asdf/tool-versions pins, note tests deferred); verify commands match Taskfile names
- [x] 4.2 Note foundation tooling completion in `doc/ROADMAP.md` (status table or footnote); verify link targets exist

## Workflow follow-up

- Archive the change after review; run `openspec validate forma-dev-tooling --strict` before archive if desired.
- Add tests and CI test step in a later change when behavior is worth asserting.
