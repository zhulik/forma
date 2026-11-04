# Design

## Context

See `proposal.md` for motivation. Today forma already has OpenSpec skills (`opsx-explore` / propose / apply / archive / finalize), Cursor rules under `.cursor/rules/`, and `doc/ROADMAP.md` as the phased delivery narrative. There is no named multi-agent factory that turns roadmap rows into stacked, human-merged PRs.

Constraints from exploration: **local-only** (no cloud agents / Automations as control plane); **markdown ROADMAP** as the single durable control surface; **Orch is a reasoning subagent**, not a heavy scheduler script; optional Node helpers use **built-ins first** and **ask before any npm dependency**.

## Goals / Non-Goals

**Goals**

- Ship `.cursor/agents/` contracts for planner, orch, architector, builder, debt; reuse `/opsx-*` skills for OpenSpec mechanics.
- Encode prefer-reuse as an always-on rule (and `dev-tooling` requirement).
- Document how roles update ROADMAP status/PR fields and how PR descriptions carry Architector decisions.
- Keep v1 operable from Cursor chat / Task / worktree without a mandatory long-running daemon.

**Non-Goals**

- Implementing product phases P3b–P3d inside this change.
- Auto-merge, cloud Automations, or a mandatory always-on GitHub poller.
- Deep builder prompt encyclopedias (v1 = agent contract; deepen later).
- Replacing OpenSpec CLI or inventing a second planning schema.

## Decisions

### 1. Agents (who) + OpenSpec skills (what)

**Choice:** Factory roles are Cursor **custom subagents** in `.cursor/agents/` (simple names, no `factory-` prefix). Procedural OpenSpec flows stay **skills/commands** (`/opsx-explore`, `/opsx-propose`, `/opsx-apply`, …). Orch delegates to architector/builder in local worktrees.

**Why:** Cursor’s own split—subagents for isolated multi-step roles; skills for repeatable procedures. Avoids duplicating opsx as role wrappers.

**Alternatives:** Role-as-skill+command (initial apply; superseded); Cursor SDK orchestrator script; cloud Automations — rejected for local-only / markdown-first stance.

### 2. ROADMAP mutable by Orch (and peers for status)

**Choice:** No `ORCH.md`. Orch (and Builder handoff) update status / branch / PR on `doc/ROADMAP.md`.

**Why:** Single document humans already read; session-local worktree/agent ids need not be durable.

**Alternatives:** Separate orch state file; conversation-only status — weaker crash visibility across days.

### 3. Architector vs Builder split

**Choice:** Architector = explore + propose, decides alone; Builder = apply → verify → archive/finalize → PR → CI babysit → wait for human merge. Both may address review (design vs code).

**Why:** Matches OpenSpec planning/apply boundary; PR is the disagreement surface for Architector decisions.

**Alternatives:** Single builder that explores and applies — rejected; blurs review of decisions.

### 4. Prefer-reuse always-on rule

**Choice:** Add `.cursor/rules/prefer-reuse-over-handroll.mdc` with `alwaysApply: true`; order third-party → codegen → hand-roll; exempt `fn/`.

**Why:** Applies to factory roles and ordinary chats; aligns with Debt propose PRs.

**Alternatives:** Skills-only policy — too easy to miss outside factory.

### 5. Debt opens propose PRs

**Choice:** Debt output is a reviewable PR (findings + recommended OpenSpec/roadmap follow-ups), not ROADMAP-only notes.

**Why:** Same comment loop as feature work.

**Alternatives:** ROADMAP bullets only; immediate implement PRs — weaker review or too aggressive.

### 6. Optional Node helpers

**Choice:** Defer any `.mjs` helpers until a skill proves markdown reasoning insufficient; Node built-ins only; ask before deps.

**Why:** Keeps apply scope small; matches exploration policy.

## Risks / Trade-offs

- **[Orch mis-orders or overlaps]** → Mitigation: ask human when unsure; human merge gate; serial default when unclear after ask.
- **[Architector wrong autonomous decision]** → Mitigation: decisions mandatory in PR body; Architector must react to review comments.
- **[Worktree / local cwd footguns]** → Mitigation: document one worktree (and preferably one process) per in-flight Builder; note SDK cwd caveat if SDK is introduced later.
- **[Role agents drift from OpenSpec]** → Mitigation: architector/builder MUST invoke existing `/opsx-*` skills rather than reimplement OpenSpec CLI flows.
- **[Debt noise PRs]** → Mitigation: ranked findings; propose-oriented; no auto-apply of mass refactors.

## Migration Plan

1. Land `.cursor/agents/*`, prefer-reuse rule, and a short operator note.
2. Optionally annotate ROADMAP OpenSpec workflow section with factory roles and status field expectations.
3. No runtime migration; rollback = delete `.cursor/agents/`, prefer-reuse rule, and operator doc (no data migration).

## Open Questions

- Exact default for debt’s “every N merges” (N value / how counted) — operator-configurable in the **debt** agent; default can be manual-only until tuned.
- Whether v1 orch uses Cursor Task tool only vs also Cursor SDK `Agent.create({ local })` — prefer Task/worktree first; SDK optional follow-up if Task proves insufficient.
