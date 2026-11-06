## MODIFIED Requirements

### Requirement: Registry maps type strings to React components

The host SHALL resolve each render-tree type string through a component registry that merges framework builtin entries with any optional application custom registry entries. Custom entries MUST override builtins when the same type string appears in both. The todo-sized owned widget surface (Button, Input, Checkbox, Card, Text, Heading, Row, Column, Container, Fragment) MUST be covered by builtins so apps do not need a kit-generated registry for those types. Unresolved type strings MUST NOT crash the host process; the host MUST skip that node and surface a diagnostic suitable for developers.

#### Scenario: Generated widget renders via registry

- **WHEN** the tree contains a node whose type matches an application custom registry entry (hand-written third-party or override)
- **THEN** the host renders the registered React component for that type

#### Scenario: Builtin layout type renders without generated entry

- **WHEN** the tree contains a framework layout or leaf type that has a builtin registry entry and no custom override
- **THEN** the host renders the builtin mapping for that type

#### Scenario: Unknown type is skipped safely

- **WHEN** the tree contains a type string absent from builtin and custom registries
- **THEN** the host omits that node from the DOM output and emits a developer-visible diagnostic

## ADDED Requirements

### Requirement: Owned interactive widgets resolve via builtins

Builtin registry entries SHALL cover Button, Input, and Checkbox in addition to layout/text/card types so a todo-like tree renders without a kit-generated registry module.

#### Scenario: Owned interactive widget renders via builtin

- **WHEN** the tree contains a Button, Input, or Checkbox node and no custom override
- **THEN** the host renders the builtin mapping for that type

#### Scenario: Custom registry overrides builtin

- **WHEN** the application custom registry registers the same type string as a builtin
- **THEN** the host uses the custom React component for that type

### Requirement: Owned widgets apply semantic look classes

Builtin host components for owned interactive and card widgets SHALL apply the project’s semantic design-system class names for default look. Layout builtins SHALL realize gap/alignment/spacing props using the same styling stack (semantic classes and/or documented utility escape hatches), not a retired kit-utility convention as the required path.

#### Scenario: Button default classes present

- **WHEN** a Button node renders without a custom class override that removes defaults
- **THEN** the DOM element includes the mapped semantic button classes for its variant

#### Scenario: Card default classes present

- **WHEN** a Card node renders with default props
- **THEN** the DOM element includes the mapped semantic card classes

#### Scenario: Layout gap without kit codegen

- **WHEN** a Row or Column node carries a positive gap prop
- **THEN** the host applies spacing consistent with the styling stack so children are separated
