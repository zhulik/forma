# Design

## Context

See [proposal.md](./proposal.md). **Implemented:** init embeds `web/**` under the `minimal` template; sync runs npm + shadcn then Go module steps; codegen/registry log-only stubs. `formaconfig` loads full project shape. Author docs describe partial sync.

## Goals / Non-Goals

**Goals**

- After full init, disk matches `doc/user/project-layout.md` for Go + `web/` tooling (`node_modules` not committed).
- Sync is idempotent enough for re-runs: npm install + shadcn add for declared components.
- Clear operator feedback when codegen/registry steps are intentionally skipped.

**Non-Goals**

- ts-morph Go codegen (P3a).
- React host registry and GopherJS bridge (later P3).
- Implementing `forma add` / `forma dev`.
- Pinning shadcn CLI to a fixed major (still `@latest` unless CI forces a pin).

## Decisions

### 1. Web shell in init embed, materialization in sync

**Choice:** Init template writes static `web/**` (including `package.json`, `components.json`, committed `package-lock.json`, Vite + Tailwind + shadcn-ready layout). Sync runs npm + shadcn against that tree.

**Alternatives:** Generate entire web tree only in sync—rejected; authors expect files immediately after init even before network.

### 2. Sync pipeline order

**Choice:**

1. Load `forma.yaml` (full project shape).
2. Resolve `webRoot = join(projectDir, web.path)`; require `package.json`.
3. `npm install` in `webRoot` (via `cmdexec.Run`).
4. `npx --yes shadcn@latest add --yes <components...>` in one invocation when `ui.components` is non-empty (dedupe via `fn.Uniq(fn.Filter(...))`, preserve first-seen order).
5. Stub: info logs that Go codegen and JS registry update are skipped (P3a/host).
6. Go steps: require/replace forma, `go mod tidy`.

**Rationale:** Baseline deps from embed install first; shadcn `add` fetches TSX and extra radix packages. Go last avoids partial module state if JS fails.

### 3. shadcn CLI invocation

**Choice:** Single batched `npx --yes shadcn@latest add --yes <names...>` with `cmd.Dir = webRoot`.

**Flags:** `--yes` for non-interactive; no `--overwrite` by default.

**Precondition:** Init embed ships valid `components.json` and Tailwind/CSS so `add` does not prompt for project setup.

### 4. Pinning and lockfiles

**Choice:** Embed `package.json` with pinned deps; golden `package-lock.json` committed in the template. Sync may refresh lockfile after shadcn mutates deps.

### 5. Config parsing

**Choice:** `formaconfig.Project` with `UI`, `Web`, `Codegen`. Sync fails if `web.path` lacks `package.json`. Empty `ui.components` skips shadcn add but still runs npm install.

### 6. Testability and DI

**Choice:** `WebSyncRunner` port in **`internal/core/interfaces.go`** (with `GoModRunner`, `ProjectSyncer`). Production adapter: `internal/webtool.Tool` via pal. Unit tests inject recording fakes on `synccmd.Service`. Integration tests call real sync when `testutil.RequireNpm` probe succeeds (minimal `npm install` in temp dir); otherwise skip.

### 7. CI

**Choice:** GitHub Actions `setup-node` (Node 22) before `task check-init-scaffold`. Smoke runs full init → sync (npm + shadcn) → `go build`.

### 8. Collection helpers

**Choice:** Component dedupe and similar slice work use **`github.com/zhulik/forma/fn`** (`Uniq`, `Filter`); extend `fn` when a pattern is reusable (see **prefer-fn** Cursor rule).

## Risks / Trade-offs

- **[Network flakiness on init]** → Clear errors; partial directory per init contract; retry `forma sync`.
- **[shadcn CLI `@latest` drift]** → Pin in a follow-up if CI flakes.
- **[Node missing]** → Sync fails with actionable message (`npm not found on PATH`).
- **[Codegen stub]** → TSX on disk without Go bindings; hello scaffold uses framework `ui` only.

## Migration Plan

No production deployments. Existing stub-init projects: upgrade CLI and run `forma sync` after adding a web tree manually is unsupported—prefer new `forma init`.

## Open Questions

- Pin `shadcn@` to a specific release if `@latest` flakes.
- Future `forma sync --force` with shadcn `--overwrite` (defer).
