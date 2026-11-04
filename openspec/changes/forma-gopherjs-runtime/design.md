# Design

## Context

See [proposal.md](./proposal.md). P3b archived `ui-host`: scaffold `web/src/host/` renders `FormaHost` from a `tree` prop; `main.tsx.tmpl` still boots with `tree = null` and a comment reserving P3c. Framework `ui.UIMount` is an empty stub; `ui.Context` is interface-only (no concrete refresh implementation). `Node` / widget `Render` already produce the wire shape the host expects (`Type`, `Key`, `Class`, `Props`, `Children`, `Slots`). Author handlers take `ui.Context` first; host sanitizers invoke JS-facing functions with zero or one plain payload—**not** a Go context. P3d will own `forma dev`; this change only needs a loadable bridge seam.

## Goals / Non-Goals

**Goals**

- Real `UIMount` / `RefreshUI` / `Update` for the GopherJS (`js`) build.
- Push render trees into the P3b host without rewriting the recursive renderer.
- Bind author callbacks so Context is injected in Go before the host calls them.
- Cancel prior callback-context generations on refresh/unmount.
- Native (`!js`) stub + injectable host sink so `go test` / `go build` stay green without a browser.
- Scaffold `main.tsx` registers an imperative `setTree` (or equivalent) for the runtime.
- Author-doc updates for P3c vs P3d.

**Non-Goals**

- `forma dev`, Vite proxy, or watch/rebuild orchestration (P3d).
- Wasm / `syscall/js` as the v1 bridge.
- Domain state or business logic in React.
- Fine-grained/signal updates; full-tree refresh only.
- Per-subtree cancellation finer than refresh-generation scope.
- Shipping a permanent framework-repo `web/` app tree (scaffold + docs suffice; P3d may add demo orchestration).

## Decisions

### 1. Capability name `ui-runtime`

**Choice:** New OpenSpec capability **`ui-runtime`** for the Go↔host bridge (as reserved by P3b). Keep **`ui-host`** for the React layer only.

**Alternatives:** Fold bridge requirements into `ui-host`—rejected; ROADMAP and archive sync stay phase-shaped.

### 2. GopherJS on the hot path (not JSON, not Wasm)

**Choice:** Compile app/`ui` with **GopherJS**. Cross `*Node` (and nested slices/maps) as JS objects via GopherJS’s struct mapping. Do **not** `JSON.stringify` the tree each refresh. Do **not** introduce Wasm/`syscall/js` for v1.

**Reuse vs hand-roll:** Reuse GopherJS as the Go→JS compiler (DESIGN.md). Hand-roll only the small mount/refresh/bind layer in `ui/`—no third-party “Go UI bridge” package.

**Alternatives:** JSON DTO each frame—rejected (DESIGN hot-path rule). Wasm—deferred by DESIGN.

### 3. Build tags: `js` runtime vs `!js` stub

**Choice:**

| Build | Behavior |
|-------|----------|
| `js` (GopherJS) | Real `UIMount`, refresh, callback bind, host lookup |
| `!js` (native) | No-op `UIMount` / refresh (current stub behavior) so `main.go` and packages keep compiling under `go test` |

Shared types (`Context` interface, `Node`, widgets) stay untagged.

**Alternatives:** Single implementation using build-time interfaces only—still need tags or stub files to avoid pulling JS APIs into native builds.

### 4. Host seam: imperative `setTree` registered by scaffold

**Choice:** Scaffold `main.tsx` lifts tree into React state, mounts `<FormaHost registry={…} tree={tree} />`, and registers a stable global bridge object, e.g. `globalThis.formaHost = { setTree(tree) { … } }` (exact name documented in code + docs). Go `UIMount` / refresh call `setTree` with the bound `*Node` root. Initial tree may still be `null` until the first Go refresh.

**Reuse:** Existing `mountFormaHost` + `FormaHost`; only add state + registration. Prefer a tiny `createFormaBridge` helper in `web/src/host/` over scattering globals in `main.tsx`.

**Alternatives:** Go creates its own React root—duplicates P3b mount, rejected. Props-only with no imperative API—GopherJS cannot re-render React without a setter. Custom events / postMessage—unnecessary indirection.

### 5. Callback binding happens in Go before `setTree`

