# Proposal

## Why

After `/opsx-apply` (and similar implementation work), reusable coding guidance is often applied as a one-off fix and lost unless the user also asks to create a Cursor rule. Agents should offer durable-rule promotion and optional repo-wide cleanup without requiring that second explicit ask.

## What Changes

- Add an always-apply Cursor meta-rule that defines when feedback looks reusable and how the agent must ask about durable rules and occurrence sweeps
- Reinforce the same behavior with a short OpenSpec apply operation-guidance line (no hooks)

## Capabilities

### New Capabilities

<!-- none — agent workflow config only -->

### Modified Capabilities

<!-- none — no product/capability behavior change; see skip_specs in .openspec.yaml -->

## Impact

- New file under `.cursor/rules/` (always-apply meta-rule)
- One-line addition to `openspec/config.yaml` `operations.apply.guidance`
- No product DSL, CLI, runtime, main specs, or `doc/user/` changes (author-facing product UX out of scope; opt-out from user-doc update)

## Out of scope

- Automatic silent rule writes without asking
- Cursor hooks (`stop` / `beforeSubmitPrompt`) for promotion
- Changing OpenSpec CLI itself
- Spec deltas under `dev-tooling` or elsewhere
- Migrating or rewriting existing Go convention rules beyond referencing them
