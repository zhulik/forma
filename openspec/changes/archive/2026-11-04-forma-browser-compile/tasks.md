# Tasks

## 1. Spike GoScript against forma (gate)

- [x] 1.1 Install a pinned GoScript version and compile a minimal package that calls `syscall/js` to invoke `globalThis.formaHost.setTree` with a small object graph; verify the host (or a Node/Bun harness) receives a non-JSON-mandated value. Record version pin and command line in the spike notes under this change or a short `doc/` note referenced from DESIGN.
- [x] 1.2 Compile forma `ui` browser sources (or a trimmed mount+bind harness) with GoScript under Go 1.27; exercise one click-style and one string-value callback through bound handlers into Go; fail the spike if callbacks or tree delivery cannot meet design.md acceptance criteria.
- [x] 1.3 Prove a Vite-loadable (or bundler-documented) artifact path for the compiled output that P3d can wrap later; document import/bundle choice chosen during the spike. If 1.1–1.3 fail, stop and execute section 5 (Wasm contingency) instead of sections 2–3 GoScript-specific steps.

Spike evidence: `openspec/changes/forma-browser-compile/spike/NOTES.md` — **gate passed** (GoScript v0.6.3).

## 2. Bridge migration (on GoScript success)

- [x] 2.1 Rewrite `ui` browser seam from `github.com/gopherjs/gopherjs/js` to `syscall/js` (adjust build tags for GoScript’s package load); keep `formaHost.setTree` contract; verify `go test ./ui/...` on native and a GoScript (or documented) browser check for mount/refresh/callbacks.
- [x] 2.2 Remove GopherJS module/tooling references from go.mod / scaffolds / scripts that would pull GopherJS for the supported path; verify `go list` / build no longer requires `gopherjs/js` for the default browser build.
- [x] 2.3 Update scaffold host load path if the artifact name or module shape changed (e.g. still `public/app.js` vs ES module); verify a fresh `forma init` scaffold loads the Go bundle after manual compile per docs.

## 3. Docs and control plane

- [x] 3.1 Amend `doc/DESIGN.md` stack table: browser compile = GoScript (v1); GopherJS retired; official Wasm listed as contingency only; verify DESIGN no longer claims “GopherJS (v1)” as the sole path.
- [x] 3.2 Amend `doc/ROADMAP.md`: note `forma-browser-compile` as the toolchain control-plane change blocking full P3d smoke; keep #3 open; verify the forward-queue row exists or is referenced.
- [x] 3.3 Update `doc/user/` install, CLI, writing-ui, troubleshooting: supported browser compiler, no GopherJS/Go 1.21 durable path; verify wording matches ui-docs delta scenarios.

## 4. Spec sync hygiene

- [x] 4.1 Ensure delta specs under this change stay aligned with landed code/docs (ui-runtime object-graph requirement; forma-cli compiler-agnostic dev contract; ui-docs install story); verify `openspec validate --change forma-browser-compile` (or project-equivalent) passes.

## 5. Contingency — official Wasm (only if spike fails)

- [x] 5.1 Skipped: spike gate passed (GoScript v0.6.3); Wasm contingency not executed.
- [x] 5.2 Skipped: spike gate passed; docs record GoScript as chosen path (not Wasm).

## Workflow follow-up

- Human review of decisions before `/opsx-apply` (propose-only until accepted).
- Do not close GitHub #3 (`forma-dev-experience`) in this change; P3d rebases onto the new compiler after merge.
- Archive this change after implementation; sync deltas into main specs.
