import { createElement, type ComponentType, type ReactNode } from "react";

import { HybridInput } from "./hybrid-input";
import { mapNodeProps } from "./props";
import type { ComponentRegistry, FormaNode } from "./types";

function renderChildren(
  nodes: FormaNode[] | null | undefined,
  registry: ComponentRegistry,
): ReactNode[] {
  if (!nodes || nodes.length === 0) {
    return [];
  }
  return nodes.map((child) => renderNode(child, registry));
}

/**
 * Recursively map a forma node through the component registry.
 * Unknown types are skipped with a console diagnostic.
 */
export function renderNode(
  node: FormaNode | null | undefined,
  registry: ComponentRegistry,
): ReactNode {
  if (!node || !node.Type) {
    return null;
  }

  const entry = registry[node.Type];
  if (!entry) {
    console.warn(`[forma] unknown component type: ${node.Type}`);
    return null;
  }

  const mapped = mapNodeProps(node.Props, node.Class);
  const children = renderChildren(node.Children, registry);

  const slotElements: ReactNode[] = [];
  if (node.Slots && entry.slots) {
    for (const [wireKey, slotNodes] of Object.entries(node.Slots)) {
      const SlotComp = entry.slots[wireKey];
      if (!SlotComp || !slotNodes) {
        continue;
      }
      slotNodes.forEach((slotNode, index) => {
        const rendered = renderNode(slotNode, registry);
        const key = slotNode.Key || `${wireKey}-${index}`;
        slotElements.push(createElement(SlotComp, { key }, rendered));
      });
    }
  }

  const allChildren = [...slotElements, ...children];
  const props = node.Key ? { ...mapped, key: node.Key } : mapped;

  if (node.Type === "Input") {
    return createElement(
      HybridInput,
      {
        ...props,
        component: entry.component as ComponentType<Record<string, unknown>>,
      },
      ...allChildren,
    );
  }

  return createElement(entry.component, props, ...allChildren);
}
