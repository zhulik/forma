# forma-cli Specification

## Purpose

Specify the forma command-line interface and project configuration authors use to create and maintain apps. The **`cmd/forma`** binary implements init (including config-only mode), sync, generate, add, version, and the browser dev loop (Vite host + supported browser compiler rebuild).

## Requirements

### Requirement: Init scaffolds module and syncs

The CLI SHALL provide `forma init <module>` that runs `go mod init` with the given module path in an empty target directory, writes a project scaffold from the selected embedded template (including `forma.yaml`, application entry, imports of the forma library, and a web tree under the configured `web.path` with package manifest and styling assets for the owned-widgets stack), and runs `forma sync` before exiting successfully. If a step fails after earlier steps wrote files, the command MAY leave a partially initialized directory; the CLI SHALL NOT claim success unless init and sync both complete. Init/sync MUST NOT materialize shadcn component sources or generate kit Go bindings for the default widget surface.

#### Scenario: New project in empty directory

- **WHEN** an author runs `forma init github.com/user/myapp` in an empty directory
- **THEN** a Go module exists with that path, scaffold files from the template are present including the web tree under `web.path`, sync has installed npm dependencies and reconciled Go module dependencies for the forma library, and the app can use framework owned widgets without shadcn materialization or kit codegen

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

### Requirement: Sync materializes full project state

The CLI SHALL provide `forma sync` that, from `forma.yaml`, ensures dependencies are installed under the configured web root and runs `go mod tidy` so forma library imports resolve. Sync MUST NOT invoke shadcn materialization and MUST NOT run kit TSX→Go codegen as part of the default product path.

#### Scenario: Fresh clone after init

- **WHEN** an author runs `forma sync` on a clean checkout with a valid `forma.yaml` and committed web tree
- **THEN** npm dependencies are installed under the web root and Go module dependencies are tidied

#### Scenario: Sync without Node tooling

- **WHEN** an author runs `forma sync` and neither `npm` nor an equivalent install command is available on PATH
- **THEN** the command fails with a clear error before claiming success

#### Scenario: Sync runs codegen

- **WHEN** an author runs `forma sync` on a project without kit component sources
- **THEN** sync completes successfully without running kit TSX→Go codegen or claiming that a codegen step was required

### Requirement: forma.yaml declares UI desired state

Each forma application SHALL include a root `forma.yaml` that declares at minimum: the UI mount element id and web root path (`web.path`). Default scaffold MUST NOT require a codegen output directory for owned core widgets.

#### Scenario: Default scaffold

- **WHEN** init completes on a new project
- **THEN** `forma.yaml` includes a default web root path of `web` and a default UI mount id, and does not require a kit component name list or codegen directory for sync to succeed

### Requirement: Dev command is reserved for the browser loop

The CLI SHALL provide `forma dev` that runs the local browser development loop: a long-lived process that starts the web-tree development server and rebuilds the application's Go entrypoint for the browser with the project's supported browser compiler whenever relevant Go sources change. The command SHALL read `forma.yaml` for the web root and mount configuration. It SHALL assume `forma sync` has been run at least once; when required web dependencies or toolchain binaries are missing, it SHALL fail with a clear error that points authors at sync and documented install steps rather than claiming the loop started. On successful start it SHALL keep running until interrupted, forwarding child-process output so authors can see the host URL and rebuild status. Documentation SHALL describe `forma dev` as available (not planned). Documentation and error messages for the loop SHALL NOT require GopherJS or an EOL Go 1.21 toolchain as the supported browser compile path. After owned-widgets migration, CLI docs MUST present `init`, `sync`, `version`, and `dev` as the primary commands and MUST NOT present shadcn `add`/`generate` as the primary widget path.

#### Scenario: Documentation before P3d

- **WHEN** an author reads CLI reference material that previously treated `forma dev` as planned-only
- **THEN** that material lists `forma dev` as available alongside implemented product commands such as `init`, `sync`, and `version`

#### Scenario: Author starts the browser loop

- **WHEN** an author runs `forma dev` in a project with a readable `forma.yaml`, installed web dependencies, and the supported browser compiler available
- **THEN** the web development server starts and an initial browser-targeted Go build completes successfully without the command exiting on its own

#### Scenario: Dev loop contract is compiler-agnostic

- **WHEN** an author reads CLI material describing the available browser loop
- **THEN** the rebuild step is described as using the project's supported browser compiler for the module's Go version, not as a GopherJS-only contract

#### Scenario: Missing browser compiler fails clearly

- **WHEN** `forma dev` runs and the supported browser compiler is missing from the environment
- **THEN** the command fails with a clear error pointing at documented install steps rather than claiming the loop started

#### Scenario: Go source change triggers rebuild

- **WHEN** `forma dev` is running and the author saves a Go source file that belongs to the app's browser build
- **THEN** the CLI rebuilds the browser Go artifact and the running loop remains active so the author can reload and see the update

#### Scenario: Dev without forma.yaml

- **WHEN** an author runs `forma dev` in a directory without a readable `forma.yaml`
- **THEN** the command fails with a clear error and does not start a host server claimed as successful

#### Scenario: Dev without synced web dependencies

- **WHEN** an author runs `forma dev` and the configured web root has no installed npm dependencies
- **THEN** the command fails with a clear error directing the author to run `forma sync` (or equivalent install) before retrying

#### Scenario: Documentation after P3d

- **WHEN** an author reads CLI reference material after owned-widgets migration ships
- **THEN** `forma dev` is listed as available alongside `init`, `sync`, and `version`, without presenting shadcn `add`/`generate` as the primary widget path