**Choice:** After `root.Render(stdCtx)`, walk the `*Node` tree and rewrite function values on `Props` into **host-callable** closures that:

1. Capture the current generation’s `ui.Context`.
2. Accept the sanitized payloads the host already passes (`()`, `(string)`, `(bool)`, …).
3. Invoke the author handler as `handler(ctx, …payload)`.

Host sanitizers stay unchanged. Authors keep writing `func(ctx Context, …)`.

**Reuse vs hand-roll:** Use **`fn`** (`Map` / walk helpers) for Children/Slots traversal; extend `fn` only if a missing helper is clearly reusable. Reflect-based prop field rewrite is hand-rolled in `ui` (no third-party schema binder)—forma `Props` are widget structs, not a generic JSON schema.

**Alternatives:** Teach the host to invent Context—impossible in JS. Callback ID registry / RPC—rejects DESIGN “funcs on the tree.” Change author signatures to drop Context—**BREAKING** and worse ergonomics; rejected.

### 6. Concrete `ui.Context` + refresh-generation cancellation

**Choice:** Unexported concrete type implementing `Context` (embeds cancellable `context.Context`). Mount owns a parent cancel. Each `RefreshUI`:

1. Cancel the previous generation.
2. Derive a new generation context.
3. Render → bind → `setTree`.

`Update(fn)`: run `fn` under `defer`/`recover`; on normal return call `RefreshUI` once; on panic, no refresh (matches `ui-dsl`).

**Alternatives:** Never cancel—leaks goroutines vs DESIGN. Per-node subtree cancel—deferred (non-goal). Coalesce multiple RefreshUI into one microtask—optional later; v1 is synchronous immediate refresh.

### 7. Injectable host sink for tests (and js lookup)

**Choice:** Small internal port, e.g. `HostSink` with `SetTree(*Node)`, set by:

- `js` build: adapter that calls `globalThis.formaHost.setTree`
- tests: fake sink recording trees / counting refreshes

Do **not** add Init wiring guards for a nil sink on the native stub path; on `js`, missing seam is a developer-visible failure at mount/refresh use (fail closed).

**Alternatives:** Only browser integration tests—too heavy for P3c CI; rejected as the sole strategy.

### 8. Script load order (manual until P3d)

**Choice:** Document that the host must register `formaHost` **before** the GopherJS `main` runs `UIMount`. Scaffold wires registration in `main.tsx` before importing/loading the Go bundle entry (placeholder import or documented script tag order). P3d replaces manual steps with `forma dev`.

**Alternatives:** Go polls until seam appears—acceptable brief retry only if load races are observed; prefer deterministic order first.

### 9. Docs and ROADMAP status

**Choice:** Update `doc/user/` (README status, writing-ui, troubleshooting, install/toolchain notes) so the bridge is available and `forma dev` stays planned. Do not claim a polished one-command demo loop. ROADMAP P3c row may move to in progress/Done at PR/archive time (orch/control plane).

### 10. Prefer-fn / no new collection libs

**Choice:** Tree walks that build new child/slot slices use **`fn`**. TypeScript side stays on existing host helpers; no lodash-like dependency.

## Risks / Trade-offs

- **[GopherJS object ≠ plain `{}` for host `isPlainObject`]** → Mitigate by verifying Props enumeration under GopherJS early; if needed, normalize Props to a plain map during bind (still no JSON of the whole tree).
- **[Reflect callback wrapping misses exotic signatures]** → Cover DSL shapes used by codegen (`func(Context)`, `func(Context, string)`, `func(Context, bool)`); document that unsupported arities fail loudly in tests.
- **[Load-order races without forma dev]** → Document order; optional short poll only if needed; P3d owns ergonomics.
- **[Full-tree refresh cost]** → Accepted for v1 (DESIGN); no signal system.
- **[Native stub diverges from js]** → Shared tests against `HostSink` cover Update/Refresh/cancel; js adapter stays thin.

## Migration Plan

1. Land `ui` runtime (tagged) + scaffold bridge registration + docs in this change.
2. Apps: rebuild with GopherJS and load after host registration (manual until P3d).
3. No dual-write/compat shims (early-stage).
4. P3d stacks on this branch for `forma dev`.

## Open Questions

None material for builder. Debounced OnChange remains a host/widget policy concern (already P3b), not a bridge question.
