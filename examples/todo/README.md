# examples/todo

Runnable in-browser todo demo for forma: default shadcn Card / Input / Button / Checkbox, responsive composer Class, in-memory CRUD via `forma dev`.

This module’s `go.mod` already has `replace github.com/zhulik/forma => ../..` for the monorepo checkout. External copies should use a published forma module or set `FORMA_REPLACE` when running `forma sync` / `forma init`.

## Run

From this directory (`examples/todo`):

```bash
# optional if node_modules / generated artifacts are missing
forma sync

forma dev
```

Open the Vite URL printed by `forma dev` (typically `http://localhost:5173`). Confirm:

- add / toggle done / delete todos
- composer stacks Input above Add on a narrow viewport
- Input and Add sit side-by-side from the `sm` breakpoint up
