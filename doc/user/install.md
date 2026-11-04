# Install

## Prerequisites

| Tool | Why |
|------|-----|
| **Go** | Application code and the `forma` CLI. Match the version in [`.tool-versions`](../../.tool-versions) or `go.mod`. Building the CLI needs **CGO** (a C toolchain) for tree-sitter. |
| **Node.js** | Web tree (`npm install`, shadcn CLI during **`forma sync`** / **`forma add`** / **`forma init`**). Also required for **`forma dev`**. Not required for **`forma generate`** alone. |
| **GopherJS** | Compile Go to JavaScript for the browser; required on PATH for **`forma dev`**. |

Optional: [asdf](https://asdf-vm.com/) with `.tool-versions` for pinned Go, Node, and Task.

Install GopherJS:

```bash
go install github.com/gopherjs/gopherjs@v1.21.0
export GOPHERJS_GOROOT=/path/to/go1.21.x   # GopherJS still compiles against Go 1.21
gopherjs version
```

**Known ceiling:** current GopherJS (`v1.21.x`) does not yet understand Go **1.27** module language (forma’s `go.mod`). Day-to-day `go test` / CLI work on Go 1.27; a full browser GopherJS compile of the forma module needs a newer GopherJS (or a temporary lower `go` directive) until upstream catches up. `forma:dev` still orchestrates Vite + rebuild once that toolchain can produce `web/public/app.js`.

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

**Browser loop:** after sync and with GopherJS on PATH:

```bash
forma dev
```

Vite serves the host; GopherJS rebuilds `web/public/app.js` on Go saves. Host asset edits use Vite HMR; Go edits need a page reload (not preserved Go heap). See [CLI — forma dev](cli.md#forma-dev).

## Learn the DSL today (no new app)

You can also learn from this repository:

1. Read [`main.go`](../../main.go) (todo demo).
2. Run `task test` and `task lint` from the repo root.

That exercises the **authoring model**. To run interactive UI in a browser, use an initialized **app** project with **`forma sync`** + **`forma dev`** (the framework repo itself has no app `web/` tree).

## Next steps

- [CLI reference](cli.md)
- [Manual setup](manual-setup.md)
- [forma.yaml](forma-yaml.md)
- [Writing UI](writing-ui.md)
