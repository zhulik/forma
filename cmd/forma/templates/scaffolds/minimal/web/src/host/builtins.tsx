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

function gapClass(gap: number): string | undefined {
  if (gap <= 0) {
    return undefined;
  }
  return `gap-${gap}`;
}

function paddingClass(padding: number): string | undefined {
  if (padding <= 0) {
    return undefined;
  }
  return `p-${padding}`;
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
  return (
    <div
      className={cn(
        maxWidth && `max-w-${maxWidth}`,
        marginX === "auto" && "mx-auto",
        props.MarginY !== undefined && `my-${asNumber(props.MarginY)}`,
        className,
      )}
      style={style}
    >
      {children}
    </div>
  );
}

function CardBuiltin({ className, style, children, ...props }: BuiltinProps) {
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
  Card: entry(CardBuiltin),
};
