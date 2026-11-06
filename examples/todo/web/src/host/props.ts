import type { CSSProperties } from "react";

import { sanitizeFunctionProp } from "./sanitize";

const STRIP_PROP_KEYS = new Set([
  "Key",
  "Class",
  "children",
  "Children",
  "fallback",
  "Fallback",
  "BaseComponent",
  "WidgetMeta",
  "Type",
  "Slots",
  "Props",
]);

type UnknownRecord = Record<string, unknown>;

function isPlainObject(value: unknown): value is UnknownRecord {
  return value !== null && typeof value === "object" && !Array.isArray(value);
}

function toStyle(value: unknown): CSSProperties | undefined {
  if (!isPlainObject(value)) {
    return undefined;
  }
  return value as CSSProperties;
}

/**
 * Map Go wire props onto React props: Class→className, strip meta, sanitize callbacks.
 */
export function mapNodeProps(
  props: unknown,
  className?: string,
): Record<string, unknown> {
  const out: Record<string, unknown> = {};

  if (isPlainObject(props)) {
    for (const [key, value] of Object.entries(props)) {
      if (STRIP_PROP_KEYS.has(key)) {
        continue;
      }
      if (typeof value === "function") {
        out[key] = sanitizeFunctionProp(key, value as (...args: never[]) => unknown);
        // Also expose common React lowercase aliases when Go uses PascalCase.
        if (/^On[A-Z]/.test(key)) {
          const reactKey = key[0].toLowerCase() + key.slice(1);
          out[reactKey] = out[key];
        }
        continue;
      }
      if (key === "Style") {
        const style = toStyle(value);
        if (style) {
          out.style = style;
        }
        continue;
      }
      if (key === "ID" && typeof value === "string") {
        out.id = value;
        continue;
      }
      if (key === "Attrs" && isPlainObject(value)) {
        Object.assign(out, value);
        continue;
      }
      out[key] = value;
      // Lowercase first letter alias for common DOM/React props from Go (Value → value).
      if (/^[A-Z]/.test(key) && typeof value !== "object") {
        const reactKey = key[0].toLowerCase() + key.slice(1);
        if (!(reactKey in out)) {
          out[reactKey] = value;
        }
      }
    }
  }

  if (className) {
    const existing = out.className;
    out.className =
      typeof existing === "string" && existing.length > 0
        ? `${existing} ${className}`
        : className;
  }

  return out;
}
