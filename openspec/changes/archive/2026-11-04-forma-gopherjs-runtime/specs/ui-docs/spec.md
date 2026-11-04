# Spec Delta

## ADDED Requirements

### Requirement: Guide documents the Go↔host bridge

The author guide SHALL describe how the browser runtime mounts a root component, refreshes the host tree from Go state, and delivers sanitized callbacks into handlers with callback context. It SHALL state that a one-command browser rebuild loop remains planned until the dedicated dev-loop phase.

#### Scenario: Author looks up RefreshUI behavior

- **WHEN** an author reads writing-UI or troubleshooting material after the runtime ships
- **THEN** they see that mutations plus refresh (or the update helper) drive the host tree from Go

#### Scenario: Author distinguishes bridge from forma dev

- **WHEN** an author reads status or CLI docs after the runtime ships
- **THEN** they see the Go↔host bridge as available and `forma dev` as still planned

## MODIFIED Requirements

### Requirement: Implementation status is explicit

Where documented behavior depends on work not yet shipped (browser rebuild / `forma dev` loop), the guide SHALL label that material as planned or not yet available. The guide SHALL state which CLI commands are available versus planned: **`init`**, config-only init, **`add`**, **`generate`**, full **`sync`** (npm, shadcn materialization via init/add, then the same codegen path as generate, JS registry), and **`version`** are documented as implemented; the scaffold **React host** and the **Go↔host bridge** (`UIMount` / refresh / callbacks) are documented as available after this change; scaffold **`check-generate`** (`go generate` then whole-repo git diff) is documented for CI drift; **`forma dev`** remains planned until P3d. The guide SHALL NOT document **`forma generate --check`** or **`forma sync --check`**. The guide SHALL NOT imply that `forma dev` already provides a one-command live rebuild when only the bridge exists.

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page before P3d lands
- **THEN** they see which steps work today (CLI init, add, generate, full sync with codegen and registry, React host modules, Go↔host bridge, DSL reference via `main.go`) versus which require a future release for the one-command browser rebuild loop

#### Scenario: Author expects full sync today

- **WHEN** an author reads sync documentation after Go codegen and the JS registry ship
- **THEN** they see that sync installs npm dependencies, runs the same codegen as `forma generate` from committed shadcn TSX, and that shadcn fetch happens on init (starter set) and `forma add`, not via a component list in `forma.yaml`

#### Scenario: Author verifies generated output in CI

- **WHEN** an author reads project layout or CLI docs for drift checking
- **THEN** they see that the scaffold uses root `generator.go` with `go:generate forma generate` and a `check-generate` task that runs `go generate` then `git diff --exit-code` on the whole repository

#### Scenario: Author learns generate does not need Node

- **WHEN** an author reads CLI or install documentation for codegen
- **THEN** they see that `forma generate` does not require Node.js, while init, add, and sync still need Node for npm and shadcn

#### Scenario: Author expects a live todo app in the browser

- **WHEN** an author opens the scaffold web app after init with the Go bridge available but without `forma dev`
- **THEN** documentation clarifies that the bridge can drive the host from Go while the one-command browser rebuild loop remains a later phase
