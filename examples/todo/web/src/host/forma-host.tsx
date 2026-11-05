import type { ComponentRegistry, FormaNode } from "./types";
import { renderNode } from "./render";

export type FormaHostProps = {
  registry: ComponentRegistry;
  /** Replaceable tree root; Go pushes updates via `formaHost.setTree`. */
  tree: FormaNode | null;
};

/** Mount seam: registry + tree root → React elements. */
export function FormaHost({ registry, tree }: FormaHostProps) {
  if (!tree) {
    return null;
  }
  return <>{renderNode(tree, registry)}</>;
}
