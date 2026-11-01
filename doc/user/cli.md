# CLI reference

The forma CLI is the main entry point for creating and maintaining apps. Normative behavior is specified in OpenSpec **`forma-cli`**; implementation is planned in a separate change from this documentation.

## Commands

| Command | Status | Summary |
|---------|--------|---------|
| `forma init <module>` | Planned | `go mod init`, scaffold project, run `sync` |
| `forma add <component>` | Planned | Append component in `forma.yaml`, run `sync` |
| `forma sync` | Planned | Install host deps, materialize shadcn, codegen, registry, `go mod tidy` |
| `forma dev` | Planned | GopherJS watch + host dev server (P3d) |

## forma init

```bash
forma init github.com/you/myapp
```

- Runs in the **current directory** (should be empty or only contain files you intend to overwrite).
- Sets the Go module path; does **not** run `git init` or configure remotes.
- Ends with **`forma sync`** so the tree is consistent before you commit.

## forma add

```bash
forma add dialog
```

- Adds `dialog` to the component list in `forma.yaml` if missing.
- Runs **`forma sync`**.

Prefer **`forma add`** over hand-editing the component list; editing yaml is for advanced cases and codegen path overrides.

## forma sync

Single command to reconcile disk with **`forma.yaml`**:

1. Install or update host dependencies (e.g. `npm install` under `host.path`).
2. Ensure shadcn component **source** exists for each listed component.
3. Run **codegen** into `codegen.dir` (default `internal/ui`).
4. Update the host **component registry** for generated types.
5. Run **`go mod tidy`** so forma and generated imports resolve.

Run sync after pulling changes that touch `forma.yaml` or when CI reports drift. Future projects may expose `task sync` or CI `forma sync && git diff --exit-code`.

## forma dev

Starts the local browser loop: rebuild Go (GopherJS) and serve the host with fast refresh or watch.

- Assumes **`forma sync`** has been run at least once.
- Exact flags and port defaults will be documented when P3d ships.

## Related

- [forma.yaml](forma-yaml.md)
- [Adding components](adding-components.md)
