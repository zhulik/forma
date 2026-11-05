# forma author guide

forma is a Go-native UI framework: you compose widgets in Go; a React + shadcn host renders them. This guide is the **primary onboarding path** for app authors—alongside the runnable demo [`examples/todo`](../../examples/todo/) and the [`ui-dsl`](../../openspec/specs/ui-dsl/spec.md) specification.

## Documentation is first-class

If a change affects **author experience**—CLI commands, `forma.yaml`, project layout, install or sync workflow, or DSL rules you document here—that change **must update this guide in the same delivery slice**. Treat `doc/user/` like a public API: breaking UX without doc updates is incomplete work.

Contributors: see also OpenSpec capability **`ui-docs`** and `openspec/config.yaml` task/proposal rules.

## What works today

| Area | Status |
|------|--------|
| UI DSL (composition, slots, `Context`, `fn`) | **Available** — [`examples/todo`](../../examples/todo/), thin root [`main.go`](../../main.go) smoke, [`ui-dsl` spec](../../openspec/specs/ui-dsl/spec.md), `task test` |
| Author guide (this tree) | **Available** |
| `forma` CLI (`init`, `sync`, `generate`, `add`, `version`, `dev`) | **Available** — full sync + codegen + browser loop; [`forma-cli` spec](../../openspec/specs/forma-cli/spec.md) |
| React host (registry merge, renderer, sanitize, text-input policy) | **Available** — scaffold `web/src/host/` (P3b) |
| Go↔host bridge (`UIMount` / `RefreshUI` / callbacks) | **Available** — ROADMAP P3c; host registers `globalThis.formaHost.setTree` |
| Browser `forma dev` loop | **Available** — ROADMAP P3d (Vite host + GoScript rebuild; Go changes need page reload) |
| Runnable todo example | **Available** — [`examples/todo`](../../examples/todo/) (`forma sync` + `forma dev`) |
| Codegen from shadcn TSX | **Available** — `forma generate` / sync (P3a) |

Use [Install](install.md) and [CLI reference](cli.md) for new apps, or [Manual setup](manual-setup.md) without the CLI.

## Guide map

| Page | Contents |
|------|----------|
| [Install](install.md) | Toolchain, installing the CLI, first project |
| [Manual setup](manual-setup.md) | Create a project without `forma init` |
| [CLI reference](cli.md) | `init`, `add`, `sync`, `generate`, `dev` |
| [forma.yaml](forma-yaml.md) | Project config and defaults |
| [Project layout](project-layout.md) | Directories, imports, this repo vs your app |
| [Writing UI](writing-ui.md) | State, mount, refresh, dot imports, lists |
| [Adding components](adding-components.md) | shadcn palette, sync, git |
| [Advanced](advanced.md) | Overrides, custom Go/JS widgets |
| [Troubleshooting](troubleshooting.md) | Common failures |

## Architecture (short)

Read [DESIGN.md](../DESIGN.md) for the full picture. In one sentence: **Go builds a component tree; the host maps type names to React/shadcn and calls back into Go for events.**
