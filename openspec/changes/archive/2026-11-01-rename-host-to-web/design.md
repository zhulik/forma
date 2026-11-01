# Design

## Context

See [proposal.md](./proposal.md). Today init writes `host:` / `path: host` in the embedded template; author docs describe a `host/` directory. No physical `host/` tree exists in user scaffolds yet, so renaming is documentation + config shape only.

## Goals / Non-Goals

**Goals:**

- One author-facing vocabulary for the React/Vite/shadcn directory: **`web/`** on disk, **`web.path`** in `forma.yaml`.
- Specs and `doc/user/` stay aligned with the init template.
- Preserve **“React host” / “JavaScript host”** in architecture and `ui-dsl` for the runtime renderer (registry, reconcile, events).

**Non-Goals:**

- Implementing sync/dev against the web tree (P3).
- Renaming `ui-dsl` requirements that use “host” for DOM or runtime behavior.
- Aliasing or accepting legacy `host:` keys in `forma.yaml`.

## Decisions

### 1. Config section name matches directory

**Choice:** `web:` with `path: web` (default).

**Alternatives:** Keep `host.path` pointing at `web/` — rejected; keeps confusing yaml vocabulary.

### 2. Vocabulary split

| Term | Meaning |
|------|---------|
| `web/`, `web.path` | Filesystem: npm, Vite, shadcn TSX, mount HTML |
| “host”, “JavaScript host” | Runtime: React registry + tree renderer + bridge |

Author docs should say “web tree” or “under `web/`” for paths; “host” only when describing render/runtime behavior (or link to DESIGN.md).

### 3. Search-and-replace scope

**Choice:** Update path/config references in `doc/user/`, `doc/ROADMAP.md` layout lines, `doc/DESIGN.md` scaffold mentions, init template, and main specs on archive. Leave stub sync log messages optional (“web tree” preferred when touched).

**Alternatives:** Global replace of “host” — rejected; would corrupt `ui-dsl` and architecture docs.

## Risks / Trade-offs

- **[Risk] Authors confuse `web/` with framework `ui/`** → Mitigation: project-layout table already separates `internal/ui` (generated Go) from `web/` (TSX); keep that table prominent.
- **[Risk] Spec uses both “web tree” and “JavaScript host”** → Mitigation: intentional; design table above is the rule of thumb.

## Migration Plan

N/A. No released apps with `host.path`; breaking yaml change needs no shim per early-stage policy.

## Open Questions

(none)
