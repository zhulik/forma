# Spec Delta

## ADDED Requirements

### Requirement: Manual project initialization is documented

The author guide SHALL document how to create a forma app without the CLI: initializing the Go module, authoring `forma.yaml`, adding a minimal `main.go` that uses the forma library, and running `go mod tidy`, with clear pointers to when `forma sync` and host setup apply later.

#### Scenario: Author prefers manual setup

- **WHEN** an author cannot or will not use `forma init`
- **THEN** the guide provides step-by-step manual setup that matches the fields and layout produced by init

## MODIFIED Requirements

### Requirement: Implementation status is explicit

Where documented behavior depends on work not yet shipped (browser dev loop, full sync with host/codegen), the guide SHALL label that material as planned or not yet available and SHALL NOT imply a working browser app when only the DSL stub exists in the framework repository. The guide SHALL state which CLI commands are available in the current release (`init`, config-only init, version) versus planned (`add`, full `sync`, `dev`).

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page before P3d lands
- **THEN** they see which steps work today (e.g. DSL reference via `main.go`, CLI init where shipped) versus which require a future release

#### Scenario: Author expects full sync today

- **WHEN** an author reads sync documentation before host and codegen ship
- **THEN** they see that sync may only reconcile Go module dependencies until full materialization is implemented
