# Proposal

## Why

Owned widgets + 0build (`forma-owned-widgets` / PR #8) make a todo-sized UI expressible, but authors still lack a **runnable** in-browser example under `forma dev`. Root `main.go` shows the DSL; the framework repo has no app `web/` tree, so the tip P5 success metric—todo on owned widgets end-to-end—is not yet shippable.

## What Changes

- Add a self-contained **`examples/todo`** mini-app (nested module, `forma.yaml`, scaffold-derived `web/`, in-memory todo UI) runnable with `forma sync` + `forma dev`.
- Prove **CRUD** (add / list / toggle done / delete) and **minimal responsive** layout using only the owned widget surface (no new widgets, no kit codegen).
- Document how to run the example in `doc/user/`, root README, and `examples/todo/README.md`; remove “planned” labels for a dedicated todo example.
- Retarget the **canonical runnable demo** to `examples/todo`; slim root `main.go` to a short Hello-style owned-widget smoke (framework module keeps a `package main` without dual-maintaining the full todo).
- Update ROADMAP / OpenSpec config pointers that still treat root `main.go` as the only north-star demo.
- **Out of scope:** backend/persistence, desktop (P4), new widget catalog, shadcn path (#5/#6), broader owned catalog / perf.

## Capabilities

### New Capabilities

- `todo-example`: Runnable in-memory todo demo on owned widgets + 0build, served via `forma dev`, with documented run path.

### Modified Capabilities

- `ui-docs`: Author guide treats `examples/todo` as available; documents run/verify steps; drops planned-only wording for the dedicated example.
- `ui-dsl`: Canonical todo demo / dense-import scenarios point at the example app (not only root `main.go`).

## Impact

- **ROADMAP phase:** **P5** (Owned widgets + example), second tip row after `forma-owned-widgets`.
- **Stack dependency:** Assumes owned Go widgets + 0build host builtins from **`forma-owned-widgets` @ `f853abb` (PR #8)**—Button, Input, Checkbox, Card, Text, Heading, Row, Column, Container, Fragment. **Not shadcn.** Do not reopen #5/#6.
- **Code:** new `examples/todo/` tree; root `main.go` slimmed; docs under `doc/user/` + README; ROADMAP status when shipping; optional Taskfile/`go test` for the example module.
- **No** CLI command changes required; reuse `forma sync` / `forma dev` as-is.
- **Breaking (docs/mental model):** canonical runnable demo moves from “read root `main.go` + init your own app” to “`cd examples/todo && forma dev`”. Early-stage: no compat shims.
