# Tasks

## 1. Host seam (scaffold)

- [ ] 1.1 Add a small bridge helper under scaffold `web/src/host/` that holds React tree state / exposes `setTree` (or equivalent) and register it on `globalThis` from `main.tsx.tmpl`; verify a Vitest covers imperative tree update re-rendering `FormaHost` without remounting a different DOM id
- [ ] 1.2 Keep `mountFormaHost` fail-closed behavior and document the global bridge name in the host module comment; verify existing mount tests still pass via scaffold `npm test`

## 2. Go runtime core (`ui/`)

- [ ] 2.1 Introduce a `HostSink` (or equivalent) port and concrete `ui.Context` with `RefreshUI` / `Update` (panic → no refresh); verify native unit tests cover successful update (one refresh), panic skips refresh, and manual `RefreshUI`
- [ ] 2.2 Implement refresh-generation cancellation (prior callback contexts done on refresh/unmount); verify unit tests observe `ctx.Done()` after a second refresh
- [ ] 2.3 Implement callback binding: walk `*Node` Children/Slots with `fn`, rewrite function Props into host-callable closures that inject `ui.Context`; verify unit tests for click-style (`Context` only) and value-style (`Context` + string/bool) dispatch against a fake sink
- [ ] 2.4 Split `UIMount` / refresh with build tags: `!js` no-op stub (native compile safe), `js` looks up `formaHost.setTree`, mounts root, initial refresh; verify `go build ./...` and `go test ./ui/...` succeed on native, and `js` sources compile under GopherJS when the toolchain is available

## 3. Author docs (P3c vs P3d)

- [ ] 3.1 Update `doc/user/` (README status, writing-ui, troubleshooting, install/toolchain as needed) so the Go↔host bridge is available and `forma dev` stays planned; verify docs do not claim a one-command live rebuild loop
- [ ] 3.2 Align any status lines that still say the bridge is only planned with P3c; verify troubleshooting distinguishes missing host seam vs missing `forma dev`

## 4. Integration check

- [ ] 4.1 Run `task test` and `task lint` after Go/docs changes and fix regressions; verify both exit 0
- [ ] 4.2 Smoke: with host bridge registered, a GopherJS-built entry calling `UIMount` pushes a non-null tree (manual or scripted check documented in the PR); verify at least one refresh after a bound callback mutates state

## Workflow follow-up

- After implementation PR merges and archive, sync `ui-runtime` (and deltas) into `openspec/specs/` and mark ROADMAP P3c Done.
- P3d (`forma-dev-experience`) stacks on this branch for `forma dev`; do not expand this change into that work.
