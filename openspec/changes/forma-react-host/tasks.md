# Tasks

## 1. Codegen slot-key alignment

- [ ] 1.1 Update registry emit so compound slot map keys use the same wire keys as Go `SlotKeyFromField`, and verify accordion (and any other compound) goldens match those keys byte-for-byte
- [ ] 1.2 Run codegen unit tests (`task test` covering `cmd/forma/internal/codegen`) and verify they pass with the updated goldens

## 2. Host core modules (scaffold)

- [ ] 2.1 Add `web/src/host/` TypeScript types for the render-tree node shape plus registry entry types aligned with `registry.gen.ts`, and verify `tsc -b` accepts them in the template tree
- [ ] 2.2 Implement builtin registry wrappers for framework types used by the reference DSL (`Fragment`, `Text`, `Heading`, `Row`, `Column`, `Container`, `Card`, and any other framework stub types the demo needs), and verify a Vitest case renders each type without the generated registry
- [ ] 2.3 Implement registry merge (`builtins` → `generated` → `custom`) and recursive `renderNode` (children, keys, Fragment, unknown-type diagnostic skip), and verify Vitest covers nested children, keys, Fragment, and unknown-type skip
- [ ] 2.4 Implement compound slot assembly from registry `slots` + node `Slots` using aligned wire keys, and verify a Vitest fixture with trigger/content-style slots renders slot content
- [ ] 2.5 Implement event sanitization helpers and wire function props through them, and verify Vitest asserts click/change callbacks never receive event objects and string/boolean payloads are plain

## 3. Text-input policy and mount seam

- [ ] 3.1 Implement hybrid text-input controller wrapping the registry `Input` component per design.md, and verify Vitest covers caret-stable refresh with same value, unfocused prop sync, and sanitized string OnChange
- [ ] 3.2 Add `FormaHost` + update scaffold `main.tsx` to mount on the configured element with merged registry and null/fixture tree, and verify missing mount element fails closed (no alternate root) in Vitest or a small DOM smoke
- [ ] 3.3 Add Vitest (+ jsdom, Testing Library, user-event) to scaffold `package.json` / lockfile with an `npm test` script, and verify `npm test` passes in a materialized web tree (init smoke or dedicated template fixture)

## 4. Author docs (same change)

- [ ] 4.1 Update `doc/user/project-layout.md`, `advanced.md`, `troubleshooting.md`, and README status lines for host modules, custom registry merge, text-input hybrid policy, and P3c/P3d still planned; verify docs do not claim a live GopherJS todo in the browser

## 5. Integration gates

- [ ] 5.1 Extend or confirm init scaffold smoke builds the web tree (`tsc`/Vite build) with the new host entry after mock or real `registry.gen.ts`, and verify the smoke passes locally
- [ ] 5.2 Run `task test` and `task lint` at the framework/`cmd/forma` roots affected by the codegen change and verify both pass

## Workflow follow-up

- After implementation PR merges and archive, sync `ui-host` into `openspec/specs/` and mark ROADMAP P3b Done.
- P3c (`forma-gopherjs-runtime`) plugs the Go bridge into the `FormaHost` tree seam; do not expand this change into that work.
