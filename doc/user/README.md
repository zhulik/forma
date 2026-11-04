# forma author guide

forma is a Go-native UI framework: you compose widgets in Go; a React + shadcn host renders them. This guide is the **primary onboarding path** for app authors—alongside the reference demo [`main.go`](../../main.go) and the [`ui-dsl`](../../openspec/specs/ui-dsl/spec.md) specification.

## Documentation is first-class

If a change affects **author experience**—CLI commands, `forma.yaml`, project layout, install or sync workflow, or DSL rules you document here—that change **must update this guide in the same delivery slice**. Treat `doc/user/` like a public API: breaking UX without doc updates is incomplete work.

Contributors: see also OpenSpec capability **`ui-docs`** and `openspec/config.yaml` task/proposal rules.

## What works today

| Area | Status |
|------|--------|
| UI DSL (composition, slots, `Context`, `fn`) | **Available** — [`main.go`](../../main.go), [`ui-dsl` spec](../../openspec/specs/ui-dsl/spec.md), `task test` |
| Author guide (this tree) | **Available** |
| `forma` CLI (`init`, `sync`, `generate`, `add`, `version`) | **Available** — full sync + codegen; [`forma-cli` spec](../../openspec/specs/forma-cli/spec.md); `dev` planned |
| Browser dev loop, GopherJS runtime, host | **Planned** — ROADMAP P3 |
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
