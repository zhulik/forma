# Spec Delta

## MODIFIED Requirements

### Requirement: Agent-oriented Go rules are present

The repository SHALL include Cursor rules covering early-stage compatibility policy, Go error handling, testing conventions (for when tests are added later), file/type organization, and preferential use of the functional helper package over raw loops, so automated edits follow the same standards as sister zhulik Go projects. The prefer-fn rule MUST forbid raw `for` loops except an explicit allowlist (helper package internals, generated code, table-driven tests, side-effect-only each walks, non-slice/non-map iteration such as channels, indexed count APIs, short-circuit search, and index-heavy algorithms), and MUST forbid no-op proxy lambdas when a named function already matches a helper callback signature.

#### Scenario: Rules available for Go edits

- **WHEN** an agent or contributor opens Go source under this repository
- **THEN** applicable Cursor rules document early-stage breaking-change policy, error variable naming, one-type-one-file expectations, and the prefer-fn allowlist plus no-op-proxy ban

#### Scenario: Agent considers a slice transform loop

- **WHEN** an agent is about to write a raw `for` loop that only maps or filters a slice into a new collection
- **THEN** the prefer-fn rule directs them to use the functional helper package instead unless the loop matches an allowlisted case
