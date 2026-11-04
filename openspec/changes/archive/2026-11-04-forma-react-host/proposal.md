# Proposal

## Why

P3a ships Go bindings and `registry.gen.ts`, but the scaffold host still shows a placeholder and never renders a forma tree. Without a thin React host (registry merge, recursive renderer, event sanitization, text-input policy), P3c cannot plug the GopherJS bridge into a real DOM path.

## What Changes

- Embed a **static React host** in the minimal web scaffold: builtin + generated registry merge, recursive widget renderer, Fragment/slots/keys handling, and a mount seam for a tree root (fixture-driven until P3c).
- Add **event sanitization** so React/DOM event objects never reach Go-facing callbacks—only plain primitives / plain objects.
- Define and implement the **text-input hybrid policy** for `Input` (and document it for authors).
- Align generated **registry slot keys** with Go wire slot keys when they currently disagree.
- Update **`doc/user/`** so host layout, custom registry merges, and text-input behavior match shipped scaffold behavior (still label live Go↔browser bridge / `forma dev` as planned for P3c–P3d).

**Out of scope:** GopherJS `UIMount`/`RefreshUI` bridge (P3c); `forma dev` loop (P3d); domain state in React; SSR; third-party island recipes beyond the documented registry extension pattern.

## Capabilities

### New Capabilities

- `ui-host`: React host that maps forma render trees to shadcn/React via a component registry, recursively renders nodes (including Fragment and named slots), sanitizes events before callbacks, and applies the text-input hybrid policy.

### Modified Capabilities

- `ui-codegen`: Generated registry slot map keys must match Go wire slot keys produced by `SlotKeyFromField` (so the host can resolve slots without a second naming scheme).
- `ui-docs`: Author guide must describe the host tree under `web/src/`, registry merge / custom entries, and text-input policy; clarify what works without the P3c bridge.

## Impact

- **ROADMAP phase: P3b** — advances the React host row; does not claim P3c runtime or P3d dev loop.
- **Scaffold:** `cmd/forma/templates/scaffolds/minimal/web/` gains host modules; `main.tsx` mounts the host instead of the placeholder paragraph.
- **Codegen:** Small emit fix for registry slot key casing; goldens updated.
- **Docs:** `doc/user/` project layout, advanced, troubleshooting, README status lines.
- **Tests:** Vitest (+ React Testing Library) in the web scaffold for renderer/sanitizer/input policy; init smoke still builds the web tree.
- **Dependencies:** Add Vitest/RTL (and related Vite test wiring) to the scaffold `package.json` / lockfile; reuse existing React 18 + generated `registry.gen.ts`.
- **Follow-ons:** P3c wires GopherJS into the host mount seam; P3d runs host + rebuild.
