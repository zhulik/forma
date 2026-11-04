# Design

## Context

See [proposal.md](./proposal.md). Before this change, **`forma sync`** ran npm install and batched **`shadcn add`** from **`ui.components`**, then logged skipped codegen/registry steps. The framework **`ui/`** package held hand-written **`*.gen.go`** stubs; app scaffolds did not emit real generated bindings. DSL rules for leaf vs compound widgets and slot rendering already lived in **`ui-dsl`** and **`doc/DESIGN.md`**.

## Goals / Non-Goals

**Goals:**

- Deterministic **TSX → `codegen.dir/*.gen.go` + `web/src/registry.gen.ts`** driven by files on disk under **`src/components/ui/`**.
- Public **`forma generate`** that loads **`forma.yaml`** and runs codegen **in-process in Go** (no Node for generate); **`forma sync`** and **`forma add`** reuse that path.
- Implement **`forma add`** as **`shadcn add` + sync**.
- **Init** materializes **`button`, `input`, `card`** via a single CLI constant; drop **`ui.components`** from scaffold **`forma.yaml`**.
- Scaffold **`generator.go`** (`//go:generate forma generate`, **`package main`**) and **`check-generate`** (`go generate` then whole-repo **`git diff --exit-code`**).
- Golden tests for at least one leaf and one compound widget; update init smoke and **`doc/user/`**.

**Non-Goals:**

- Node.js / ts-morph / separate codegen CLI as part of forma.
- Embedding Microsoft **typescript-go** internals (no supported public Go API; prefer IPC upstream).
- **`forma generate --check`** / **`forma sync --check`** flags.
- React host renderer, event bridge, GopherJS, **`forma dev`** (P3b–P3d).
- **`examples/`** directory and removing framework **`ui/*.gen.go`** stubs (follow-up change).
- Full shadcn catalog parity; shipping generated widgets inside **`github.com/zhulik/forma/ui`**.
- LLM-first YAML component registry as source of truth.

## Decisions

### 1. Canonical input = committed TSX glob

**Choice:** Codegen enumerates **`{web.path}/src/components/ui/*.tsx`** (Vite/shadcn `@/` → `src/` alias; exclude non-widget helpers via naming or allowlist file later).

**Alternatives:** Filter by **`forma.yaml` list** (rejected—duplicate inventory); scan all of **`web/`** (too broad).

**Prune:** Remove **`*.gen.go`** whose source TSX disappeared (same basename mapping).

### 2. Author surface = `forma generate`; analyzer is in-process Go

**Choice:** **`cmd/forma`** exposes **`forma generate`**, which binds **`formaconfig`** and runs an in-process Go codegen library under **`internal/codegen`**. Parse TSX with official **`github.com/tree-sitter/go-tree-sitter`** + **`github.com/tree-sitter/tree-sitter-typescript`** (TSX language; **CGO**). Classification: value exports (including `export { … }`), root vs prefixed parts, `*Props` members, heuristic event/scalar mapping. Authors, CI, and **`//go:generate`** call **`forma generate`** only.

**`forma sync`** after npm install calls the same runner, then **`go mod tidy`**.

**Alternatives:**

| Option | Why not |
|--------|---------|
| Node + ts-morph | Forces Node into forma CLI for generate |
| Embed typescript-go | No public Go API; `internal/` only; upstream prefers IPC |
| bonsai-tsx (pure Go) | Obscure one-maintainer module; prefer official tree-sitter |
| smacker/go-tree-sitter | Popular but stale; community moving to official bindings |
| Full typechecker fidelity | Overkill for v1 export/props heuristics |

**Override policy:** Optional per-file YAML/JSON beside TSX or central **`codegen.overrides`** in **`forma.yaml`** (minimal v1: per-file in repo only if needed for dialog/select spikes).

### 3. One `registry.gen.ts` at `web/src/registry.gen.ts`

**Choice:** Emit a single module under **`{web.path}/src/registry.gen.ts`** mapping Go **`Node.Type`** strings to React components + slot assembly metadata. Host **`main.tsx`** imports it; authors merge custom registry entries in a small hand-written file if needed (document in advanced guide).

**Alternatives:** Patch **`main.tsx`** each sync (fragile); one registry entry per file (harder to import); registry at web root (rejected—scaffold host lives under **`src/`**).

### 4. When to run `shadcn add`

**Choice:**

| Trigger | shadcn CLI | Codegen |
|---------|------------|---------|
| **`forma init` → sync** | Starter triple (CLI constant) | All TSX on disk (via generate path) |
| **`forma add X`** | Add **X** if missing | All TSX on disk (via sync → generate) |
| **`forma sync`** alone | No (unless future flag) | All TSX on disk (via generate path) |
| **`forma generate`** alone | No | All TSX on disk |

Fresh clones rely on **committed TSX**. **`forma generate`** needs no Node; init/add/sync still need Node for npm/shadcn only.

### 5. Defaults

**Choice:** **`codegen.dir` = `internal/ui`**, **`codegen.package` = `ui`**. Generated Go imports **`github.com/zhulik/forma/ui`** for **`BaseComponent`**, **`WidgetMeta`**, slot helpers.

### 6. Drift check via `go generate` + whole-repo git diff

**Choice:** Init template includes root **`generator.go`**:

```go
//go:generate forma generate
package main
```

Scaffold **Taskfile** target **`check-generate`** runs **`go generate`** then **`git diff --exit-code`** on the **whole repository**. No CLI **`--check`** flag.

## Risks / Trade-offs

- **[shadcn `@latest` drift]** → Pin shadcn CLI version in a follow-up if CI flakes; golden tests pin fixture TSX.
- **[Non-shadcn TSX in `components/ui/`]** → Heuristic classification may emit junk; document “CLI-managed widgets only” and add skip marker later.
- **[Compound heuristics wrong]** → Override file + golden test per problematic widget (accordion, dialog).
- **[Tree-sitter vs checker]** → Accept syntactic heuristics for v1; deepen later if needed.
- **[Breaking yaml]** → Early-stage allowed; document migration in **`doc/user/forma-yaml.md`**.
- **[`forma` on PATH for `go generate`]** → Document that CI and local check-generate need the forma binary available.
- **[CGO / C toolchain]** → Official tree-sitter needs a C compiler for `go build` / CI; documented as a forma CLI build prerequisite.
- **[tree-sitter binary size]** → Grammar + runtime linked in; acceptable for CLI.

## Migration Plan

1. Ship in-process Go **`forma generate`** + wired sync/add; remove Node codegen tool; update init template.
2. Update **`doc/user/`** and archive OpenSpec deltas.
3. Existing experiments: delete **`ui.components`**, run **`forma sync`** or **`forma generate`**, commit gen output.
4. Framework repo: keep **`ui/`** stubs until **`examples/`** change.
