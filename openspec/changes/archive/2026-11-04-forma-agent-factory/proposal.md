# Proposal

## Why

Manual OpenSpec delivery does not scale to “ROADMAP → many stacked PRs” while still encoding forma’s worldview. We need a **local, markdown-driven agent factory** so Planner / Orch / Architector / Builder / Debt roles turn `doc/ROADMAP.md` into human-reviewed PRs without cloud agents or auto-merge.

## What Changes

- Add Cursor **custom subagents** under `.cursor/agents/` for five roles (**planner**, **orch**, **architector**, **builder**, **debt**). Keep OpenSpec **skills/commands** (`/opsx-*`) as the “what to do”; agents are the “who”.
- Treat **`doc/ROADMAP.md` as the control plane**: intent plus live status / branch / PR fields that Orch (and related roles) may update.
- Orch **interprets** roadmap prose for dependencies, stacking, and safe parallelism; **asks the human when unsure** before spawning.
- Split planning vs shipping: **Architector** (`/opsx-explore` → propose, decides alone) vs **Builder** (apply → verify → archive → PR → babysit CI → wait for human merge).
- PR descriptions **MUST** surface Architector decisions; both Architector and Builder may address review comments (design vs code).
- **Debt auditor** (manual / every N merges): analyze for reuse opportunities; **open a propose-oriented PR** (not silent ROADMAP-only).
- Add always-on Cursor rule **prefer reuse over hand-roll** (well-known third party → existing codegen → hand-write; **`fn/` exempt**).
- Optional tiny **Node built-in** helpers only if markdown reasoning is insufficient; **ask before adding any npm dependency**.
- Document factory conventions in-repo (short operator note); **no** `doc/user/` author-guide changes (agent/operator tooling, not forma app-author UX) — explicit opt-out.

## Capabilities

### New Capabilities

- `agent-factory`: Local roadmap-driven agent roles, handoffs, PR/review/merge gates, and debt propose PRs.

### Modified Capabilities

- `dev-tooling`: Extend agent-oriented Cursor rules to require prefer-reuse-over-handroll (third party / codegen before NIH; `fn/` exception).

## Impact

- **Touch:** `.cursor/agents/`, `.cursor/rules/`, `doc/AGENT_FACTORY.md`, ROADMAP factory section; OpenSpec `/opsx-*` skills unchanged as mechanics. No runtime/UI/CLI product code required for v1.
- **ROADMAP phase:** Meta / contributor automation (orthogonal to P3b–P3d product phases); noted in ROADMAP OpenSpec workflow section.
- **Out of scope:** Cloud agents, Cursor Automations as the control plane, auto-merge, Graphite-as-a-service, deep builder prompt polish beyond the agent contract, mandatory always-on review poller (on-demand / orch-driven review advance is enough for v1).
- **Dependencies:** Existing OpenSpec skills (`opsx-explore` / propose / apply / archive / finalize), local `gh`, git worktrees; Cursor SDK local spawn only if needed later — v1 uses Task/worktree delegation from **orch**.
