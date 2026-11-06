import type { ComponentRegistry } from "./host";

/**
 * Hand-written registry overrides / third-party wrappers.
 * Merged last (wins over builtins).
 */
export const customRegistry: ComponentRegistry = {};
