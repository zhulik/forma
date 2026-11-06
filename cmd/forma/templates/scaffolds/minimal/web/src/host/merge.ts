import type { ComponentRegistry } from "./types";

/**
 * Merge registries. Later maps win on key collision:
 * builtins → custom (no generated kit registry).
 */
export function mergeRegistry(
  builtins: ComponentRegistry,
  custom: ComponentRegistry = {},
): ComponentRegistry {
  return {
    ...builtins,
    ...custom,
  };
}
