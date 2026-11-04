# Proposal

## Why

Hand-written slice/map loops still appear beside `fn` helpers, and the prefer-fn Cursor rule only soft-suggests conversion. Tightening the rule and extending `fn` (including fallible collect helpers) keeps author and agent code consistent and reviewable.

## What Changes

- Split `fn` into `slices.go` / `maps.go`; add **`MapErr`** / **`FilterMapErr`** (fail-fast) and map-entry analogues (`MapValues`, `MapEntries`, `FilterEntries`, `FilterMapEntries`, plus Err variants as needed).
- Convert fallible **collecting** loops (e.g. codegen analyze) to `MapErr` / `FilterMapErr`; leave side-effect **each** walks as raw `for`.
- Convert remaining replaceable slice transforms where a helper fits.
- **Stricten** `.cursor/rules/prefer-fn.mdc`: forbid raw `for` except an explicit allowlist; ban no-op proxy lambdas when passing named functions into `fn`.
- Light author-guide touch so `fn` catalog mentions stay accurate (`MapErr`, map helpers).

## Capabilities

### New Capabilities

- (none)

### Modified Capabilities

- `ui-dsl`: Functional helper package contract — slice/map transforms and fallible collect helpers authors and the framework share.
- `dev-tooling`: Agent-oriented Go rules MUST include the strict prefer-fn allowlist and no-op-proxy ban.
- `ui-docs`: Author guide `fn` material MUST mention fallible collect helpers and map transforms at the same level as `Map` / `Filter`.

## Impact

- Package `github.com/zhulik/forma/fn` (API additive; early-stage, no compat shims).
- Call sites under `cmd/forma/internal/codegen` (and any other replaceable loops).
- `.cursor/rules/prefer-fn.mdc`.
- `doc/user/` writing-UI / project-layout mentions of `fn`.
- Out of scope: golangci custom analyzer for `for`; `EachErr`; product runtime behavior; ROADMAP phase changes.
