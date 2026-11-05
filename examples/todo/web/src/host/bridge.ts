import type { FormaNode } from "./types";

/**
 * Global bridge name the GopherJS runtime looks up (`globalThis.formaHost`).
 * Register via {@link registerFormaHost} before the Go bundle calls UIMount.
 */
export const FORMA_HOST_GLOBAL = "formaHost" as const;

export type FormaHostBridge = {
  /** Replace the tree root FormaHost renders; null clears the host. */
  setTree: (tree: FormaNode | null) => void;
};

type FormaHostGlobal = typeof globalThis & {
  formaHost?: FormaHostBridge;
};

type Listener = (tree: FormaNode | null) => void;

/**
 * Imperative tree seam for the Go runtime. Create and {@link registerFormaHost}
 * synchronously before loading the GopherJS bundle so UIMount can call setTree.
 */
export function createFormaBridge(): FormaHostBridge & {
  getTree: () => FormaNode | null;
  subscribe: (listener: Listener) => () => void;
} {
  let tree: FormaNode | null = null;
  const listeners = new Set<Listener>();

  return {
    getTree: () => tree,
    setTree(next: FormaNode | null) {
      tree = next;
      for (const listener of listeners) {
        listener(next);
      }
    },
    subscribe(listener: Listener) {
      listeners.add(listener);
      listener(tree);
      return () => {
        listeners.delete(listener);
      };
    },
  };
}

/** Publish the bridge on `globalThis.formaHost` for the Go runtime. */
export function registerFormaHost(bridge: FormaHostBridge): void {
  (globalThis as FormaHostGlobal).formaHost = bridge;
}

/** Remove `globalThis.formaHost` when the host unmounts. */
export function unregisterFormaHost(): void {
  delete (globalThis as FormaHostGlobal).formaHost;
}
