## REMOVED Requirements

### Requirement: Codegen reads local shadcn component sources

**Reason:** ROADMAP P5 abandons shadcn TSX→Go as the widget source of truth; owned hand-written widgets replace kit bindings for the product path.

**Migration:** Use framework `ui` owned widgets. Do not add shadcn sources under `web/src/components/ui` expecting `forma generate` to emit app bindings. Historical codegen may remain only as non-product parked code until deleted.

### Requirement: Generated Go bindings match DSL leaf and compound rules

**Reason:** Core leaf/layout widgets are hand-maintained; catalog compounds are no longer a codegen product deliverable.

**Migration:** Hand-write Go widget types following `ui-dsl` / `ui-widgets`. Future compounds are owned widgets, not generated shadcn graphs.

### Requirement: Single generated host registry module

**Reason:** Todo-sized surface registers in framework builtins; apps no longer require `registry.gen.ts` from shadcn for core widgets.

**Migration:** Rely on scaffold builtin registry (plus optional hand-written custom registry entries for third-party islands).

### Requirement: Codegen output is deterministic

**Reason:** Author-facing generate path is removed from the product CLI surface.

**Migration:** None for authors. Maintainers deleting parked codegen remove its determinism tests with the package.

### Requirement: Per-component override hook

**Reason:** Locked direction forbids forma-shipped per-widget YAML overrides as catalog SoT; the override hook existed to paper over codegen failures.

**Migration:** Adjust hand-written Go/host pairs in reviewable PRs. Do not add override YAML matrices.

### Requirement: Registry slot keys match Go wire keys

**Reason:** Tied to generated compound registry emission; no longer a product codegen requirement.

**Migration:** When owned compounds exist later, keep slot wire keys aligned between Go `Node.Slots` and host slot maps by hand (same rule as today for Fragment/slots DSL).

### Requirement: Committed output is verifiable by regenerate and git diff

**Reason:** Scaffold no longer commits generated kit bindings as the sync contract.

**Migration:** CI verifies init/build and host tests; drop generate-and-diff against shadcn sources.

## ADDED Requirements

### Requirement: Kit codegen is not the product widget path

The product SHALL NOT require authors to run a shadcn (or other kit) TSX→Go codegen pipeline to obtain the todo-sized widget surface. Any remaining kit-codegen implementation in the repository MUST NOT be invoked by default `init` / `sync` and MUST NOT be documented as the primary way to add core widgets.

#### Scenario: Fresh init has no kit-codegen step for core widgets

- **WHEN** an author runs full project init successfully
- **THEN** the resulting app can import and render owned core widgets without a generate-from-kit step having produced application bindings

#### Scenario: Docs do not prescribe kit codegen for Button

- **WHEN** an author follows the user guide to use a Button
- **THEN** the prescribed path is the framework UI package, not regenerating bindings from a kit source file
