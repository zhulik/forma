import type { ComponentType, ReactNode } from "react";

/** Forma render-tree node (matches Go `ui.Node` wire shape). */
export type FormaNode = {
  Type: string;
  Key?: string;
  Class?: string;
  Props?: unknown;
  Children?: FormaNode[] | null;
  Slots?: Record<string, FormaNode[] | null | undefined> | null;
};

/** One registry entry — same shape as generated `registry.gen.ts`. */
export type RegistryEntry = {
  component: ComponentType<Record<string, unknown>>;
  slots?: Record<string, ComponentType<Record<string, unknown>>>;
};

export type ComponentRegistry = Record<string, RegistryEntry>;

export type RenderNode = (node: FormaNode, registry: ComponentRegistry) => ReactNode;
