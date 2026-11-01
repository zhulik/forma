# Spec Delta

## MODIFIED Requirements

### Requirement: Add registers a component and syncs

The CLI SHALL provide `forma add <component>` that adds the component name to the project's declared component list in `forma.yaml` (if not already present) and runs `forma sync`.

#### Scenario: Add dialog

- **WHEN** an author runs `forma add dialog`
- **THEN** `dialog` appears in the configured component list and sync materializes shadcn sources under the configured web root and generated bindings for that component

### Requirement: Sync materializes full project state

The CLI SHALL provide `forma sync` that, from `forma.yaml`, ensures dependencies are installed under the configured web root, shadcn component sources exist for each listed component, Go widget bindings are generated into the configured codegen directory, the JavaScript component registry under the web root reflects generated types, and `go mod tidy` runs so forma library imports resolve. Until the web tree and codegen pipeline ships, sync SHALL still run `go mod tidy` and SHALL ensure the forma library module requirement is present; documentation and command output SHALL distinguish stub sync from full materialization.

#### Scenario: Fresh clone after init

- **WHEN** an author runs `forma sync` on a clean checkout with a valid `forma.yaml`
- **THEN** web tree install and codegen complete without manual shadcn CLI steps and Go module dependencies are tidied

#### Scenario: Stub sync after init before P3 host

- **WHEN** an author runs `forma sync` before the web tree and codegen are implemented
- **THEN** Go module dependencies including the forma library are reconciled and the command exits successfully without requiring manual npm or shadcn steps

### Requirement: forma.yaml declares UI desired state

Each forma application SHALL include a root `forma.yaml` that declares at minimum: the UI mount element id; the list of shadcn component names to include; web root path (`web.path`); codegen output directory and Go package name for generated widgets.

#### Scenario: Default scaffold

- **WHEN** init completes on a new project
- **THEN** `forma.yaml` includes a default codegen directory of `internal/ui` and a default web root path of `web` unless the author overrides them before sync

### Requirement: Dev command is reserved for the browser loop

The CLI SHALL eventually provide `forma dev` to run the local browser development loop (GopherJS rebuild and dev server for the web tree). Until that command is implemented, documentation SHALL describe it as planned and SHALL NOT claim it is available.

#### Scenario: Documentation before P3d

- **WHEN** an author reads CLI reference material during P3d planning
- **THEN** `forma dev` is listed with planned status distinct from implemented commands such as `init`, stub `sync`, and `version`
