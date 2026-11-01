# forma roadmap

forma is a **Go-native UI framework**: authors write composable components in Go; a fixed **React + shadcn/ui** host renders them. The roadmap prioritizes **DSL quality and author docs before runtime implementation**, so OpenSpec artifacts and code stay aligned with how people (and agents) actually use the framework.

**Current status:** **P0–P2** complete for specs and author docs. **CLI:** nested **`cmd/forma`** ships **`init`** (embedded **`minimal`** template + **`web/`** tree), **partial **`sync`** (npm + shadcn; codegen/registry stub), and **`version`**; **`add`** / full sync / **`dev`** remain P3. Canonical specs: [`dev-tooling`](../openspec/specs/dev-tooling/spec.md), [`ui-dsl`](../openspec/specs/ui-dsl/spec.md), [`forma-cli`](../openspec/specs/forma-cli/spec.md); author guide [`doc/user/`](user/README.md). [`main.go`](../main.go) is the reference todo demo; browser runtime still stub. **Next:** finish CLI breadth as needed, then P3a codegen → host → GopherJS → dev loop.

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
| **P0** | **Dev tooling** — golangci-lint, Taskfile, GitHub CI (lint/generate/fix), Cursor rules; `task test` wired | **`forma-dev-tooling`** → `openspec/specs/dev-tooling/` | **Done** |
| **P1** | **UI DSL v1** — semantics, composition, events, keys, inputs, escape hatches, **Fragment**, field-only compound **slots**; `main.go` reference; stub `ui` runtime | **`forma-ui-dsl`**, **`forma-fragment-slots`** (archived) → `openspec/specs/ui-dsl/` | **Done** |
| **P2** | **User documentation** — CLI-first install (`forma init` / `add` / `sync`), `forma.yaml`, project layout, writing UI, components, advanced, troubleshooting; docs-first policy | **`forma-user-docs`** → `doc/user/`, OpenSpec **`ui-docs`** + **`forma-cli`** (contract) | **Done** (guide + specs; CLI binary deferred) |
| **P3a** | **Codegen pipeline** — ts-morph (or equivalent) from shadcn TSX → `ui/*.gen.go`; CI check for drift | **`forma-ui-codegen`** | Planned |
| **P3b** | **React host** — registry, recursive renderer, event sanitization, text-input policy | **`forma-react-host`** | Planned |
| **P3c** | **GopherJS runtime** — real `UIMount`/`RefreshUI`, tree crossing, callback dispatch, `ui.Context` propagation | **`forma-gopherjs-runtime`** | Planned |
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
- **Main specs home:** `openspec/specs/` — **`dev-tooling`**, **`ui-dsl`** today; more capabilities (e.g. `ui-runtime`, `ui-codegen`) land when P3+ changes archive.
- **Active work:** `openspec/changes/<change-name>/` with `proposal.md`, `specs/.../spec.md`, `design.md`, `tasks.md`.
- **Archive:** completed changes move to `openspec/changes/archive/`; sync capabilities to `openspec/specs/`.

### Suggested changes (by phase)

| Order | Change name | Purpose | Status |
|-------|-------------|---------|--------|
| 1 | `forma-dev-tooling` | Lint, Taskfile, CI, Cursor rules | **Done** (archived) |
| 2 | `forma-ui-dsl` | Lock DSL behavior; `Component`, `Node`, `Context`, `.Or`, `fn`, events | **Done** (archived) |
| 3 | `forma-fragment-slots` | Fragment + field-only slots; tree `Slots` map | **Done** (archived) |
| 4 | `forma-user-docs` | Author guide under `doc/user/` + ui-docs / forma-cli specs | **Done** (archive pending) |
| 5 | `forma-cli-scaffold` | CLI module, init, stub sync, CI init smoke | **Done** |
| 6 | `forma-ui-codegen` | Generator + first real generated components | Planned |
| 7 | `forma-react-host` | Static host app + registry | Planned |
| 8 | `forma-gopherjs-runtime` | Wire Go to host | Planned |
| 9 | `forma-dev-experience` | Taskfile/scripts; demo runs in browser | Planned |

Proposals must reference **ROADMAP phase** (P1–P5) in the Impact section.

---

## P1 — DSL focus (complete)

Deliverables authors and agents can rely on **before** runtime exists (implemented and specified in `openspec/specs/ui-dsl/`):

- **`UIMount` / `RefreshUI`** contract and when handlers must refresh.
- **Imports:** dot-import `ui` and `fn` in apps; `fn.Map` / `fn.Filter` for lists and state.
- **Component** composition: `.With`, **`.Or`** fallbacks, string children, splatting dynamic slices (see `renderTodoList` in `main.go`).
- **Layout vs leaf** components; naming and prop conventions for generated shadcn types.
- **Events:** `ui.Context` on handlers; no DOM types in Go API.
- **Lists:** keys; **inputs:** controlled vs deferred sync policy.
- **Escape hatches:** `Class`, `Style`, `Attrs` — documented limits.
- **Open questions** captured in design (not left implicit): optional vs required props, variant typing strictness, custom component naming.

Acceptance: a reader can implement a new screen **only from specs + `main.go`**, without reading React code.

---

## P2 — User documentation focus

Target audience: Go developers who may never open the React host.

Topics (minimum):

1. Prerequisites and repo layout (Go app + `web/` or equivalent).
2. Creating an app from scratch (clone template → `main.go` → mount id in `index.html`).
3. State + `ctx.RefreshUI()` pattern; async goroutines.
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
| Signal-based fine-grained updates | Alternative to full-tree RefreshUI; breaks simple mental model |
| Desktop webview | P4 |
| Official CLI binary (`cmd/forma`) | Spec in **`forma-cli`**; implement with P3 (after or alongside dev loop) |
| Agent-specific prompts/skills | Optional; after DSL + docs stable |

---

## Success metrics (informal)

- New widget: **registry + codegen + Go usage** in minutes, not days.
- Agent-generated UI diff is **Go-only** and compiles.
- Authors never need to read React source for standard shadcn widgets.
