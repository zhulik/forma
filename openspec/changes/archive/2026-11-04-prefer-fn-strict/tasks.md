# Tasks

## 1. Extend `fn` package

- [x] 1.1 Split `fn/fn.go` into `slices.go` and `maps.go` (move existing slice helpers unchanged) and verify `go test ./fn/...` still passes
- [x] 1.2 Add slice `MapErr` and `FilterMapErr` (fail-fast) with unit tests covering success and first-error paths
- [x] 1.3 Add map helpers `MapValues`, `MapEntries`, `FilterEntries`, `FilterMapEntries` (and Err variants as designed) with unit tests

## 2. Convert call sites and tighten rule

- [x] 2.1 Convert fallible collecting loops (at least `AnalyzeComponentsDir`) to `FilterMapErr` / `MapErr` and verify codegen tests still pass
- [x] 2.2 Convert other replaceable slice transforms (e.g. `ident.ToPascal` if it fits) while leaving allowlisted loops (runner each, tree-sitter index, `pickRoot`, table tests) as raw `for`
- [x] 2.3 Rewrite `.cursor/rules/prefer-fn.mdc` with forbid-by-default, explicit allowlist, and no-op proxy ban; verify the rule file documents all design allowlist cases

## 3. Author docs

- [x] 3.1 Update `doc/user/` (`writing-ui.md` and/or `project-layout.md`) so `fn` mentions include fallible collect and map helpers; verify the guide pages name those helpers

## 4. Verification

- [x] 4.1 Run `task test` and `task lint` and verify both succeed
