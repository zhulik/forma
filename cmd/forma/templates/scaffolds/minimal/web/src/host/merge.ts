import type { ComponentRegistry } from "./types";

/**
 * Merge registries. Later maps win on key collision:
 * builtins → generated → custom.
 */
export function mergeRegistry(
  builtins: ComponentRegistry,
  generated: ComponentRegistry = {},
  custom: ComponentRegistry = {},
): ComponentRegistry {
  return {
    ...builtins,
    ...generated,
    ...custom,
  };
}
