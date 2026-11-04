# Spec Delta

## ADDED Requirements

### Requirement: Author guide covers functional helpers catalog

The author guide SHALL document that the functional helper package supports slice transforms (`Map`, `Filter`, and related helpers), fallible collect helpers that stop on the first error, and map-entry transform helpers, at the same level of discoverability as existing `Map` / `Filter` mentions.

#### Scenario: Author looks up fallible list helpers

- **WHEN** an author reads writing-UI or project-layout material about `fn`
- **THEN** they learn that fallible collect helpers and map transforms exist alongside the basic slice helpers
