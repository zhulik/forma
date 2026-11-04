# ui-codegen Specification

## Purpose

Define how forma turns committed shadcn component sources in an application web tree into typed Go widget bindings and a generated JavaScript component registry, deterministically and verifiably in CI.

## Requirements

### Requirement: Codegen reads local shadcn component sources

The codegen pipeline SHALL treat each eligible shadcn component source file under the configured web root `src/components/ui/` directory as input and SHALL derive widget metadata from that source without requiring a parallel component name list in project configuration.

#### Scenario: Widget present on disk

- **WHEN** a shadcn button source file exists under the configured web components directory after sync, add, or generate
- **THEN** codegen emits a corresponding Go binding file under the configured codegen directory

#### Scenario: Widget removed from disk

- **WHEN** an author deletes a shadcn source file and runs generate or sync
- **THEN** codegen removes or stops emitting the Go binding that corresponded to that source so generated output matches the TSX set

### Requirement: Generated Go bindings match DSL leaf and compound rules

Generated Go widget types SHALL implement the application component interface, use value-receiver composition helpers for leaf widgets where the DSL requires primary children, use field-based slots for compound widgets, and set render tree type names that match registry entries.

#### Scenario: Leaf widget

- **WHEN** codegen processes a simple shadcn leaf export such as a button
- **THEN** the emitted struct supports primary-child composition and renders a node whose type identifies the button widget

#### Scenario: Compound widget

- **WHEN** codegen processes a shadcn export graph with a root and prefixed part exports
- **THEN** the emitted struct exposes one field per author-facing slot region and renders named slots rather than separate Go types per part export

### Requirement: Single generated host registry module

Codegen SHALL emit one generated registry module under the configured web root that maps render tree type strings to the React components required to render generated widgets, and SHALL regenerate that module wholesale on each successful codegen run.

#### Scenario: Registry after generate

- **WHEN** generate or sync completes codegen for a project with at least one shadcn source file
- **THEN** a generated registry module exists under the web root and lists every generated widget type present in the Go bindings

### Requirement: Codegen output is deterministic

Running codegen twice against an unchanged web components tree and configuration SHALL produce byte-identical generated Go and registry outputs.

#### Scenario: Repeat generation

- **WHEN** an author runs generate twice without changing shadcn sources or codegen settings
- **THEN** the second run makes no diff to committed generated files

### Requirement: Per-component override hook

The pipeline MAY consult optional per-source override data when automatic classification of props, events, or compound slots fails, and SHALL document the override format for maintainers.

#### Scenario: Override applied

- **WHEN** a maintainer adds an override for a source file that automatic classification cannot handle
- **THEN** codegen emits Go and registry entries consistent with the override without hand-editing generated files

### Requirement: Registry slot keys match Go wire keys

For compound widgets, the generated host registry slot map keys SHALL match the wire slot keys written into the render tree for the corresponding Go slot fields (the same mapping used when emitting named slots from generated Go bindings). Registry slot keys MUST NOT use a separate capitalization scheme from those wire keys.

#### Scenario: Compound registry keys align with tree slots

- **WHEN** codegen emits a compound widget with slot fields and a registry entry that lists slot components
- **THEN** each registry slot map key equals the wire key used for that field in the generated Go render output

#### Scenario: Accordion-style parts stay consistent

- **WHEN** codegen processes a compound with trigger/content-style parts
- **THEN** the registry slot keys and the Go `Slots` map keys for those parts are identical strings

### Requirement: Committed output is verifiable by regenerate and git diff

Projects SHALL be able to verify that committed generated Go and registry files match current shadcn sources by regenerating via `forma generate` (including through `go generate`) and confirming the whole repository has no git diff, without requiring a dedicated forma check flag.

#### Scenario: Clean tree in CI

- **WHEN** automation regenerates on a clean checkout whose generated files match the current shadcn sources and then runs a whole-repo git diff check
- **THEN** the check exits successfully with no file changes

#### Scenario: Stale generated files

- **WHEN** generated Go or registry files are out of date relative to committed shadcn sources and automation regenerates then runs a whole-repo git diff check
- **THEN** the check fails with a non-zero exit status
