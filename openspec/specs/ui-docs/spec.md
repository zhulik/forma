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

The author guide SHALL include dedicated material for: prerequisites and toolchain (including the supported browser Go→JS/TS compiler aligned with the module's Go version); CLI onboarding (`init`, `sync`, and available `dev`); `forma.yaml` and project layout; writing UI (state, mount, refresh, dot imports, `fn`); using owned widgets and the styling stack; custom Go and third-party React extension; troubleshooting.

#### Scenario: Author adds a shadcn-backed widget

- **WHEN** an author needs Button, Input, Checkbox, or another todo-sized owned widget
- **THEN** the guide explains importing the framework UI package and does not prescribe shadcn `forma add` + kit codegen as the primary path

#### Scenario: Author finds browser compiler prerequisites

- **WHEN** an author opens the install or prerequisites material
- **THEN** they see the supported browser compiler for the current Go module version, not an EOL Go 1.21-only GopherJS path presented as the durable install story

### Requirement: Manual project initialization is documented

The author guide SHALL document how to create a forma app without the CLI: initializing the Go module, authoring `forma.yaml`, adding a minimal `main.go` that uses the forma library, and running `go mod tidy`, with clear pointers to when `forma sync` and web tree setup apply later.

#### Scenario: Author prefers manual setup

- **WHEN** an author cannot or will not use `forma init`
- **THEN** the guide provides step-by-step manual setup that matches the fields and layout produced by init

### Requirement: Implementation status is explicit

Where documented behavior still depends on work not yet shipped (for example desktop packaging or a dedicated `examples/todo` package), the guide SHALL label that material as planned or not yet available. After this change, the guide SHALL treat **`forma dev`**, the live Go↔host bridge, the scaffold React host, and the **owned-widgets + semantic CSS** path as available for the todo-sized surface. The guide SHALL document **`init`**, config-only init, **`sync`**, **`version`**, and **`dev`** as implemented for that path. The guide MUST NOT document shadcn **`add`**, kit **`generate`**, or scaffold **`check-generate`** as the primary author workflow. The guide SHALL NOT document **`forma generate --check`** or **`forma sync --check`**.

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
- **THEN** documentation directs them to open the browser URL from the dev loop and confirms a Go-driven interactive UI (reference `main.go` patterns; dedicated `examples/todo` may still be labeled planned until that change ships)

### Requirement: Guide documents the Go↔host bridge

The author guide SHALL describe how the browser runtime mounts a root component, refreshes the host tree from Go state, and delivers sanitized callbacks into handlers with callback context. It SHALL describe `forma dev` as the available one-command browser rebuild loop. It SHALL describe the browser compile/rebuild story in terms of the project's supported compiler and SHALL NOT require authors to install GopherJS against Go 1.21 as the supported path.

#### Scenario: Author looks up RefreshUI behavior

- **WHEN** an author reads writing-UI or troubleshooting material after the runtime ships
- **THEN** they see that mutations plus refresh (or the update helper) drive the host tree from Go

#### Scenario: Author distinguishes bridge from forma dev

- **WHEN** an author reads status or CLI docs after the runtime and dev loop ship
- **THEN** they see both the Go↔host bridge and `forma dev` as available

#### Scenario: Author reads bridge without GopherJS 1.21 ceiling

- **WHEN** an author reads bridge or install material after this change
- **THEN** they are not instructed that GopherJS plus a Go 1.21 `GOROOT` is the supported way to run the app in the browser

### Requirement: Guide documents the forma dev browser loop

The author guide SHALL document how to run `forma dev`, including prerequisites, host-vs-Go update semantics, and a concrete browser verification path for the reference todo demo (or equivalent interactive UI from `main.go` patterns).

#### Scenario: Author learns rebuild vs HMR

- **WHEN** an author reads CLI or install documentation for the browser loop
- **THEN** they learn that host asset changes use the web tool's HMR while Go application changes rebuild the browser Go artifact and need a page reload (or remount), not preserved in-memory Go state across rebuilds

#### Scenario: Author runs the todo demo in the browser

- **WHEN** an author follows the documented verification path after `forma sync`
- **THEN** the guide tells them how to open the running app in a browser and confirm interactive todo (or equivalent) behavior driven by Go via the live bridge

### Requirement: Guide documents the React host and text-input policy

The author guide SHALL describe the web host modules under the project web tree (registry merge, recursive renderer location), how authors add custom registry entries for third-party wrappers, and the text-input hybrid policy (ephemeral React state while focused; sanitized string callbacks on edit; value prop sync when unfocused).

#### Scenario: Author extends the registry

- **WHEN** an author needs a custom React wrapper registered beside generated widgets
- **THEN** the guide explains where to merge a hand-written registry entry without editing generated registry output

#### Scenario: Author hits text-input focus issues

- **WHEN** an author reads troubleshooting or writing-UI material about text fields
- **THEN** they learn the host hybrid policy and that domain state still lives in Go

### Requirement: Author guide covers functional helpers catalog

The author guide SHALL document that the functional helper package supports slice transforms (`Map`, `Filter`, and related helpers), fallible collect helpers that stop on the first error, and map-entry transform helpers, at the same level of discoverability as existing `Map` / `Filter` mentions.

#### Scenario: Author looks up fallible list helpers

- **WHEN** an author reads writing-UI or project-layout material about `fn`
- **THEN** they learn that fallible collect helpers and map transforms exist alongside the basic slice helpers

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

### Requirement: Guide documents semantic look and escape hatches

The author guide SHALL describe that default widget look comes from semantic design-system classes applied by the host, and that `Class` / utilities are escape hatches—not the primary authoring model.

#### Scenario: Author looks up styling

- **WHEN** an author opens writing-UI or advanced styling material
- **THEN** they learn to prefer widget props and semantic defaults, and to use class escape hatches only when needed
