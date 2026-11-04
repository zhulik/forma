# Spec Delta — ui-docs

## MODIFIED Requirements

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
