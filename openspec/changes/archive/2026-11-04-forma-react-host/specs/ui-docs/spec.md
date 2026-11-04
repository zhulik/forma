# Spec Delta

## ADDED Requirements

### Requirement: Guide documents the React host and text-input policy

The author guide SHALL describe the web host modules under the project web tree (registry merge, recursive renderer location), how authors add custom registry entries for third-party wrappers, and the text-input hybrid policy (ephemeral React state while focused; sanitized string callbacks on edit; value prop sync when unfocused).

#### Scenario: Author extends the registry

- **WHEN** an author needs a custom React wrapper registered beside generated widgets
- **THEN** the guide explains where to merge a hand-written registry entry without editing generated registry output

#### Scenario: Author hits text-input focus issues

- **WHEN** an author reads troubleshooting or writing-UI material about text fields
- **THEN** they learn the host hybrid policy and that domain state still lives in Go

## MODIFIED Requirements

### Requirement: Implementation status is explicit

Where documented behavior depends on work not yet shipped (GopherJS bridge, browser rebuild/`forma dev` loop), the guide SHALL label that material as planned or not yet available and SHALL NOT imply a live Go app in the browser when only the host + DSL stubs exist. The guide SHALL state which CLI commands are available versus planned: **`init`**, config-only init, **`add`**, **`generate`**, full **`sync`** (npm, shadcn materialization via init/add, then the same codegen path as generate, JS registry), and **`version`** are documented as implemented; the scaffold **React host** (registry + recursive renderer, fixture/placeholder tree until the bridge) is documented as available after this change; scaffold **`check-generate`** (`go generate` then whole-repo git diff) is documented for CI drift; **`forma dev`** and the live Go↔host bridge remain planned until P3c–P3d. The guide SHALL NOT document **`forma generate --check`** or **`forma sync --check`**.

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page before P3d lands
- **THEN** they see which steps work today (CLI init, add, generate, full sync with codegen and registry, React host modules, DSL reference via `main.go`) versus which require a future release for the live browser Go loop

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

- **WHEN** an author opens the scaffold web app after init without the Go bridge
- **THEN** documentation clarifies that the host can render a tree but the GopherJS-driven todo demo requires later runtime and `forma dev` phases
