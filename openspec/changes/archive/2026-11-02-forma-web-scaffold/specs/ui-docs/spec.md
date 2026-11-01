# Spec Delta — ui-docs

## MODIFIED Requirements

### Requirement: Implementation status is explicit

Where documented behavior depends on work not yet shipped (browser dev loop, Go codegen from shadcn TSX, JS component registry), the guide SHALL label that material as planned or not yet available and SHALL NOT imply a working browser app when only the DSL stub exists in the framework repository. The guide SHALL state which CLI commands are available in the current release (`init`, config-only init, `sync` with web npm and shadcn materialization, version) versus planned (`add`, full sync including Go codegen and registry, `dev`).

#### Scenario: Author expects forma dev today

- **WHEN** an author reads the development workflow page before P3d lands
- **THEN** they see which steps work today (e.g. DSL reference via `main.go`, CLI init and web tree sync) versus which require a future release

#### Scenario: Author expects full sync today

- **WHEN** an author reads sync documentation before Go codegen and the JS registry ship
- **THEN** they see that sync installs npm dependencies and shadcn sources but may skip Go codegen and registry update until those phases land
