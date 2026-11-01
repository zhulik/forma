# Tasks

## 1. Web embed scaffold

- [x] 1.1 Add `web/**` under `cmd/forma/templates/scaffolds/minimal/` (Vite + React + TS + Tailwind, `components.json`, `index.html` with `app-root`, `package.json` with pinned deps); align `forma.yaml.tmpl` default `ui.components` with button, input, card
- [x] 1.2 Extend template tests to assert web paths and mount id; verify `go test ./cmd/forma/templates/...`

## 2. Config and sync core

- [x] 2.1 Extend `formaconfig.Project` for `ui`, `web`, `codegen`; load in sync; validate `web.path`
- [x] 2.2 Add injectable npm/shadcn runners (core port + `cmdexec` implementation); implement sync steps: npm install, shadcn add for `ui.components`, stub logs for codegen/registry, then existing Go mod steps
- [x] 2.3 Unit tests with fakes for npm/shadcn; update `sync_test.go` and init tests as needed; verify `go test -race ./...` in `cmd/forma`

## 3. CI and smoke

- [x] 3.1 Ensure Node.js on CI job running `task check-init-scaffold`; optionally commit generated `package-lock.json` in embed after golden `npm install`
- [x] 3.2 Confirm init integration smoke passes with network/npm; verify `task check-init-scaffold` and root `task test`

## 4. Author docs

- [x] 4.1 Update `doc/user/install.md`, `cli.md`, `project-layout.md` (and README status if needed) for partial sync vs codegen/dev; verify no doc claims working browser app

## 5. Verify

- [x] 5.1 Run `openspec validate forma-web-scaffold --strict`
- [x] 5.2 Run `task test` and `task lint` for full repo
