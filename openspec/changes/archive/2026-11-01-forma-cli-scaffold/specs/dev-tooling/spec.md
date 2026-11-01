# Spec Delta

## ADDED Requirements

### Requirement: Init scaffold smoke check in automation

The project SHALL verify in continuous integration that building the forma CLI and running full init against an empty directory produces a Go module that compiles with `go build` for the scaffolded application entrypoint.

#### Scenario: CI init smoke

- **WHEN** automation runs on a clean checkout after CLI scaffold changes
- **THEN** it builds the forma CLI, runs init with a test module path in a temporary empty directory, and succeeds compiling the generated application without manual fixes
