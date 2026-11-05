import type { CSSProperties, ReactNode } from "react";
import { Fragment as ReactFragment } from "react";

import { cn } from "@/lib/utils";

import type { ComponentRegistry, RegistryEntry } from "./types";

type BuiltinProps = Record<string, unknown> & {
  className?: string;
  style?: CSSProperties;
  children?: ReactNode;
};

function asNumber(value: unknown, fallback = 0): number {
  return typeof value === "number" && Number.isFinite(value) ? value : fallback;
}

function asString(value: unknown): string {
  return typeof value === "string" ? value : "";
}

function asBool(value: unknown): boolean {
  return value === true;
}

/** Closed Tailwind spacing scale (0–12) used by layout props. */
const SPACING_CLASS: Record<number, string> = {
  0: "gap-0",
  1: "gap-1",
  2: "gap-2",
  3: "gap-3",
  4: "gap-4",
  5: "gap-5",
  6: "gap-6",
  7: "gap-7",
  8: "gap-8",
  9: "gap-9",
  10: "gap-10",
  11: "gap-11",
  12: "gap-12",
};

const PADDING_CLASS: Record<number, string> = {
  0: "p-0",
  1: "p-1",
  2: "p-2",
  3: "p-3",
  4: "p-4",
  5: "p-5",
  6: "p-6",
  7: "p-7",
  8: "p-8",
  9: "p-9",
  10: "p-10",
  11: "p-11",
  12: "p-12",
};

const MARGIN_Y_CLASS: Record<number, string> = {
  0: "my-0",
  1: "my-1",
  2: "my-2",
  3: "my-3",
  4: "my-4",
  5: "my-5",
  6: "my-6",
  7: "my-7",
  8: "my-8",
  9: "my-9",
  10: "my-10",
  11: "my-11",
  12: "my-12",
};

const MAX_WIDTH_CLASS: Record<string, string> = {
  sm: "max-w-sm",
  md: "max-w-md",
  lg: "max-w-lg",
  xl: "max-w-xl",
  "2xl": "max-w-2xl",
  full: "max-w-full",
  none: "max-w-none",
  prose: "max-w-prose",
};

function gapClass(gap: number): string | undefined {
  if (gap <= 0) {
    return undefined;
  }
  return SPACING_CLASS[gap];
}

function paddingClass(padding: number): string | undefined {
  if (padding <= 0) {
    return undefined;
  }
  return PADDING_CLASS[padding];
}

function marginYClass(marginY: number): string | undefined {
  return MARGIN_Y_CLASS[marginY];
}

function maxWidthClass(token: string): string | undefined {
  if (!token) {
    return undefined;
  }
  return MAX_WIDTH_CLASS[token];
}

function TextBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const text = asString(props.Text);
  return (
    <span
      className={cn(
        asBool(props.Muted) && "text-muted-foreground",
        asBool(props.Strikethrough) && "line-through",
        className,
      )}
      style={style}
    >
      {text}
      {children}
    </span>
  );
}

function HeadingBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const level = Math.min(6, Math.max(1, asNumber(props.Level, 1)));
  const Tag = `h${level}` as "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
  const text = asString(props.Text);
  return (
    <Tag className={cn("font-semibold tracking-tight", className)} style={style}>
      {text}
      {children}
    </Tag>
  );
}

function RowBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const align = asString(props.Align);
  const justify = asString(props.Justify);
  return (
    <div
      className={cn(
        "flex flex-row",
        gapClass(asNumber(props.Gap)),
        paddingClass(asNumber(props.Padding)),
        asBool(props.Border) && "border",
        align === "center" && "items-center",
        justify === "between" && "justify-between",
        className,
      )}
      style={style}
    >
      {children}
    </div>
  );
}

function ColumnBuiltin({ className, style, children, ...props }: BuiltinProps) {
  return (
    <div
      className={cn("flex flex-col", gapClass(asNumber(props.Gap)), className)}
      style={style}
    >
      {children}
    </div>
  );
}

function ContainerBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const maxWidth = asString(props.MaxWidth);
  const marginX = asString(props.MarginX);
  const hasMarginY = props.MarginY !== undefined;
  return (
    <div
      className={cn(
        maxWidthClass(maxWidth),
        marginX === "auto" && "mx-auto",
        hasMarginY && marginYClass(asNumber(props.MarginY)),
        className,
      )}
      style={style}
    >
      {children}
    </div>
  );
}

function PanelBuiltin({ className, style, children, ...props }: BuiltinProps) {
  return (
    <div
      className={cn(
        "rounded-lg border bg-card text-card-foreground shadow-sm",
        paddingClass(asNumber(props.Padding, 4)),
        className,
      )}
      style={style}
    >
      {children}
    </div>
  );
}

function entry(
  component: RegistryEntry["component"],
): RegistryEntry {
  return { component };
}

/** Framework layout/leaf types not present in generated shadcn registry. */
export const builtinRegistry: ComponentRegistry = {
  Fragment: entry(ReactFragment as RegistryEntry["component"]),
  Text: entry(TextBuiltin),
  Heading: entry(HeadingBuiltin),
  Row: entry(RowBuiltin),
  Column: entry(ColumnBuiltin),
  Container: entry(ContainerBuiltin),
  Panel: entry(PanelBuiltin),
};
