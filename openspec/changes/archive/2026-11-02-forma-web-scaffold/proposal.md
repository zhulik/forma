# Proposal

## Why

`forma init` today scaffolds only Go sources and a stub `forma sync` that tidies the module. Author docs and `forma-cli` already describe a `web/` tree (Vite, React, shadcn) materialized by sync, but nothing creates it—authors and CI cannot exercise the documented layout. P3 host and Go codegen still depend on committed shadcn TSX under `web.path`; we should land the web scaffold and real JS-side materialization now, with codegen and registry explicitly stubbed.

## What Changes

- Embed a **minimal Vite + React + Tailwind + shadcn-ready** tree under the **`minimal`** init template at `web/` (`package.json`, `components.json`, `index.html` with mount id, toolchain configs, thin host entry).
- Default **`ui.components`** in init template to **button, input, card** (aligned with author guide).
- Extend **`formaconfig`** parsing for `ui`, `web`, and `codegen` fields used by sync.
- **`forma sync`**: run **`npm install`** in `web.path`, run **`shadcn add`** for each name in `ui.components` (non-interactive), **stub** Go codegen and JS registry update (logged, no failure), then existing Go module reconciliation.
- **`forma init`** unchanged in flow (still ends with sync); failures during npm/shadcn fail init.
- **CI / Task**: init smoke requires **Node.js** so full init can complete sync.
- **`doc/user/`**: install, CLI, sync status—web + shadcn materialization available; codegen/registry/dev still planned.

## Capabilities

### New Capabilities

_(none)_

### Modified Capabilities

- **`forma-cli`**: Init embeds web tree; sync runs npm install and shadcn materialization; partial sync (codegen/registry stub); default component list in scaffold.
- **`dev-tooling`**: Init scaffold smoke assumes Node for sync npm/shadcn steps.
- **`ui-docs`**: Implementation status for sync (web/shadcn vs codegen/dev).

## Impact

- **ROADMAP**: Advances P3 web-tree slice; P3a codegen and P3d dev loop remain separate.
- **Author prerequisites**: Node/npm required for init and sync (already documented for future P3).
- **Dependencies**: Scaffold pins npm packages + committed lockfile; sync hits npm and shadcn registries on init/sync.
- **CLI internals**: `WebSyncRunner` in `core/interfaces.go`, `webtool` adapter, `fn` helpers for component dedupe.
- **Out of scope**: Go widget codegen, host ComponentRegistry wiring, `forma add` implementation, `forma dev`, browser runtime.
