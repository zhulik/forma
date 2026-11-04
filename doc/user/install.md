# Install

## Prerequisites

| Tool | Why |
|------|-----|
| **Go** | Application code and the `forma` CLI. Match the version in [`.tool-versions`](../../.tool-versions) or `go.mod`. Building the CLI needs **CGO** (a C toolchain) for tree-sitter. |
| **Node.js** | Web tree (`npm install`, shadcn CLI during **`forma sync`** / **`forma add`** / **`forma init`**). Not required for **`forma generate`** alone. |
| **GopherJS** | Compile Go to JavaScript for the browser (bridge available; **`forma dev`** watch loop is P3d). |

Optional: [asdf](https://asdf-vm.com/) with `.tool-versions` for pinned Go, Node, and Task.

## Install the forma CLI

From a clone of this repository (until the first release tag is published):

```bash
go install github.com/zhulik/forma/cmd/forma@latest
```

Or build from source:

```bash
cd cmd/forma && go build -o forma .
```

Verify:

```bash
forma version
```

## Create your first app

```bash
mkdir myapp && cd myapp
forma init github.com/you/myapp
```

What **`forma init`** does today:

1. Requires an **empty** directory (or creates one with `-d`).
2. **`go mod init`** with your module path (no git remote setup).
3. Applies the embedded **`minimal`** template: `main.go`, `forma.yaml`, **`generator.go`**, **`Taskfile.yaml`**, a **`web/`** Vite + React + shadcn-ready tree, and a hello layout using `github.com/zhulik/forma/ui`.
4. Runs **`forma sync`**: **`npm install`**, **`shadcn add`** for starter components (button, input, card), **codegen** into **`internal/ui/`** plus **`web/src/registry.gen.ts`**, Go module tidy, and a **`require`** on the forma library.

Use your real module path (e.g. `github.com/you/myapp`), not a placeholder.

**Local forma checkout:** before the library is on the public module proxy, set **`FORMA_REPLACE`** to the absolute path of the forma repo when running init (CI in this repository does this automatically):

```bash
FORMA_REPLACE=/path/to/forma forma init github.com/you/myapp
```

**Config only:** in an existing tree, generate **`forma.yaml`** without touching Go sources:

```bash
forma init --only-config github.com/you/myapp
```

Fails if **`forma.yaml`** already exists. See [Manual setup](manual-setup.md) for the equivalent hand-written steps.

**Browser loop:** **`forma dev`** is not available yet (ROADMAP P3d).

## Learn the DSL today (no new app)

You can also learn from this repository:

1. Read [`main.go`](../../main.go) (todo demo).
2. Run `task test` and `task lint` from the repo root.

That exercises the **authoring model**. Browser builds use the GopherJS bridge (`UIMount` → `formaHost.setTree`); native `go test` keeps a no-op stub. One-command **`forma dev`** remains P3d.

## Next steps

- [CLI reference](cli.md)
- [Manual setup](manual-setup.md)
- [forma.yaml](forma-yaml.md)
- [Writing UI](writing-ui.md)
