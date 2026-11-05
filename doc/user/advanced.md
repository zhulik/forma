# Advanced

## Override codegen output path

In **`forma.yaml`**:

```yaml
codegen:
  dir: pkg/widgets
  package: widgets
```

Use when a monorepo or import policy requires generated code outside **`internal/ui`**. After changing paths, run **`forma generate`** or **`forma sync`**.

Framework types (`Component`, `Context`, layout) stay in **`github.com/zhulik/forma/ui`**.

## Per-source codegen overrides (maintainers)

When TSX analysis misses a prop or classifies an event incorrectly, place a **`<file>.override.yaml`** next to the component source (e.g. `input.tsx` → `input.override.yaml`). Overrides merge on top of analysis and can set Go field names/types (including `func(ctx ui.Context, value string)` and `func(ctx ui.Context, checked bool)`).

forma also ships **embedded maintainer defaults** for shadcn `input` and `checkbox` (forwardRef sources without a Prop interface). Sidecar files win over embeds. Format and merge rules are documented in `cmd/forma/internal/codegen/override.go`. Prefer overrides over hand-editing `*.gen.go`.

## Custom Go components

Any struct that implements **`Component`** (`Render(context.Context) *Node`) can appear in the tree. Use layout helpers (`Row`, `Column`, `Fragment`, …) and compose without shadcn.

Register a stable **`Type`** string in the host registry when the renderer must map to a custom React wrapper (see below).

## Custom / third-party registry entries

The host merges registries in this order (later keys win):

1. **Builtins** — framework types (`Fragment`, `Text`, `Row`, …) under `web/src/host/`
2. **Generated** — `web/src/registry.gen.ts` (codegen-owned; never hand-edit)
3. **Custom** — `web/src/registry.custom.ts` (hand-written merges)

To add a third-party React wrapper:

1. Write a thin **JS wrapper** in `web/` that sanitizes vendor events to plain data (or rely on the host prop-wiring sanitizers).
2. Export an entry from **`registry.custom.ts`** with the same shape as generated entries (`{ component, slots? }`).
3. Expose a Go struct with props + callbacks (hand-written or via codegen for shadcn files).

`main.tsx` already merges `customRegistry` last. Do **not** patch `registry.gen.ts`.

## Text-input hybrid policy

For **`Input`** (and similar string fields), the host keeps ephemeral React state while the field is focused so typing survives tree refreshes. Domain state still lives in Go: the sanitized **`OnChange`** callback receives a plain string. When the field is unfocused, the displayed value syncs from the tree’s value prop. See [Troubleshooting](troubleshooting.md) if caret jumps persist.

## Edit shadcn TSX

shadcn is **source in your repo** so you can tailor Tailwind and behavior. After edits, run **`forma generate`** (or **`forma sync`**). Prefer CLI-managed files under **`components/ui/`**.

## Agents and automation

Agents should prefer outputting **Go** that composes generated structs and framework layout—not raw HTML/JS. When CLI exists, prefer **`forma add`** + documented patterns in [Writing UI](writing-ui.md) over inventing registry entries.

## Related

- [DESIGN.md](../DESIGN.md)
- [Troubleshooting](troubleshooting.md)
