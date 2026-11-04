# Spec Delta

## ADDED Requirements

### Requirement: Functional helpers for collections

The framework SHALL provide a generic functional helper package for transforming slices and maps. Authors and framework code MUST be able to map, filter, and optionally keep elements; fallible collect helpers MUST stop on the first error and return that error with any partial result discarded. Side-effect-only walks are not required to use that package.

#### Scenario: Author maps a slice to components

- **WHEN** an author turns a slice of domain values into UI components
- **THEN** the functional helper package supplies a map operation that produces a new slice without a hand-written accumulation loop

#### Scenario: Fallible collect stops on first error

- **WHEN** a fallible collect helper encounters an error for one element
- **THEN** it returns that error and does not return a successfully built collection for the remaining elements

#### Scenario: Map transforms without reinventing loops

- **WHEN** an author remaps or filters map entries into another map or slice of results
- **THEN** the functional helper package supplies map-entry analogues to the slice helpers
