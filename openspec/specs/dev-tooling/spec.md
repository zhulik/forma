# dev-tooling Specification

## Purpose

Define the engineering gates and contributor conventions forma uses so local development and automation enforce the same quality bar before product features land.

## Requirements

### Requirement: Default local check runs static analysis

The project SHALL provide a single default developer command that runs static analysis on all Go packages.

#### Scenario: Contributor runs default check

- **WHEN** a contributor runs the project's default task target from the repository root
- **THEN** static analysis completes with zero reported issues

### Requirement: Static analysis rejects issues

The project SHALL treat static analysis findings as blocking: no linter category may report unresolved issues on committed Go code in the main branch workflow.

#### Scenario: Lint clean tree

- **WHEN** static analysis runs on a clean checkout with no uncommitted Go changes
- **THEN** the process exits successfully with zero findings

### Requirement: Generated output must be committed

When code generation is configured, the project SHALL verify that running generation produces no diff against the committed tree.

#### Scenario: Generate check in automation

- **WHEN** the automation job runs the generate-and-compare step on a clean checkout
- **THEN** the working tree remains unchanged after generation completes

### Requirement: Modernizer fix output must be committed

The project SHALL verify that running the configured Go modernizer pass produces no diff against the committed tree.

#### Scenario: Fix check in automation

- **WHEN** the automation job runs the fix-and-compare step on a clean checkout
- **THEN** the working tree remains unchanged after the modernizer pass completes

### Requirement: Continuous integration mirrors local lint gates

The repository SHALL run, on every push, the same generate check, static analysis, and fix check that contributors are expected to run locally before merge.

#### Scenario: Push triggers backend CI

- **WHEN** a change is pushed to any branch
- **THEN** automation runs generate verification, static analysis, and fix verification
- **AND** the workflow fails if any step fails

### Requirement: Init scaffold smoke check in automation

The project SHALL verify in continuous integration that building the forma CLI and running full init against an empty directory produces a Go module that compiles with `go build` for the scaffolded application entrypoint, with init running sync that installs npm dependencies, materializes default shadcn components under the web root, and generates Go widget bindings under the default codegen directory from those sources. The scaffold under test SHALL include root `generator.go` for `go generate` / `forma generate`. The automation environment SHALL provide a C toolchain sufficient for CGO when building the forma CLI, and Node.js and npm compatible with the scaffold web tree.

#### Scenario: CI init smoke

- **WHEN** automation runs on a clean checkout after CLI scaffold or codegen changes
- **THEN** it builds the forma CLI with CGO, runs init with a test module path in a temporary empty directory, sync completes npm, default shadcn materialization, and codegen during init, `generator.go` is present, and the generated application compiles with `go build` without manual fixes

### Requirement: Examples todo module smoke in automation

The project SHALL verify in continuous integration (or an equivalent Task entry wired into CI) that the committed `examples/todo` Go module builds successfully against the local forma replace path, and that the example’s web tree can run its package test and production build scripts after dependencies are present (matching the spirit of the init-scaffold web checks without re-running `forma init` into a temp dir).

#### Scenario: Example Go entrypoint compiles

- **WHEN** CI runs the examples-todo smoke check
- **THEN** `go build` for the `examples/todo` module entrypoint succeeds with the module’s committed `replace` of the forma library to the repository root

#### Scenario: Example web host builds

- **WHEN** CI runs the examples-todo smoke check with Node available
- **THEN** the example’s `web/` tree completes its configured unit test and build commands (after `npm` dependencies are installed if not already present in the check)

### Requirement: Agent-oriented Go rules are present

The repository SHALL include Cursor rules covering early-stage compatibility policy, Go error handling, testing conventions (for when tests are added later), file/type organization, and preferential use of the functional helper package over raw loops, so automated edits follow the same standards as sister zhulik Go projects. The prefer-fn rule MUST forbid raw `for` loops except an explicit allowlist (helper package internals, generated code, table-driven tests, side-effect-only each walks, non-slice/non-map iteration such as channels, indexed count APIs, short-circuit search, and index-heavy algorithms), and MUST forbid no-op proxy lambdas when a named function already matches a helper callback signature.

#### Scenario: Rules available for Go edits

- **WHEN** an agent or contributor opens Go source under this repository
- **THEN** applicable Cursor rules document early-stage breaking-change policy, error variable naming, one-type-one-file expectations, and the prefer-fn allowlist plus no-op-proxy ban

#### Scenario: Agent considers a slice transform loop

- **WHEN** an agent is about to write a raw `for` loop that only maps or filters a slice into a new collection
- **THEN** the prefer-fn rule directs them to use the functional helper package instead unless the loop matches an allowlisted case

### Requirement: Prefer reuse over hand-roll rule is present

The repository SHALL include an always-applied Cursor rule that prefers well-known third-party modules, then existing repository codegen, then hand-written code—before planning or writing new solving code. Package `fn` MUST remain exempt from third-party replacement.

#### Scenario: Agent about to invent a parser

- **WHEN** an agent is about to plan or implement a problem already solved by a well-known third-party module (and the need is not in the `fn` domain)
- **THEN** the prefer-reuse rule directs them to use that module instead of a parallel in-tree implementation

#### Scenario: Agent about to hand-maintain generated twin

- **WHEN** an agent is about to hand-write code that an existing repository generator is contracted to emit
- **THEN** the prefer-reuse rule directs them to extend or run that generator instead

### Requirement: Prefer-reuse rule covers debt proposals

The prefer-reuse Cursor rule SHALL require debt-style or planning reviews to propose replacing hand-rolled code that a well-known module or existing codegen should own.

#### Scenario: Debt finds NIH util

- **WHEN** a debt or planning pass finds hand-rolled code that a well-known module or existing codegen should own
- **THEN** the prefer-reuse rule requires proposing that replacement rather than leaving NIH as the default
