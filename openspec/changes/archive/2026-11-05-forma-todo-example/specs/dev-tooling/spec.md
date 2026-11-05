# Spec Delta

## ADDED Requirements

### Requirement: Examples todo module smoke in automation

The project SHALL verify in continuous integration (or an equivalent Task entry wired into CI) that the committed `examples/todo` Go module builds successfully against the local forma replace path, and that the example’s web tree can run its package test and production build scripts after dependencies are present (matching the spirit of the init-scaffold web checks without re-running `forma init` into a temp dir).

#### Scenario: Example Go entrypoint compiles

- **WHEN** CI runs the examples-todo smoke check
- **THEN** `go build` for the `examples/todo` module entrypoint succeeds with the module’s committed `replace` of the forma library to the repository root

#### Scenario: Example web host builds

- **WHEN** CI runs the examples-todo smoke check with Node available
- **THEN** the example’s `web/` tree completes its configured unit test and build commands (after `npm` dependencies are installed if not already present in the check)
