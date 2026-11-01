# Tasks

P1 — UI DSL extension (Fragment + field-only named slots). Verify with `go build ./...`, then `task test` and `task lint`.

## 1. Tree protocol

- [x] 1.1 Add `Slots map[string][]*Node` to `ui/node.go` and verify `go build ./...` succeeds
- [x] 1.2 Add exported slot helpers in `ui/slot.go` and unit tests (compoundProbe + Fragment; no hand catalog types)

## 2. Fragment and meta split

- [x] 2.1 Implement `Fragment` in `ui/fragment.go`; split `WidgetMeta` from `BaseComponent` in `ui/widget_meta.go` / `component.go`

## 3. No hand catalog compounds

- [x] 3.1 Remove hand-written `Dialog`/`Tabs`; tests exercise `RenderComponentSlot` / `PutSlot` / `SlotKeyFromField` only
- [x] 3.2 Keep `main.go` todo-only (no Dialog demo until generated type exists)

## 4. Documentation

- [x] 4.1 Update `doc/DESIGN.md` and change `design.md`: TSX export analysis for compounds; overrides optional; no `CompoundMeta`

## 5. Verification

- [x] 5.1 Run `task test` and `task lint` and verify both pass

## Workflow follow-up

- Archive the change and sync `openspec/specs/ui-dsl/spec.md` after review.
- P3 **`forma-ui-codegen`**: compound template, TSX export heuristics, registry slot map from same pass.
- P3 host: render `Node.Slots`.
- P2: user guide compounds section.
