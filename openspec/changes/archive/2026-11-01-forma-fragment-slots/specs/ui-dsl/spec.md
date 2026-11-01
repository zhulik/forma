# Spec Delta

## ADDED Requirements

### Requirement: Fragment grouping without layout host

The UI package SHALL provide a fragment widget that groups child components for composition and rendering without introducing a layout or DOM host element of its own. Fragment MUST support the same primary-child and fallback composition methods as other composable widgets that embed the shared base component.

#### Scenario: Multiple siblings from one component value

- **WHEN** an author needs several sibling nodes inside a single parent field or child list without adding flex, gap, or padding
- **THEN** the author composes a fragment widget whose primary children are those siblings

#### Scenario: Host rendering

- **WHEN** the render tree includes a fragment node
- **THEN** the host reconciles the fragment’s children as direct siblings with no wrapper element corresponding to the fragment

#### Scenario: Fallback on empty primary children

- **WHEN** a fragment has zero primary children after composition and defines fallback children
- **THEN** the fallback children are rendered with the same no-wrapper semantics as primary children

### Requirement: Named slots on the render tree

The render tree node type SHALL support an optional map of named slots, each slot holding an ordered list of child nodes. Layout and leaf widgets SHALL populate only the primary child list. Compound widgets that declare named regions SHALL populate named slots and MUST NOT use the primary child list for slot payload in the same render output.

#### Scenario: Layout widget unchanged

- **WHEN** a row or column widget renders
- **THEN** its output uses the primary child list only and leaves named slots unset or empty

#### Scenario: Compound with two regions

- **WHEN** a compound widget renders with two named regions supplied by the author
- **THEN** each region’s subtree appears under the corresponding slot key in the render node, not mixed into the primary child list

### Requirement: Field-only named slots for compound widgets

Compound widgets (including generated shadcn-style widgets and custom registry widgets) SHALL expose each author-facing named region as an exported struct field whose type is the component interface or a plain Go collection type defined for that widget. Authors MUST supply slot content through those fields only, not through primary-child composition on the compound widget.

#### Scenario: Dialog-shaped compound

- **WHEN** an author configures a dialog-like compound with separate trigger and body regions
- **THEN** the author sets the trigger and body fields to component values such as a button and a column, without using separate author-facing types that mirror Radix or shadcn sub-parts

#### Scenario: Multiple nodes in one slot field

- **WHEN** a single slot field must contain more than one root node
- **THEN** the author supplies a fragment widget as the field value whose children are those roots

#### Scenario: Optional slot omitted

- **WHEN** an optional slot field is zero or unset
- **THEN** the render output omits that slot key or supplies an empty slot list, and the host skips that region

### Requirement: No author-facing Radix-shaped subcomponent types

The UI package and generated widget catalog MUST NOT require authors to use separate component types whose sole purpose is to mirror Radix or shadcn sub-parts such as triggers, content panels, or list chrome. Mapping from named slots to underlying React structure is a host and registry concern, not part of the Go authoring surface.

#### Scenario: Adding a new compound from the catalog

- **WHEN** a new shadcn-style compound is added to the framework
- **THEN** authors receive one compound struct with named fields and scalar props, not a set of nested subcomponent Go types for each Radix part

### Requirement: Repeated compound items without per-item subcomponent types

When a compound widget represents a variable-length collection of linked items (such as tabs or accordion sections), the widget SHALL accept a slice of a plain Go item struct defined for that widget. Each item struct MUST carry the data and component values needed for that entry (such as value, label, and panel content) without introducing separate author-facing trigger or panel component types per item.

#### Scenario: Tab strip with two panels

- **WHEN** an author renders two tabs with distinct panels
- **THEN** the author passes a slice of two item values on the compound widget, each item supplying label and panel components and a stable value identifier

#### Scenario: Keys on repeated items

- **WHEN** items in such a slice are rendered as repeated UI
- **THEN** each item’s rendered nodes include stable keys derived from the item’s value identifier or an explicit key field on the item struct

### Requirement: Catalog compound widgets are generated

Shadcn-style compound widget structs MUST be emitted by shadcn-to-Go codegen into `*.gen.go` files. The UI package MUST NOT ship hand-maintained catalog compound types alongside generated leaf widgets.

#### Scenario: No hand-written dialog type

- **WHEN** a reader inspects the UI package for catalog widgets
- **THEN** dialog-like compounds appear only as generated sources, not as hand-maintained Go files next to generated buttons

### Requirement: Compound shape from shadcn exports

The codegen pipeline MUST classify shadcn compound exports automatically when possible: a root export plus same-prefix part exports become one Go struct with field-based slots, not separate Go types per part. Per-source overrides MAY exist only when automatic classification fails.

#### Scenario: Prefixed part exports

- **WHEN** a shadcn source file exports a root component and sibling parts such as trigger and content with a shared name prefix
- **THEN** codegen emits one compound struct whose fields correspond to stripped part names and a Render method that fills named slots, not separate Go types for each part export

#### Scenario: Adding dialog from shadcn

- **WHEN** a maintainer adds shadcn dialog source to the host and runs codegen
- **THEN** the tool emits a single generated dialog struct with field-based slots, without adding hand-written dialog types in the UI package

#### Scenario: Slot helpers for generated Render

- **WHEN** generated compound Render methods run
- **THEN** they use shared UI package slot helpers so slot keys and omitted empty slots behave consistently across compounds
