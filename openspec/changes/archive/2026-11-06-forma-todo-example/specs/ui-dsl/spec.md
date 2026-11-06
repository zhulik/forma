# Spec Delta

## MODIFIED Requirements

### Requirement: Application mount entry

Apps SHALL attach a root component to a DOM mount identifier through the single exported mount entry point. On browser builds, that mount MUST activate the live runtime bridge (initial refresh into the host). Non-browser builds MAY keep a no-op stub so native compile and test remain possible.

#### Scenario: Demo app startup

- **WHEN** the canonical todo example application starts
- **THEN** it registers its root component against a fixed mount id

#### Scenario: Browser mount is live

- **WHEN** the application is built for the browser target and calls the mount entry
- **THEN** the runtime performs an initial refresh that drives the host tree seam

### Requirement: Dense imports in application code

Application and screen packages MUST be allowed to use dot-import style for the UI and functional helper packages so composition reads as a DSL. Library packages that re-export forma MUST use normal qualified imports.

#### Scenario: Reference demo imports

- **WHEN** a reader opens the canonical todo example application sources
- **THEN** it uses dot imports for UI and list helpers with linter exceptions documented inline
