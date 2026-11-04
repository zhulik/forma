# Tasks

## 1. Prefer-reuse rule

- [x] 1.1 Add always-on `.cursor/rules/prefer-reuse-over-handroll.mdc` (third-party → codegen → hand-roll; `fn/` exempt; debt/planning propose replacements) and verify it appears under `.cursor/rules/` with `alwaysApply: true`
- [x] 1.2 Cross-link from `prefer-fn.mdc` or rule text so `fn/` exemption stays consistent, and verify both rules coexist without contradictory “add collection deps” guidance

## 2. Factory agents (who) + OpenSpec skills (what)

- [x] 2.1 Add `.cursor/agents/planner.md` — interviews human, updates ROADMAP intent; no product implementation
- [x] 2.2 Add `.cursor/agents/orch.md` — roadmap interpretation, ask-when-unsure, local worktrees, status/PR updates, never-merge; forbids cloud Automations as control plane; delegates to architector/builder
- [x] 2.3 Add `.cursor/agents/architector.md` — explore→propose via `/opsx-*`, autonomous decisions, PR decision recording, design-comment reaction; forbids feature apply
- [x] 2.4 Add `.cursor/agents/builder.md` — apply→verify→archive→PR with mandatory Architector decisions→CI→wait for human merge; forbids merge
- [x] 2.5 Add `.cursor/agents/debt.md` — analysis + propose-oriented PR (not ROADMAP-only); never merge
- [x] 2.6 Remove obsolete `.cursor/skills/factory-*` and `.cursor/commands/factory-*` role wrappers

## 3. Operator documentation and ROADMAP conventions

- [x] 3.1 Add a short operator note (e.g. under `doc/` outside `doc/user/`) describing roles, local-only, ROADMAP as control plane, review/merge gates, and optional Node helper policy; verify `doc/user/` is unchanged
- [x] 3.2 Update `doc/ROADMAP.md` OpenSpec workflow section with factory role pointers and expected live status/PR fields; verify the product phase tables remain coherent

## 4. Integration check

- [x] 4.1 Walk a dry-run checklist (Planner → Orch → Architector → Builder → review → Debt) against the skills/docs without implementing a product change, and verify each role’s forbidden actions and handoffs are explicit in the written contracts

## Workflow follow-up

- Archive the change after review; sync `agent-factory` and `dev-tooling` deltas into main specs.
- Optionally tune Debt’s default N and whether Orch should gain Cursor SDK local spawn later.
