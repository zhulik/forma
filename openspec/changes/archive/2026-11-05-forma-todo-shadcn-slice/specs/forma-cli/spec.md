## MODIFIED Requirements

### Requirement: Init scaffolds module and syncs

The CLI SHALL provide `forma init <module>` that runs `go mod init` with the given module path in an empty target directory, writes a project scaffold from the selected embedded template (including `forma.yaml`, application entry, imports of the forma library, and a web tree under the configured `web.path` with package manifest and shadcn configuration), and runs `forma sync` before exiting successfully. If a step fails after earlier steps wrote files, the command MAY leave a partially initialized directory; the CLI SHALL NOT claim success unless init and sync both complete. First sync on a new project SHALL materialize a default starter set of shadcn components defined by the CLI (button, input, card, and checkbox) via the shadcn CLI even when those names are not listed in `forma.yaml`.

#### Scenario: New project in empty directory

- **WHEN** an author runs `forma init github.com/user/myapp` in an empty directory
- **THEN** a Go module exists with that path, scaffold files from the template are present including the web tree tooling under `web.path`, and sync has installed npm dependencies, materialized the default starter shadcn components under the web root, generated Go bindings and the host registry from those sources, and reconciled Go module dependencies for the forma library

#### Scenario: Starter set includes checkbox

- **WHEN** init’s first sync materializes the default starter set
- **THEN** the web components tree includes button, input, card, and checkbox sources and codegen emits bindings for each
