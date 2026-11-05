## ADDED Requirements

### Requirement: Builtin spacing and max-width classes are statically enumerable

Framework layout builtins that map numeric gap/padding/margin or token max-width props to Tailwind utility classes SHALL resolve those props through a closed set of literal class strings (or equivalent static enumeration). The host MUST NOT build those utility class names solely by concatenating a prefix with a runtime number or token in a way that leaves no complete class literal for the stylesheet scanner.

#### Scenario: Gap and padding map to known utilities

- **WHEN** a Row or Column supplies a positive gap or padding within the supported scale
- **THEN** the rendered element’s class list includes a full spacing utility string that exists as a literal in the host sources

#### Scenario: Max-width token maps to known utility

- **WHEN** a Container supplies a supported max-width token
- **THEN** the rendered element’s class list includes a full max-width utility string that exists as a literal in the host sources

#### Scenario: Unsupported scale values omit the utility

- **WHEN** a layout prop value is outside the supported static scale
- **THEN** the host omits that spacing or max-width utility rather than emitting a dynamic class string

### Requirement: Scaffold Tailwind content includes author Go Class sources

The minimal project scaffold’s Tailwind content configuration SHALL include paths that cover application Go sources where authors place Class escape-hatch utility strings, so those utilities are eligible for inclusion in the built stylesheet without requiring a hand-maintained safelist of every Class string.

#### Scenario: Responsive Class utilities survive build

- **WHEN** an author sets a Class field on a layout or widget to a static Tailwind responsive utility string present in a Go source under the configured content paths
- **THEN** the production stylesheet build retains that utility

### Requirement: Builtin boxed layout type is Panel

The framework builtin registry SHALL expose the padded bordered boxed layout surface under the type string Panel. The builtin registry MUST NOT register a layout entry under the type string Card, so a generated shadcn Card registry entry can own Card without colliding with a layout builtin.

#### Scenario: Panel renders without generated entry

- **WHEN** the tree contains a Panel node and no generated override for Panel
- **THEN** the host renders the builtin boxed layout mapping

#### Scenario: Generated Card is not shadowed by a layout builtin

- **WHEN** the generated registry includes Card and the tree contains a Card node
- **THEN** the host renders the generated Card mapping, not a framework layout Card builtin
