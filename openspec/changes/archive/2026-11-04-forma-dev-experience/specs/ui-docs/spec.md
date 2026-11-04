# Spec Delta

## ADDED Requirements

### Requirement: Guide documents the forma dev browser loop

The author guide SHALL document how to run `forma dev`, including prerequisites, host-vs-Go update semantics, and a concrete browser verification path for the reference todo demo (or equivalent interactive UI from `main.go` patterns).

#### Scenario: Author learns rebuild vs HMR

- **WHEN** an author reads CLI or install documentation for the browser loop
- **THEN** they learn that host asset changes use the web tool’s HMR while Go application changes rebuild the browser Go artifact and need a page reload (or remount), not preserved in-memory Go state across rebuilds

#### Scenario: Author runs the todo demo in the browser

- **WHEN** an author follows the documented verification path after `forma sync`
- **THEN** the guide tells them how to open the running app in a browser and confirm interactive todo (or equivalent) behavior driven by Go via the live bridge

## MODIFIED Requirements

### Requirement: Guide covers the P2 topic set

The author guide SHALL include dedicated material for: prerequisites and toolchain; CLI onboarding (`init`, `add`, `sync`, and `dev`); `forma.yaml` and project layout; writing UI (state, mount, refresh, dot imports, `fn`); adding components; custom Go and third-party React extension; troubleshooting.

#### Scenario: Author adds a shadcn-backed widget

- **WHEN** an author needs to add a new visual component to their app
- **THEN** the guide explains `forma add`, the role of `forma sync`, and what belongs in version control

### Requirement: Implementation status is explicit

Where documented behavior still depends on work not yet shipped (for example desktop packaging), the guide SHALL label that material as planned or not yet available. After this change, the guide SHALL treat **`forma dev`**, the live Go↔host bridge, and the scaffold React host as available: **`init`**, config-only init, **`add`**, **`generate`**, full **`sync`**, **`version`**, and **`dev`** are documented as implemented; scaffold **`check-generate`** remains documented for CI drift. The guide SHALL NOT document **`forma generate --check`** or **`forma sync --check`**.

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page after P3d lands
- **THEN** they see `forma dev` documented as the one-command browser loop (host + Go rebuild), not as planned

#### Scenario: Author expects full sync today

- **WHEN** an author reads sync documentation after Go codegen and the JS registry ship
- **THEN** they see that sync installs npm dependencies, runs the same codegen as `forma generate` from committed shadcn TSX, and that shadcn fetch happens on init (starter set) and `forma add`, not via a component list in `forma.yaml`

#### Scenario: Author verifies generated output in CI

- **WHEN** an author reads project layout or CLI docs for drift checking
- **THEN** they see that the scaffold uses root `generator.go` with `go:generate forma generate` and a `check-generate` task that runs `go generate` then `git diff --exit-code` on the whole repository

#### Scenario: Author learns generate does not need Node

- **WHEN** an author reads CLI or install documentation for codegen
- **THEN** they see that `forma generate` does not require Node.js, while init, add, sync, and `forma dev` still need Node for the web tree

#### Scenario: Author expects a live todo app in the browser

- **WHEN** an author follows the documented `forma dev` verification path with the live bridge available
- **THEN** documentation directs them to open the browser URL from the dev loop and confirms the Go-driven todo (or equivalent) demo is the expected success check
