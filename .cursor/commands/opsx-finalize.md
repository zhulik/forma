---
name: "/opsx-finalize"
id: "opsx-finalize"
category: "Workflow"
description: "Finalize a change: update artifacts, archive with sync, and commit (no prompts)"
---

Finalize an OpenSpec change non-interactively.

**Read and follow** `.cursor/skills/openspec-finalize/SKILL.md` for the full workflow.

Under the hood:

1. **`/opsx-update`** — reconcile planning artifacts; apply all revisions **without asking**.
2. **`/opsx-archive`** — archive with **delta-spec sync**, **no confirmation prompts** (warn and proceed on incomplete artifacts/tasks; stop if sync is blocked or fails).
3. **Commit** — one conventional git commit for the resulting tree (no push unless separately requested).

**Input**: Optional change name after `/opsx-finalize` (e.g. `/opsx-finalize refactor-codegen-emit`). If omitted, infer from context or the sole active change; ask only when selection is ambiguous.

**Store selection / project check**: Same rules as other `/opsx-*` commands — see the finalize skill (and nested update/archive skills) for `--store`, `openspec list --json` root checks, and never creating `openspec/` as a side effect.
