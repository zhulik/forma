# forma roadmap

forma is a **Go-native UI framework**: authors write composable components in Go; a fixed **React + shadcn/ui** host renders them. The roadmap prioritized **DSL quality and author docs before runtime implementation**, so OpenSpec artifacts and code stay aligned with how people (and agents) actually use the framework.

**Current status:** **P0–P3d** complete (P3b host + P3c GopherJS bridge + P3d `forma-dev-experience`). Nested **`cmd/forma`** ships **`init`**, **`sync`**, **`generate`**, **`add`**, **`version`**, and **`dev`**. Canonical specs include [`ui-host`](../openspec/specs/ui-host/spec.md) and [`ui-runtime`](../openspec/specs/ui-runtime/spec.md). Author guide: [`doc/user/`](user/README.md). [`main.go`](../main.go) is the reference todo demo. **Next:** P4 desktop (deferred) or P5 breadth.

**Canonical design:** [DESIGN.md](./DESIGN.md). **Agent factory (operators):** [AGENT_FACTORY.md](./AGENT_FACTORY.md).

---

## Goals

- **Best-in-class Go DSL** for UI: dense, extensible, expressive—minimal boilerplate vs raw HTML or div wrappers.
- **Zero required HTML/CSS/JS** for app authors; optional escape hatches only.
- **Deterministic shadcn → Go codegen** so the widget surface stays maintainable.
- **Third-party React** available via registry + thin JS wrappers, not forked frameworks.
- **Agent-friendly:** review Go structs and handlers, not generated markup.

## Non-goals (initial releases)

- SSR/SSG as a supported product path.
- Wasm-first UI bridge (GopherJS is v1).
- Electron desktop; full native webview story comes after web dev loop works.
- Complete shadcn catalog before core loop is stable.
- Competing state managers in React (Redux, etc.) for domain state.

---

## Phased delivery

OpenSpec changes should name the ROADMAP phase they belong to. Early phases (DSL → docs → codegen) are complete; remaining work is host, runtime, and the browser dev loop.

| Phase | Scope | OpenSpec / deliverables | Status |
|-------|--------|-------------------------|--------|
| **P0** | **Dev tooling** — golangci-lint, Taskfile, GitHub CI (lint/generate/test), Cursor rules; `task test` wired | **`forma-dev-tooling`** → `openspec/specs/dev-tooling/` | **Done** |
| **P1** | **UI DSL v1** — semantics, composition, events, keys, inputs, escape hatches, **Fragment**, field-only compound **slots**; `main.go` reference; stub `ui` runtime | **`forma-ui-dsl`**, **`forma-fragment-slots`** → `openspec/specs/ui-dsl/` | **Done** |
| **P2** | **User documentation** — CLI-first install (`forma init` / `add` / `sync`), `forma.yaml`, project layout, writing UI, components, advanced, troubleshooting | **`forma-user-docs`** → `doc/user/`, OpenSpec **`ui-docs`** + **`forma-cli`** | **Done** |
| **P3a** | **Codegen pipeline** — pure-Go TSX parse → app `*.gen.go` + registry; `forma generate`; scaffold `check-generate` | **`forma-ui-codegen`** → `openspec/specs/ui-codegen/` | **Done** |
| **P3b** | **React host** — registry, recursive renderer, event sanitization, text-input policy | **`forma-react-host`** → `openspec/specs/ui-host/` | **Done** |
| **P3c** | **GopherJS runtime** — real `UIMount`/`RefreshUI`, tree crossing, callback dispatch, `ui.Context` propagation | **`forma-gopherjs-runtime`** → `openspec/specs/ui-runtime/` | **Done** |
| **P3d** | **End-to-end dev loop** — one command to run host + rebuild Go; todo demo works in browser | **`forma-dev-experience`** | **Done** |
| **P4** | **Desktop packaging** — native webview wrapper, same bundle | **`forma-desktop`** | Deferred |
| **P5** | **Breadth** — more shadcn widgets, 3rd-party recipes (maps, payments), performance notes | Multiple small changes | Deferred |

### Phase dependencies

```
P1 DSL ──► P2 User docs ──► P3a Codegen ──┐
                           P3b Host ─────┼──► P3d Dev loop
                           P3c Runtime ──┘
                                         └──► P4 Desktop ──► P5 Breadth
```

P0–P3a are done. Schedule **P3b** and **P3c** before **P3d**; ask before parallelizing if overlap is unclear. **P4** / **P5** stay deferred until the browser loop works.

