# Tasks

## 1. Shared helpers

- [x] 1.1 Add `cmd/forma/pkg/ident` with `ToPascal`, `ToExportedGoName`, and `ReceiverName` (behavior matching current analyze/emit helpers) and verify unit tests in `ident_test` pass
- [x] 1.2 Replace local `uniqSorted` with `fn.Uniq` + `slices.Sort` (or add `fn.UniqSorted` if preferred) and verify `fn` tests still pass
- [x] 1.3 Wire codegen to `pkg/ident` / `fn` for naming and dedupe; convert eligible analyze/runner slice loops to `fn`; verify `go test ./internal/codegen/...` under `cmd/forma` still passes

## 2. Emit rewrite

- [x] 2.1 Add `github.com/dave/jennifer` to `cmd/forma` and rewrite Go leaf/compound emission with jennifer; verify button/accordion `*.gen.go` goldens still match (refresh only if formatting change is intentional and reviewed)
- [x] 2.2 Rewrite `emitRegistry` with `text/template` (split files as needed); verify registry goldens still match
- [x] 2.3 Apply `fn` to remaining emit/runner slice transforms where applicable; verify no leftover local string-concat Go emitter remains

## 3. Verification

- [x] 3.1 Run `task test` and `task lint` and verify both succeed

## Workflow follow-up

- Archived via `/opsx-finalize` (update + archive + commit).
