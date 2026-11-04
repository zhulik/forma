# ui-runtime Specification

## Purpose

Define the Go↔browser runtime bridge: mounting a root component, refreshing the render tree into the host, dispatching sanitized callbacks into Go handlers with callback context, and cancelling prior handler work across refresh generations.

## Requirements

### Requirement: Browser mount drives the host tree

When the application runs in the browser target, the mount entry point SHALL register the root component against the configured mount identifier, obtain the host tree-update seam, perform an initial refresh, and SHALL NOT leave the host tree permanently null after a successful mount.

#### Scenario: First paint after mount

- **WHEN** a browser-built app calls the mount entry with a root component and the host seam is available
- **THEN** the host receives a non-null render tree produced from that root

#### Scenario: Missing host seam fails closed

- **WHEN** the mount entry runs in the browser but the host has not registered a tree-update seam
- **THEN** the runtime surfaces a developer-visible failure and does not silently attach elsewhere

### Requirement: Refresh rebuilds from current Go state

An explicit UI refresh SHALL re-render the mounted root component with the standard library context type, produce a new render tree, and replace the host’s current tree root with that result. Domain state remains in Go; the host MUST NOT own application business state.

#### Scenario: State change then refresh

- **WHEN** Go application state changes and a refresh runs
- **THEN** the host tree reflects a fresh render of the root against that state

#### Scenario: Render uses standard context

- **WHEN** a refresh rebuilds the tree
- **THEN** the root component’s render entry receives the standard library context, not the callback context type

### Requirement: Update runs mutation then one refresh

The callback context update helper SHALL run the supplied synchronous closure and, when the closure returns without panicking, trigger exactly one refresh. When the closure panics, that update call MUST NOT trigger a refresh.

#### Scenario: Successful update

- **WHEN** a handler invokes the update helper with a closure that mutates state and returns
- **THEN** the UI refreshes once after the closure completes

#### Scenario: Panic skips refresh

- **WHEN** the update closure panics
- **THEN** no refresh is triggered by that update call

### Requirement: Callbacks receive bound callback context

Callbacks that cross from the host into Go SHALL invoke the author-supplied handler with a live callback context as the first argument, followed by any sanitized plain payloads the host provides. The host MUST NOT be required to construct or understand the callback context type.

#### Scenario: Click-style handler

- **WHEN** the host invokes a click-style callback with no payload
- **THEN** the author handler runs with a callback context and no DOM or framework event object

#### Scenario: Value-style handler

- **WHEN** the host invokes a change-style callback with a plain string or boolean
- **THEN** the author handler runs with a callback context plus that plain value

### Requirement: Refresh cancels prior callback contexts

Each refresh generation SHALL provide callback contexts derived from a cancellable parent. Starting a new refresh MUST cancel callback contexts from the previous generation so in-flight async work can observe cancellation. Unmount or replacement of the mount MUST likewise cancel outstanding callback contexts.

#### Scenario: Stale generation cancelled

- **WHEN** a handler starts async work under a callback context and a later refresh runs before that work finishes
- **THEN** the earlier callback context becomes done

#### Scenario: Unmount cancels

- **WHEN** the mounted app is torn down or replaced
- **THEN** outstanding callback contexts from that mount become done

### Requirement: Native builds keep a non-browser stub

On non-browser builds, the mount and refresh entry points MUST remain safe no-ops (or test doubles) so ordinary native compile and test of application packages continue to succeed without a DOM or host seam.

#### Scenario: Native compile of the reference app

- **WHEN** the reference application is built for a non-browser target
- **THEN** the mount entry compiles and does not require a live host
