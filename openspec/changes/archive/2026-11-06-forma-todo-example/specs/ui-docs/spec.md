# Spec Delta

## MODIFIED Requirements

### Requirement: Implementation status is explicit

Where documented behavior still depends on work not yet shipped (for example desktop packaging), the guide SHALL label that material as planned or not yet available. The guide SHALL treat **`forma dev`**, the live Go↔host bridge, the scaffold React host, the **owned-widgets + semantic CSS** path, and the dedicated **`examples/todo`** runnable demo as available for the todo-sized surface. The guide SHALL document **`init`**, config-only init, **`sync`**, **`version`**, and **`dev`** as implemented for that path. The guide MUST NOT document shadcn **`add`**, kit **`generate`**, or scaffold **`check-generate`** as the primary author workflow. The guide SHALL NOT document **`forma generate --check`** or **`forma sync --check`**.

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page after P3d lands
- **THEN** they see `forma dev` documented as the one-command browser loop (host + Go rebuild), not as planned

#### Scenario: Author expects full sync today

- **WHEN** an author reads sync documentation after owned-widgets migration
- **THEN** they see that sync installs npm dependencies and tidies Go modules, without requiring kit codegen or shadcn fetch as part of the default path

#### Scenario: Author verifies generated output in CI

- **WHEN** an author reads project layout or CLI docs for CI verification after owned-widgets migration
- **THEN** they see that new scaffolds do not rely on `generator.go` / `check-generate` for kit bindings, and that compile/test/init smoke cover the product path

#### Scenario: Author learns generate does not need Node

- **WHEN** an author reads CLI or install documentation looking for kit codegen
- **THEN** they learn that kit `forma generate` is not part of the primary owned-widgets workflow, while init, sync, and `forma dev` still need Node for the web tree

#### Scenario: Author expects a live todo app in the browser

- **WHEN** an author follows the documented `forma dev` verification path with the live bridge available
- **THEN** documentation directs them to the dedicated `examples/todo` run path (sync + `forma dev` from that project root) and confirms a Go-driven interactive todo UI—not labeled planned

### Requirement: Guide documents the forma dev browser loop

The author guide SHALL document how to run `forma dev`, including prerequisites, host-vs-Go update semantics, and a concrete browser verification path for the dedicated `examples/todo` demo.

#### Scenario: Author learns rebuild vs HMR

- **WHEN** an author reads CLI or install documentation for the browser loop
- **THEN** they learn that host asset changes use the web tool's HMR while Go application changes rebuild the browser Go artifact and need a page reload (or remount), not preserved in-memory Go state across rebuilds

#### Scenario: Author runs the todo demo in the browser

- **WHEN** an author follows the documented verification path after `forma sync` in `examples/todo`
- **THEN** the guide tells them how to open the running app in a browser and confirm interactive todo behavior driven by Go via the live bridge

## ADDED Requirements

### Requirement: Guide documents the examples/todo run path

The author guide and repository entry points SHALL document how to sync and run the dedicated todo example with `forma dev`, including that the example is a nested app module with its own web tree (not the framework root).

#### Scenario: Install or README reader

- **WHEN** an author opens install material or the root README looking for a browser todo demo
- **THEN** they find clear steps to enter `examples/todo`, sync, and run `forma dev`

#### Scenario: Project layout distinguishes framework vs example

- **WHEN** an author reads project-layout material about this repository
- **THEN** they learn that `examples/todo` is the runnable demo app while the framework root remains library + CLI source without requiring an app `web/` at repo root
