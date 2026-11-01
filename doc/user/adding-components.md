# Adding components

## Normal path: forma add

> **Planned** until the CLI ships.

```bash
forma add dialog
```

1. Adds `dialog` to **`ui.components`** in `forma.yaml`.
2. Runs **`forma sync`**, which:
   - Installs or updates shadcn **source** under `web/…/components/ui/`
   - Regenerates Go bindings in **`internal/ui/`** (or your `codegen.dir`)
   - Updates the host registry so the tree type `Dialog` renders correctly

You should not need to run **`npx shadcn add`** in the happy path.

## What to commit

| Yes | No |
|-----|-----|
| Updated `forma.yaml` | `node_modules/` |
| New or changed TSX under `web/` | |
| New or changed `internal/ui/*.gen.go` | |
| Lockfile under `web/` if sync changed deps | |

Generated Go is **committed** so CI can verify `forma sync` produces a clean tree (same idea as `task check-generate` in the framework repo).

## Using a new widget in Go

After sync, import your generated package and use the struct the codegen emitted—for example `Button`, `Input`, or compound types with slot fields. Props and events are typed on the struct; events use **`func(ctx Context)`** (or variants documented for that widget).

## Starter set

**`forma init`** will include a small default list (e.g. button, input, card) so hello world works without an immediate **`forma add`**.

## Manual / advanced

If you add TSX by hand or patch shadcn files:

- Run **`forma sync`** to refresh codegen and registry.
- Sync overwrite policy for edited TSX will be documented with P3a codegen (prefer CLI-managed files until then).

See [Advanced](advanced.md).

## Related

- [CLI reference](cli.md)
- [forma.yaml](forma-yaml.md)
