# Spec Delta — dev-tooling

## MODIFIED Requirements

### Requirement: Init scaffold smoke check in automation

The project SHALL verify in continuous integration that building the forma CLI and running full init against an empty directory produces a Go module that compiles with `go build` for the scaffolded application entrypoint, with init running sync that installs npm dependencies and materializes default shadcn components under the web root. The automation environment SHALL provide Node.js and npm compatible with the scaffold web tree.

#### Scenario: CI init smoke

- **WHEN** automation runs on a clean checkout after CLI scaffold changes
- **THEN** it builds the forma CLI, runs init with a test module path in a temporary empty directory, sync completes npm and shadcn steps during init, and the generated application compiles with `go build` without manual fixes
