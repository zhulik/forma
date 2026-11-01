# ui-dsl Specification

## Purpose

Define the Go authoring model for forma UIs: how authors compose widgets, attach keys, branch on conditions, and refresh the tree after state changes, independent of the React host implementation.

## Requirements

### Requirement: Children are components only

Composable widgets SHALL accept child nodes only as other components. Plain string values MUST NOT be accepted as direct children in composition APIs.

#### Scenario: Label on a button

- **WHEN** an author adds visible label text to a button-like widget
- **THEN** the label is supplied as a text widget struct or as a child via composition, not as a bare string child argument

#### Scenario: Dynamic list splat

- **WHEN** an author builds a slice of repeated row components and splats them into a parent
- **THEN** the slice element type is the component interface, not an untyped slot

### Requirement: Text content spelling

Authors MUST represent simple inline copy with the text widget struct using an explicit **`Text` field** (e.g. **`Text{Text: "hello"}`**). Field-based text on other widgets that expose a dedicated text property MUST remain valid where the widget defines it.

#### Scenario: Inline label

- **WHEN** an author writes short copy such as a button label
- **THEN** the label is a text widget value with the copy in the **`Text` field**, composed as a child (not a bare string)

#### Scenario: Heading copy

- **WHEN** an author sets a heading’s primary copy via a struct field
- **THEN** the rendered tree reflects that field without requiring a child text node

### Requirement: Typed fallback composition

Every composable widget type SHALL provide a fallback composition method that returns the same widget type as its primary composition method, so layout and widget-specific properties are preserved when fallbacks run.

#### Scenario: Empty list with layout wrapper

- **WHEN** a column with gap or padding has zero primary children and defines a fallback child
- **THEN** the rendered node type and layout properties match the column, not a generic anonymous wrapper

#### Scenario: Multiple fallback calls

- **WHEN** an author registers fallbacks in more than one call on the same node
- **THEN** all registered fallback children are considered in order when primary children are empty

### Requirement: Primary-empty fallback selection

When a composable node has zero primary children after composition, the tree builder SHALL render its fallback children instead of rendering an empty container (unless no fallback was defined).

#### Scenario: Zero todos

- **WHEN** a list region has no mapped items and defines a fallback message
- **THEN** the fallback message appears in the tree

### Requirement: Stable keys for repeated siblings

Repeated sibling nodes built from collections MUST carry a stable string key on each instance. Keys SHOULD be omitted only for nodes that are not part of a dynamic sibling list.

#### Scenario: Todo rows

- **WHEN** an app renders one row per item in a mutable list
- **THEN** each row’s node includes a key derived from a stable item identifier

### Requirement: Lazy conditional composition in the UI package

The UI package SHALL provide conditional composition that accepts a boolean and one or two lazy thunks returning a component. Only the selected thunk MAY be invoked. This helper MUST NOT live in the generic functional helper package.

#### Scenario: Single branch false

- **WHEN** the condition is false and only one thunk is provided
- **THEN** the result contributes no visible node where it is composed

#### Scenario: Two branches true

- **WHEN** the condition is true and two thunks are provided
- **THEN** only the first thunk runs and its component is used

#### Scenario: Else branch

- **WHEN** the condition is false and two thunks are provided
- **THEN** only the second thunk runs and its component is used

#### Scenario: Invalid arity

- **WHEN** more than two thunks are supplied
- **THEN** the library fails at call time rather than silently ignoring extras

#### Scenario: Conditional label inside a button

- **WHEN** an author wraps a conditional text branch as a child of a button via composition
- **THEN** the button properties stay on the button while the lazy branch selects the label component

### Requirement: Callback context versus render context

Event and interaction callbacks SHALL receive the UI callback context type embedding cancellation and refresh. Tree rendering and internal forma lifecycle SHALL use the standard library context type only.

#### Scenario: Handler refresh

- **WHEN** a user action runs an app handler
- **THEN** the handler receives the UI callback context, not the render-only context type

#### Scenario: Render pass

- **WHEN** the framework rebuilds the component tree
- **THEN** the render entry point accepts the standard library context

### Requirement: Atomic update and refresh

The UI callback context SHALL support running a synchronous mutation closure followed by exactly one UI refresh when the closure returns without panicking.

#### Scenario: Successful mutation

- **WHEN** a handler calls the update helper with a closure that mutates app state
- **THEN** the closure runs to completion and the UI refreshes once afterward

#### Scenario: Panic during mutation

- **WHEN** the closure panics
- **THEN** no refresh is triggered by that update call

### Requirement: Manual refresh remains available

The UI callback context MUST expose explicit refresh in addition to the update helper. Authors MAY call explicit refresh when no mutation closure is used, such as after asynchronous work completes.

#### Scenario: Async completion

- **WHEN** a goroutine finishes work and mutates app state outside a synchronous update closure
- **THEN** the handler or goroutine may trigger refresh explicitly once state is consistent

### Requirement: Application mount entry

Apps SHALL attach a root component to a DOM mount identifier through the single exported mount entry point. Behavior of the live bridge MAY remain stubbed until runtime phases.

#### Scenario: Demo app startup

- **WHEN** the reference todo application starts
- **THEN** it registers its root component against a fixed mount id

### Requirement: Dense imports in application code

Application and screen packages MUST be allowed to use dot-import style for the UI and functional helper packages so composition reads as a DSL. Library packages that re-export forma MUST use normal qualified imports.

#### Scenario: Reference demo imports

- **WHEN** a reader opens the canonical todo demo
- **THEN** it uses dot imports for UI and list helpers with linter exceptions documented inline
