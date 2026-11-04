# ui-docs Specification

## Purpose

Define where forma's author-facing documentation lives, what it must cover, and the project obligation to keep it aligned with author experience.

## Requirements

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

### Requirement: Manual project initialization is documented

The author guide SHALL document how to create a forma app without the CLI: initializing the Go module, authoring `forma.yaml`, adding a minimal `main.go` that uses the forma library, and running `go mod tidy`, with clear pointers to when `forma sync` and web tree setup apply later.

#### Scenario: Author prefers manual setup

- **WHEN** an author cannot or will not use `forma init`
- **THEN** the guide provides step-by-step manual setup that matches the fields and layout produced by init

### Requirement: Implementation status is explicit

Where documented behavior depends on work not yet shipped (browser dev loop, GopherJS rebuild integration), the guide SHALL label that material as planned or not yet available and SHALL NOT imply a working browser app when only the DSL stub exists in the framework repository. The guide SHALL state which CLI commands are available versus planned: **`init`**, config-only init, **`add`**, **`generate`**, full **`sync`** (npm, shadcn materialization via init/add, then the same codegen path as generate, JS registry), and **`version`** are documented as implemented once this change ships; scaffold **`check-generate`** (`go generate` then whole-repo git diff) is documented for CI drift; **`forma dev`** remains planned until P3d. The guide SHALL NOT document **`forma generate --check`** or **`forma sync --check`**.

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page before P3d lands
- **THEN** they see which steps work today (CLI init, add, generate, full sync with codegen and registry, DSL reference via `main.go`) versus which require a future release for the browser loop

#### Scenario: Author expects full sync today

- **WHEN** an author reads sync documentation after Go codegen and the JS registry ship
- **THEN** they see that sync installs npm dependencies, runs the same codegen as `forma generate` from committed shadcn TSX, and that shadcn fetch happens on init (starter set) and `forma add`, not via a component list in `forma.yaml`

#### Scenario: Author verifies generated output in CI

- **WHEN** an author reads project layout or CLI docs for drift checking
- **THEN** they see that the scaffold uses root `generator.go` with `go:generate forma generate` and a `check-generate` task that runs `go generate` then `git diff --exit-code` on the whole repository

#### Scenario: Author learns generate does not need Node

- **WHEN** an author reads CLI or install documentation for codegen
- **THEN** they see that `forma generate` does not require Node.js, while init, add, and sync still need Node for npm and shadcn

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
