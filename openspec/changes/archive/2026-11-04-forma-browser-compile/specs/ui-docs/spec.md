# Spec Delta — ui-docs

## MODIFIED Requirements

### Requirement: Guide covers the P2 topic set

The author guide SHALL include dedicated material for: prerequisites and toolchain (including the supported browser Go→JS/TS compiler aligned with the module's Go version); CLI onboarding (`init`, `add`, `sync`, and available `dev`); `forma.yaml` and project layout; writing UI (state, mount, refresh, dot imports, `fn`); adding components; custom Go and third-party React extension; troubleshooting.

#### Scenario: Author adds a shadcn-backed widget

- **WHEN** an author needs to add a new visual component to their app
- **THEN** the guide explains `forma add`, the role of `forma sync`, and what belongs in version control

#### Scenario: Author finds browser compiler prerequisites

- **WHEN** an author opens the install or prerequisites material
- **THEN** they see the supported browser compiler for the current Go module version, not an EOL Go 1.21-only GopherJS path presented as the durable install story

### Requirement: Guide documents the Go↔host bridge

The author guide SHALL describe how the browser runtime mounts a root component, refreshes the host tree from Go state, and delivers sanitized callbacks into handlers with callback context. It SHALL describe `forma dev` as the available one-command browser rebuild loop. It SHALL describe the browser compile/rebuild story in terms of the project's supported compiler and SHALL NOT require authors to install GopherJS against Go 1.21 as the supported path.

#### Scenario: Author looks up RefreshUI behavior

- **WHEN** an author reads writing-UI or troubleshooting material after the runtime ships
- **THEN** they see that mutations plus refresh (or the update helper) drive the host tree from Go

#### Scenario: Author distinguishes bridge from forma dev

- **WHEN** an author reads status or CLI docs after the runtime and dev loop ship
- **THEN** they see both the Go↔host bridge and `forma dev` as available

#### Scenario: Author reads bridge without GopherJS 1.21 ceiling

- **WHEN** an author reads bridge or install material after this change
- **THEN** they are not instructed that GopherJS plus a Go 1.21 `GOROOT` is the supported way to run the app in the browser
