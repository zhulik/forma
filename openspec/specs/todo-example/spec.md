# todo-example Specification

## Purpose

Define the product contract for a runnable in-memory todo example that proves the owned-widgets + 0build path end-to-end in the browser via the local development loop.

## Requirements

### Requirement: Dedicated runnable todo example exists

The product SHALL ship a dedicated in-repository todo example application that authors can run locally with the standard project sync and browser development loop commands. The example MUST NOT depend on a remote backend or durable persistence.

#### Scenario: Author finds the example in-repo

- **WHEN** an author looks for the canonical interactive todo demo in the repository
- **THEN** a dedicated example application path is present and documented as the runnable demo

#### Scenario: No persistence required

- **WHEN** the example application runs in the browser
- **THEN** task state lives only in process memory for that session and is not written to a database, file store, or remote API

### Requirement: In-memory todo CRUD on owned widgets

The example SHALL implement add, list, toggle-done, and delete for text tasks using only the todo-sized owned widget surface (interactive button, text input, checkbox; structural card, text, heading, and layout primitives). It MUST NOT introduce new framework widget types or require kit codegen.

#### Scenario: Author adds a task

- **WHEN** the author enters non-empty text and activates the add control
- **THEN** a new incomplete task appears in the list and the input clears

#### Scenario: Author toggles and deletes

- **WHEN** the author toggles a task’s done control or activates its delete control
- **THEN** the list reflects the updated done state or removes that task without a page-driven persistence step

#### Scenario: Empty list fallback

- **WHEN** there are zero tasks
- **THEN** the UI shows an empty-state message instead of an empty list chrome

### Requirement: Minimal responsive layout without new widgets

The example SHALL present a centered, max-width constrained layout that remains usable at narrow viewport widths using existing layout props and optional class escape hatches. It MUST NOT add a new responsive or breakpoint widget API.

#### Scenario: Narrow viewport remains usable

- **WHEN** the example is viewed at a typical phone-width viewport
- **THEN** the primary card and controls remain readable without requiring horizontal scrolling of the page chrome for the default content

### Requirement: Example runs via the standard browser loop

From the example’s project root, after sync prerequisites are met, the standard browser development loop command SHALL start the host and rebuild the example’s Go entry for the browser so authors can open the printed URL and interact with the todo UI.

#### Scenario: Successful local run

- **WHEN** an author completes documented sync for the example and starts the browser development loop from the example project root
- **THEN** they can open the served URL and perform add, toggle, and delete against the live Go-driven UI
