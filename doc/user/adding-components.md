# Adding components

## Normal path: forma add

```bash
forma add dialog
```

1. Materializes **`dialog`** via the shadcn CLI under the configured web root (if missing).
2. Runs **`forma sync`**, which:
   - Installs npm dependencies if needed
   - Regenerates Go bindings in **`internal/ui/`** (or your `codegen.dir`)
   - Regenerates **`web/src/registry.gen.ts`**
   - Runs **`go mod tidy`**

You should not need to run **`npx shadcn add`** in the happy path. **`forma.yaml`** does not maintain a component name list.

## What to commit

| Yes | No |
|-----|-----|
| New or changed TSX under `web/` | `node_modules/` |
| New or changed `internal/ui/*.gen.go` | |
| Updated `web/src/registry.gen.ts` | |
| Lockfile under `web/` if sync changed deps | |

Generated Go and registry files are **committed** so CI can verify **`task check-generate`** (`go generate` then whole-repo **`git diff --exit-code`**) stays clean.

## Using a new widget in Go

After sync, import your generated package and use the struct the codegen emitted—for example `Button`, `Input`, or compound types with slot fields. Props and events are typed on the struct; events use **`func(ctx Context)`** (or variants documented for that widget).

## Starter set

**`forma init`** materializes a CLI-hardcoded starter set (`button`, `input`, `card`) during the first sync so hello world works without an immediate **`forma add`**.

## Manual / advanced

If you add or edit TSX by hand:

- Run **`forma generate`** (codegen only) or **`forma sync`** (npm + codegen + tidy).
- Prefer CLI-managed files under **`components/ui/`**.

See [Advanced](advanced.md).

## Related

- [CLI reference](cli.md)
- [forma.yaml](forma-yaml.md)
