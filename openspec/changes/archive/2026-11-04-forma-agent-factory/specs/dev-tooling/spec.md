## ADDED Requirements

### Requirement: Prefer reuse over hand-roll rule is present

The repository SHALL include an always-applied Cursor rule that prefers well-known third-party modules, then existing repository codegen, then hand-written code—before planning or writing new solving code. Package `fn` MUST remain exempt from third-party replacement.

#### Scenario: Agent about to invent a parser

- **WHEN** an agent is about to plan or implement a problem already solved by a well-known third-party module (and the need is not in the `fn` domain)
- **THEN** the prefer-reuse rule directs them to use that module instead of a parallel in-tree implementation

#### Scenario: Agent about to hand-maintain generated twin

- **WHEN** an agent is about to hand-write code that an existing repository generator is contracted to emit
- **THEN** the prefer-reuse rule directs them to extend or run that generator instead

### Requirement: Prefer-reuse rule covers debt proposals

The prefer-reuse Cursor rule SHALL require debt-style or planning reviews to propose replacing hand-rolled code that a well-known module or existing codegen should own.

#### Scenario: Debt finds NIH util

- **WHEN** a debt or planning pass finds hand-rolled code that a well-known module or existing codegen should own
- **THEN** the prefer-reuse rule requires proposing that replacement rather than leaving NIH as the default
