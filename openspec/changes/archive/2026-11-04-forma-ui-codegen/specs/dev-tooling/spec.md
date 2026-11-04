# Spec Delta — dev-tooling

## MODIFIED Requirements

### Requirement: Init scaffold smoke check in automation

The project SHALL verify in continuous integration that building the forma CLI and running full init against an empty directory produces a Go module that compiles with `go build` for the scaffolded application entrypoint, with init running sync that installs npm dependencies, materializes default shadcn components under the web root, and generates Go widget bindings under the default codegen directory from those sources. The scaffold under test SHALL include root `generator.go` for `go generate` / `forma generate`. The automation environment SHALL provide a C toolchain sufficient for CGO when building the forma CLI, and Node.js and npm compatible with the scaffold web tree.

#### Scenario: CI init smoke

- **WHEN** automation runs on a clean checkout after CLI scaffold or codegen changes
- **THEN** it builds the forma CLI with CGO, runs init with a test module path in a temporary empty directory, sync completes npm, default shadcn materialization, and codegen during init, `generator.go` is present, and the generated application compiles with `go build` without manual fixes
