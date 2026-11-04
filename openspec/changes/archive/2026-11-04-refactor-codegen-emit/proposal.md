# Proposal

## Why

The new in-process `forma generate` codegen landed as a dense `analyze.go` / string-joined `emit.go` pair. Helpers that belong in `fn` or CLI `pkg/` are local, slice transforms ignore existing `fn` helpers, and Go emission via concatenation is harder to maintain than a typed emitter. Refactor now while the surface is small and goldens already pin behavior.

## What Changes

- Move generic naming helpers (`ToPascal`, exported Go ident helpers) into `cmd/forma/pkg/ident` with unit tests
- Replace local `uniqSorted` with `fn.UniqSorted`; extend `fn` with `FlatMap` / `GroupBy` where registry aggregation needs them; convert other eligible slice loops to `fn`
- Keep analyze in `internal/codegen`; split emit into `emit_go.go` (jennifer) and `emit_registry.go` (`text/template`)
- Keep generated output byte-identical to existing goldens

## Capabilities

### New Capabilities

<!-- none — pure internal refactor -->

### Modified Capabilities

<!-- none — ui-codegen behavior unchanged; see skip_specs in .openspec.yaml -->

## Impact

- Code: `cmd/forma/internal/codegen/*`, new `cmd/forma/pkg/ident`, possibly `fn/`
- Dependencies: add `github.com/dave/jennifer` to `cmd/forma`
- Author-facing CLI / forma.yaml / docs: none (opt out of doc/user updates)
- Out of scope: TSX analysis heuristics, override hooks, new widget kinds, CLI UX
