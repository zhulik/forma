# ui-host Specification

## Purpose

Define the React host that turns forma render trees into DOM via a component registry, recursive rendering, sanitized callbacks, and a stable text-input policy—independent of the Go↔browser bridge.

## Requirements

### Requirement: Registry maps type strings to React components

The host SHALL resolve each render-tree type string through a component registry that merges framework builtin entries with the generated application registry. Generated entries MUST override builtins when the same type string appears in both. Unresolved type strings MUST NOT crash the host process; the host MUST skip that node and surface a diagnostic suitable for developers.

#### Scenario: Generated widget renders via registry

- **WHEN** the tree contains a node whose type matches a generated registry entry
- **THEN** the host renders the registered React component for that type

#### Scenario: Builtin layout type renders without generated entry

- **WHEN** the tree contains a framework layout or leaf type that has a builtin registry entry and no generated override
- **THEN** the host renders the builtin mapping for that type

#### Scenario: Unknown type is skipped safely

- **WHEN** the tree contains a type string absent from both builtin and generated registries
- **THEN** the host omits that node from the DOM output and emits a developer-visible diagnostic

### Requirement: Recursive tree rendering

The host SHALL recursively render each node’s primary children and named slots into the React tree. Nodes with a key MUST pass that key to React. Fragment-typed nodes MUST introduce no extra DOM host element of their own. Slot content MUST be rendered using the registry slot mapping for the parent type when slots are declared.

#### Scenario: Nested children render in order

- **WHEN** a parent node has an ordered list of primary child nodes
- **THEN** the host renders those children in the same order under the parent component

#### Scenario: Fragment has no DOM host

- **WHEN** a Fragment node wraps sibling children
- **THEN** those siblings appear in the parent without an intermediate Fragment DOM element

#### Scenario: Named slots attach to compound widgets

- **WHEN** a compound node carries named slot lists whose keys match the registry slot map for that type
- **THEN** the host supplies the rendered slot content to the corresponding React slot components

#### Scenario: Stable keys reach React

- **WHEN** a node carries a non-empty key
- **THEN** the rendered React element uses that key for reconciliation

### Requirement: Event payloads are sanitized

Callback props invoked from the host MUST receive only plain data (strings, booleans, numbers, plain objects/arrays of those). The host MUST NOT pass framework or DOM event objects into those callbacks.

#### Scenario: Click handler receives no event object

- **WHEN** a registered click-style callback fires from a rendered widget
- **THEN** the callback is invoked without a DOM or synthetic event object as an argument

#### Scenario: Change handler receives a primitive value

- **WHEN** a text or checked-change style callback fires
- **THEN** the callback receives a plain string or boolean payload rather than an event object

### Requirement: Text-input hybrid policy

For text input widgets that expose a string value and a change callback, the host SHALL keep ephemeral string state in the React layer while the field is focused so typing remains stable across tree refreshes, SHALL sync displayed value from the tree’s value prop when the field is not focused or when the prop changes from outside while unfocused, and SHALL invoke the sanitized change callback with the current string as the user edits.

#### Scenario: Typing survives a same-value refresh

- **WHEN** the user is focused in a text input and the tree refreshes with the same value prop the host last reported
- **THEN** the caret and in-progress text remain intact

#### Scenario: External value updates when unfocused

- **WHEN** the text input is unfocused and the tree supplies a new value prop
- **THEN** the displayed field shows the new value

#### Scenario: Edits notify via sanitized string

- **WHEN** the user changes text in the input
- **THEN** the change callback receives the updated string without an event object

### Requirement: Host mount seam accepts a tree root

The scaffold host entry SHALL mount into the configured DOM mount element and render from a replaceable tree-root input. Until the Go bridge ships, that input MAY be a fixture or null placeholder, but the mount path MUST be the same one the bridge will drive later.

#### Scenario: Mount element hosts the renderer

- **WHEN** the configured mount element exists in the document
- **THEN** the host attaches the React root there and renders from the current tree-root input

#### Scenario: Missing mount element fails closed

- **WHEN** the configured mount element is absent
- **THEN** the host does not create a React root elsewhere and leaves a developer-visible indication of the failure
