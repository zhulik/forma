# Proposal

## Why

P1 locked the UI DSL; authors still have only `doc/DESIGN.md`, `doc/ROADMAP.md`, and `main.go`—no install path or day-to-day workflow written for Go developers. ROADMAP **P2** requires author-facing guides before runtime work so CLI, host, and codegen implement what users (and agents) already expect. User documentation is a **first-class product surface**: if it changes author experience, the guide must change in the same delivery slice.

## What Changes

- Add **`doc/user/`** plain Markdown author guide (install, CLI contract, project layout, writing UI, components, advanced, troubleshooting).
- Add OpenSpec capabilities **`ui-docs`** (documentation completeness and docs-first policy) and **`forma-cli`** (normative CLI + `forma.yaml` contract; **implementation deferred** to a later change).
- Update **`README.md`**, **`doc/ROADMAP.md`**, and **`doc/DESIGN.md`** where they still describe clone-template onboarding instead of **`forma init` / `forma add` / `forma sync`**.
- Add contributor-facing policy in **`openspec/config.yaml`** so future changes treat user docs as mandatory when UX changes.

**Out of scope (this change):** `cmd/forma` binary, React host, GopherJS runtime, codegen pipeline, hosted doc site (GitBook/MkDocs).

## Capabilities

### New Capabilities

- `ui-docs`: Author guide location, required topics, honesty about implementation status, and docs-first obligation for author-visible changes.
- `forma-cli`: Target commands (`init`, `add`, `sync`, later `dev`) and `forma.yaml` desired-state model (no binary in this change).

### Modified Capabilities

- (none)

## Impact

- **ROADMAP phase:** P2 (user documentation).
- **Docs:** `doc/user/**`, README, ROADMAP, DESIGN; `openspec/specs/ui-docs/`, `openspec/specs/forma-cli/` after archive.
- **Code:** none (documentation and specs only).
