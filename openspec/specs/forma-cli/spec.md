# forma-cli Specification

## Purpose

Specify the forma command-line interface and project configuration authors use to create and maintain apps. The **`cmd/forma`** binary implements init (including config-only mode), sync, generate, add, version, and the browser dev loop (Vite host + GopherJS rebuild).

## Requirements

### Requirement: Init scaffolds module and syncs

The CLI SHALL provide `forma init <module>` that runs `go mod init` with the given module path in an empty target directory, writes a project scaffold from the selected embedded template (including `forma.yaml`, application entry, imports of the forma library, and a web tree under the configured `web.path` with package manifest and shadcn configuration), and runs `forma sync` before exiting successfully. If a step fails after earlier steps wrote files, the command MAY leave a partially initialized directory; the CLI SHALL NOT claim success unless init and sync both complete. First sync on a new project SHALL materialize a default starter set of shadcn components defined by the CLI (currently button, input, and card) via the shadcn CLI even when those names are not listed in `forma.yaml`.

#### Scenario: New project in empty directory

- **WHEN** an author runs `forma init github.com/user/myapp` in an empty directory
- **THEN** a Go module exists with that path, scaffold files from the template are present including the web tree tooling under `web.path`, and sync has installed npm dependencies, materialized the default starter shadcn components under the web root, generated Go bindings and the host registry from those sources, and reconciled Go module dependencies for the forma library

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

The CLI SHALL provide `forma add <component>` that materializes the named shadcn component under the configured web root when not already present and runs `forma sync`.

#### Scenario: Add dialog

- **WHEN** an author runs `forma add dialog`
- **THEN** shadcn sources for dialog exist under the configured web root and sync regenerates Go bindings and the host registry so dialog is available to the application

### Requirement: Sync materializes full project state

The CLI SHALL provide `forma sync` that, from `forma.yaml`, ensures dependencies are installed under the configured web root, runs the same codegen path as `forma generate` (regenerating Go widget bindings and the JavaScript component registry from shadcn sources present under the web root), and runs `go mod tidy` so forma library imports resolve. Sync SHALL NOT require a component name list in `forma.yaml` to drive codegen. Sync SHALL run shadcn materialization for the init starter set only during the init-driven first sync, and SHALL run shadcn materialization for explicitly added component names when invoked via `forma add`; ordinary sync on an existing project SHALL regenerate from committed TSX without re-invoking shadcn for every widget unless add or init materialization is in progress.

#### Scenario: Fresh clone after init

- **WHEN** an author runs `forma sync` on a clean checkout with a valid `forma.yaml` and committed web tree sources including shadcn TSX and generated bindings
- **THEN** npm dependencies are installed under the web root, Go bindings and the host registry match the committed shadcn sources, and Go module dependencies are tidied

#### Scenario: Sync without Node tooling

- **WHEN** an author runs `forma sync` and neither `npm` nor an equivalent install command is available on PATH
- **THEN** the command fails with a clear error before claiming success

#### Scenario: Sync runs codegen

- **WHEN** an author runs `forma sync` on a project with shadcn sources
- **THEN** codegen and registry update run to completion and the command does not log that those steps were skipped

### Requirement: forma.yaml declares UI desired state

Each forma application SHALL include a root `forma.yaml` that declares at minimum: the UI mount element id; web root path (`web.path`); codegen output directory and Go package name for generated widgets.

#### Scenario: Default scaffold

- **WHEN** init completes on a new project
- **THEN** `forma.yaml` includes a default codegen directory of `internal/ui`, a default web root path of `web`, and a default UI mount id, and does not require a shadcn component name list for sync to succeed

### Requirement: Generate regenerates bindings from config

The CLI SHALL provide `forma generate` that reads the project's `forma.yaml`, regenerates Go widget bindings from shadcn sources under the configured web root into the configured codegen directory, and regenerates the JavaScript component registry under the web root. Generate SHALL NOT install npm dependencies, SHALL NOT invoke the shadcn CLI, SHALL NOT run `go mod tidy`, and SHALL NOT require Node.js or npm on PATH. Generate SHALL NOT require a component name list in `forma.yaml`.

#### Scenario: Generate after editing committed TSX

- **WHEN** an author edits a committed shadcn source under the web components directory and runs `forma generate`
- **THEN** Go bindings and the host registry are updated to match the TSX without reinstalling npm packages or fetching new shadcn components

#### Scenario: Generate without Node.js

- **WHEN** an author runs `forma generate` with a valid `forma.yaml` and committed TSX while Node.js is unavailable on PATH
- **THEN** codegen still completes successfully

#### Scenario: Generate without forma.yaml

- **WHEN** an author runs `forma generate` in a directory without a readable `forma.yaml`
- **THEN** the command fails with a clear error and writes no generated output claimed as successful

### Requirement: Init scaffold includes go generate and drift check

The minimal init template SHALL include a root `generator.go` whose `go:generate` directive invokes `forma generate`, and a task target that runs `go generate` then fails when `git diff --exit-code` reports any change in the whole repository, so application CI can verify committed generated output is current without a dedicated forma check flag.

#### Scenario: Author runs check in new project

- **WHEN** an author runs the scaffold's generate check task after init on a clean tree whose generated files match current shadcn sources
- **THEN** `go generate` runs successfully and the whole-repo git diff check exits successfully

#### Scenario: Stale generated files fail check

- **WHEN** generated Go or registry files are out of date relative to committed shadcn sources and the author runs the scaffold's generate check task
- **THEN** the task fails with a non-zero exit status after regeneration leaves a dirty whole-repo git diff

### Requirement: Dev command is reserved for the browser loop

The CLI SHALL provide `forma dev` that runs the local browser development loop: a long-lived process that starts the web-tree development server and rebuilds the application's Go entrypoint for the browser whenever relevant Go sources change. The command SHALL read `forma.yaml` for the web root and mount configuration. It SHALL assume `forma sync` has been run at least once; when required web dependencies or toolchain binaries are missing, it SHALL fail with a clear error that points authors at sync and documented install steps rather than claiming the loop started. On successful start it SHALL keep running until interrupted, forwarding child-process output so authors can see the host URL and rebuild status. Documentation SHALL describe `forma dev` as available (not planned).

#### Scenario: Documentation before P3d

- **WHEN** an author reads CLI reference material that previously treated `forma dev` as planned-only
- **THEN** after this change that material lists `forma dev` as available alongside implemented commands such as `init`, `generate`, `sync`, `add`, and `version`

#### Scenario: Author starts the browser loop

- **WHEN** an author runs `forma dev` in a project with a readable `forma.yaml`, installed web dependencies, and the GopherJS toolchain available
- **THEN** the web development server starts and an initial browser-targeted Go build completes successfully without the command exiting on its own

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

- **WHEN** an author reads CLI reference material after this change ships
- **THEN** `forma dev` is listed as available alongside `init`, `generate`, `sync`, `add`, and `version`
