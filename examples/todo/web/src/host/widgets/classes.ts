import type { CSSProperties } from "react";

import { cn } from "@/lib/utils";

/** Map a spacing int to a 0build shorthand utility (e.g. gap=2). */
export function spaceUtil(prefix: string, value: number): string | undefined {
  if (value <= 0) {
    return undefined;
  }
  return `${prefix}=${value}`;
}

const MAX_WIDTH: Record<string, string> = {
  sm: "24rem",
  md: "28rem",
  lg: "32rem",
  xl: "36rem",
  "2xl": "42rem",
  full: "100%",
};

export function maxWidthClass(maxWidth: string): string | undefined {
  if (!maxWidth) {
    return undefined;
  }
  const resolved = MAX_WIDTH[maxWidth] ?? maxWidth;
  return `[max-width]=${resolved}`;
}

export function buttonClasses(variant: string, size: string, className?: string): string {
  const v = variant || "default";
  const variantClass =
    v === "ghost"
      ? "z-button-ghost"
      : v === "destructive"
        ? "z-button-destructive"
        : v === "secondary"
          ? "z-button-secondary"
          : v === "outline"
            ? "z-button-outline"
            : "z-button-primary";

  const sizeClass =
    size === "icon"
      ? "z-button-icon"
      : size === "sm" || size === "small"
        ? "z-form-small"
        : size === "lg" || size === "large"
          ? "z-form-large"
          : undefined;

  return cn("z-button", variantClass, sizeClass, className);
}

export function mergeStyles(
  style: CSSProperties | undefined,
  extra: CSSProperties | undefined,
): CSSProperties | undefined {
  if (!style && !extra) {
    return undefined;
  }
  return { ...style, ...extra };
}
