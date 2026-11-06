# CLI reference

The forma CLI lives in **`cmd/forma/`** (nested Go module). Normative behavior is in OpenSpec **`forma-cli`**.

## Commands

| Command | Status | Summary |
|---------|--------|---------|
| `forma init <module>` | **Available** | Empty dir, `minimal` template, sync (needs Node) |
| `forma init --only-config <module>` | **Available** | Write `forma.yaml` if missing |
| `forma sync` | **Available** | `npm install` under `web.path`, then `go mod tidy` |
| `forma version` | **Available** | Print CLI version |
| `forma dev` | **Available** | Vite host + GoScript rebuild loop |

There is **no** author-facing **`forma add`** or **`forma generate`**. Widgets come from the framework `ui` package and scaffold host builtins.

Global flags:

| Flag | Env | Purpose |
|------|-----|---------|
| `--log-level` | `FORMA_LOG_LEVEL` | slog level (debug, info, warn, error) |

Sync/init honor **`FORMA_REPLACE`**: local path for `go mod edit -replace=github.com/zhulik/forma=…` (development and CI before publish).

## forma init

```bash
forma init github.com/you/myapp
forma init github.com/you/myapp -d ./myapp
forma init --only-config github.com/you/myapp
forma init --template minimal github.com/you/myapp
```

Full init:

- Target directory must be **empty** (only `.` and `..`).
- Does **not** run `git init`.
- Ends with **`forma sync`** (npm install + Go module reconcile). Does **not** materialize shadcn or run kit codegen.
- Scaffold includes 0build CDN links in **`web/index.html`** and host builtins under **`web/src/host/`**.

`--only-config`:

- Writes **`forma.yaml`** only when the file does not exist; **fails** if it already exists.
- Does not run `go mod init` or sync.

## forma sync

1. Install npm dependencies under **`web.path`** (`npm install`).
2. Add **`require github.com/zhulik/forma`**, optional replace from **`FORMA_REPLACE`**, **`go mod tidy`**.

Requires **Node.js** and **`npm`** on PATH. Without them, sync fails with a clear error.

Sync does **not** invoke shadcn or regenerate kit bindings.

## forma dev

Starts the local browser loop: Vite for the web tree and GoScript rebuilds when Go sources change. Read **`forma.yaml`** for web root and mount. Assume **`forma sync`** has been run at least once.

**Verify with the in-repo demo:** from [`examples/todo`](../../examples/todo/), run **`forma sync`** then **`forma dev`**, open the printed URL, and confirm empty state plus add / toggle-done / delete. See [examples/todo/README.md](../../examples/todo/README.md) and [Install](install.md).

Host asset edits use Vite HMR; Go rebuilds rewrite `web/public/app.js` and need a **page reload** (Go heap is not preserved).

## Related

- [forma.yaml](forma-yaml.md)
- [Project layout](project-layout.md)
- [Owned widgets](adding-components.md)
