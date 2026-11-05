# forma roadmap

forma is a **Go-native UI framework**: authors write composable components in Go; a fixed **React + shadcn/ui** host renders them. The roadmap prioritized **DSL quality and author docs before runtime implementation**, so OpenSpec artifacts and code stay aligned with how people (and agents) actually use the framework.

**Current status:** **P0–P3d** complete and archived (P3b host + P3c bridge + GoScript browser compile via **`forma-browser-compile`** + P3d `forma dev` on [#3](https://github.com/zhulik/forma/pull/3)). Nested **`cmd/forma`** ships **`init`**, **`sync`**, **`generate`**, **`add`**, **`version`**, and **`dev`**. Canonical specs include [`ui-host`](../openspec/specs/ui-host/spec.md) and [`ui-runtime`](../openspec/specs/ui-runtime/spec.md). Author guide: [`doc/user/`](user/README.md). [`main.go`](../main.go) remains the framework-root smoke todo. **Next:** **P5** — prove forma can run a **single in-browser todo list** as an example project with a **minimal, responsive UI on default shadcn** (gap-fill catalog → ship `examples/todo`). **P4** desktop stays deferred.

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
- Wasm-first UI bridge (GoScript is v1; official Wasm is contingency only).
- Electron desktop; full native webview story is **P4**.
- Full shadcn catalog or third-party recipes before the **responsive todo example** proves the browser path.
- Competing state managers in React (Redux, etc.) for domain state.

---

## Phased delivery

OpenSpec changes should name the ROADMAP phase they belong to. Early phases through the browser dev loop are complete; the tip is **P5** breadth aimed at the todo example. Desktop packaging remains deferred.

| Phase | Scope | OpenSpec / deliverables | Status |
|-------|--------|-------------------------|--------|
| **P0** | **Dev tooling** — golangci-lint, Taskfile, GitHub CI (lint/generate/test), Cursor rules; `task test` wired | **`forma-dev-tooling`** → `openspec/specs/dev-tooling/` | **Done** |
| **P1** | **UI DSL v1** — semantics, composition, events, keys, inputs, escape hatches, **Fragment**, field-only compound **slots**; `main.go` reference; stub `ui` runtime | **`forma-ui-dsl`**, **`forma-fragment-slots`** → `openspec/specs/ui-dsl/` | **Done** |
| **P2** | **User documentation** — CLI-first install (`forma init` / `add` / `sync`), `forma.yaml`, project layout, writing UI, components, advanced, troubleshooting | **`forma-user-docs`** → `doc/user/`, OpenSpec **`ui-docs`** + **`forma-cli`** | **Done** |
| **P3a** | **Codegen pipeline** — pure-Go TSX parse → app `*.gen.go` + registry; `forma generate`; scaffold `check-generate` | **`forma-ui-codegen`** → `openspec/specs/ui-codegen/` | **Done** |
| **P3b** | **React host** — registry, recursive renderer, event sanitization, text-input policy | **`forma-react-host`** → `openspec/specs/ui-host/` | **Done** |
| **P3c** | **Browser runtime** — real `UIMount`/`RefreshUI`, tree crossing, callback dispatch, `ui.Context` propagation | **`forma-gopherjs-runtime`** → `openspec/specs/ui-runtime/` (interop now GoScript/`syscall/js`) | **Done** |
| **P3c′** | **Browser compile control plane** — retire GopherJS; GoScript under Go 1.27; amend DESIGN | **`forma-browser-compile`** | **Done** |
| **P3d** | **End-to-end dev loop** — one command to run host + rebuild Go; todo demo works in browser | **`forma-dev-experience`** | **Done** |
| **P4** | **Desktop packaging** — native webview wrapper, same bundle | **`forma-desktop`** | Deferred |
| **P5** | **Breadth (todo example tip)** — gap-fill default shadcn + host/codegen as needed; ship responsive in-browser `examples/todo`. Later: more widgets, 3rd-party recipes, perf notes | **`forma-todo-shadcn-slice`**, **`forma-todo-example`** | **Planned** |

### Phase dependencies

```
P1 DSL ──► P2 User docs ──► P3a Codegen ──┐
                           P3b Host ─────┼──► P3d Dev loop
                           P3c Runtime ──┘
                                         ├──► P5 Breadth (todo example) ──► …further P5
                                         └──► P4 Desktop (deferred; same bundle)
```

P0–P3d are done. **Orch** schedules **P5** tip rows below; **P4** stays deferred until the operator un-defers it. P5 does **not** wait on P4.

---

## OpenSpec workflow

- **Schema:** `spec-driven` (see [`openspec/config.yaml`](../openspec/config.yaml)).
- **Main specs home:** `openspec/specs/` — **`dev-tooling`**, **`agent-factory`**, **`ui-dsl`**, **`ui-docs`**, **`forma-cli`**, **`ui-codegen`**, **`ui-host`**, **`ui-runtime`**.
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

### Completed through P3d

Shipped and archived: **P0** `forma-dev-tooling`; **P1** `forma-ui-dsl`, `forma-fragment-slots`; **P2** `forma-user-docs`, `forma-cli-scaffold`; web tree via `forma-web-scaffold`; **P3a** `forma-ui-codegen`; **P3b** `forma-react-host`; **P3c** `forma-gopherjs-runtime`; **P3c′** `forma-browser-compile`; **P3d** `forma-dev-experience`. Meta: **`forma-agent-factory`** (local roles + this control-plane convention). Debt/refactor archives are intentionally omitted from the queue below.

### Suggested changes (forward queue)

Orch schedules from this table only. Leave **branch** / **PR** empty until a change is in flight.

| Order | Change name | Phase | Purpose | Status | branch | PR |
|-------|-------------|-------|---------|--------|--------|-----|
| 1 | `forma-todo-shadcn-slice` | P5 | Gap-analyze a minimal responsive shadcn todo UI; add any missing default-shadcn widgets, layout/responsive primitives, codegen, and host/registry wiring forma still lacks | complete | `forma-todo-shadcn-slice` | |
| 2 | `forma-todo-example` | P5 | Ship `examples/todo`: in-browser todo list on that slice, runnable with `forma dev`; document how to run it | Planned | | |
| 3 | `forma-desktop` | P4 | Native webview packaging, same bundle | Deferred | | |

Proposals must reference **ROADMAP phase** (P4–P5) in the Impact section. Further P5 items (extra catalog, 3rd-party recipes, perf notes) stay out of this queue until the todo example lands.

---

## Phase notes

### P1 — DSL (complete)

Authors rely on `openspec/specs/ui-dsl/` + [`main.go`](../main.go): `UIMount`/`RefreshUI`, dot-import `ui`/`fn`, `.With`/`.Or`, events via `ui.Context`, keys/inputs policy, escape hatches. Browser runtime is in `openspec/specs/ui-runtime/` (P3c).

### P2 — User docs (complete)

Guide under [`doc/user/`](user/README.md) for Go authors who may never open the React host. Docs track shipped CLI (`init`/`add`/`sync`/`generate`/`dev`), host, and bridge. Spec contradictions → update specs first.

### P3 — Implementation (complete)

- **P3a:** codegen owns prop parity with TSX; overrides explicit.
- **P3b–P3c:** host stays thin; runtime implements P1 contracts via `ui-runtime` (GoScript/`syscall/js`).
- **P3c′:** GopherJS retired; browser builds use GoScript under Go 1.27.
- **P3d:** `forma dev` runs the Vite host, watches Go, and rebuilds the browser bundle; user guide documents the loop.

### P5 — Breadth / todo example (planned)

**Product goal:** forma can run a **single in-browser todo list** as an **example project**—minimal UI, **default shadcn**, **responsive**.

1. **`forma-todo-shadcn-slice`** — invent the target UI once (architector), diff against today’s generated/default set (`button` / `input` / `card` / `checkbox` / layout, etc.), then extend forma only where the example needs it. Not a full catalog sweep.
2. **`forma-todo-example`** — add **`examples/todo`** (own module + web tree as needed), implement the UI, prove CRUD + layout at narrow/wide viewports via `forma dev`. Framework-root [`main.go`](../main.go) may stay a thin smoke until this example owns the narrative; migration of docs/pointers is in scope of the example change if useful.

Later P5 (not scheduled yet): wider shadcn coverage, third-party recipes, performance notes.

---

## Follow-ups (deferred)

| Item | Notes |
|------|--------|
| SSR / hydration | See DESIGN.md; only if product requires SEO |
| Signal-based fine-grained updates | Alternative to full-tree RefreshUI; breaks simple mental model |
| Desktop webview | P4 |
| Deeper agent prompts/skills | Optional; factory v1 ships roles + OpenSpec skills |
| Broader P5 catalog / 3rd-party / perf | After `examples/todo` ships |

---

## Success metrics (informal)

- New widget: **registry + codegen + Go usage** in minutes, not days.
- Agent-generated UI diff is **Go-only** and compiles.
- Authors never need to read React source for standard shadcn widgets.
- **P5 tip:** `examples/todo` runs in the browser with a responsive default-shadcn todo UI end-to-end.
