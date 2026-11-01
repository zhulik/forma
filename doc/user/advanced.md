# Advanced

## Override codegen output path

In **`forma.yaml`**:

```yaml
codegen:
  dir: pkg/widgets
  package: widgets
```

Use when a monorepo or import policy requires generated code outside **`internal/ui`**. After changing paths, run **`forma sync`**.

Framework types (`Component`, `Context`, layout) stay in **`github.com/zhulik/forma/ui`**.

## Custom Go components

Any struct that implements **`Component`** (`Render(context.Context) *Node`) can appear in the tree. Use layout helpers (`Row`, `Column`, `Fragment`, …) and compose without shadcn.

Register a stable **`Type`** string in the host registry when the renderer must map to a custom React wrapper (P3 host docs).

## Third-party React libraries

Pattern (see [DESIGN.md](../DESIGN.md)):

1. Thin **JS wrapper** in `web/` (refs, vendor events → plain data).
2. Entry in **ComponentRegistry**.
3. Go struct with props + callbacks (hand-written or narrow codegen).

Go never imports DOM types; the wrapper calls back into Go with sanitized payloads.

## Edit shadcn TSX

shadcn is **source in your repo** so you can Tailor Tailwind and behavior. After edits, run **`forma sync`**. Future codegen may distinguish CLI-managed vs owner-managed files; until then, treat TSX as yours and re-sync deliberately.

## Agents and automation

Agents should prefer outputting **Go** that composes generated structs and framework layout—not raw HTML/JS. When CLI exists, prefer **`forma add`** + documented patterns in [Writing UI](writing-ui.md) over inventing registry entries.

## Related

- [DESIGN.md](../DESIGN.md)
- [Troubleshooting](troubleshooting.md)