---

## OpenSpec workflow

- **Schema:** `spec-driven` (see [`openspec/config.yaml`](../openspec/config.yaml)).
- **Main specs home:** `openspec/specs/` — **`dev-tooling`**, **`agent-factory`**, **`ui-dsl`**, **`ui-docs`**, **`forma-cli`**, **`ui-codegen`**; more (e.g. `ui-runtime`) land when P3+ changes archive.
- **Active work:** `openspec/changes/<change-name>/` with `proposal.md`, `specs/.../spec.md`, `design.md`, `tasks.md`.
- **Archive:** completed changes move to `openspec/changes/archive/`; sync capabilities to `openspec/specs/`.

### Agent factory (local)

Operator overview: [`AGENT_FACTORY.md`](./AGENT_FACTORY.md). **Who:** Cursor agents in `.cursor/agents/` — `planner`, `orch`, `architector`, `builder`, `debt`. **What:** OpenSpec skills/commands (`/opsx-*`).

This roadmap is the factory **control plane**: **orch** (and peers) may update live fields on suggested-change rows as work advances—at least **Status**, and when known **branch** / **PR** (link or number). Humans alone merge. Factory runs locally (worktrees); cloud Automations are not the scheduler.

| Live field | Meaning |
|------------|---------|
| Status | Planned / in progress / PR open / Done (etc.) |
| branch | Feature branch when opened |
| PR | Pull request identity when opened |

### Completed through P3a

Shipped and archived: **P0** `forma-dev-tooling`; **P1** `forma-ui-dsl`, `forma-fragment-slots`; **P2** `forma-user-docs`, `forma-cli-scaffold`; web tree via `forma-web-scaffold`; **P3a** `forma-ui-codegen`. Meta: **`forma-agent-factory`** (local roles + this control-plane convention). Debt/refactor archives are intentionally omitted from the queue below.

### Suggested changes (forward queue)

Orch schedules from this table only. Leave **branch** / **PR** empty until a change is in flight.

| Order | Change name | Phase | Purpose | Status | branch | PR |
|-------|-------------|-------|---------|--------|--------|-----|
| 1 | `forma-react-host` | P3b | Static host app + ComponentRegistry + recursive renderer | Done | | |
| 2 | `forma-gopherjs-runtime` | P3c | Wire Go `UIMount`/`RefreshUI` to host; callbacks + `ui.Context` | Done | | |
| 3 | `forma-dev-experience` | P3d | `forma dev` (replace stub); demo runs in browser | Done | `change/forma-dev-experience` | [#3](https://github.com/zhulik/forma/pull/3) |
| 4 | `forma-desktop` | P4 | Native webview packaging, same bundle | Deferred | | |
| 5 | — | P5 | Breadth: more widgets, 3rd-party recipes, perf notes | Deferred | | |

Proposals must reference **ROADMAP phase** (P3b–P5) in the Impact section.

---

## Phase notes

### P1 — DSL (complete)

Authors rely on `openspec/specs/ui-dsl/` + [`main.go`](../main.go): `UIMount`/`RefreshUI`, dot-import `ui`/`fn`, `.With`/`.Or`, events via `ui.Context`, keys/inputs policy, escape hatches. Browser runtime is in `openspec/specs/ui-runtime/` (P3c).

### P2 — User docs (complete)

Guide under [`doc/user/`](user/README.md) for Go authors who may never open the React host. Docs track shipped CLI (`init`/`add`/`sync`/`generate`/`dev`), host, and bridge. Spec contradictions → update specs first.

### P3 — Implementation

- **P3a (done):** codegen owns prop parity with TSX; overrides explicit.
- **P3b–P3c (done):** host stays thin; runtime implements P1 contracts via `ui-runtime`.
- **P3d done when:** todo demo runs in a browser with HMR or fast rebuild documented in the user guide (`forma dev` no longer stub).

---

## Follow-ups (deferred)

| Item | Notes |
|------|--------|
| SSR / hydration | See DESIGN.md; only if product requires SEO |
| Signal-based fine-grained updates | Alternative to full-tree RefreshUI; breaks simple mental model |
| Desktop webview | P4 |
| Deeper agent prompts/skills | Optional; factory v1 ships roles + OpenSpec skills |

---

## Success metrics (informal)

- New widget: **registry + codegen + Go usage** in minutes, not days.
- Agent-generated UI diff is **Go-only** and compiles.
- Authors never need to read React source for standard shadcn widgets.
