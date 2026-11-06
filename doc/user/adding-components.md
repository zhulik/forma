# Owned widgets

forma ships a **todo-sized** owned widget surface. Import from **`github.com/zhulik/forma/ui`**—there is no shadcn `forma add` / kit codegen path.

## Built-in types

| Go type | Role |
|---------|------|
| `Button`, `Input`, `Checkbox` | Interactive |
| `Card`, `Text`, `Heading` | Content |
| `Row`, `Column`, `Container`, `Fragment` | Layout / grouping |

The scaffold host **`builtinRegistry`** maps each type to a thin React component that applies **0build** semantic classes (for example `.z-button`, `.z-button-primary`, `.z-card`, `.z-input`).

## Look: semantic first

Prefer semantic `.z-*` classes applied by the host from typed props (`Variant`, `Size`, …). Use **0build utilities** or the Go **`Class`** escape hatch for one-offs—not as the default authoring model.

Pinned CDN assets (scaffold `web/index.html`):

- CSS: `https://cdn.0build.dev/0build.css`
- Runtime: `https://cdn.0build.dev/0build.js` (`zRuntime`)

## Extending

1. **Compose** existing widgets in Go when possible.
2. **New primitive:** hand-write a Go struct in framework or app code and a matching host builtin under `web/src/host/widgets/`; register it in `builtinRegistry` (framework) or **`registry.custom.ts`** (app override / third-party).
3. **Third-party React:** thin JS wrapper → `registry.custom.ts` (merged last) → Go props/callbacks only.

## What not to do

- Do not expect **`forma add`** or **`forma generate`**.
- Do not treat shadcn TSX or `*.gen.go` as the widget source of truth.
- Do not add a Tailwind/PostCSS kit pipeline to the default scaffold.

## Related

- [CLI reference](cli.md)
- [Advanced](advanced.md)
- [Writing UI](writing-ui.md)
