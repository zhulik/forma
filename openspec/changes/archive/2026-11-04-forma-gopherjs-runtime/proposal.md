# Proposal

## Why

P3b ships a React host that can render a forma tree, but Go `UIMount` / `RefreshUI` remain stubs and never drive that seam. Without a GopherJS bridge—tree crossing, callback dispatch, and real `ui.Context`—authors cannot run interactive Go UI in the browser.

## What Changes

- Implement the **browser runtime** behind `UIMount` and `ui.Context` (`RefreshUI`, `Update`): mount a root `Component`, rebuild the `Node` tree from Go state, and push it into the P3b host mount seam.
- Define **tree crossing** via GopherJS (structs + function values on the hot path—not JSON) and **callback binding** so host-sanitized JS invocations reach author handlers with a live `ui.Context` as the first argument.
- Propagate **cancellation** on the callback context across refresh generations so async work can observe `ctx.Done()`.
- Wire the **scaffold host entry** (`main.tsx`) to expose an imperative tree setter the Go runtime calls; keep React dumb (no domain state).
- Keep **native `go build` / `go test`** working via build-tagged stubs and an injectable host sink for tests.
- Update **`doc/user/`** so the bridge is documented as available (manual load); leave **`forma dev`** labeled planned until P3d.

**Out of scope:** `forma dev` / Vite+GopherJS watch orchestration (P3d); Wasm/`syscall/js` bridge; SSR; domain state in React; rewriting the P3b recursive renderer or text-input hybrid policy; desktop packaging.

## Capabilities

### New Capabilities

- `ui-runtime`: Go↔browser bridge that mounts a root component, refreshes the render tree into the host, dispatches sanitized host callbacks into Go handlers with `ui.Context`, and cancels prior handler contexts on refresh/unmount.

### Modified Capabilities

- `ui-dsl`: Application mount and refresh requirements no longer allow a permanently stubbed live bridge for browser builds; runtime behavior must match the authoring contracts (`Update` / `RefreshUI` / callback context).
- `ui-host`: Mount seam must expose a replaceable tree-root API the Go bridge can drive (imperative set), not only a static fixture/`null` prop at boot.
- `ui-docs`: Author guide must describe the GopherJS bridge path and how trees/callbacks cross; keep `forma dev` planned until P3d.

## Impact

- **ROADMAP phase: P3c** — advances the GopherJS runtime row; does not claim P3d `forma dev` or a one-command browser loop.
- **Go framework:** `ui/` gains real mount/refresh/context implementations (build-tagged `js` vs native stub) plus callback binding and host-sink abstraction.
- **Scaffold:** `cmd/forma/templates/scaffolds/minimal/web/src/` host entry registers a bridge `setTree` (or equivalent) for GopherJS; `FormaHost` remains the renderer.
- **Docs:** `doc/user/` status/troubleshooting/writing-ui (and related) updated for P3c vs P3d.
- **Tests:** Native Go tests against a fake host sink; scaffold Vitest only if the new host seam needs coverage; full in-browser todo demo via `forma dev` remains P3d.
- **Dependencies:** GopherJS as the browser compile toolchain for apps (document; do not introduce Wasm). Prefer `fn` for tree walks; no new JS bridge libraries.
- **Follow-ons:** P3d (`forma-dev-experience`) stacks on this branch to run host + Go rebuild as one command.
