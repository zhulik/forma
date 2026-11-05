## ADDED Requirements

### Requirement: Framework boxed layout is Panel

The UI package SHALL provide a framework layout widget for a padded bordered box whose render-tree type string is Panel. Authors MUST NOT rely on a framework layout widget whose type string is Card; Card is reserved for generated shadcn Card bindings in application projects after the default starters are materialized.

#### Scenario: Panel composition

- **WHEN** an author wraps children in the framework boxed layout widget
- **THEN** the rendered node type is Panel and layout properties such as padding are preserved on that node

#### Scenario: Card means generated shadcn

- **WHEN** an application includes a generated Card binding from the default starter set
- **THEN** author Card usage refers to that generated compound (or leaf) binding, not a framework layout Card type

### Requirement: Generated value and checked handlers are author-callable with payloads

Author-facing generated widgets that expose text change or checked-change interactions SHALL type those handlers so authors receive callback context plus the corresponding plain string or boolean payload. Authors MUST be able to update controlled Value or Checked fields from those handlers without casting or host-specific event types.

#### Scenario: Controlled input authoring

- **WHEN** an author wires an input Value field and an OnChange handler
- **THEN** the handler signature accepts callback context and the new string value

#### Scenario: Controlled checkbox authoring

- **WHEN** an author wires a checkbox Checked field and an OnCheckedChange handler
- **THEN** the handler signature accepts callback context and the new boolean value
