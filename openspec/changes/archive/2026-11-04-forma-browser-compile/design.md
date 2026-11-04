# Design

## Context

See proposal.md — Why. P3b (React host) and P3c (GopherJS bridge: `formaHost.setTree`, `//go:build js`, `github.com/gopherjs/gopherjs/js`) are on `main`. P3d (`forma-dev-experience`, #3) orchestrates Vite + GopherJS rebuild but cannot fully smoke against Go **1.27.3**.

Upstream GopherJS (as of 2026-11): latest release **v1.21.0** for Go **1.21.13**; open WIP scratch pad for **go1.22** only ([gopherjs#1462](https://github.com/gopherjs/gopherjs/pull/1462)). Historical pace is roughly one Go minor every many months—catching to 1.27 is not a near-term product schedule.

Observed failure compiling forma with GopherJS 1.21 (not merely `go.mod` directive mismatch): `reflect.TypeFor`, `range` over `int`, `reflect.TypeAssert` in `ui/bind.go`—idioms required by the project’s modern-Go posture.

DESIGN.md today: GopherJS v1; Wasm “Not v1” (payload / serialization). This change deliberately revisits that table; early-stage no-compat applies.

## Goals / Non-Goals

**Goals:**

- Decision-ready options analysis; prefer a **non-Wasm** path if viable.
- Unblock browser smoke of the todo demo under Go **1.27** without EOL Go 1.21 as the authoring language.
- Preserve the product intent: Go owns domain state; React is a dumb host; tree + callbacks cross without **requiring** JSON-encoding the full tree every refresh.
- Leave P3d (#3) open; this change amends toolchain + bridge so P3d can finish.

**Non-Goals:**

- Shipping the complete `forma dev` Vite loop (P3d ownership).
- Becoming GopherJS maintainers / multi-year upstream catch-up as the only plan.
- Host-driven LiveView as the default product (evaluated; not selected for v1 web).
- Rubber-stamping official Go Wasm as the winner without a failed non-Wasm path.

## Options matrix

| ID | Option | Go 1.27? | Wasm? | Keeps object/callback bridge intent? | Unblocks P3d soon? | Verdict |
|----|--------|----------|-------|--------------------------------------|--------------------|---------|
| A | Stay GopherJS + dual toolchain / pin browser `go` / `GOPHERJS_GOROOT` 1.21 | No (browser stuck on 1.21 + rewrite modern idioms) | No | Yes | Partial (after language downgrade) | **Reject** as primary — papers over the showstopper; fights `modern-go` |
| B | Fork / vendor / contribute GopherJS past 1.21 | Eventually | No | Yes | No (years; 1.22 WIP only) | **Reject** as primary; optional community contrib later |
| C | GoScript (Go→TypeScript, `s4wave/goscript`) | Yes (tooling on go 1.27) | No (emits TS; not app Wasm) | Yes (Go values as JS objects; `syscall/js`) | Yes after spike | **Recommend** |
| C′ | Other Go→JS (goany, Joy, …) | Partial / dormant | No | Weak / subset | No | **Discard** — not full-Go app compilers |
| D | TinyGo (or other) Wasm | Often | **Yes** | Via JS imports | Maybe | **Reject** for “non-Wasm” goal; still Wasm-adjacent |
| E | Official Go `js/wasm` + `syscall/js` | Yes | **Yes** | Weaker (encode/copy at boundary unless carefully designed) | Yes | **Contingency** if C fails spike — not default winner |
| F | Host-driven / JSON / WS bridge (Go native process; browser = React only) | Yes | No | No (serialized tree / RPC) | Yes with redesign | **Defer** — viable escape hatch, wrong default for DESIGN’s client-side Go app |

### Option notes (serious paths only)

**A — Dual GopherJS.** `GOPHERJS_GOROOT=go1.21` alone does not fix it: app sources already use post-1.21 language. Pinning `go 1.21` in `go.mod` (or a browser submodule) forces a language freeze/rewrite across `ui/` and author apps. Still ships an EOL browser stdlib. Unacceptable as the durable answer to a version-gap showstopper.

**B — Advance GopherJS.** Upstream is healthy but lagged; forma should not bet P3d on finishing 1.22→1.27. Contributing is welcome but orthogonal.

**C — GoScript.** Actively maintained (v0.6.x, MIT, dogfooded by Spacewave). Compiles with the Go toolchain at **1.27**; emits readable TypeScript modules; documents Go values as ordinary JS objects and `syscall/js` / `js.FuncOf` interop; claims much smaller payloads than Go Wasm. Reuse-over-handroll: adopt the mature-enough third-party compiler rather than inventing one. Risks: young ecosystem, reflect/callback edge cases, Vite integration details—mitigated by a time-boxed spike with hard acceptance criteria.

**C′ — goany / Joy.** Subset or dormant; discard quickly.

**D — TinyGo Wasm.** Still Wasm; smaller than gc Wasm sometimes, but does not satisfy “besides Wasm.”

**E — Official `GOOS=js GOARCH=wasm`.** Only mature path that matches go.mod today if GoScript is rejected. Requires rewriting P3c seam to `syscall/js` and accepting DESIGN’s prior concerns (runtime size, boundary copying). Keep as explicit fallback, not the recommended default.

**F — Host-driven.** Native Go 1.27, no browser Go compile. Changes the product to “local/server process drives React” (LiveView-shaped). Natural for some desktop stories later; abandons in-browser Go for web v1. Prefer only if both C and E are unacceptable.

## Decisions

1. **Primary browser compile target: GoScript (non-Wasm Go→TypeScript).**  
   Rationale: only serious option that simultaneously (a) supports Go 1.27, (b) is not Wasm, (c) preserves object-graph / callback crossing closer to DESIGN’s GopherJS rationale than JSON/Wasm defaults. Prefer reuse of `github.com/s4wave/goscript` over hand-rolling a transpiler or waiting on GopherJS.

2. **Interop API: migrate the thin JS seam from `gopherjs/js` to `syscall/js` (build tags appropriate to GoScript’s `GOOS=js` package load).**  
   Rationale: GoScript’s documented host API; also the contingency Wasm path—one seam rewrite serves both C and E. Host contract stays `globalThis.formaHost.setTree`. Callback binding via reflect/`MakeFunc` remains Go-side.

3. **Reject dual-toolchain GopherJS and GopherJS-fork-as-plan** for unblocking P3d.  
   Rationale: A does not clear the showstopper; B is not schedule-viable. Optional later: watch GopherJS 1.22+; do not block product on it.

4. **Official Go Wasm is contingency, not default.**  
   Spike acceptance for GoScript (must all pass):
   - Compile reference `main.go` / minimal UIMount app under Go 1.27.
   - `setTree` receives a host-usable tree (object graph or equivalent) without JSON-encoding the full tree as the required hot path.
   - Click / string callbacks round-trip into Go handlers with `ui.Context`.
   - Artifact loadable from the Vite scaffold (or documented equivalent) with a rebuild story P3d can wrap.
   If spike fails within the task budget, **pivot to E** and update DESIGN accordingly—do not reopen A as durable.

5. **Host-driven JSON/WS is not v1 web default.** Record as deferred architectural escape if browser-Go compile is abandoned entirely later.

6. **Docs / ROADMAP:** Amend DESIGN stack table; note this change as blocking full P3d verification; do not close #3 in this change.

7. **Reuse vs hand-roll:** Prefer GoScript (third-party) over inventing a compiler; prefer adapting the existing P3c mount/bind design over a new runtime architecture; keep React host unchanged aside from any load-path tweaks for the new artifact name/format.

## Risks / Trade-offs

- [GoScript immaturity / breaking releases] → Pin a known-good version in docs/CI; spike before deleting GopherJS paths; Wasm contingency.
- [Reflect / `MakeFunc` / async function semantics differ from GopherJS] → Spike must exercise bindTree callbacks; adjust wrapper if GoScript makes blocking handlers async.
- [Build-tag / GOOS confusion (`js` vs `js,wasm`)] → Document exact `goscript compile` flags and tags in design tasks; one package layout for browser vs native.
- [Authors already reading GopherJS install docs on P3d branch] → Single doc pass; **BREAKING** toolchain message explicit.
- [Payload size still large for big apps] → Better than gc Wasm per GoScript claims; measure on todo demo; defer fine-grained updates (ROADMAP follow-up).

## Migration Plan

1. Spike GoScript against current `ui/` + scaffold host (this change’s early tasks).
2. On success: rewrite `state_js.go` (and tags) to `syscall/js`; remove `gopherjs/js` imports; update DESIGN/ROADMAP/user docs; leave rebuild CLI to P3d or a thin follow-up if #3 lands after.
3. On spike failure: implement E (official Wasm) with the same `syscall/js` seam; document DESIGN deviation; still avoid returning to EOL GopherJS.
4. Rollback: early-stage — revert the change branch; GopherJS code remains on older commits until this merges.

## Open Questions

- Exact Vite integration shape for GoScript output modules (import map vs bundled `public/app.js`) — resolve during spike without changing capability requirements.
- Whether CI installs GoScript via `go install` or npm `goscript` package — resolve in tasks/apply; either meets author-guide clarity.
