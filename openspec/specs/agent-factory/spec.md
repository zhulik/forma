# agent-factory Specification

## Purpose

Define the local, roadmap-driven agent factory that turns `doc/ROADMAP.md` into human-reviewed pull requests via named roles, OpenSpec workflows, and explicit merge gates—without cloud agent control planes or auto-merge.

## Requirements

### Requirement: Roadmap is the factory control plane

The repository SHALL treat `doc/ROADMAP.md` as the durable plan and live progress surface for the agent factory. Factory roles that advance work MUST update status, branch, and pull-request fields on the roadmap when those become known, without requiring a separate orchestration state document.

#### Scenario: Orch records an opened PR

- **WHEN** a Builder opens a pull request for a roadmap-named change
- **THEN** the Orch role updates the corresponding roadmap entry with at least status and pull-request identity
- **AND** no separate required orch-only state file is mandated for that bookkeeping

### Requirement: Named factory agents

The project SHALL provide Cursor custom subagents under `.cursor/agents/` named `planner`, `orch`, `architector`, `builder`, and `debt`. Each agent MUST state that role’s inputs, outputs, and forbidden actions. OpenSpec mechanics SHALL remain skills/commands (what to do); these agents are the who.

#### Scenario: Operator advances via orch

- **WHEN** an operator asks to advance the roadmap through the factory
- **THEN** the `orch` agent is available and describes reading the roadmap, delegating to `architector` / `builder`, and never merging

### Requirement: Local-only execution

Factory agent work SHALL run against local checkouts (including git worktrees). The factory MUST NOT require Cursor cloud agents or Cursor Automations as the control plane for scheduling or building.

#### Scenario: Parallel changes use isolation

- **WHEN** Orch runs more than one in-flight change concurrently
- **THEN** each change uses an isolated local working tree (for example a git worktree)
- **AND** scheduling does not depend on a Cursor-hosted cloud VM

### Requirement: Orch interprets roadmap and asks when unsure

The Orch role SHALL infer change order, branch stacking bases, and safe concurrency from roadmap prose, phase notes, and repository context. When order or overlap is unclear, Orch MUST ask the human before spawning further builders rather than guessing.

#### Scenario: Ambiguous parallelism

- **WHEN** two planned changes might touch overlapping areas and Orch cannot determine safe concurrency
- **THEN** Orch asks the human before starting both
- **AND** does not silently start both in parallel

### Requirement: Architector plans autonomously

The Architector role SHALL run OpenSpec explore and propose for exactly one roadmap-assigned change, decide material design questions without blocking on human Q&A for ordinary ambiguity, and leave planning artifacts ready for Builder apply. Material decisions MUST be recorded so they can appear in the eventual pull-request description.

#### Scenario: Ambiguous design choice

- **WHEN** Architector faces a material design alternative during explore/propose
- **THEN** it chooses, records the decision in planning artifacts
- **AND** does not wait for human approval before finishing the proposal unless blocked by missing credentials, broken tools, or an explicit operator stop

### Requirement: Builder ships one OpenSpec change to a PR

The Builder role SHALL apply the Architector’s OpenSpec change in a local worktree, run the project’s required verification gates, complete archive/finalize as defined by project OpenSpec habit, open a pull request against the base Orch assigned, babysit continuous integration failures within the change’s scope, and wait until a human approves and merges. Builder MUST NOT merge the pull request.

#### Scenario: Change reaches review

- **WHEN** Builder finishes apply and verification for its assigned change
- **THEN** it opens a pull request
- **AND** the pull request remains unmerged by the Builder
- **AND** Builder continues addressing in-scope CI failures until green or escalates a blocker

### Requirement: Pull request surfaces Architector decisions

Every factory feature pull request MUST include a description section that lists the Architector’s important decisions for that change.

#### Scenario: Reviewer reads the PR

- **WHEN** a human opens a Builder-created feature pull request
- **THEN** the description contains an explicit decisions section derived from the Architector’s planning artifacts
- **AND** those decisions are visible without reading the OpenSpec files alone

### Requirement: Review comments engage Architector and Builder

When a human leaves review comments on a factory pull request, the Architector SHALL respond to design or decision challenges (revise plan or defend), and the Builder SHALL implement agreed code changes and continue CI babysitting. Neither role MAY merge.

#### Scenario: Human disagrees with a recorded decision

- **WHEN** a human comments that an Architector decision in the PR description is wrong
- **THEN** the Architector reacts on that thread (revise or defend with rationale)
- **AND** if the plan changes, the Builder updates the implementation and verification as needed
- **AND** merge remains a human action

### Requirement: Debt auditor opens propose PRs

The Debt auditor role, when invoked manually or on the project’s every-N-merges cadence, SHALL analyze the tree for reuse opportunities (well-known third-party modules, existing codegen, dead code, DRY, modern language use) and open a propose-oriented pull request presenting findings and recommended OpenSpec or roadmap follow-ups. Debt MUST NOT merge and MUST NOT silently apply large refactors without that review path.

#### Scenario: Operator runs a debt pass

- **WHEN** an operator invokes the Debt auditor
- **THEN** Debt opens a pull request containing ranked findings and proposed follow-ups
- **AND** does not merge that pull request
