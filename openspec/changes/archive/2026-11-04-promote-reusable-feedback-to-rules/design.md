# Design

## Context

See proposal.md — Why. Agent-oriented conventions already live as `.cursor/rules/*.mdc` (including always-apply and glob-scoped Go rules). `.cursor/hooks.json` is empty; OpenSpec apply/finalize skills do not currently harvest post-apply style feedback into rules. This change is agent-behavior tooling only (P1-adjacent contributor experience), not DSL/runtime.

## Goals / Non-Goals

**Goals:**

- Encode the confirmed ask-first promotion workflow so every agent session sees it without the user repeating “make this a rule”
- Keep questions sequential and compatible with the existing one-clarifying-question rule
- Allow explicit one-shot phrasing to skip both asks

**Non-Goals:**

- Silent auto-write of rules
- Cursor hooks or stop-loop follow-ups
- OpenSpec CLI changes or new slash commands
- Spec deltas (agent chat workflow is not product capability behavior; change uses `skip_specs: true`)
- Rewriting existing Go convention rule bodies

## Decisions

### 1. Always-apply meta-rule as the primary mechanism

**Choice:** Add `.cursor/rules/promote-reusable-feedback.mdc` (or similar kebab name) with `alwaysApply: true`.

**Why:** Post-apply feedback is usually freeform chat, not `/opsx-apply`. Skills and `operations.apply.guidance` only fire when those workflows are active. An always-apply rule covers all sessions.

**Alternatives considered:**

- OpenSpec skill-only guidance → misses freeform follow-ups
- `stop` hook with `followup_message` → noisy extra user bubble; poor classifier home
- User-global hooks/rules only → not shared with the repo

### 2. Ask-first with sequential Q1 then Q2

**Choice:** After a local fix for reusable-looking feedback: ask durable-rule promotion; only if yes, write/update the rule, then ask about occurrence sweep. One question per assistant message.

**Why:** Matches user preference and existing `one-clarifying-question` rule. Avoids auto-sprawl and avoids batching two decisions.

### 3. Classifier stays heuristic in the rule text

**Choice:** Document positive signals (naming, forbidden patterns, preferred APIs, package/file layout) and negative signals (typos, this-file-only renames, temporary workarounds) in the meta-rule. No separate classifier script.

**Why:** Judgment belongs in the agent; a hook cannot reliably classify. Explicit user phrasing covers false negatives.

### 4. Optional apply-guidance one-liner

**Choice:** Add one advisory line under `openspec/config.yaml` `operations.apply.guidance` pointing agents at the meta-rule during apply/review — not a second source of truth.

**Why:** Reinforces the habit right after `/opsx-apply` without duplicating the full workflow in skills.

### 5. Rule file authoring conventions

**Choice:** When promoting, create or merge into a focused `.mdc` (prefer updating an overlapping existing rule). Use `create-rule` frontmatter conventions (`description`, `globs` or `alwaysApply`). Briefly acknowledge which rule path was created/updated.

**Why:** Matches existing forma rule style (`prefer-fn.mdc`, etc.) and keeps the ruleset navigable.

## Risks / Trade-offs

- [False-positive Q1 noise] → Mitigation: tight reusable vs one-off examples in the meta-rule; user can answer no quickly
- [False-negative missed promotion] → Mitigation: explicit “add durable rule …” phrasing; apply-guidance reminder after apply
- [Rule sprawl from many small yes answers] → Mitigation: instruct merge-into-existing when topics overlap; keep rules focused
- [Sweep scope ambiguity] → Mitigation: Q2 asks before sweeping; on yes, search proportionally and report what changed

## Migration Plan

1. Add the always-apply meta-rule
2. Optionally append apply-guidance one-liner
3. No runtime deploy; takes effect on next agent session that loads project rules
4. Rollback: delete the meta-rule (and guidance line) if the workflow proves too noisy

## Open Questions

None — promotion confirmation model, classifier threshold, and sequential asks were decided in explore.
