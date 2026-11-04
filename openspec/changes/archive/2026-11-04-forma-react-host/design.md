# Design

## Context

See [proposal.md](./proposal.md). P3a already emits app `*.gen.go` and `web/src/registry.gen.ts`; the minimal scaffold `main.tsx` imports the registry but only renders a placeholder. Framework layout/leaf types (`Fragment`, `Row`, `Text`, …) live in `github.com/zhulik/forma/ui` and are **not** in the generated registry. Go wire slot keys use `SlotKeyFromField` (`Trigger` → `trigger`), while current registry goldens still key slots with Go field names (`"Trigger"`)—the host cannot assemble compounds until those agree. P3c owns the GopherJS bridge; this change leaves a clear mount seam.

## Goals / Non-Goals

**Goals**

- Thin host in the init scaffold: registry merge, recursive renderer, sanitization, text-input hybrid, mount seam.
- Builtin registry covering framework `Node.Type` values used by the reference DSL (layout + text + Fragment at minimum; Card/Heading/etc. as shipped in framework stubs).
- Vitest coverage for renderer/sanitizer/input policy without needing GopherJS.
- Slot-key alignment in codegen goldens so host and Go share one wire alphabet.
- Author-doc updates for host layout and text-input policy.

**Non-Goals**

- Implementing `UIMount` / `RefreshUI` / callback `ui.Context` in Go (P3c).
- `forma dev`, Vite proxy to GopherJS, or live `main.go` in the browser (P3d).
- Domain state or business logic in React.
- Replacing or rewriting shadcn components; host only maps and wraps.
- Moving framework `ui/*.gen.go` stubs into `examples/` (separate follow-up).

## Decisions

### 1. Capability name `ui-host` (not `ui-runtime`)

**Choice:** New OpenSpec capability **`ui-host`** for the React layer. Reserve **`ui-runtime`** language for P3c Go/GopherJS bridge contracts.

**Alternatives:** Single `ui-runtime` for host+bridge—rejected; ROADMAP splits P3b/P3c and archive sync should stay phase-shaped.

### 2. Host lives in scaffold `web/src/host/`

**Choice:** Write-once TypeScript/TSX under **`cmd/forma/templates/scaffolds/minimal/web/src/host/`** (renderer, sanitize, builtins, `FormaHost` root). Scaffold **`main.tsx`** creates the React root on `forma.yaml` mount id and renders `<FormaHost registry={…} tree={…} />`.

**Alternatives:** Publish an npm `@forma/host` package—deferred until multiple consumers exist; keep copy-in-scaffold like the rest of the web tree. Framework-repo-only host not copied to apps—rejected; authors’ apps need the host after `forma init`.

### 3. Registry merge: builtins + generated (+ optional custom)

**Choice:**

```
effectiveRegistry = { ...builtinRegistry, ...generatedRegistry, ...customRegistry }
```

Later keys win (generated overrides framework stubs when both define e.g. future shared names; custom always last). Generated `registry.gen.ts` stays codegen-owned; custom merges live in a hand-written module (e.g. `web/src/registry.custom.ts`) imported by `main.tsx`, never patched into `registry.gen.ts`.

**Alternatives:** Patch `main.tsx` during sync—fragile. Codegen emits framework builtins—wrong ownership (framework types are not under app `components/ui/`).

### 4. Builtin primitives are thin React wrappers

**Choice:** Hand-roll small wrappers for framework types (`Fragment` → `React.Fragment`, `Text`/`Heading` → semantic elements + Tailwind class mapping from props, `Row`/`Column`/`Container`/`Card` → layout `div`s with gap/align/padding classes). Reuse scaffold `cn` helper. Do **not** invent a second CSS system.

**Reuse vs hand-roll:** Reuse React 18, existing Tailwind tokens, and generated shadcn entries. Hand-roll the recursive renderer and builtins—no third-party “JSON→React schema” library (forma `Node` shape + slots are specific; prefer-reuse stops at mature modules that match the problem).

### 5. Recursive renderer algorithm

**Choice:** Single `renderNode(node, registry)`:

1. Look up `registry[node.Type]`; on miss → diagnostic + `null`.
2. Strip non-DOM Go meta from props (children/fallback slices, nested `BaseComponent` blobs as needed); map `Class` → `className`; forward scalar props; wrap function props with sanitizers.
3. Recursively render `Children`; for compounds, render each `Slots[wireKey]` list and pass into the slot assembly expected by the registered root (compose root + slot components per registry metadata).
4. Pass `key={node.Key}` when non-empty.

**Alternatives:** Per-widget generated renderers—duplicates codegen; rejected for v1 thin host.

