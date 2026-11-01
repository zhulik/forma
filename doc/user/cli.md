# CLI reference

The forma CLI lives in **`cmd/forma/`** (nested Go module). Normative behavior is in OpenSpec **`forma-cli`**.

## Commands

| Command | Status | Summary |
|---------|--------|---------|
| `forma init <module>` | **Available** | Empty dir, `minimal` template, stub `sync` |
| `forma init --only-config <module>` | **Available** | Write `forma.yaml` if missing |
| `forma sync` | **Stub** | Go module + forma `require` + `go mod tidy` only |
| `forma version` | **Available** | Print CLI version |
| `forma add <component>` | Planned | Append component in `forma.yaml`, run `sync` |
| `forma dev` | Planned | GopherJS watch + host dev server (P3d) |

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
- Ends with **stub **`sync`**.

`--only-config`:

- Writes **`forma.yaml`** only when the file does not exist; **fails** if it already exists.
- Does not run `go mod init` or sync.

## forma sync

**Today (stub):** reads **`forma.yaml`**, adds **`require github.com/zhulik/forma`**, optional replace from **`FORMA_REPLACE`**, runs **`go mod tidy`**. Host install, shadcn sources, codegen, and registry update are **planned** (P3).

**Target (full sync):**

1. Install or update host dependencies under `host.path`.
2. Materialize shadcn sources per `ui.components`.
3. Codegen into `codegen.dir`.
4. Update the host component registry.
5. `go mod tidy`.

## forma add

```bash
forma add dialog
```

Not implemented yet; returns a clear error. Prefer waiting for **`forma add`** over hand-editing the component list when it ships.

## forma dev

Planned P3d. Will assume **`forma sync`** has been run at least once.

## Related

- [Install](install.md)
- [Manual setup](manual-setup.md)
- [forma.yaml](forma-yaml.md)
- [Adding components](adding-components.md)
