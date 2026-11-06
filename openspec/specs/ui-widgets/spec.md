# ui-widgets Specification

## Purpose
Define the hand-written, forma-owned widget surface authors use for todo-like UIs: stable Go APIs, thin host rendering, and semantic look classes—without kit codegen or per-widget override files as the source of truth.

## Requirements

### Requirement: Todo-sized owned widget set

The framework SHALL provide hand-written Go widget types sufficient to express a todo-like UI: interactive Button, Input, and Checkbox; structural Card, Text, Heading, Row, Column, and Container; plus Fragment. These types MUST live in the framework UI package as reviewed source (not generated kit bindings).

#### Scenario: Todo demo types resolve from the UI package

- **WHEN** an author composes a screen that adds text tasks, toggles done state, and lays out a card with heading, input row, and list
- **THEN** every required widget type is available from the framework UI package without generating bindings from a third-party component kit

#### Scenario: No kit-generated binding required for core widgets

- **WHEN** a new project is initialized with the default scaffold
- **THEN** Button, Input, Checkbox, Card, and the layout/text types above are usable without an application-local generated widget package

### Requirement: Semantic look via named classes

Owned interactive and card widgets SHALL obtain their default look from semantic design-system class names (for example button and card primitives). Authors MUST be able to append escape-hatch class strings on widgets without replacing the semantic defaults unless they intentionally override presentation.

#### Scenario: Default button look without author CSS

- **WHEN** an author renders a Button with a primary-style variant and no custom class
- **THEN** the host applies the framework’s mapped semantic button classes so the control is visibly styled

#### Scenario: Class escape hatch composes

- **WHEN** an author sets an additional class string on a Card or Text
- **THEN** those classes are present on the rendered element alongside the widget’s default semantic or layout classes

### Requirement: Host applies look; Go owns API

Go widget structs SHALL expose a small, reviewable prop surface (variants, values, callbacks, layout fields). The React host MUST map those props to elements and semantic (or utility) classes. The product MUST NOT treat third-party React kit sources or per-widget YAML override matrices as the source of truth for this surface.

#### Scenario: Variant maps in the host

- **WHEN** a Button carries a documented variant value such as default/primary
- **THEN** the host selects the corresponding semantic class mapping without reading a per-widget override file

#### Scenario: Checkbox and input callbacks stay plain

- **WHEN** Checkbox checked-change or Input value-change fires
- **THEN** the Go callback receives a boolean or string payload consistent with the host sanitization policy

### Requirement: Layout props remain expressive without kit codegen

Row, Column, and Container SHALL keep author-facing layout fields (gap, alignment, max width, margins, padding, border as already used by the reference demo) and the host MUST realize them with the project’s styling stack without requiring generated kit wrappers for layout.

#### Scenario: Reference demo layout fields still work

- **WHEN** an author uses Container max-width and margin fields with Row/Column gap and alignment as in the reference todo demo
- **THEN** the host renders a corresponding layout structure without importing a generated layout binding
