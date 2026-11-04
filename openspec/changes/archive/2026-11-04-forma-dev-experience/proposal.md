# Proposal

## Why

P3b and P3c deliver a React host and a live GopherJS `UIMount`/`RefreshUI` bridge, but authors still have no one-command browser loop—`forma dev` remains a stub. Without orchestration of the Vite host plus Go rebuild, the reference todo demo cannot be exercised end-to-end in a browser.

## What Changes

- Replace the **`forma dev` stub** with a real long-running command that starts the web host and rebuilds the Go app for the browser whenever sources change.
- Document the **author workflow**: prerequisites (`forma sync`, Node, GopherJS), how Vite HMR vs Go rebuild/reload work, and how to run the **todo demo** in the browser.
- Update **`doc/user/`** (and related status lines) so `forma dev` is shipped, not planned.
- Add focused CLI tests for orchestration seams (missing config/deps, process lifecycle) without claiming a full browser E2E harness in CI unless a cheap smoke path is already available.

**Out of scope:** Re-implementing the React host (P3b) or GopherJS `UIMount`/`RefreshUI` bridge and tree crossing (P3c)—this change consumes those seams. Desktop packaging (P4). Wasm. Auto-running full `forma sync` on every `dev` start (recommend once; fail clearly if deps missing). True Go HMR / preserving in-memory Go state across rebuilds.

## Capabilities

### New Capabilities

- (none)

### Modified Capabilities

- `forma-cli`: `forma dev` becomes an implemented browser development loop (Vite host + GopherJS rebuild) instead of a planned stub.
- `ui-docs`: Author guide documents the live `forma dev` workflow, rebuild/HMR semantics, and browser todo-demo verification; remove “planned until P3d” labels for this command.

## Impact

- **ROADMAP phase: P3d** — completes the end-to-end browser path (`forma-dev-experience`); stacked PR base is `change/forma-gopherjs-runtime` (P3c), not `main`.
- **CLI:** `cmd/forma/internal/cli/subcommands/dev/` and supporting orchestration (watch + child processes) under `cmd/forma/internal/`.
- **Scaffold / host:** Minimal wiring only as needed to load P3c’s GopherJS artifact under Vite (no host rewrite); reuse existing `npm run dev`.
- **Docs:** `doc/user/cli.md`, `install.md`, `troubleshooting.md`, README status tables; optional ROADMAP status row when the shipping PR opens.
- **Dependencies:** Reuse Vite, `cmdexec`, existing webtool patterns; add a well-known Go file watcher (e.g. `fsnotify`) rather than a hand-rolled poller or a third-party “dev server” framework.
- **Assumes P3c:** Real `UIMount`/`RefreshUI`, host mount seam, and a documented GopherJS build → JS artifact path served by the web tree.
