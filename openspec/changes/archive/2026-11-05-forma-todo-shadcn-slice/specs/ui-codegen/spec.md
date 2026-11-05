## ADDED Requirements

### Requirement: Value-style and checked-style callbacks keep payload types

Generated Go callback fields for change-style and checked-change-style props SHALL take the UI callback context as the first argument and a sanitized plain payload as the second: string for text/value change handlers and boolean for checked-change handlers. Click-style handlers without a payload SHALL remain context-only. Codegen MUST NOT collapse every event prop to a context-only signature.

#### Scenario: Text input OnChange

- **WHEN** codegen processes a text input source that exposes a string value change handler
- **THEN** the emitted Go field is a two-argument callback with callback context then string

#### Scenario: Checkbox OnCheckedChange

- **WHEN** codegen processes a checkbox source that exposes a checked-change handler
- **THEN** the emitted Go field is a two-argument callback with callback context then boolean

#### Scenario: Click stays context-only

- **WHEN** codegen processes a button-style click handler with no value payload
- **THEN** the emitted Go field remains a single-argument callback with callback context only

### Requirement: Controlled value props are emitted as ordinary fields

Generated bindings for widgets that expose controlled string or boolean state in their TypeScript props SHALL include corresponding Go fields (such as value, checked, and placeholder) with matching scalar Go types, not omit them in favor of callbacks alone.

#### Scenario: Input value and placeholder

- **WHEN** codegen processes an input source with string value and placeholder props
- **THEN** the emitted struct includes string fields for those props alongside the change callback

#### Scenario: Checkbox checked

- **WHEN** codegen processes a checkbox source with a boolean checked prop
- **THEN** the emitted struct includes a boolean checked field alongside the checked-change callback

## MODIFIED Requirements

### Requirement: Per-component override hook

The pipeline SHALL consult optional per-source override data when automatic classification of props, events, or compound slots is insufficient, and SHALL document the override format for maintainers. Overrides MUST be able to set Go field names and Go types for props and callbacks (including value-style and checked-style signatures) and compound slot classification without hand-editing generated output.

#### Scenario: Override applied

- **WHEN** a maintainer adds an override for a source file that automatic classification cannot handle
- **THEN** codegen emits Go and registry entries consistent with the override without hand-editing generated files

#### Scenario: Override forces payload callback type

- **WHEN** an override specifies a checked-change or value-change callback Go type for a prop that heuristics would otherwise emit as context-only
- **THEN** the generated binding uses the override’s Go type for that field
