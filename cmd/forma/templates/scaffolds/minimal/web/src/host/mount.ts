import { createRoot, type Root } from "react-dom/client";
import { createElement, StrictMode, type ReactNode } from "react";

/**
 * Attach a React root to the configured mount element.
 * Fails closed when the element is missing — never falls back to document.body.
 *
 * The Go runtime drives updates through `globalThis.formaHost.setTree`
 * (see bridge.ts / FormaBridgeHost); this helper only owns the React root.
 */
export function mountFormaHost(
  mountId: string,
  element: ReactNode,
): Root | null {
  const mount = document.getElementById(mountId);
  if (!mount) {
    console.error(`[forma] mount element #${mountId} not found`);
    return null;
  }
  const root = createRoot(mount);
  root.render(createElement(StrictMode, null, element));
  return root;
}
