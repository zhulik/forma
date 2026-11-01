# Tasks

Roadmap: CLI scaffold (before P3a). Verify Go changes with `task test` and `task lint` from repo root and `task -d cmd/forma check` (or equivalent) once the nested module exists.

## 1. CLI module skeleton

- [x] 1.1 Create `cmd/forma/go.mod`, `main.go`, and `internal/cli` root command with urfave/cli v3; verify `go build -o /dev/null .` from `cmd/forma`
- [x] 1.2 Add pal dependency and `internal/cli/app` runner (configure slog from `--log-level` / env bound in main); verify a no-op subcommand runs through pal without panic
- [x] 1.3 Implement `version` subcommand and root `Command.Version` from build metadata; verify `forma version` prints non-empty output after build

## 2. Templates embed pipeline

- [x] 2.1 Add `internal/templates` with `//go:embed`, `Names`, `Exists`, `Apply` (`.tmpl` strip, `text/template`, `Data` with Module/Package); verify unit tests for Apply output paths and content
- [x] 2.2 Add `templates/minimal/` scaffold (`forma.yaml.tmpl`, `main.go.tmpl`, …) producing buildable minimal app; verify template tests assert key files render

## 3. Init and sync commands

- [x] 3.1 Implement `initcmd` with empty-dir check, `--directory`, `--template` (default `minimal`), `--only-config`; verify unit tests for ErrNotEmpty and config-only fail-when-exists
- [x] 3.2 Wire full init: `go mod init`, template Apply in place, invoke sync stub; verify integration test or runner test with temp dir and fake executor where needed
- [x] 3.3 Implement sync stub (read `forma.yaml`, require forma module, `go mod tidy`); verify runner test with temp module
- [x] 3.4 Register `sync` subcommand; stub `add`/`dev` with clear not-implemented errors if exposed; verify `forma sync --help` and stub behavior via tests

## 4. cmdexec and sentinels

- [x] 4.1 Add `internal/cmdexec` for `go` with live stdio; verify test with `go version` subprocess
- [x] 4.2 Define init/sync sentinel errors and user-facing messages; verify errors.Is in tests

## 5. CI and Task

- [x] 5.1 Add Task target `check-init-scaffold` (build CLI, temp dir init, replace, `go build`); verify `task check-init-scaffold` passes locally
- [x] 5.2 Wire CI workflow step or job to run init smoke; verify workflow YAML references the Task target
- [x] 5.3 Extend root or nested Taskfile for lint/test of `cmd/forma`; verify `golangci-lint` runs clean on CLI module

## 6. Agent rules and author docs

- [x] 6.1 Add `.cursor/rules/forma-cli.mdc` (urfave v3, pal, templates, init in-place, nested module); verify file present and alwaysApply or globs set
- [x] 6.2 Update `doc/user/install.md`, `cli.md`, add manual setup section/page; verify documented steps match `minimal` template output
- [x] 6.3 Update `doc/ROADMAP.md` CLI/status table for scaffold phase; verify no contradiction with proposal

## 7. CLI module tests

- [x] 7.1 Add subprocess init smoke test (built binary, `--no-interactive` if added or plain flags, replace, `go build`); verify `go test ./...` in `cmd/forma` passes
- [x] 7.2 Enforce external test packages where project convention requires; verify linter testpackage settings for CLI module if configured

## Workflow follow-up

- Archive change and sync deltas to `openspec/specs/` after review.
- Update `forma-cli` Purpose in main spec away from “contract only” on archive.
