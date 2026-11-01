# Proposal

## Why

Compound shadcn/Radix widgets need multiple distinct regions (trigger, content, panels) without bloating the Go API with mirror types such as `DialogTrigger`. Authors also need grouping without extra layout DOM. P1 deferred `Fragment` and slot protocol details; extending the UI DSL now keeps the author surface small before host and codegen work.

## What Changes

- Add **`Fragment`**: layout-neutral grouping with no extra host element.
- Add **explicit named slots** on the render tree (`Node.Slots`); compound **catalog** widgets are **`*.gen.go`** with **field-only** slot fields (same as `Button` is generated, not hand-written).
- Add **`WidgetMeta`**, slot helpers (`RenderComponentSlot`, `PutSlot`, `SlotKeyFromField`) for generated compound `Render` methods.
- **Codegen (P3):** derive compound shape from **shadcn TSX export analysis** (root + prefixed parts → one struct); optional per-file override when heuristics fail—not a hand-maintained widget list.
- **Out of scope:** React host, full codegen implementation, P2 user docs, `Slot("name", c)` helper.

## Capabilities

### New Capabilities

_(none)_

### Modified Capabilities

- `ui-dsl`: Fragment; `Node.Slots`; field-only compound widgets; generated catalog compounds; `WidgetMeta` / slot helpers.

## Impact

- **`ui/`**: `Fragment`, `Node.Slots`, `WidgetMeta`, slot helpers; **no** hand catalog compounds.
- **`doc/DESIGN.md`**: composition + codegen table for compounds.
- **P3 codegen/host**: TSX export analysis, compound template, registry slot map from same analysis.
