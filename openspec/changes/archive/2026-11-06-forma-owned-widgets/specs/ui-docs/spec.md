## MODIFIED Requirements

### Requirement: Guide covers the P2 topic set

The author guide SHALL include dedicated material for: prerequisites and toolchain (including the supported browser Go→JS/TS compiler aligned with the module's Go version); CLI onboarding (`init`, `sync`, and available `dev`); `forma.yaml` and project layout; writing UI (state, mount, refresh, dot imports, `fn`); using owned widgets and the styling stack; custom Go and third-party React extension; troubleshooting.

#### Scenario: Author adds a shadcn-backed widget

- **WHEN** an author needs Button, Input, Checkbox, or another todo-sized owned widget
- **THEN** the guide explains importing the framework UI package and does not prescribe shadcn `forma add` + kit codegen as the primary path

#### Scenario: Author finds browser compiler prerequisites

- **WHEN** an author opens the install or prerequisites material
- **THEN** they see the supported browser compiler for the current Go module version, not an EOL Go 1.21-only GopherJS path presented as the durable install story

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

## ADDED Requirements

### Requirement: Guide documents semantic look and escape hatches

The author guide SHALL describe that default widget look comes from semantic design-system classes applied by the host, and that `Class` / utilities are escape hatches—not the primary authoring model.

#### Scenario: Author looks up styling

- **WHEN** an author opens writing-UI or advanced styling material
- **THEN** they learn to prefer widget props and semantic defaults, and to use class escape hatches only when needed
