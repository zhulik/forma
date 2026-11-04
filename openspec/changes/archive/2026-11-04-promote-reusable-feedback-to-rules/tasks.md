# Tasks

## 1. Always-apply meta-rule

- [x] 1.1 Add `.cursor/rules/promote-reusable-feedback.mdc` with `alwaysApply: true`, covering reusable vs one-off classifier, sequential Q1 (durable rule) then Q2 (occurrence sweep), explicit skip phrasing, merge-into-existing guidance, and a brief acknowledgment when a rule is written — verify the file exists with valid frontmatter and those sections
- [x] 1.2 Cross-check the meta-rule against `one-clarifying-question.mdc` so Q1/Q2 stay one question per turn — verify the meta-rule states sequential asks and does not instruct batching both into one message

## 2. OpenSpec apply reinforcement

- [x] 2.1 Append one advisory line to `openspec/config.yaml` under `operations.apply.guidance` pointing agents at the promote-reusable-feedback meta-rule after reusable implementation feedback — verify `openspec instructions apply --change promote-reusable-feedback-to-rules --json` surfaces the new guidance entry when apply instructions load project operation guidance

## Workflow follow-up

- Archive the change after review (`/opsx-archive` or `/opsx-finalize`). No main-spec sync expected (`skip_specs: true`).
