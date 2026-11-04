# CLI reference

The forma CLI lives in **`cmd/forma/`** (nested Go module). Normative behavior is in OpenSpec **`forma-cli`**.

## Commands

| Command | Status | Summary |
|---------|--------|---------|
| `forma init <module>` | **Available** | Empty dir, `minimal` template, sync with starter shadcn set (needs Node) |
| `forma init --only-config <module>` | **Available** | Write `forma.yaml` if missing |
| `forma sync` | **Available** | npm install, codegen via generate path, `go mod tidy` (no shadcn unless init/add) |
| `forma generate` | **Available** | Codegen only from `forma.yaml` (TSX → Go + registry) |
| `forma add <component>` | **Available** | `shadcn add` then sync |
| `forma version` | **Available** | Print CLI version |
| `forma dev` | Planned | GopherJS watch + host dev server (P3d) |

Global flags:

| Flag | Env | Purpose |
|------|-----|---------|
| `--log-level` | `FORMA_LOG_LEVEL` | slog level (debug, info, warn, error) |

Sync/init/add honor **`FORMA_REPLACE`**: local path for `go mod edit -replace=github.com/zhulik/forma=…` (development and CI before publish).

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
- Ends with **`forma sync`** that materializes starter components (`button`, `input`, `card`) then runs codegen.
- Writes root **`generator.go`** (`//go:generate forma generate`) and a **`Taskfile.yaml`** with **`check-generate`**.

`--only-config`:

- Writes **`forma.yaml`** only when the file does not exist; **fails** if it already exists.
- Does not run `go mod init` or sync.

## forma sync

1. Install npm dependencies under **`web.path`** (`npm install`).
2. Optionally run **`shadcn add`** when invoked from **init** (starter set) or **`forma add`** (named component). Plain **`forma sync`** does not re-fetch shadcn.
3. Run the same codegen path as **`forma generate`** (Go bindings + `web/src/registry.gen.ts`).
4. Add **`require github.com/zhulik/forma`**, optional replace from **`FORMA_REPLACE`**, **`go mod tidy`**.

Requires **Node.js** and **`npm`** on PATH. Without them, sync fails with a clear error.

## forma generate

```bash
forma generate
forma generate -d ./myapp
```

Reads **`forma.yaml`**, regenerates **`codegen.dir/*.gen.go`** and **`web/src/registry.gen.ts`** from **`web/src/components/ui/*.tsx`**. Does **not** run npm, shadcn, or `go mod tidy`, and **does not require Node.js** (codegen is pure Go).

Scaffolded apps expose this via **`go generate`** (see **`generator.go`**). CI drift check:

```bash
task check-generate   # go generate ./... && git diff --exit-code (whole repo)
```

There is no **`forma generate --check`** or **`forma sync --check`**.

## forma add

```bash
forma add dialog
```

Runs **`shadcn add dialog`** under the web root, then **`forma sync`**. Does not edit a component list in **`forma.yaml`**—the catalog is committed TSX plus generated output.

## forma dev

Planned P3d. Will assume **`forma sync`** has been run at least once.

## Related

- [Install](install.md)
- [Manual setup](manual-setup.md)
- [forma.yaml](forma-yaml.md)
- [Adding components](adding-components.md)
