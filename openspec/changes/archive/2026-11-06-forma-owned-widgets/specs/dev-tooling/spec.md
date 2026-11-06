## MODIFIED Requirements

### Requirement: Init scaffold smoke check in automation

The project SHALL verify in continuous integration that building the forma CLI and running full init against an empty directory produces a Go module that compiles with `go build` for the scaffolded application entrypoint, with init running sync that installs npm dependencies under the web root without requiring shadcn materialization or kit Go binding generation. The automation environment SHALL provide a C toolchain sufficient for CGO when building the forma CLI, and Node.js and npm compatible with the scaffold web tree.

#### Scenario: CI init smoke

- **WHEN** automation runs on a clean checkout after CLI scaffold or owned-widgets changes
- **THEN** it builds the forma CLI with CGO, runs init with a test module path in a temporary empty directory, sync completes npm install without shadcn kit codegen, and the application compiles with `go build` without manual fixes
