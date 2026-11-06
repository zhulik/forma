# forma author guide

forma is a Go-native UI framework: you compose **owned widgets** in Go; a thin React host renders them with **[0build](https://0build.dev)** semantic CSS (`.z-*`). This guide is the **primary onboarding path** for app authors—alongside the reference demo [`main.go`](../../main.go) and the [`ui-dsl`](../../openspec/specs/ui-dsl/spec.md) specification.

## Documentation is first-class

If a change affects **author experience**—CLI commands, `forma.yaml`, project layout, install or sync workflow, or DSL rules you document here—that change **must update this guide in the same delivery slice**. Treat `doc/user/` like a public API: breaking UX without doc updates is incomplete work.

Contributors: see also OpenSpec capability **`ui-docs`** and `openspec/config.yaml` task/proposal rules.

## What works today

| Area | Status |
|------|--------|
| UI DSL (composition, slots, `Context`, `fn`) | **Available** — [`main.go`](../../main.go), [`ui-dsl` spec](../../openspec/specs/ui-dsl/spec.md), `task test` |
| Author guide (this tree) | **Available** |
| `forma` CLI (`init`, `sync`, `version`, `dev`) | **Available** — owned-widgets scaffold; [`forma-cli` spec](../../openspec/specs/forma-cli/spec.md) |
| React host (builtins, registry merge, renderer, sanitize, text-input policy) | **Available** — scaffold `web/src/host/` |
| Go↔host bridge (`UIMount` / `RefreshUI` / callbacks) | **Available** — host registers `globalThis.formaHost.setTree` |
| Browser `forma dev` loop | **Available** — Vite host + GoScript rebuild (Go changes need page reload) |
| Owned widgets + 0build look | **Available** — todo-sized surface (Button, Input, Checkbox, Card, layout, Text, Heading, Fragment) |

Use [Install](install.md) and [CLI reference](cli.md) for new apps, or [Manual setup](manual-setup.md) without the CLI.

## Guide map

| Page | Contents |
|------|----------|
| [Install](install.md) | Toolchain, installing the CLI, first project |
| [Manual setup](manual-setup.md) | Create a project without `forma init` |
| [CLI reference](cli.md) | `init`, `sync`, `dev`, `version` |
| [forma.yaml](forma-yaml.md) | Project config and defaults |
| [Project layout](project-layout.md) | Directories, imports, this repo vs your app |
| [Writing UI](writing-ui.md) | State, mount, refresh, dot imports, lists |
| [Owned widgets](adding-components.md) | Framework widgets, 0build look, extending |
| [Advanced](advanced.md) | Custom Go/JS widgets, escape hatches |
| [Troubleshooting](troubleshooting.md) | Common failures |

## Architecture (short)

Read [DESIGN.md](../DESIGN.md) for the full picture. In one sentence: **Go builds a component tree; the host maps type names to thin React builtins styled with 0build and calls back into Go for events.**