### 6. Event sanitization module

**Choice:** Central helpers, e.g. `sanitizeClick(fn)`, `sanitizeStringValue(fn)`, `sanitizeChecked(fn)`, used when wiring props whose values are functions. Strip `SyntheticEvent` / `Event`; pass `event.target.value` / `.checked` or no payload for void clicks. Never pass the event object through.

**Alternatives:** Per-shadcn wrapper only—rejected for leaves that use generated components directly; sanitization must sit at the host prop-wiring boundary. Full allowlist serialization—overkill for v1.

### 7. Text-input hybrid policy

**Choice:** Host **renderer policy**, not a registry override war: when rendering type `"Input"` (and later any widget marked text-field), wrap the **registry component** (usually generated shadcn `Input`) in a hybrid controller:

- Local React state holds the string while focused.
- On change: update local state + call Go-facing `OnChange` with sanitized string (matches `main.go` / DSL `OnChange` API).
- Sync from `Value` prop when not focused, or when unfocused prop changes.
- While focused, ignore prop updates that merely echo the last emitted value (avoids caret jump on RefreshUI).

Generated registry still supplies the visual `Input`; custom registry may replace the visual component; hybrid wrapping stays in the renderer so merge order (decision 3) does not drop the policy.

**Alternatives:** Blur-only notify—breaks current demo handlers. Fully controlled from Go every keystroke—known caret issues (DESIGN.md). Builtin registry entry that overrides generated `Input`—fragile if custom replaces builtins. Debounced-only—extra latency; can add later per-widget without changing the hybrid shell.

### 8. Mount seam for P3c (no Go bridge yet)

**Choice:** `FormaHost` props: `{ registry, tree: FormaNode | null }`. `main.tsx` supplies `tree={null}` or a tiny committed fixture for smoke visuals. Export a stable imperative helper (e.g. `setHostTree`) **only if** needed for P3c discovery; prefer props/state lifted in `main.tsx` that P3c will replace with bridge-driven updates. Document the seam in design/tasks; do not implement GopherJS interop here.

**Alternatives:** Fake Go callbacks in JS for the todo demo—scope creep into P3c/P3d.

### 9. Slot keys: fix codegen, don’t dual-map in host

**Choice:** Change registry emit so slot map keys use the same wire keys as Go (`SlotKeyFromField`). Update accordion (and any) goldens. Host looks up `node.Slots` keys directly in `entry.slots`.

**Alternatives:** Host lowercases first character at runtime—hides the bug and invites drift. **BREAKING** for any checked-in app registries—acceptable (early-stage, no compat shims).

### 10. Testing: Vitest + Testing Library in scaffold

**Choice:** Add Vitest, jsdom, `@testing-library/react` (+ user-event) to scaffold `package.json` / lockfile; `npm test` script. Tests ship under `web/src/host/*.test.tsx` in the template. Init smoke may run a focused host test or rely on `tsc -b` + unit tests in CI when Node is present.

**Reuse:** Vite ecosystem Vitest over a custom harness or Playwright for unit behavior.

### 11. Docs and ROADMAP

**Choice:** Update `doc/user/` project-layout, advanced, troubleshooting, README status; note P3b host vs P3c bridge. Optionally mark ROADMAP P3b row in progress when builder opens the PR (orch/control plane)—not required for planning artifacts.

### 12. Prefer-fn / Go surface

**Choice:** No new Go host packages. Codegen slot-key fix may reuse existing Go helpers; any new slice remaps in codegen stay on **`fn`**. TypeScript host does not pull a lodash-like dependency.

## Risks / Trade-offs

- **[Builtin prop↔class mapping incomplete]** → Cover types used by `main.go` / framework stubs first; expand when demos need more.
- **[Custom Input replaces visuals]** → Hybrid policy remains in the renderer; custom entries only swap the inner component.
- **[Slot-key BREAKING for existing gen files]** → Early-stage rewrite; regenerate via `forma generate`.
- **[Fixture-only main until P3c]** → Docs must not claim live Go UI; tests use explicit fixture trees.
- **[Compound slot assembly shape varies by shadcn]** → v1 supports registry `slots` map composition; exotic Select patterns may need overrides in P5.

## Migration Plan

1. Land host modules + codegen slot-key fix + docs in this change.
2. Apps: re-run `forma generate` / `forma sync` to refresh `registry.gen.ts`.
3. P3c: drive `FormaHost` tree (and callbacks) from GopherJS without rewriting the renderer.
4. P3d: `forma dev` runs Vite + Go rebuild against the same host.

## Open Questions

None material for builder. Debounced OnChange as an opt-in prop can wait for a later widget-policy tweak.
