# Tasks

## 1. Init scaffold and CLI copy

- [x] 1.1 Change `cmd/forma/templates/scaffolds/minimal/forma.yaml.tmpl` to `web:` / `path: web` and update `cmd/forma/templates/templates_test.go` (or init tests) if they assert the old yaml shape; verify with `go test ./cmd/forma/...`
- [x] 1.2 Reword stub sync user-facing strings and comments in `cmd/forma/internal/cli/subcommands/sync/` and `cmd/forma/internal/core/interfaces.go` to say “web tree” where they mean the filesystem/npm side; verify `task test` and `task lint`

## 2. Author guide (`doc/user/`)

- [x] 2.1 Update `project-layout.md`, `forma-yaml.md`, `cli.md`, `manual-setup.md`, `adding-components.md`, and `advanced.md` path/config references (`web/`, `web.path`, `web:`); keep “React host” only for runtime/registry prose per design.md; verify `rg 'host/' doc/user` and `rg 'host\.path' doc/user` return no path/config hits (runtime-only “host” in README/troubleshooting/writing-ui is OK)
- [x] 2.2 Adjust `doc/user/README.md`, `writing-ui.md`, and `troubleshooting.md` mount/HTML wording to point at the web tree (`web/index.html` or equivalent) without renaming runtime concepts; verify examples in manual-setup yaml match the init template

## 3. Canonical docs

- [x] 3.1 Replace layout mentions of `host/` with `web/` in `doc/ROADMAP.md` and `doc/DESIGN.md` where they describe user app directories (not “JavaScript host” architecture); verify with targeted search for `` `host/` `` in those files

## 4. Verification

- [x] 4.1 Run `openspec validate rename-host-to-web --strict` and fix any spec delta issues; run `task test` and `task lint` for the full repo

## Workflow follow-up

- Archive the change and sync deltas into `openspec/specs/forma-cli/spec.md` and `openspec/specs/ui-docs/spec.md` when review is complete.
