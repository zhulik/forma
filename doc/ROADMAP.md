# forma roadmap

forma is a **Go-native UI framework**: authors write composable components in Go; a fixed **React + shadcn/ui** host renders them. The roadmap prioritizes **DSL quality and author docs before runtime implementation**, so OpenSpec artifacts and code stay aligned with how people (and agents) actually use the framework.

**Current status:** Greenfield. [`main.go`](../main.go) is the DSL reference demo; [`ui/`](../ui/) has stub `Mount`/`Refresh` and placeholder `*.gen.go` files. No working GopherJS bundle or React host in-repo yet.

**Canonical design:** [DESIGN.md](./DESIGN.md)

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

Work proceeds in **strict order**: DSL → user documentation → implementation. OpenSpec changes should name the phase they belong to.

| Phase | Scope | OpenSpec / deliverables | Status |
|-------|--------|-------------------------|--------|
| **P1** | **UI DSL v1** — semantics, composition, events, keys, inputs, escape hatches; evolve `main.go` as reference; stub `ui` types only | Change e.g. **`forma-ui-dsl`**: proposal, specs (`ui-dsl`), design, tasks; may add `examples/` variants | **Next** |
| **P2** | **User documentation** — install toolchain (Go, Node, GopherJS, host app), project layout, `Mount`/`Refresh`, state pattern, codegen workflow, adding shadcn components, custom Go/JS components, troubleshooting | Change e.g. **`forma-user-docs`**: specs for doc completeness OR doc-only tasks; publish under `doc/user/` (or site later) | Planned |
| **P3a** | **Codegen pipeline** — ts-morph (or equivalent) from shadcn TSX → `ui/*.gen.go`; CI check for drift | **`forma-ui-codegen`** | Planned |
| **P3b** | **React host** — registry, recursive renderer, event sanitization, text-input policy | **`forma-react-host`** | Planned |
| **P3c** | **GopherJS runtime** — real `Mount`/`Refresh`, tree crossing, callback dispatch, `context` propagation | **`forma-gopherjs-runtime`** | Planned |
| **P3d** | **End-to-end dev loop** — one command to run host + rebuild Go; todo demo works in browser | **`forma-dev-experience`** | Planned |
| **P4** | **Desktop packaging** — native webview wrapper, same bundle | **`forma-desktop`** | Deferred |
| **P5** | **Breadth** — more shadcn widgets, 3rd-party recipes (maps, payments), performance notes | Multiple small changes | Deferred |

### Phase dependencies

```
P1 DSL ──► P2 User docs ──► P3a Codegen ──┐
                           P3b Host ─────┼──► P3d Dev loop
                           P3c Runtime ──┘
                                         └──► P4 Desktop ──► P5 Breadth
```

**Rule:** Do not merge implementation-heavy changes (P3+) until P1 specs are approved and P2 doc outline exists (even if P2 prose is still draft). Exception: spikes behind feature flags must say so in the proposal.

---

## OpenSpec workflow

- **Schema:** `spec-driven` (see [`openspec/config.yaml`](../openspec/config.yaml)).
- **Main specs home:** `openspec/specs/` — populated when changes archive (e.g. `ui-dsl`, `ui-runtime`, `ui-codegen`).
- **Active work:** `openspec/changes/<change-name>/` with `proposal.md`, `specs/.../spec.md`, `design.md`, `tasks.md`.
- **Archive:** completed changes move to `openspec/changes/archive/`; sync capabilities to `openspec/specs/`.

### Suggested first changes

| Order | Change name | Purpose |
|-------|-------------|---------|
| 1 | `forma-ui-dsl` | Lock DSL behavior; extend `main.go`; define `Component`, `Node`, `Refresh`, events |
| 2 | `forma-user-docs` | Author-facing guide structure + content tasks |
| 3 | `forma-ui-codegen` | Generator + first real generated components |
| 4 | `forma-react-host` | Static host app + registry |
| 5 | `forma-gopherjs-runtime` | Wire Go to host |
| 6 | `forma-dev-experience` | Taskfile/scripts; demo runs in browser |

Proposals must reference **ROADMAP phase** (P1–P5) in the Impact section.

---

## P1 — DSL focus (immediate)

Deliverables authors and agents can rely on **before** runtime exists:

- **Mount / Refresh** contract and when to call `Refresh`.
- **Component** composition: `.With`, string children, splatting dynamic slices (see `renderTodoList` in `main.go`).
- **Layout vs leaf** components; naming and prop conventions for generated shadcn types.
- **Events:** `context.Context` first parameter; no DOM types in Go API.
- **Lists:** keys; **inputs:** controlled vs deferred sync policy.
- **Escape hatches:** `Class`, `Style`, `Attrs` — documented limits.
- **Open questions** captured in design (not left implicit): optional vs required props, variant typing strictness, custom component naming.

Acceptance: a reader can implement a new screen **only from specs + `main.go`**, without reading React code.

---

## P2 — User documentation focus

Target audience: Go developers who may never open the React host.

Topics (minimum):

1. Prerequisites and repo layout (Go app + `host/` or equivalent).
2. Creating an app from scratch (clone template → `main.go` → mount id in `index.html`).
3. State + `Refresh` pattern; async goroutines.
4. Adding a shadcn component (CLI → registry → codegen → Go usage).
5. Custom composite components in Go; custom JS wrapper for a npm React lib.
6. Debugging (tree not updating, callback not firing, input focus).

Docs must match P1 specs; contradictions are resolved by updating specs first.

---

## P3 — Implementation focus

Implementation is **downstream** of P1 + P2:

- **Codegen** owns prop parity with TSX; manual overrides are explicit.
- **Host** stays thin; all business rules stay in Go.
- **Runtime** implements exactly what P1 specifies; gaps become spec amendments, not silent behavior.

**Definition of done for P3d:** `main.go` todo app runs in a browser with HMR or fast rebuild documented in user guide.

---

## Follow-ups (deferred)

| Item | Notes |
|------|--------|
| SSR / hydration | See DESIGN.md; only if product requires SEO |
| Signal-based fine-grained updates | Alternative to full-tree Refresh; breaks simple mental model |
| Desktop webview | P4 |
| Official CLI (`forma init`) | After dev loop stable |
| Agent-specific prompts/skills | Optional; after DSL + docs stable |

---

## Success metrics (informal)

- New widget: **registry + codegen + Go usage** in minutes, not days.
- Agent-generated UI diff is **Go-only** and compiles.
- Authors never need to read React source for standard shadcn widgets.
