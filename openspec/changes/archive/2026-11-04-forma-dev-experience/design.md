# Design

## Context

See `proposal.md` — Why. Stack below this change: P3b React host (`web/src/host/`, Vite `npm run dev`) and P3c GopherJS runtime (real `UIMount`/`RefreshUI`, host mount seam, documented browser Go build → JS artifact). Current tip may still equal P3b until P3c commits land; design against that P3c contract. `forma dev` today is a stub in `cmd/forma/internal/cli/subcommands/dev`.

## Goals / Non-Goals

**Goals**

- One author command: host up + Go rebuild on change.
- Clear docs for HMR vs Go reload and a browser todo-demo verification path.
- Thin orchestration only—consume P3b/P3c seams.

**Non-Goals (design-level)**

- No second host renderer or alternate bridge.
- No Wasm path; no desktop packaging.
- No preserving Go heap across rebuilds (not true Go HMR).
- No auto `forma sync` on every start.

## Decisions

Copy into the shipping PR body.

1. **`forma dev` = process supervisor, not a new runtime**  
   Long-lived CLI starts (a) the existing Vite script under `web.path` (`npm run dev`) and (b) a GopherJS rebuild loop for the app’s browser entry. Does not reimplement host or `UIMount`.  
   *Alternatives:* separate `task` targets only (rejected—ROADMAP requires one CLI command); embed a custom HTTP server (rejected—reuse Vite).

2. **Reuse Vite for host HMR; Go gets fast rebuild + reload**  
   Host TS/CSS keep Vite HMR. Go changes: rebuild browser artifact, then full page reload (or P3c remount if exposed). Document explicitly—do not claim Go HMR.  
   *Alternatives:* custom websocket “soft HMR” for Go (deferred—complexity, loses state honesty).

3. **Consume P3c’s build + load contract—do not invent a second compile path**  
   Invoke the same GopherJS build entrypoint and output path P3c documents (artifact under the web tree Vite serves, or the import P3c wires in `main.tsx`). Preserve P3c’s **host-before-Go** order: scaffold registers `formaHost.setTree` (or equivalent) before the GopherJS `main` runs `UIMount`. `forma dev` only rebuilds that Go bundle and keeps the Vite host up—it does not invent a parallel `-o` path or a second bridge.  
   *Alternatives:* P3d defines its own output path (rejected—stack drift); poll until `formaHost` appears (P3c already prefers deterministic order).

4. **Watch with `moddwatch` (not air-as-library); serialize rebuilds**  
   Watch project Go packages relevant to the browser build (app module + local replace of forma when present). Use [`github.com/cortesi/moddwatch`](https://github.com/cortesi/moddwatch) (MIT; powers `modd`) for recursive watch + lull batching; serialize rebuilds (no overlapping GoScript runs). Authors need only Node + GoScript—no global live-reload binary.  
   *Reuse:* `cmdexec` for child processes; `golang.org/x/sync/errgroup` for supervising Vite + watcher.  
   *Rejected — air as library (`github.com/air-verse/air/runner`):* GPL-3.0 (viral link into `forma`); `Engine.Run` calls `os.Exit` on watch setup failure (unsafe under pal); always `runBin` after build (wrong shape next to long-lived Vite); pulls Hugo and a large unrelated module graph.  
   *Alternatives:* hand-rolled `fsnotify` tree walk (replaced here); require a global `air`/`watchexec` binary (extra toolchain); `go run` air as a subprocess (possible but slower/offline-fragile vs linking MIT `moddwatch`).

5. **Fail fast on missing prerequisites; do not auto-sync**  
   Require readable `forma.yaml`. If `node_modules` (or equivalent) missing under `web.path`, error → run `forma sync`. If `gopherjs` (or P3c-documented tool) missing on PATH, error → install docs. Optional `-d/--directory` like other commands.  
   *Alternatives:* silent `npm install` on start (rejected—surprising, slow).

6. **Lifecycle: shared cancel, forward stdio, exit on fatal child death**  
   Context cancel (SIGINT/SIGTERM) stops Vite and watcher. Forward child stdout/stderr. If Vite exits unexpectedly, `forma dev` exits non-zero. Failed Go rebuilds log and keep watching (edit-fix loop); initial build failure before listen MAY exit non-zero so authors notice a broken toolchain immediately.

7. **Todo demo success path**  
   Acceptance: with P3c stacked underneath, an initialized app using the reference todo patterns (`main.go`) runs under `forma dev` and is interactive in the browser. Keep scaffold `main.go.tmpl` as the small hello unless P3c already upgraded it; document copying/adapting the repo todo into the app (or a short recipe). Do not add a second permanent demo app tree in the framework root in this change.  
   *Alternatives:* replace scaffold with full todo (nice later; avoid blocking P3d on catalog widget parity in the template).

8. **Specs / docs owned here; `ui-runtime` stays P3c**  
   Delta only `forma-cli` + `ui-docs`. No new capability. No host behavior rewrite beyond whatever tiny Vite/static wiring P3c already requires for the artifact.

9. **prefer-reuse / prefer-fn**  
   Reuse Vite, npm script, `cmdexec`, webtool require-npm patterns, pal/`app.Run` if the subcommand needs DI. Add `moddwatch` (+ errgroup). Do not add a lodash-like JS util. Any Go slice remaps in new CLI code use `fn`. Do not replace `fn/`.

10. **Stacked PR base**  
    Feature branch `change/forma-dev-experience`; PR targets `change/forma-gopherjs-runtime`. Rebase onto P3c tip before merge-ready apply.

11. **No hand-rolled process signal subscription in CLI commands**  
    Pal’s `Run` already `signal.NotifyContext`s `SIGINT`/`SIGTERM`. Subcommands pass the urfave context into `app.Run` / `pal.Run` and must not wrap it again with `signal.NotifyContext` (double-subscribe / competing stop).

## Risks / Trade-offs

- **[P3c artifact path]** → Concrete contract: `gopherjs build -o <web.path>/public/app.js .`; scaffold `main.tsx` loads `/app.js` after `formaHost` registration.
- **[GopherJS vs Go 1.27]** → Current GopherJS `v1.21.x` cannot compile modules with `go 1.27.x`. Documented in install docs; orchestration still ships. Full interactive browser smoke waits on a GopherJS that understands forma's Go version.
- **[GopherJS rebuild latency]** → Debounce + clear rebuilding logs; document reload expectation.
- **[Watching too much of module cache]** → Limit watch roots to app dir; ignore `web/` / `node_modules` / `.git`.
- **[Framework repo has no `web/`]** → Success is app-project oriented; docs must not imply forma dev at framework root without a scaffolded app.
- **[Initial build vs Vite race]** → Initial GopherJS build runs before Vite start; host may briefly load before `/app.js` is fetched—document reload.

## Migration Plan

1. Land P3c on the stack; rebase this branch.  
2. Implement `forma dev` + docs; remove stub/`ErrNotImplemented`.  
3. Authors: install GopherJS per guide, `forma sync` once, `forma dev`.  
4. No compat shims (early-stage).

## Open Questions

None material for builder. Exact GopherJS CLI flags/output path come from P3c artifacts after rebase—adapter only, not a new product decision.
