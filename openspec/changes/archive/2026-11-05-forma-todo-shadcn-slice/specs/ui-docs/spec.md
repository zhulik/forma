## ADDED Requirements

### Requirement: Guide documents todo-slice authoring surface

The author guide SHALL document the init starter set including checkbox, the rename of the framework boxed layout from Card to Panel, typed value and checked callback signatures on generated Input and Checkbox bindings, and that Tailwind content scanning includes Go Class sources so static utility strings in Class fields (including responsive variants) remain available after CSS build.

#### Scenario: Starter and checkbox

- **WHEN** an author reads CLI or adding-components material about init
- **THEN** they see that the default starter set includes button, input, card, and checkbox

#### Scenario: Panel and Card distinction

- **WHEN** an author reads writing-UI or layout material
- **THEN** they see that Panel is the framework boxed layout and Card refers to the generated shadcn widget after init

#### Scenario: Callback and Class scanning

- **WHEN** an author reads writing-UI or troubleshooting material about events and styling
- **THEN** they see OnChange/OnCheckedChange payload shapes and that Class utilities must be static strings covered by the scaffold Tailwind content paths
