# Spec Delta

## Purpose

Specify the target forma command-line interface and project configuration file authors will use to create and maintain apps. This change documents the contract only; the binary is implemented in a later roadmap phase.

## ADDED Requirements

### Requirement: Init scaffolds module and syncs

The CLI SHALL provide `forma init <module>` that runs `go mod init` with the given module path, writes a minimal project scaffold (including `forma.yaml`, application entry, host tree placeholder, and imports of the forma library), and runs `forma sync` before exiting successfully.

#### Scenario: New project in empty directory

- **WHEN** an author runs `forma init github.com/user/myapp` in an empty directory
- **THEN** a Go module exists with that path, scaffold files are present, and sync has applied the default configuration

### Requirement: Add registers a component and syncs

The CLI SHALL provide `forma add <component>` that adds the component name to the project's declared component list in `forma.yaml` (if not already present) and runs `forma sync`.

#### Scenario: Add dialog

- **WHEN** an author runs `forma add dialog`
- **THEN** `dialog` appears in the configured component list and sync materializes host sources and generated bindings for that component

### Requirement: Sync materializes full project state

The CLI SHALL provide `forma sync` that, from `forma.yaml`, ensures host dependencies are installed, shadcn component sources exist for each listed component, Go widget bindings are generated into the configured codegen directory, the host component registry reflects generated types, and `go mod tidy` runs so forma library imports in the scaffold resolve.

#### Scenario: Fresh clone after init

- **WHEN** an author runs `forma sync` on a clean checkout with a valid `forma.yaml`
- **THEN** host install and codegen complete without manual shadcn CLI steps and Go module dependencies are tidied

### Requirement: forma.yaml declares UI desired state

Each forma application SHALL include a root `forma.yaml` that declares at minimum: the UI mount element id; the list of shadcn component names to include; host root path; codegen output directory and Go package name for generated widgets.

#### Scenario: Default scaffold

- **WHEN** init completes on a new project
- **THEN** `forma.yaml` includes a default codegen directory of `internal/ui` unless the author overrides it before sync

### Requirement: Dev command is reserved for the browser loop

The CLI SHALL eventually provide `forma dev` to run the local browser development loop (GopherJS rebuild and host dev server). Until that command is implemented, documentation SHALL describe it as planned and SHALL NOT claim it is available.

#### Scenario: Documentation before P3d

- **WHEN** an author reads CLI reference material during P2
- **THEN** `forma dev` is listed with planned status distinct from `init`, `add`, and `sync`
