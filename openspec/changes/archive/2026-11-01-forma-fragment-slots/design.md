# Design

## Context

See [proposal.md](./proposal.md). P1 locked composition via `.With` / `.Or` on layout widgets and `Component.Render` → `Node` with a primary child list. Compound shadcn/Radix widgets need multiple regions; the agreed author model is **field-only explicit slots**—no `DialogTrigger`-style Go types. `Fragment` closes the gap where authors need several roots in one slot or one `Component` value without using `Row`/`Column`.

Framework code in `ui/` provides `Node.Slots`, slot helpers, `Fragment`, and **`WidgetMeta`** / **`BaseComponent`**. **Catalog widgets (including `Dialog`) are `*.gen.go`**, like `Button`.

## Goals / Non-Goals

**Goals:**

- `Fragment` + `Node.Slots` + exported slot helpers for generated `Render` methods.
- **`WidgetMeta`** shared escape hatches; **`BaseComponent`** = `WidgetMeta` + `.With` / `.Or` (layout/leaves only).
- **Codegen (P3)** derives compound shape from shadcn **TSX export analysis** (root + prefixed parts → one struct, slot fields); emit matching registry slot map from the same analysis.
- Optional **per-file override** only when heuristics fail (Select, Tabs-class widgets).

**Non-Goals:**

- Hand-maintained catalog compounds (`dialog.go`) in `ui/`.
- Default **per-widget YAML manifest** (overrides are exception-only).
- React host implementation (P3 `forma-react-host`).
- Composition API `Slot("name", c)` or implicit slots via child node types.

## Decisions

### 1. Wire format: `Slots` alongside `Children`

**Choice:** `Slots map[string][]*Node` on `Node`. Compounds populate `Slots` only; layout/leaves use `Children` only.

### 2. Slot keys

**Choice:** `SlotKeyFromField("Trigger")` → `"trigger"`. Codegen uses the same rule for field names it emits from stripped export names (`DialogTrigger` → field `Trigger`).

### 3. Slot rendering helpers (hand-maintained, used by codegen)

**Choice:** Exported **`RenderComponentSlot`**, **`PutSlot`**, **`SlotKeyFromField`** in `ui/slot.go`. Generated compound `Render` builds `map[string][]*Node` and returns `Node{ Slots: …, Children: nil }`.

### 4. `WidgetMeta` vs `BaseComponent`

**Choice:** **`WidgetMeta`**: Key, Class, Style, ID, Attrs. **`BaseComponent`**: embeds `WidgetMeta` + children + `With`/`Or`. **Generated compounds** embed **`WidgetMeta` only** so `.With` is not promoted.

**Alternatives:** `CompoundMeta` duplicate — rejected; one meta type for all widgets.

### 5. Compound discovery (codegen, P3)

**Choice:** For each shadcn `.tsx` file:

1. Identify **root** export (e.g. `Dialog`).
2. Treat same-prefix sibling exports (`DialogTrigger`, `DialogContent`, …) as **slot sources**, not separate Go types.
3. Emit Go field names by stripping prefix (`Trigger`, `Content`); types `Component` unless collection heuristic applies.
4. Emit **`dialog.gen.go`** with compound template (embed `WidgetMeta`, slot fields, scalar props, `Render` using slot helpers).
5. Emit **registry** fragment: wire slot key → TSX subcomponent wrapper (`asChild` on trigger when convention says so).

**Alternatives:** Manual manifest for every widget — rejected as default; **override file** only for failures.

### 6. Collection compounds (Tabs, Accordion)

**Choice:** Heuristic or override: emit `Items []TabItem` (widget-specific item struct) when export pattern matches list/trigger/content; `Render` expands to slot node lists (e.g. `triggers`, `panels`). Exact expansion in P3 codegen spec.

### 7. `Fragment`

**Choice:** Hand-maintained `ui/fragment.go`; embeds full `BaseComponent`; `Type: "Fragment"`; host uses React Fragment.

### 8. Tests

**Choice:** External tests use a **`compoundProbe`** type in `slot_test.go` to exercise public slot helpers—**not** a shipped `Dialog` type.

## Risks / Trade-offs

- **[Risk] TSX heuristics wrong on edge components** → Mitigation: override config; fail codegen loudly.
- **[Risk] Go/registry slot key drift** → Mitigation: single codegen pass emits both.
- **[Risk] Agents call `.With` on compounds** → Mitigation: compounds do not embed `BaseComponent`.

## Migration Plan

Remove hand `Dialog`/`Tabs`/`CompoundMeta` from `ui/`. Archive merges spec delta. P3 codegen change implements TSX analysis + compound template.

## Open Questions

- Required-slot dev behavior — P3 runtime spec.
- Exact Tabs item struct naming per widget — codegen convention doc in P3.
