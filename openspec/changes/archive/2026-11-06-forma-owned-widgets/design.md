# Design — forma-owned-widgets

## Context

See `proposal.md` (Why / What). Main today still runs the P3a path: shadcn starters (`button`/`input`/`card`), Tailwind scaffold, TSX→Go codegen into app `internal/ui`, and builtins only for layout/Text/Heading/Card. ROADMAP P5 + DESIGN tip lock **React thin host**, **hand-written Go widgets**, **[0build](https://0build.dev)** semantic `.z-*` (utilities escape hatch). #5/#6 superseded. This change is the **architecture migration** so a todo-like UI is expressible; shipping `examples/todo` stays **`forma-todo-example`**.

Observed product surface to migrate:

| Area | Today | Target |
|------|-------|--------|
| Widget SoT | shadcn TSX + `*.gen.go` | Hand Go in framework `ui/` |
| Host leaves | Generated registry + partial builtins | Builtins for full todo surface |
| CSS | Tailwind + shadcn tokens | 0build CDN CSS + runtime utils |
| CLI | `add` / `generate` / sync codegen | Init + sync (npm + tidy); no kit codegen |
| Overrides | Spec’d per-source YAML (unimplemented) | Never ship |

## Goals / Non-Goals

**Goals**

- Express `main.go`-shaped todo UI on owned widgets + 0build without kit codegen.
- Make init/sync/docs match that path (**BREAKING**, early-stage OK).
- Clear retire vs keep so Builder can delete without reopen.

**Non-Goals (design-level)**

- Full catalog, Dialog/Tabs, desktop (P4), Mantine codegen.
- Shipping `examples/todo` in this change.
- Publishing a separate npm host package (keep scaffold-embedded host).
- Perfect visual parity with shadcn New York theme.

## Decisions

Copy-ready for PR body:

1. **Slice vs example** — Keep **`forma-todo-example` separate**. This change only migrates architecture + todo-sized surface so the example can follow. Rationale: smaller reviewable PR; example proves end-to-end after stack stabilizes.

2. **Locked stack (do not reopen)** — React thin host; 0build semantic `.z-*` first; hand-written Go widgets; **no** shadcn/Mantine codegen or forma-shipped per-widget YAML as SoT.

3. **Go package layout** — Hand-write todo-sized widgets in framework **`ui/`** as ordinary `*.go` (one type per file). **Rename off `*.gen.go`**. Apps import `github.com/zhulik/forma/ui` for Button/Input/… — **no** default app `internal/ui` codegen package.

4. **Todo-sized surface (in scope)** — Button, Input, Checkbox, Card, Text, Heading, Row, Column, Container, Fragment. Checkbox is in (used by `main.go`; was missing from shadcn starters). No Dialog/Select/etc. here.

5. **Host components** — Extend scaffold **`builtinRegistry`** (prefer `web/src/host/widgets/*` + re-export) so interactive + layout types all resolve without `registry.gen.ts`. Keep **custom registry last-wins** merge for third-party islands. Host stays dumb: map props → elements + classes; sanitize events; keep hybrid text-input policy.

6. **0build integration (reuse)** — Prefer **CDN** stylesheet + **zRuntime** (and related scripts per 0build install docs) in scaffold `index.html` / entry — **do not** hand-roll a parallel CSS kit. Pin documented CDN URLs/versions in scaffold comments. Load utility runtime so `Class` escape hatches and layout spacing can use 0build utilities when needed. **Do not** adopt 0build/shadcn React kit components as SoT.

7. **Class mapping** — Small explicit maps in the host (e.g. Button `Variant` `default`/`primary` → `.z-button` + `.z-button-primary`; Card → `.z-card`). Go keeps a **small string/enum-like prop surface** close to today’s `main.go` (`Variant`, `Size`, `Placeholder`, `Value`, `OnChange`, `Checked`, …). Refine names only if 0build class vocabulary forces it; document the table in `doc/user/`.

8. **Retire Tailwind-as-kit** — Remove scaffold dependence on Tailwind PostCSS pipeline, `components.json`, `tailwind.config.js`, `@tailwind` entry CSS, and `tailwind-merge`/`cva` as required kit deps. Replace `cn`/`twMerge` with a trivial class join. Layout builtins today use Tailwind (`flex`, `gap-N`); **reimplement with 0build semantic/utility patterns** (e.g. flex + `gap`/`p` utilities with CSS variables or shorthand)—not a second utility framework.

9. **Retire shadcn product path** — Delete init starters materialization, `ShadcnWrapper` from sync/add, and author-facing **`forma add`** / **`forma generate`**. Remove scaffold `generator.go` and check-generate task. Drop `codegen:` from default `forma.yaml`. **Delete** `cmd/forma/internal/codegen` and shadcn CLI wiring from the product binary (early-stage: no dual path). Parked Mantine/kit codegen remains a ROADMAP note only—not code kept “just in case.”

10. **Registry.gen.ts** — Stop generating/committing kit registry. Scaffold `main.tsx` merges **builtins + `registry.custom.ts` only**. Custom file stays for hand wrappers.

11. **DSL slots / compounds** — Keep Fragment, `Node.Slots`, field-only slot rules for future owned compounds. Remove product requirements that catalog compounds are **generated from shadcn**. No compound widgets in this change.

12. **Reuse vs hand-roll** — **Reuse** 0build CSS/runtime, existing React host (render/sanitize/bridge/hybrid-input), `fn/`, CLI init/sync/dev structure. **Hand-write** Go widgets + thin host wrappers + class maps. **Do not** replace `fn/` or invent a new CSS system.

13. **Docs / DESIGN / config context** — Same change updates `doc/user/*`, `doc/DESIGN.md` (remove tip-path contradictions that still describe shadcn codegen as current), and `openspec/config.yaml` context/rules so agents stop prescribing shadcn SoT. ROADMAP row already points here; Builder only adjusts status/PR when shipping.

14. **Root `main.go`** — Keep as north-star; swap Tailwind-ish `Class` strings for 0build-friendly escape hatches where needed. Do **not** move the demo into `examples/todo` in this change.

15. **Breaking changes accepted** — No migration shims, dual-read, or feature flags (`early-stage-no-compat`).

### Alternatives considered

| Topic | Rejected | Why |
|-------|----------|-----|
| Keep codegen for “escape” widgets | Dual SoT, agent confusion | Tip abandoned kit codegen |
| Ship example in same change | Larger blast radius | Prefer `forma-todo-example` |
| Embed 0build React kits | Violates locked direction | CSS (+ runtime utils) only |
| Per-widget YAML matrices | Locked “never” | Hand Go/host PRs instead |
| Keep Tailwind beside 0build | Two utility systems | Retire Tailwind-as-kit |
| npm-publish host package | Extra product surface | Scaffold embed is enough for P5 |

## Risks / Trade-offs

- **[Risk] 0build CDN/API churn** → Pin versions; document upgrade in one scaffold place; utilities via runtime may need `z-not-ready` / LSH snippets from 0build install—copy from upstream docs, don’t invent.
- **[Risk] Visual gap vs former shadcn demo** → Accept for migration; tune class maps; example change can polish.
- **[Risk] Deleting codegen breaks in-flight mental models** → Docs + CLI help text must be explicit; supersede #5/#6 already signaled.
- **[Risk] Layout spacing semantics differ (Tailwind scale vs 0build vars)** → Map `Gap`/`Padding` ints through a single host helper; adjust `main.go` only if demos break badly.
- **[Trade-off] Smaller prop surface than shadcn** → Intentional for agents; escape hatch `Class`/`Attrs` remains.

## Migration Plan

1. Specs/docs/design (this change’s planning) → Builder applies code.
2. Implement owned Go + host + 0build scaffold; delete shadcn/codegen product path.
3. Update CI init smoke; `task test` / `task lint`.
4. Human review of planning/PR; **do not** start `forma-todo-example` until this merges (or is explicitly sequenced after).
5. No rollback shim—revert the PR if needed.

## Open Questions

None that block specs/tasks. Deferrable polish (exact CDN pin string, final Variant string set) is Builder-time within the mapping decision above.
