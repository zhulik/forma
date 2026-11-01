# Tasks

## 1. OpenSpec and policy

- [x] 1.1 Add docs-first guidance to `openspec/config.yaml` (`rules.proposal`, `rules.tasks`)
- [x] 1.2 Run `openspec validate forma-user-docs --strict`

## 2. Author guide (`doc/user/`)

- [x] 2.1 Add `doc/user/README.md` (map, docs-first note, availability table)
- [x] 2.2 Add `doc/user/install.md` (toolchain, `go install` CLI when available, init flow)
- [x] 2.3 Add `doc/user/forma-yaml.md` (`forma.yaml` fields, defaults including `internal/ui`)
- [x] 2.4 Add `doc/user/cli.md` (`init`, `add`, `sync`, planned `dev`)
- [x] 2.5 Add `doc/user/project-layout.md` (user app vs this framework repo)
- [x] 2.6 Add `doc/user/writing-ui.md` (state, `UIMount`, `RefreshUI`, dot imports, `fn`, link `main.go`)
- [x] 2.7 Add `doc/user/adding-components.md` (`forma add`, sync, git expectations)
- [x] 2.8 Add `doc/user/advanced.md` (codegen dir override, custom Go/JS)
- [x] 2.9 Add `doc/user/troubleshooting.md`

## 3. Canonical doc alignment

- [x] 3.1 Update `README.md` with author guide link
- [x] 3.2 Update `doc/ROADMAP.md` (P2 in progress/done, CLI-first onboarding, defer site)
- [x] 3.3 Update `doc/DESIGN.md` “How users work with forma” for CLI + `forma.yaml`

## Workflow follow-up

- Archive after review; sync specs to `openspec/specs/ui-docs/` and `openspec/specs/forma-cli/`.
- Implement `cmd/forma` in a separate **`forma-cli`** change against this spec.
