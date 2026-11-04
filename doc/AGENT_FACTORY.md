# Agent factory (operators)

Local, markdown-driven factory that turns [`ROADMAP.md`](./ROADMAP.md) into human-reviewed PRs. Not part of the forma app-author guide (`doc/user/`).

## Who vs what

| Kind | Path | Meaning |
|------|------|---------|
| **Agents** (who) | `.cursor/agents/*.md` | Named roles with isolated context; Orch delegates to them |
| **Skills / commands** (what) | `.cursor/skills/`, `.cursor/commands/` | Procedures—especially OpenSpec `/opsx-*` |

## Roles (agents)

| Agent | Does | Must not |
|-------|------|----------|
| **planner** | Interview you; edit ROADMAP **intent** | Implement product code; merge |
| **orch** | Interpret ROADMAP; stack/sequence; delegate to architector/builder; update status/PR | Merge; cloud Automations as control plane; guess when unsure |
| **architector** | Explore→propose (`/opsx-*`); decide alone; record decisions; design review replies | Apply/implement feature code; merge |
| **builder** | Apply→`task test`/`lint`→archive→PR (**Architector decisions** required)→CI→wait for your merge | Merge; invent other roadmap items |
| **debt** | Analyze; **open propose PR**; comment loop | ROADMAP-only output; merge; silent mass refactors |

## Control plane

- **`doc/ROADMAP.md`** is the plan and live progress surface (status / branch / PR).
- **Human merge only.** Agents wait; they do not merge.

## Runtime

- **Local only** (checkouts / git worktrees). No Cursor cloud agents or Automations as the factory scheduler.
- Optional Node helpers: built-ins first; **ask before adding npm dependencies**.
- Prefer-reuse is always-on: `.cursor/rules/prefer-reuse-over-handroll.mdc` (`fn/` exempt; see also `prefer-fn`).

## Happy path

1. **planner** — shape the roadmap.
2. **orch** — pick ready work; ask you if concurrency/order is unclear.
3. **architector** → **builder** per change → you review (challenge decisions in PR comments) → you merge.
4. Periodically **debt** → review the propose PR → promote items via planner/orch.
