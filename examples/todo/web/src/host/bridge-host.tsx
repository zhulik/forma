import { useEffect, useState } from "react";

import {
  createFormaBridge,
  registerFormaHost,
  unregisterFormaHost,
  type FormaHostBridge,
} from "./bridge";
import { FormaHost } from "./forma-host";
import type { ComponentRegistry, FormaNode } from "./types";

export type FormaBridgeHostProps = {
  registry: ComponentRegistry;
  /** Optional pre-built bridge (tests); otherwise one is created and registered. */
  bridge?: FormaHostBridge & {
    getTree: () => FormaNode | null;
    subscribe: (listener: (tree: FormaNode | null) => void) => () => void;
  };
};

/**
 * Owns React tree state, registers `globalThis.formaHost`, and renders FormaHost.
 * Go drives updates through formaHost.setTree — do not create a second React root.
 */
export function FormaBridgeHost({ registry, bridge: bridgeProp }: FormaBridgeHostProps) {
  const [bridge] = useState(() => bridgeProp ?? createFormaBridge());
  const [tree, setTree] = useState<FormaNode | null>(() => bridge.getTree());

  useEffect(() => {
    registerFormaHost(bridge);
    const unsubscribe = bridge.subscribe(setTree);
    return () => {
      unsubscribe();
      unregisterFormaHost();
    };
  }, [bridge]);

  return <FormaHost registry={registry} tree={tree} />;
}
