# Spec Delta

## Purpose

Define where forma's author-facing documentation lives, what it must cover, and the project obligation to keep it aligned with author experience.

## ADDED Requirements

### Requirement: Author guide lives in plain Markdown

The project SHALL maintain the primary author guide as plain Markdown under `doc/user/` until a separate published site replaces it as the source of truth.

#### Scenario: Contributor finds install instructions

- **WHEN** a new author opens the repository looking for how to start with forma
- **THEN** `doc/user/README.md` or linked pages under `doc/user/` describe prerequisites, onboarding, and pointers to deeper topics

### Requirement: Guide covers the P2 topic set

The author guide SHALL include dedicated material for: prerequisites and toolchain; CLI onboarding (`init`, `add`, `sync`, and planned `dev`); `forma.yaml` and project layout; writing UI (state, mount, refresh, dot imports, `fn`); adding components; custom Go and third-party React extension; troubleshooting.

#### Scenario: Author adds a shadcn-backed widget

- **WHEN** an author needs to add a new visual component to their app
- **THEN** the guide explains `forma add`, the role of `forma sync`, and what belongs in version control

### Requirement: Implementation status is explicit

Where documented behavior depends on work not yet shipped (CLI binary, browser dev loop, codegen, host), the guide SHALL label that material as planned or not yet available and SHALL NOT imply a working browser app when only the DSL stub exists in the framework repository.

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page before P3d lands
- **THEN** they see which steps work today (e.g. DSL reference via `main.go`) versus which require a future release

### Requirement: User documentation is first-class

Any change that materially affects author experience (CLI commands, `forma.yaml` shape, project layout, DSL authoring rules exposed to users, install or sync workflow) SHALL update the author guide in the same delivery slice unless the change explicitly opts out in its proposal with a documented reason.

#### Scenario: New CLI subcommand ships

- **WHEN** a change introduces or renames an author-facing CLI command
- **THEN** that change includes updates to the relevant `doc/user/` pages (or a tracked follow-up task in the same change if docs-only merge is impossible, with proposal justification)

### Requirement: Entry points link to the author guide

The repository root README SHALL link to `doc/user/` for authors. ROADMAP and DESIGN SHALL treat `doc/user/` as the author onboarding path alongside `main.go` and OpenSpec `ui-dsl`.

#### Scenario: README reader

- **WHEN** someone reads README.md for forma
- **THEN** they find a clear link to the user documentation index
