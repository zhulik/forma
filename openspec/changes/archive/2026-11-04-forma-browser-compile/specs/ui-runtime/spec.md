# Spec Delta — ui-runtime

## ADDED Requirements

### Requirement: Browser build preserves object-graph tree updates

When the application is built for the browser target, the refresh path SHALL deliver the render tree to the host seam as a host-callable value graph (nodes, props, and function callbacks) and SHALL NOT require JSON-encoding the full render tree as the sole supported update mechanism. Domain state remains in Go; the host remains a dumb renderer.

#### Scenario: Refresh without mandatory full-tree JSON

- **WHEN** a browser-built app refreshes after a state change
- **THEN** the host receives an updated tree through the tree-update seam without the runtime mandating that the entire tree be serialized as JSON for that refresh

#### Scenario: Callbacks remain host-invocable functions

- **WHEN** the refreshed tree includes author handlers
- **THEN** the host can invoke those handlers as callbacks with sanitized plain payloads (no DOM event objects required)

## MODIFIED Requirements

### Requirement: Native builds keep a non-browser stub

On non-browser builds, the mount and refresh entry points MUST remain safe no-ops (or test doubles) so ordinary native compile and test of application packages continue to succeed without a DOM, host seam, or browser-only compiler toolchain.

#### Scenario: Native compile of the reference app

- **WHEN** the reference application is built for a non-browser target
- **THEN** the mount entry compiles and does not require a live host or a browser Go→JS compiler on PATH
