# Spec Delta — forma-cli

## MODIFIED Requirements

### Requirement: Init scaffolds module and syncs

The CLI SHALL provide `forma init <module>` that runs `go mod init` with the given module path in an empty target directory, writes a project scaffold from the selected embedded template (including `forma.yaml`, application entry, imports of the forma library, and a web tree under the configured `web.path` with package manifest and shadcn configuration), and runs `forma sync` before exiting successfully. If a step fails after earlier steps wrote files, the command MAY leave a partially initialized directory; the CLI SHALL NOT claim success unless init and sync both complete.

#### Scenario: New project in empty directory

- **WHEN** an author runs `forma init github.com/user/myapp` in an empty directory
- **THEN** a Go module exists with that path, scaffold files from the template are present including the web tree tooling under `web.path`, and sync has installed npm dependencies, materialized declared shadcn components, and reconciled Go module dependencies for the forma library

### Requirement: Sync materializes full project state

The CLI SHALL provide `forma sync` that, from `forma.yaml`, ensures dependencies are installed under the configured web root, shadcn component sources exist for each listed component, Go widget bindings are generated into the configured codegen directory, the JavaScript component registry under the web root reflects generated types, and `go mod tidy` runs so forma library imports resolve. Until the Go codegen and registry pipeline ships, sync SHALL still install npm dependencies under the web root, materialize shadcn sources for each listed component via the shadcn CLI, run `go mod tidy`, and ensure the forma library module requirement is present; sync SHALL log that codegen and registry steps are skipped rather than failing, and documentation SHALL distinguish partial sync from full materialization.

#### Scenario: Fresh clone after init

- **WHEN** an author runs `forma sync` on a clean checkout with a valid `forma.yaml` and committed web tree sources
- **THEN** npm dependencies are installed under the web root, shadcn component sources exist for each listed component without manual shadcn CLI steps, and Go module dependencies are tidied

#### Scenario: Stub sync after init before P3 host

- **WHEN** an author runs `forma sync` before Go widget codegen and the JS component registry are implemented
- **THEN** npm install and shadcn materialization complete successfully, Go module dependencies including the forma library are reconciled, and the command exits successfully while reporting that codegen and registry update were skipped

#### Scenario: Sync without Node tooling

- **WHEN** an author runs `forma sync` and neither `npm` nor an equivalent install command is available on PATH
- **THEN** the command fails with a clear error before claiming success

### Requirement: forma.yaml declares UI desired state

Each forma application SHALL include a root `forma.yaml` that declares at minimum: the UI mount element id; the list of shadcn component names to include; web root path (`web.path`); codegen output directory and Go package name for generated widgets.

#### Scenario: Default scaffold

- **WHEN** init completes on a new project
- **THEN** `forma.yaml` includes a default codegen directory of `internal/ui`, a default web root path of `web`, and a default starter set of shadcn component names unless the author overrides them before sync
