# Spec Delta

## ADDED Requirements

### Requirement: Registry slot keys match Go wire keys

For compound widgets, the generated host registry slot map keys SHALL match the wire slot keys written into the render tree for the corresponding Go slot fields (the same mapping used when emitting named slots from generated Go bindings). Registry slot keys MUST NOT use a separate capitalization scheme from those wire keys.

#### Scenario: Compound registry keys align with tree slots

- **WHEN** codegen emits a compound widget with slot fields and a registry entry that lists slot components
- **THEN** each registry slot map key equals the wire key used for that field in the generated Go render output

#### Scenario: Accordion-style parts stay consistent

- **WHEN** codegen processes a compound with trigger/content-style parts
- **THEN** the registry slot keys and the Go `Slots` map keys for those parts are identical strings
