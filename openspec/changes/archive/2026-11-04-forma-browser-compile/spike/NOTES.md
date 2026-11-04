# GoScript spike notes (forma-browser-compile)

**Date:** 2026-11-04  
**GoScript:** `github.com/s4wave/goscript/cmd/goscript@v0.6.3` (`go install …@latest`)  
**Go:** 1.27.3  
**Runner:** Bun 1.3.5

## Commands

```bash
# Minimal setTree
cd openspec/changes/forma-browser-compile/spike/minimal
goscript compile --package . --output ./output
bun run run.mjs

# MakeFunc + string (isolated)
cd ../makefunc && goscript compile --package . --output ./output && bun run run.mjs

# forma ui mount + bind + callbacks
cd ../../../../spikeforma   # worktree: forma-browser-compile/spikeforma
goscript compile --package ./spikeforma --output ./spikeforma/output
bun run run.mjs
bunx vite build   # needs vite.config target: esnext → dist/app.js (~486KB / ~114KB gzip)
```

Vite alias: `"@goscript" → output/@goscript`. Entry imports `main` and calls `boot()` (avoid relying on GoScript `isMainScript` top-level await alone under older browser targets).

## Acceptance (design.md)

| Criterion | Result |
|-----------|--------|
| Compile under Go 1.27 | **Pass** — `ui` + `spikeforma` including `bind.go` (`TypeFor`, `TypeAssert`, `MakeFunc`, `FuncOf`) |
| `setTree` host-usable object graph (not JSON-mandated) | **Pass with seam** — raw `*Node` arrives as GoScript `VariableRef`; seam must convert to maps/slices/`js.Func` (see `ui/state_js.go` spike) |
| Click + string callbacks + `ui.Context` | **Pass** — OnClick refresh; OnChange string round-trips into author handler / `Update` |
| Vite-loadable artifact | **Pass** — ES module graph via `@goscript/*`; Vite lib build `dist/app.js` with `build.target: "esnext"` |

## Gaps vs GopherJS (must plan for)

1. **Not a drop-in:** GopherJS exposed Go structs as JS objects; GoScript needs an explicit host conversion (`nodeToHost` / `propsToHost` / `wrapJSFunc`).
2. **Callbacks:** `reflect.MakeFunc` values are not automatically host-callable with JS string/bool args — wrap with `syscall/js.FuncOf` at the seam (verified; isolated MakeFunc+FuncOf also passes).
3. **Async:** many generated functions are `async`; host must `await` promise returns from handlers / `UIMount`.
4. **Payload:** spikeforma Vite bundle ~486KB JS (~114KB gzip) — much smaller than typical gc Wasm, larger than TinyGo Wasm toys.
5. **Tooling:** prefers Bun for quick runs; path alias `@goscript/*` required; `package main` auto-await needs esnext or explicit `boot()`.

## Verdict

**GoScript can replace GopherJS for forma’s browser path**, provided we keep a thin `syscall/js` seam that (a) materializes a plain object tree for React and (b) wraps bound handlers as `js.Func`. Spike gate **passed** — proceed to sections 2–4 of `tasks.md` (not Wasm contingency).

Do **not** treat as “zero bridge work”: the conversion layer is required product code, not a spike hack.
