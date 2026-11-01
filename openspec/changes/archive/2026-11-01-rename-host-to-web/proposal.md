# Proposal

## Why

Author docs and `forma.yaml` use **`host/`** and **`host.path`** for the Vite + React + shadcn tree. Go-first authors read that as deployment or runtime jargon, and it collides mentally with DSL terms like “layout host.” Renaming to **`web/`** and **`web.path`** makes the on-disk layout obvious before P3 materializes the tree.

## What Changes

- **BREAKING:** Replace `forma.yaml` section **`host:`** / **`host.path`** with **`web:`** / **`web.path`** (default path **`web`**).
- Update init scaffold template, author guide, and ROADMAP/DESIGN references to the **`web/`** directory name.
- Align **`forma-cli`** and **`ui-docs`** specs with the new config field and default layout.
- Keep architecture prose “React host” / “JavaScript host” where it means the **runtime renderer**, not a filesystem path.

**Out of scope:** P3 implementation of sync/dev against the web tree; `ui-dsl` runtime semantics; migration shims for hypothetical existing `host.path` projects (none shipped yet).

## Capabilities

### New Capabilities

(none)

### Modified Capabilities

- `forma-cli`: `forma.yaml` declares **`web.path`**; sync/dev/docs refer to web dependencies and registry under that path.
- `ui-docs`: project layout, `forma.yaml`, CLI, and manual setup document **`web/`** instead of **`host/`**.

## Impact

- `cmd/forma/templates/scaffolds/minimal/forma.yaml.tmpl`
- `doc/user/*.md` (layout, yaml, cli, adding-components, advanced, manual-setup, README as needed)
- `doc/ROADMAP.md`, `doc/DESIGN.md` (layout mentions only)
- `openspec/specs/forma-cli/spec.md`, `openspec/specs/ui-docs/spec.md` (via delta sync on archive)
- Comments in stub sync CLI (optional wording only)
