# Spec Delta

## MODIFIED Requirements

### Requirement: Host mount seam accepts a tree root

The scaffold host entry SHALL mount into the configured DOM mount element and render from a replaceable tree-root input. The entry MUST expose an imperative tree-update API the Go runtime can call after boot (including replacing a null initial tree). The mount path and renderer MUST remain the same whether the tree is set from Go or from a temporary fixture.

#### Scenario: Mount element hosts the renderer

- **WHEN** the configured mount element exists in the document
- **THEN** the host attaches the React root there and renders from the current tree-root input

#### Scenario: Missing mount element fails closed

- **WHEN** the configured mount element is absent
- **THEN** the host does not create a React root elsewhere and leaves a developer-visible indication of the failure

#### Scenario: Imperative tree update after boot

- **WHEN** the Go runtime pushes a new tree root through the host’s tree-update API
- **THEN** the mounted host re-renders from that tree without remounting a different DOM element
