# Proposal

## Why

Shadcn TSX→Go codegen as the widget source of truth (#5/#6 tip) was abandoned: kit prop dumps and generated bindings are a poor review surface for humans and agents. ROADMAP **P5** locks React thin host + **hand-written Go widgets** + **[0build](https://0build.dev)** semantic CSS. Main still ships the P3a shadcn/Tailwind scaffold path; this change migrates architecture so a todo-like UI is expressible on the locked stack before a separate example ships.

## What Changes

- **Owned todo-sized widget surface** in framework `ui/` (hand-written Go: Button, Input, Checkbox, Card, Text, Heading, Row, Column, Container, Fragment) with matching thin React host primitives that apply 0build `.z-*` classes (utilities / `Class` as escape hatch only).
- **Scaffold styling** switches from Tailwind-as-kit + shadcn TSX to **0build** CSS (+ utility runtime as required); Vite/React host remains.
- **BREAKING:** Retire author dependence on shadcn starters-as-SoT, app-local `*.gen.go` kit bindings, `forma add` / `forma generate` / sync codegen for that surface, `forma.yaml` `codegen:` as the default widget path, and embedded per-widget YAML overrides (never ship).
- **BREAKING:** Remove or disconnect P3a shadcn CLI materialization from `forma init` / `sync`; registry for the todo surface is **framework builtins**, not `registry.gen.ts` from shadcn.
- Update **DESIGN.md**, **openspec/config.yaml** context, and **doc/user/** to the owned-widgets + 0build author model.
- Keep **`forma-todo-example`** as a **separate later queue row** (do not ship `examples/todo` here). Root `main.go` stays the DSL north-star and must remain expressible on the new surface.

## Capabilities

### New Capabilities

- `ui-widgets`: Hand-written Go widget APIs and host rendering contracts for the todo-sized owned surface (semantic props → 0build classes; no kit codegen / YAML override SoT).

### Modified Capabilities

- `ui-codegen`: Retire product requirements for shadcn TSX→Go as the widget SoT; mark the historical pipeline non-product (parked/removed from author path).
- `forma-cli`: Init/sync no longer materialize shadcn or regenerate kit bindings; remove or redefine `add`/`generate` for the owned-widgets product path; drop `codegen:` from default `forma.yaml`.
- `ui-host`: Builtin registry covers the todo-sized interactive + layout surface with 0build classes; generated registry is not required for core widgets.
- `ui-dsl`: Drop “catalog compounds are generated from shadcn” as product rules; compounds remain a future owned-widget concern; leaf/layout owned types are hand-maintained.
- `ui-docs`: Author guide documents owned widgets, 0build look, and the post-migration CLI/layout (no shadcn add/codegen as primary).
- `dev-tooling`: Full-init CI verifies owned-widgets scaffold (0build + builtins) without shadcn materialization/codegen.

## Impact

- **ROADMAP phase:** **P5** (`forma-owned-widgets`); unlocks later `forma-todo-example`. **P4** desktop unchanged/deferred.
- **Locked stack:** React thin host; 0build semantic `.z-*`; hand-written Go widgets; no shadcn/Mantine codegen or forma-shipped per-widget YAML as SoT.
- **Code:** `ui/` (rename off `.gen.go`), scaffold `web/` (host builtins, CSS, package deps), `cmd/forma` (starters, sync/init/add/generate, shadcn wrapper, codegen package product wiring), root `main.go` class strings, docs/specs/config context.
- **Out of scope:** Shipping `examples/todo`; full widget catalog; Dialog/Tabs/etc.; desktop (P4); Mantine `.d.ts` codegen; reopening React-host or 0build vs utilities direction.
