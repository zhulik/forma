# forma roadmap

forma is a **Go-native UI framework**: authors write composable components in Go; a fixed **React** host reconciles them. Visual look comes from **[0build](https://0build.dev)** (semantic CSS classes; successor to Franken UI); **widget APIs are hand-written Go** in forma—not generated from shadcn TSX, 0build React kits, or third-party `.d.ts`. Early phases prioritized DSL quality, docs, and the browser runtime loop; OpenSpec and code stay aligned with how people (and agents) actually use the framework.

**Current status:** **P0–P3d** complete and archived (host + GoScript runtime + `forma dev` on [#3](https://github.com/zhulik/forma/pull/3)). Nested **`cmd/forma`** ships **`init`**, **`sync`**, **`version`**, and **`dev`** (no kit `add`/`generate` on the tip path). Canonical specs still include [`ui-host`](../openspec/specs/ui-host/spec.md) and [`ui-runtime`](../openspec/specs/ui-runtime/spec.md); author guide: [`doc/user/`](user/README.md). **Direction locked:** React thin host; **own Go widgets** + **0build** semantic classes (e.g. `.z-button-primary`, `.z-card`) for look; **no** per-widget YAML overrides or generated kit bindings as SoT. Mantine-from-`.d.ts` codegen is **parked**. Shadcn-codegen-as-primary is **retired** from the author path (historical through P3a). **Tip:** **P5** — `forma-owned-widgets` ([#8](https://github.com/zhulik/forma/pull/8)), then `forma-todo-example` (`examples/todo`). **P4** desktop stays deferred.

**Canonical design:** [DESIGN.md](./DESIGN.md). **Agent factory (operators):** [AGENT_FACTORY.md](./AGENT_FACTORY.md).

---

## Goals

- **Best-in-class Go DSL** for UI: dense, extensible, expressive—minimal boilerplate vs raw HTML or div wrappers.
- **Zero required HTML/CSS/JS** for app authors; optional escape hatches only.
- **Hand-maintained Go widget surface** (Button, Input, Checkbox, Card, …) with thin React host components—stable APIs authors and agents review in Go.
- **[0build](https://0build.dev)** for look: semantic CSS first; utilities / `Class` as escape hatch.
- **Third-party React** still available later via registry + thin JS wrappers when needed—not the default widget path.
- **Agent-friendly:** review Go structs and handlers, not generated markup or kit prop dumps.

## Non-goals (initial releases)

- SSR/SSG as a supported product path.
- Wasm-first UI bridge (GoScript is v1; official Wasm is contingency only).
- Electron desktop; full native webview story is **P4**.
- A Go DOM / non-React renderer.
- **shadcn → Go codegen** or **Mantine `.d.ts` → Go codegen** as the product source of truth (both parked / historical).
- Per-widget YAML override matrices as the way to ship a catalog.
- Full widget catalog before a **todo-like example** proves the owned-widget path.
- Competing state managers in React (Redux, etc.) for domain state.

---

## Phased delivery

OpenSpec changes should name the ROADMAP phase they belong to. Browser loop through P3d is done. Tip is **P5**: owned widgets + **0build**, then example. Desktop remains deferred.

| Phase | Scope | OpenSpec / deliverables | Status |
|-------|--------|-------------------------|--------|
| **P0** | **Dev tooling** — golangci-lint, Taskfile, GitHub CI (lint/generate/test), Cursor rules; `task test` wired | **`forma-dev-tooling`** → `openspec/specs/dev-tooling/` | **Done** |
| **P1** | **UI DSL v1** — semantics, composition, events, keys, inputs, escape hatches, **Fragment**, field-only compound **slots**; `main.go` reference; stub `ui` runtime | **`forma-ui-dsl`**, **`forma-fragment-slots`** → `openspec/specs/ui-dsl/` | **Done** |
| **P2** | **User documentation** — CLI-first install (`forma init` / `add` / `sync`), `forma.yaml`, project layout, writing UI, components, advanced, troubleshooting | **`forma-user-docs`** → `doc/user/`, OpenSpec **`ui-docs`** + **`forma-cli`** | **Done** |
| **P3a** | **Codegen pipeline** — pure-Go TSX parse → app `*.gen.go` + registry; `forma generate`; scaffold `check-generate` | **`forma-ui-codegen`** → `openspec/specs/ui-codegen/` | **Done** (historical; not tip SoT) |
| **P3b** | **React host** — registry, recursive renderer, event sanitization, text-input policy | **`forma-react-host`** → `openspec/specs/ui-host/` | **Done** |
| **P3c** | **Browser runtime** — real `UIMount`/`RefreshUI`, tree crossing, callback dispatch, `ui.Context` propagation | **`forma-gopherjs-runtime`** → `openspec/specs/ui-runtime/` (interop now GoScript/`syscall/js`) | **Done** |
| **P3c′** | **Browser compile control plane** — retire GopherJS; GoScript under Go 1.27; amend DESIGN | **`forma-browser-compile`** | **Done** |
| **P3d** | **End-to-end dev loop** — one command to run host + rebuild Go; todo demo works in browser | **`forma-dev-experience`** | **Done** |
| **P4** | **Desktop packaging** — native webview wrapper, same bundle | **`forma-desktop`** | Deferred |
| **P5** | **Owned widgets + example** — **0build** + hand-written Go widgets (todo surface); in-browser todo-like example. Later: more widgets, optional 3rd-party recipes, perf notes | **`forma-owned-widgets`**, **`forma-todo-example`** | **In progress** (`forma-owned-widgets` #8 + `forma-todo-example` PR open) |

### Phase dependencies

```
P1 DSL ──► P2 User docs ──► P3a Codegen ──┐
                           P3b Host ─────┼──► P3d Dev loop
                           P3c Runtime ──┘
                                         ├──► P5 Owned widgets + example ──► …further P5
                                         └──► P4 Desktop (deferred; same bundle)
```

P0–P3d are done. CSS kit is **0build** — **Orch** may schedule **P5** tip rows below. **P4** stays deferred. P5 does **not** wait on P4.

---

## OpenSpec workflow

- **Schema:** `spec-driven` (see [`openspec/config.yaml`](../openspec/config.yaml)).
- **Main specs home:** `openspec/specs/` — **`dev-tooling`**, **`agent-factory`**, **`ui-dsl`**, **`ui-docs`**, **`forma-cli`**, **`ui-codegen`**, **`ui-host`**, **`ui-runtime`** (capabilities evolve with owned widgets; architector amends deltas).
- **Active work:** `openspec/changes/<change-name>/` with `proposal.md`, `specs/.../spec.md`, `design.md`, `tasks.md`.
- **Archive:** completed changes move to `openspec/changes/archive/`; sync capabilities to `openspec/specs/`.

### Agent factory (local)

Operator overview: [`AGENT_FACTORY.md`](./AGENT_FACTORY.md). **Who:** Cursor agents in `.cursor/agents/` — `planner`, `orch`, `architector`, `builder`, `debt`. **What:** OpenSpec skills/commands (`/opsx-*`).

This roadmap is the factory **control plane**: **orch** (and peers) may update live fields on suggested-change rows as work advances—at least **Status**, and when known **branch** / **PR** (link or number). Humans alone merge. Factory runs locally (worktrees); cloud Automations are not the scheduler.

| Live field | Meaning |
|------------|---------|
| Status | Planned / in progress / PR open / Done / Superseded (etc.) |
| branch | Feature branch when opened |
| PR | Pull request identity when opened |

### Completed through P3d

Shipped and archived: **P0** `forma-dev-tooling`; **P1** `forma-ui-dsl`, `forma-fragment-slots`; **P2** `forma-user-docs`, `forma-cli-scaffold`; web tree via `forma-web-scaffold`; **P3a** `forma-ui-codegen`; **P3b** `forma-react-host`; **P3c** `forma-gopherjs-runtime`; **P3c′** `forma-browser-compile`; **P3d** `forma-dev-experience`. Meta: **`forma-agent-factory`**. Debt/refactor archives omitted from the queue below.

### Suggested changes (forward queue)

Orch schedules from this table only. Leave **branch** / **PR** empty until a change is in flight.

| Order | Change name | Phase | Purpose | Status | branch | PR |
|-------|-------------|-------|---------|--------|--------|-----|
| 1 | `forma-owned-widgets` | P5 | Hand-written Go widgets + thin React host primitives for a todo-sized surface; scaffold loads [0build](https://0build.dev) styles (+ runtime utils if required); map Go variants → semantic `.z-*` classes; drop generated kit bindings / per-widget YAML as SoT | PR open | `forma-owned-widgets` | [#8](https://github.com/zhulik/forma/pull/8) |
| 2 | `forma-todo-example` | P5 | In-browser todo-like example on owned widgets (`examples/todo`), runnable with `forma dev`; docs point here; stacks on owned-widgets/#8 | PR open | `forma-todo-example` | [#9](https://github.com/zhulik/forma/pull/9) |
| 3 | `forma-todo-shadcn-slice` | P5 | Prior tip: gap-fill default shadcn via codegen path | **Superseded** | `forma-todo-shadcn-slice` | [#5](https://github.com/zhulik/forma/pull/5) |
| 4 | `forma-todo-example` (shadcn stack) | P5 | Prior tip: example on shadcn slice | **Superseded** | `forma-todo-example` | [#6](https://github.com/zhulik/forma/pull/6) |
| 5 | `forma-desktop` | P4 | Native webview packaging, same bundle | Deferred | | |

Proposals must reference **ROADMAP phase** (P4–P5) in the Impact section. **Parked (not queued):** Mantine-from-`.d.ts` codegen; further shadcn catalog codegen. Further P5 (wider owned catalog, 3rd-party recipes, perf) after the todo example lands.

**Handoff:** [#5](https://github.com/zhulik/forma/pull/5) and [#6](https://github.com/zhulik/forma/pull/6) are **closed (not merged)** and marked **Superseded** by the owned-widgets direction. Historical branches may remain for reference only.

---

## Phase notes

### P1 — DSL (complete)

Authors rely on `openspec/specs/ui-dsl/` + the interactive demo under [`examples/todo`](../examples/todo/): `UIMount`/`RefreshUI`, dot-import `ui`/`fn`, `.With`/`.Or`, events via `ui.Context`, keys/inputs policy, escape hatches. Browser runtime is in `openspec/specs/ui-runtime/` (P3c).

### P2 — User docs (complete)

Guide under [`doc/user/`](user/README.md) for Go authors who may never open the React host. Docs track shipped CLI and browser loop; update when owned widgets replace generated kit bindings in author-facing flows.

### P3 — Implementation (complete)

- **P3a:** shipped shadcn TSX → Go codegen (historical). Tip does **not** extend this as the widget SoT.
- **P3b–P3c:** host stays thin; runtime implements P1 contracts via `ui-runtime` (GoScript/`syscall/js`). Host registry will map to **forma-owned** React components under P5.
- **P3c′:** GopherJS retired; browser builds use GoScript under Go 1.27.
- **P3d:** `forma dev` runs the Vite host, watches Go, and rebuilds the browser bundle.

### P5 — Owned widgets + example (planned)

**Product goal:** nice in-browser UI via the Go DSL; prove with a **todo-like example**.

**Stack:** React (reconciler only) + **[0build](https://0build.dev)** + **hand-written Go widgets** and matching thin host components. No Go DOM framework. No generated kit prop surfaces as SoT. 0build may market shadcn/coss *looks*; forma still does **not** adopt those as generated React kit SoT—only semantic CSS (and kit runtime utils if required).

1. **`forma-owned-widgets`** — define the todo-sized Go API (Button, Input, Checkbox, Card, layout, …); host renders native elements or minimal wrappers applying `.z-*` (e.g. `.z-button-primary`, `.z-card`); scaffold loads 0build CSS (CDN-friendly / layers OK; utilities as escape hatch). Retire author dependence on `*.gen.go` kit bindings / YAML overrides for that surface. Architector owns package layout and how much of P3a CLI/`generate` remains.
2. **`forma-todo-example`** — ship the example on owned widgets; prove CRUD + responsive layout via `forma dev`.

**Spike note (Mantine):** `.d.ts` flatten/filters, Button `onClick` gaps, and Card slot mismatch reinforced **own Go widgets** over kit codegen as tip.

---

## Follow-ups (deferred / parked)

| Item | Notes |
|------|--------|
| SSR / hydration | See DESIGN.md; only if product requires SEO |
| Signal-based fine-grained updates | Alternative to full-tree RefreshUI; breaks simple mental model |
| Desktop webview | P4 |
| Deeper agent prompts/skills | Optional; factory v1 ships roles + OpenSpec skills |
| Mantine / kit `.d.ts` → Go codegen | Parked; optional later if a kit binding track returns |
| Broader owned catalog / 3rd-party / perf | After todo example ships |

---

## Success metrics (informal)

- New core widget: **hand Go API + thin host component + 0build `.z-*` classes** in a reviewable PR—not a codegen dump.
- Agent-generated UI diff is **Go-only** and compiles.
- Authors never need to read React or CSS-framework internals for the default todo surface.
- **P5 tip:** todo-like example runs in the browser on owned widgets end-to-end.
