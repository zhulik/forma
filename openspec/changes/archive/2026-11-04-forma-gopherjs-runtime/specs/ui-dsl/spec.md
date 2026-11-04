# Spec Delta

## MODIFIED Requirements

### Requirement: Application mount entry

Apps SHALL attach a root component to a DOM mount identifier through the single exported mount entry point. On browser builds, that mount MUST activate the live runtime bridge (initial refresh into the host). Non-browser builds MAY keep a no-op stub so native compile and test remain possible.

#### Scenario: Demo app startup

- **WHEN** the reference todo application starts
- **THEN** it registers its root component against a fixed mount id

#### Scenario: Browser mount is live

- **WHEN** the application is built for the browser target and calls the mount entry
- **THEN** the runtime performs an initial refresh that drives the host tree seam
