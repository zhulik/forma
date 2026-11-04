# Proposal

## Why

Authors can already scaffold a web tree and materialize shadcn TSX via **`forma sync`**, but Go widget bindings and the JS **ComponentRegistry** are stubbed—apps cannot use generated shadcn types or a host registry aligned with the DSL. P3a needs a deterministic **TSX → Go + registry** pipeline exposed as **`forma generate`** (and reused by **`forma sync`** / **`forma add`**), with init supplying a small default widget set without maintaining a duplicate component inventory in **`forma.yaml`**. Codegen itself MUST run in pure Go so the forma CLI has **no Node.js dependency** for generation (Node remains only for app web tooling: npm / shadcn).

## What Changes

- **Codegen pipeline** (pure Go) scans shadcn TSX under the configured web root (`src/components/ui/`), emits `*.gen.go` into `codegen.dir` (default `internal/ui`, package `ui`), and regenerates a single **`registry.gen.ts`** for the host. Parsing uses official **tree-sitter** Go bindings + **tree-sitter-typescript** (TSX; CGO). Microsoft’s Go TypeScript compiler is **not** embedded (no public Go API).
- **`forma generate`** is the public codegen command: it reads **`forma.yaml`** and runs full codegen/registry regeneration from on-disk TSX (prune stale generated Go when TSX is removed). Authors and `go generate` invoke this command; there is no separate codegen binary and **no Node subprocess for generate**.
- **`forma sync`** runs npm install, invokes the **same generate path**, and runs `go mod tidy`; it does **not** re-fetch shadcn sources on every sync unless driven by init or **`forma add`**.
- **`forma add <component>`** runs **`shadcn add`** for that name then **`forma sync`**; it does **not** append to **`forma.yaml`**.
- **`forma init`** materializes a **CLI-hardcoded** starter set (`button`, `input`, `card`) via shadcn during first sync; scaffold **`forma.yaml`** drops **`ui.components`**.
- **BREAKING:** **`forma.yaml`** no longer declares **`ui.components`**; the widget catalog is **committed TSX + generated outputs**.
- Minimal init template gains root **`generator.go`** with **`//go:generate forma generate`**, plus a **`check-generate`** task that runs **`go generate`** then **`git diff --exit-code`** on the **whole repo**. No **`forma generate --check`** or **`forma sync --check`**.
- **`doc/user/`** updated for generate, full sync, add workflow, yaml shape, and commit expectations.
- Framework repo **`ui/*.gen.go` stubs** remain until a follow-up **`examples/`** change; this change does not add **`examples/`** or the React host runtime (P3b).

## Capabilities

### New Capabilities

- `ui-codegen`: Deterministic generation of app widget bindings and host registry from local shadcn TSX; project-level drift verification via regenerate + whole-repo git diff.

### Modified Capabilities

- `forma-cli`: **`forma generate`**, init starter components, **`forma add`**, full **`forma sync`** (calls generate), **`forma.yaml`** shape, scaffold **`generator.go`** and **`check-generate`**.
- `ui-docs`: Implementation status, generate/sync/add/yaml documentation for codegen and registry.
- `dev-tooling`: Init scaffold smoke expectations when generated Go must compile after sync.

## Impact

- **`cmd/forma`**: new **`generate`** subcommand, sync service calls shared **in-process Go** codegen runner, add command, init template (`forma.yaml.tmpl`, **`generator.go`**, **`Taskfile.yaml`**).
- **Codegen library** under **`cmd/forma/internal/codegen`** (analyze + emit); **delete** any Node **`tools/forma-codegen`** package; CLI tests do not require Node for generate unit/golden coverage.
- **Scaffold `web/`**: import **`registry.gen.ts`** from host entry; generated files committed after init.
- **OpenSpec main specs** archive when change completes; **`doc/ROADMAP.md`** P3a status.
- **Tests**: generate/sync/add fakes, codegen golden tests (fixtures in Go testdata), CI init smoke may assert presence of `internal/ui/*.gen.go`.
