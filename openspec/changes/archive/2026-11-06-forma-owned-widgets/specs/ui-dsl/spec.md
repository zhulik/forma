## REMOVED Requirements

### Requirement: Catalog compound widgets are generated

**Reason:** Product tip is hand-written owned widgets; shadcn-to-Go catalog compounds are not the SoT.

**Migration:** When compounds (Dialog, Tabs, …) return, add them as hand-maintained Go + host pairs under `ui-widgets` rules. Do not expect `*.gen.go` compounds from kit sources.

### Requirement: Compound shape from shadcn exports

**Reason:** Classification of shadcn export graphs is part of the retired codegen product path.

**Migration:** Keep field-only slots and “no Radix-shaped subcomponent types” DSL rules. Implement future compounds by hand with the same slot wire model.

## ADDED Requirements

### Requirement: Core catalog widgets are hand-maintained

Todo-sized leaf and layout widgets in the framework UI package MUST be hand-maintained Go source reviewed like ordinary library code. The DSL MUST NOT require authors to obtain those types from a kit codegen emit step.

#### Scenario: Button is ordinary library source

- **WHEN** a reader inspects the framework UI package for Button
- **THEN** they find a hand-maintained type implementing the component interface, not a kit-generated binding file presented as the product SoT

#### Scenario: Authors import framework UI for leaves

- **WHEN** an application composes Button, Input, and Checkbox
- **THEN** those types are imported from the framework UI package without a project-local generated widget package
