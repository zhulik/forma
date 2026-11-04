# Spec Delta

## MODIFIED Requirements

### Requirement: Dev command is reserved for the browser loop

The CLI SHALL provide `forma dev` that runs the local browser development loop: a long-lived process that starts the web-tree development server and rebuilds the application’s Go entrypoint for the browser whenever relevant Go sources change. The command SHALL read `forma.yaml` for the web root and mount configuration. It SHALL assume `forma sync` has been run at least once; when required web dependencies or toolchain binaries are missing, it SHALL fail with a clear error that points authors at sync and documented install steps rather than claiming the loop started. On successful start it SHALL keep running until interrupted, forwarding child-process output so authors can see the host URL and rebuild status. Documentation SHALL describe `forma dev` as available (not planned).

#### Scenario: Documentation before P3d

- **WHEN** an author reads CLI reference material that previously treated `forma dev` as planned-only
- **THEN** after this change that material lists `forma dev` as available alongside implemented commands such as `init`, `generate`, `sync`, `add`, and `version`

#### Scenario: Author starts the browser loop

- **WHEN** an author runs `forma dev` in a project with a readable `forma.yaml`, installed web dependencies, and the GopherJS toolchain available
- **THEN** the web development server starts and an initial browser-targeted Go build completes successfully without the command exiting on its own

#### Scenario: Go source change triggers rebuild

- **WHEN** `forma dev` is running and the author saves a Go source file that belongs to the app’s browser build
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
