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

Apps SHALL attach a root component to a DOM mount identifier through the single exported mount entry point. On browser builds, that mount MUST activate the live runtime bridge (initial refresh into the host). Non-browser builds MAY keep a no-op stub so native compile and test remain possible.

#### Scenario: Demo app startup

- **WHEN** the reference todo application starts
- **THEN** it registers its root component against a fixed mount id

#### Scenario: Browser mount is live

- **WHEN** the application is built for the browser target and calls the mount entry
- **THEN** the runtime performs an initial refresh that drives the host tree seam

### Requirement: Dense imports in application code

Application and screen packages MUST be allowed to use dot-import style for the UI and functional helper packages so composition reads as a DSL. Library packages that re-export forma MUST use normal qualified imports.

#### Scenario: Reference demo imports

- **WHEN** a reader opens the canonical todo demo
- **THEN** it uses dot imports for UI and list helpers with linter exceptions documented inline

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
