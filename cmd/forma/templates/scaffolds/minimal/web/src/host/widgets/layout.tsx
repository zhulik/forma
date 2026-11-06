import type { CSSProperties, ReactNode } from "react";

import { cn } from "@/lib/utils";

import { maxWidthClass, spaceUtil } from "./classes";

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

export function TextBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const text = asString(props.Text);
  return (
    <span
      className={cn(
        asBool(props.Muted) && "opacity=60",
        asBool(props.Strikethrough) && "[text-decoration]=line-through",
        className,
      )}
      style={style}
    >
      {text}
      {children}
    </span>
  );
}

export function HeadingBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const level = Math.min(6, Math.max(1, asNumber(props.Level, 1)));
  const Tag = `h${level}` as "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
  const text = asString(props.Text);
  return (
    <Tag className={cn("z-heading", className)} style={style}>
      {text}
      {children}
    </Tag>
  );
}

export function RowBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const align = asString(props.Align);
  const justify = asString(props.Justify);
  return (
    <div
      className={cn(
        "display-flex",
        "flex-row",
        spaceUtil("gap", asNumber(props.Gap)),
        spaceUtil("p", asNumber(props.Padding)),
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

export function ColumnBuiltin({ className, style, children, ...props }: BuiltinProps) {
  return (
    <div
      className={cn(
        "display-flex",
        "flex-col",
        spaceUtil("gap", asNumber(props.Gap)),
        className,
      )}
      style={style}
    >
      {children}
    </div>
  );
}

export function ContainerBuiltin({ className, style, children, ...props }: BuiltinProps) {
  const maxWidth = asString(props.MaxWidth);
  const marginX = asString(props.MarginX);
  return (
    <div
      className={cn(
        maxWidthClass(maxWidth),
        marginX === "auto" && "mx=auto",
        props.MarginY !== undefined && spaceUtil("my", asNumber(props.MarginY)),
        className,
      )}
      style={style}
    >
      {children}
    </div>
  );
}

export function CardBuiltin({ className, style, children, ...props }: BuiltinProps) {
  return (
    <div
      className={cn("z-card", spaceUtil("p", asNumber(props.Padding, 4)), className)}
      style={style}
    >
      {children}
    </div>
  );
}
