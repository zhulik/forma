# Spec Delta

## ADDED Requirements

### Requirement: Guide documents the runnable examples/todo project

The author guide SHALL document the repository’s `examples/todo` app as the runnable in-browser todo example: how to install dependencies (`forma sync` or equivalent), how to run `forma dev` from that project directory, and that the UI is a single default-shadcn todo list proving create/toggle/delete and responsive composer layout.

#### Scenario: Author finds how to run the example

- **WHEN** an author opens install, CLI, or project-layout material looking for a live browser demo
- **THEN** they are directed to `examples/todo` with concrete run steps (sync prerequisites, `forma dev`, open the Vite URL)

#### Scenario: Author understands example vs framework smoke

- **WHEN** an author compares the framework repository root to an initialized app
- **THEN** the guide states that root `main.go` is thin DSL smoke without a web tree, while `examples/todo` is the full app module used for browser verification

## MODIFIED Requirements

### Requirement: Guide documents the forma dev browser loop

The author guide SHALL document how to run `forma dev`, including prerequisites, host-vs-Go update semantics, and a concrete browser verification path for the **`examples/todo`** reference todo (or equivalent interactive UI from that example’s patterns).

#### Scenario: Author learns rebuild vs HMR

- **WHEN** an author reads CLI or install documentation for the browser loop
- **THEN** they learn that host asset changes use the web tool's HMR while Go application changes rebuild the browser Go artifact and need a page reload (or remount), not preserved in-memory Go state across rebuilds

#### Scenario: Author runs the todo demo in the browser

- **WHEN** an author follows the documented verification path after syncing `examples/todo`
- **THEN** the guide tells them how to open the running app in a browser and confirm interactive todo behavior (add, toggle done, delete) driven by Go via the live bridge, including that the composer stacks on narrow viewports and sits Input+Add side-by-side from the `sm` breakpoint up

### Requirement: Implementation status is explicit

Where documented behavior still depends on work not yet shipped (for example desktop packaging), the guide SHALL label that material as planned or not yet available. After this change, the guide SHALL treat **`forma dev`**, the live Go↔host bridge, the scaffold React host, and the **`examples/todo`** runnable example as available: **`init`**, config-only init, **`add`**, **`generate`**, full **`sync`**, **`version`**, and **`dev`** are documented as implemented; scaffold **`check-generate`** remains documented for CI drift. The guide SHALL NOT document **`forma generate --check`** or **`forma sync --check`**.

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
- **THEN** documentation directs them to `examples/todo` and confirms the Go-driven todo demo is the expected success check

### Requirement: Entry points link to the author guide

The repository root README SHALL link to `doc/user/` for authors. ROADMAP and DESIGN SHALL treat `doc/user/` as the author onboarding path alongside OpenSpec `ui-dsl`, with the runnable reference UI pointed at **`examples/todo`** (framework-root `main.go` MAY remain cited only as thin DSL smoke).

#### Scenario: README reader

- **WHEN** someone reads README.md for forma
- **THEN** they find a clear link to the user documentation index and a pointer to the runnable `examples/todo` demo

#### Scenario: DESIGN north-star pointer

- **WHEN** a reader opens DESIGN for the authoring model north-star
- **THEN** the north-star runnable example is `examples/todo`, not an implication that root `main.go` alone is the browser product demo
