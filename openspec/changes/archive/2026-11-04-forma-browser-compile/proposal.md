# Proposal

## Why

GopherJS latest (`v1.21.0`) still targets Go **1.21** (EOL) while forma’s `go.mod` is **1.27.3**. That gap blocks a real browser smoke of P3c/P3d: GopherJS cannot compile the shipped bridge (it rejects Go 1.22+ language used in `ui/`, e.g. `reflect.TypeFor`, `range` over int, `reflect.TypeAssert`). DESIGN.md currently pins GopherJS as v1 and Wasm as “Not v1”; that assumption is a control-plane showstopper and must be amended before `forma-dev-experience` (#3) can fully ship.

## What Changes

- **BREAKING (toolchain):** Replace GopherJS as the browser compile target with a path that supports Go **1.27** without pinning authors or CI to EOL Go 1.21.
- Amend `doc/DESIGN.md` stack choice: browser compile is no longer “GopherJS forever”; record the chosen compiler and why Wasm remains non-default.
- Amend `doc/ROADMAP.md`: insert this control-plane change ahead of completing P3d smoke; keep #3 open until the toolchain decision lands and P3d can rebuild against it.
- Rewrite the thin `//go:build js` host seam (today `github.com/gopherjs/gopherjs/js`) to the chosen interop API (`syscall/js` or equivalent) while preserving `formaHost.setTree` + callback semantics.
- Update author install / CLI docs so prerequisites name the new compiler and remove “GopherJS + GOPHERJS_GOROOT Go 1.21” as the supported path.
- Time-box a spike; if the preferred non-Wasm path fails acceptance criteria, fall back per design (official `js/wasm` contingency)—do not leave P3d blocked indefinitely.

## Capabilities

### New Capabilities

_(none)_

### Modified Capabilities

- `ui-runtime`: Require the browser build to produce a host-callable artifact that updates the React tree via object/callback crossing without mandating JSON encoding of the full render tree on each refresh; keep native stubs; stay toolchain-name-agnostic in requirements.
- `ui-docs`: Update install/toolchain and bridge wording so authors are not guided to GopherJS / Go 1.21; document the chosen browser compiler and verification path.
- `forma-cli`: Stop hardcoding GopherJS in the `forma dev` / browser-loop contract; the loop rebuilds the app for the browser with the project’s supported compiler.

## Impact

- **ROADMAP phase:** Control-plane amendment affecting **P3c** (bridge interop) and **P3d** (`forma dev` rebuild). Does not close #3.
- **Code:** `ui/*_js.go` (and any build-tag layout), scaffold/docs that mention GopherJS, future `forma dev` rebuild command (P3d).
- **Dependencies:** Drop runtime/tool dependence on `github.com/gopherjs/gopherjs`; add the chosen compiler as a documented install prerequisite.
- **Out of scope:** Implementing the full P3d Vite loop itself (owned by `forma-dev-experience`); desktop packaging; SSR; forking GopherJS to 1.27 as the primary plan.
