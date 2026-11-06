# Install

## Prerequisites

| Tool | Why |
|------|-----|
| **Go** | Application code and the `forma` CLI. Match the version in [`.tool-versions`](../../.tool-versions) or `go.mod`. Building the CLI needs **CGO** (a C toolchain) when required by dependencies. |
| **Node.js** | Web tree (`npm install` during **`forma sync`** / **`forma init`**). Also required for **`forma dev`**. |
| **GoScript** | Compile Go to TypeScript for the browser; required on PATH for **`forma dev`**. Pin **v0.6.3**. |

Optional: [asdf](https://asdf-vm.com/) with `.tool-versions` for pinned Go, Node, and Task.

Install GoScript (pinned):

```bash
go install github.com/s4wave/goscript/cmd/goscript@v0.6.3
goscript compile --help
```

GoScript targets the same Go version as `go.mod` (currently **1.27**). Official Go Wasm remains a documented contingency only—not the supported author path.

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
3. Applies the embedded **`minimal`** template: `main.go`, `forma.yaml`, **`Taskfile.yaml`**, and a **`web/`** Vite + React host with **0build** CDN styling and owned-widget builtins.
4. Runs **`forma sync`**: **`npm install`**, Go module tidy, and a **`require`** on the forma library. No shadcn materialization or kit codegen.

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

**Browser loop:** after sync and with GoScript on PATH:

```bash
forma dev
```

Vite serves the host; GoScript + Vite lib rebuild write `web/public/app.js` on Go saves. Host asset edits use Vite HMR; Go edits need a page reload (not preserved Go heap). See [CLI — forma dev](cli.md#forma-dev).

## Run the in-repo todo demo

The canonical interactive demo is [`examples/todo`](../../examples/todo/) (nested module + `forma.yaml` + scaffold `web/`). It is **available** today—not planned:

```bash
cd examples/todo
forma sync
forma dev
```

Open the printed URL and confirm add / toggle-done / delete (empty state with zero tasks). Details: [examples/todo/README.md](../../examples/todo/README.md).

Root [`main.go`](../../main.go) is a short Hello-style smoke for the framework module only.

## Next steps


- [CLI reference](cli.md)
- [Manual setup](manual-setup.md)
- [forma.yaml](forma-yaml.md)
- [Writing UI](writing-ui.md)
