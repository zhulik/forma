# forma-cli Specification

## Purpose

Specify the forma command-line interface and project configuration authors use to create and maintain apps. The **`cmd/forma`** binary implements init (including config-only mode), stub sync, and version; add, full sync, and dev remain planned for later phases.

## Requirements

### Requirement: Init scaffolds module and syncs

The CLI SHALL provide `forma init <module>` that runs `go mod init` with the given module path in an empty target directory, writes a project scaffold from the selected embedded template (including `forma.yaml`, application entry, and imports of the forma library), and runs `forma sync` before exiting successfully. If a step fails after earlier steps wrote files, the command MAY leave a partially initialized directory; the CLI SHALL NOT claim success unless init and sync both complete.

#### Scenario: New project in empty directory

- **WHEN** an author runs `forma init github.com/user/myapp` in an empty directory
- **THEN** a Go module exists with that path, scaffold files from the template are present, and sync has applied at least module dependency reconciliation for the forma library

### Requirement: Full init requires empty target directory

For a full project scaffold (without `--only-config`), the CLI SHALL require the init target directory to contain no files or subdirectories other than `.` and `..`, and SHALL fail before writing when the directory is not empty.

#### Scenario: Init in non-empty directory

- **WHEN** an author runs full `forma init` in a directory that already contains files
- **THEN** the command fails with a clear error and does not claim success

#### Scenario: Init in empty directory

- **WHEN** an author runs full `forma init` in an empty directory
- **THEN** the command proceeds with module initialization and scaffold writes

### Requirement: Init supports config-only mode

The CLI SHALL provide `forma init --only-config` that writes `forma.yaml` when that file does not yet exist in the target directory and SHALL fail when `forma.yaml` already exists. Config-only init SHALL NOT run `go mod init`, SHALL NOT write application source files, and SHALL NOT run sync.

#### Scenario: Config only when missing

- **WHEN** an author runs `forma init --only-config` in a directory without `forma.yaml`
- **THEN** `forma.yaml` is created with default project fields

#### Scenario: Config only when present

- **WHEN** an author runs `forma init --only-config` and `forma.yaml` already exists
- **THEN** the command fails without overwriting the file

### Requirement: Init selects embedded project template

Full init SHALL accept a template name (default `minimal`) that selects an embedded scaffold. The CLI SHALL render template files with standard Go text templating and SHALL ship at least the `minimal` template in the first release that implements init.

#### Scenario: Default template

- **WHEN** an author runs full init without specifying a template
- **THEN** the `minimal` template is applied

#### Scenario: Unknown template

- **WHEN** an author specifies a template name that does not exist
- **THEN** the command fails before mutating the target directory

### Requirement: CLI reports version

The CLI SHALL provide a version string via `--version` / `-v` and a `forma version` command that print the same value derived from build metadata.

#### Scenario: User checks version

- **WHEN** an author runs `forma version`
- **THEN** they receive a non-empty version or development identifier string

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
