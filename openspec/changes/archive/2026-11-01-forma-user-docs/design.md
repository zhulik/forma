# Design

## Context

See [proposal.md](./proposal.md). The framework repo still has stub `UIMount`, no `host/`, and no CLI. Exploration fixed the **target** author journey: install **`forma` CLI**, **`forma init <module>`** (scaffold + sync), **`forma add`**, **`forma sync`** (host deps, shadcn sources, codegen, registry, **`go mod tidy`**), generated widgets default **`internal/ui`**, config in **`forma.yaml`**.

## Goals / Non-Goals

**Goals:**

- Single Markdown tree under **`doc/user/`** as the authoritative author guide until a viewer/site exists.
- Document CLI + yaml as the **contract** future P3 work implements; label what works **today** vs **planned**.
- Encode **docs-first**: author-visible behavior changes require user-doc updates in the same change (spec + OpenSpec rules).
- Align ROADMAP/DESIGN onboarding with CLI-first model.

**Non-Goals:**

- Implementing CLI, host, or codegen.
- GitBook, Docusaurus, or CI doc publishing.
- Duplicating full `ui-dsl` spec prose—link to specs and `main.go` instead.

## Decisions

### 1. Plain Markdown in-repo

**Choice:** `doc/user/*.md` with `doc/user/README.md` as the map.

**Alternatives:** GitBook (deferred—content churn, extra ops); generated site (defer until CLI/runtime stable).

### 2. Two capabilities: `ui-docs` + `forma-cli`

**Choice:** `ui-docs` owns guide structure and docs-first policy; `forma-cli` owns command/yaml behavior authors can rely on when the binary lands.

**Alternatives:** `skip_specs` docs-only change—rejected because CLI/yaml contract should be normative before implementation.

### 3. “Available today” callouts

**Choice:** Install and **`forma dev`** pages state clearly: DSL is usable via this repo’s `main.go` + `task test`; browser loop and CLI are **planned** per ROADMAP P3.

**Alternatives:** Document only current reality—rejected; P2 purpose is to lock target UX before build.

### 4. Docs-first enforcement

**Choice:** Requirement in **`ui-docs`** spec; mirror in **`openspec/config.yaml`** `rules.tasks` and `rules.proposal`; repeat in **`doc/user/README.md`** for human readers.

**Alternatives:** Cursor rule only—too easy to miss in proposals.

### 5. Default codegen path in docs

**Choice:** Document default **`codegen.dir: internal/ui`**, overridable in **`forma.yaml`**.

## Risks / Trade-offs

- **[Risk] Docs describe unimplemented CLI** → Mitigation: status table on guide index; `forma-cli` spec is explicit “target”; tasks link implementation change.
- **[Risk] DESIGN/ROADMAP drift from guide** → Mitigation: tasks update both; archive sync; docs-first requirement on future changes.
- **[Risk] Framework repo layout vs user app layout** → Mitigation: `project-layout.md` describes **user app** scaffold; note this repo is library + reference demo until split/examples.

## Migration Plan

N/A (greenfield docs). After **`forma-cli`** implements, remove “planned” banners section-by-section and add verification steps (e2e) in that change.

## Open Questions

- Exact default starter **`ui.components`** list in scaffold (recommend button, input, card)—finalize in **`forma-cli`** implementation change.
- Whether **`forma dev`** auto-runs sync when stale—document as recommendation in cli spec; implementation can refine.
