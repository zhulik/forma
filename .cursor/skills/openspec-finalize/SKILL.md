---
name: openspec-finalize
description: >-
  Finalize an OpenSpec change non-interactively: reconcile planning artifacts
  (opsx-update, no prompts), archive with delta-spec sync (opsx-archive, no
  prompts), then create a git commit. Use when the user says "openspec finalize",
  "opsx finalize", or "/opsx-finalize".
allowed-tools: Bash(openspec:*), Bash(git:*)
license: MIT
compatibility: Requires openspec CLI and a git working tree.
metadata:
  author: forma
  version: "1.0.0"
---

Finalize a change end-to-end **without asking for confirmation**.

Pipeline (always in this order):

1. **Update** — run the `/opsx-update` workflow with auto-apply (no confirmations).
2. **Archive** — run the `/opsx-archive` workflow with **sync** and no prompts.
3. **Commit** — create one git commit for the resulting tree.

Do **not** pause between steps for approval. Do **not** ask yes/no, sync-or-skip, or “proceed despite incomplete …?” questions. The only allowed question is **change selection** when the target change cannot be resolved (see below).

Read and follow the project skills (or command bodies) for update and archive for mechanics, CLI flags, store selection, and project-check rules — then apply the **non-interactive overrides** in this skill. Prefer reading:

- `.cursor/skills/openspec-update-change/SKILL.md` (or `/opsx-update`)
- `.cursor/skills/openspec-archive-change/SKILL.md` (or `/opsx-archive`)
- `.cursor/skills/openspec-sync-specs/SKILL.md` when syncing delta specs

Keep the same `--store <id>` stickiness and OpenSpec project-check behavior as those workflows.

---

## Input

Optionally a change name after `/opsx-finalize` (e.g. `/opsx-finalize refactor-codegen-emit`).

## Select the change

1. If a name is provided, use it.
2. Else infer from conversation context.
3. Else if exactly one active change exists (`openspec list --json`), use it.
4. Else **ask once** which change to finalize (this is the only allowed prompt), then continue with no further questions.

Announce: `Using change: <name>` and how to override (`/opsx-finalize <other>`).

---

## 1. Update (no questions)

Execute `/opsx-update` for the selected change with these overrides:

- Treat the request as a full coherence review against the implemented work and conversation decisions (not a partial “propose only” pass).
- **Do not** show revisions for user confirmation. Apply every needed edit to existing planning artifacts immediately.
- Still obey update guardrails that are not about prompting:
  - Planning artifacts only — do not edit implementation code in this step.
  - Edit only concrete `existingOutputPaths` (plus a missing file under an already partially populated glob artifact when reconciliation requires it — create it without asking).
  - Leave `skipped` artifacts untouched.
  - Do not create brand-new frontier artifacts that have empty `existingOutputPaths` and status `ready`/`blocked`; if those block a coherent finalize, stop and report what is missing (no soft-prompt loop).
- If the change is already coherent, write nothing and continue.
- Do **not** stop to suggest `/opsx-apply` or `/opsx-archive` — continue to archive in this same run.

Briefly note which artifacts were revised (or that none needed changes), then proceed.

---

## 2. Archive with sync (no questions)

Execute `/opsx-archive` for the same change with these overrides:

### Completion warnings

- If artifacts are incomplete (not `done`/`skipped`) or tasks are incomplete: **print a warning** and **proceed anyway**. Do not ask.
- If task/status lookup fails or counts are invalid: **stop** and report (hard failure, not a prompt).

### Delta spec sync

- If there are **no** delta specs (`specs` missing or `existingOutputPaths` empty): archive with **Specs: No delta specs**.
- If delta specs exist:
  - Run the same sync assessment as archive.
  - If any capability is **sync-blocked**: **stop**, explain the blocker, do not archive. (Cannot fulfill “archive with sync”.)
  - Otherwise, if changes are needed **or** already synced: **always sync** (equivalent to choosing “Sync now” / ensure main specs match), run `openspec-sync-specs` inline, verify the post-sync contract from the archive skill, then archive.
  - Never choose “Archive without syncing” or “Cancel”.
- On sync failure or post-sync mismatch: **stop**, do not move `changeRoot`, do not commit.

### Perform archive

Same as archive skill: `mkdir -p` archive dir, date-prefix target name, fail if target exists, `mv` `changeRoot` into archive. Preserve `.openspec.yaml` with the directory.

Show the archive success summary (include any incomplete-artifact/task warnings).

---

## 3. Commit

After a successful archive, create a **single git commit** for the finalize outcome (implementation + updated/archived OpenSpec files + any synced main specs).

Follow `skillset-git-github` commit defaults and the user’s committing-changes-with-git rules:

1. Run in parallel: `git status`, `git diff` / `git diff --cached`, `git log -5 --oneline` (style).
2. Stage relevant files (do not stage secrets). Prefer explicit paths; avoid wiping unrelated dirty files — if unrelated changes are present, include only finalize-related paths or stop and report the conflict (no interactive “stage everything?” prompt: prefer staging the change’s touched paths; if unsure what is unrelated, stop with a clear list).
3. Draft a conventional commit message focused on **why** (1–2 short sentences). Prefer deriving subject from the change name / proposal Why (e.g. `refactor(cli): tidy forma generate codegen helpers`).
4. Commit via HEREDOC:

```bash
git commit -m "$(cat <<'EOF'
Subject line here.

Optional body.

EOF
)"
```

5. Run `git status` after and show the new commit briefly.

Do **not** push unless the user separately asked to push.

If there is nothing to commit after archive, report that and stop successfully (archive still happened).

---

## Output

```markdown
## Finalize Complete

**Change:** <name>
**Update:** <revised artifacts | no artifact changes>
**Archive:** <archive path>
**Specs:** <synced | no delta specs>
**Commit:** <hash and subject, or none>

<Warnings if incomplete artifacts/tasks were overridden>
```

## Guardrails

- No confirmation prompts except unavoidable change-selection ambiguity.
- Always sync when delta specs exist and sync is possible; never skip sync silently.
- Stop on sync-blocked, sync failure, archive-target collision, invalid task counts, or unrelated dirty-tree ambiguity.
- Update step does not edit product code; commit may include code already changed outside this skill.
- Never force-push; never amend unless the user’s git amend rules are fully satisfied (finalize creates a **new** commit by default).
