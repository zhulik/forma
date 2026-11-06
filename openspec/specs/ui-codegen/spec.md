# ui-codegen Specification

## Purpose

Define how forma turns committed shadcn component sources in an application web tree into typed Go widget bindings and a generated JavaScript component registry, deterministically and verifiably in CI.

## Requirements

### Requirement: Kit codegen is not the product widget path

The product SHALL NOT require authors to run a shadcn (or other kit) TSX→Go codegen pipeline to obtain the todo-sized widget surface. Any remaining kit-codegen implementation in the repository MUST NOT be invoked by default `init` / `sync` and MUST NOT be documented as the primary way to add core widgets.

#### Scenario: Fresh init has no kit-codegen step for core widgets

- **WHEN** an author runs full project init successfully
- **THEN** the resulting app can import and render owned core widgets without a generate-from-kit step having produced application bindings

#### Scenario: Docs do not prescribe kit codegen for Button

- **WHEN** an author follows the user guide to use a Button
- **THEN** the prescribed path is the framework UI package, not regenerating bindings from a kit source file
