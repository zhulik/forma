import type { CSSProperties, ReactNode } from "react";

import { buttonClasses } from "./classes";

type ButtonProps = Record<string, unknown> & {
  className?: string;
  style?: CSSProperties;
  children?: ReactNode;
  Variant?: unknown;
  Size?: unknown;
  OnClick?: unknown;
  onClick?: unknown;
};

function asString(value: unknown): string {
  return typeof value === "string" ? value : "";
}

export function ButtonBuiltin({
  className,
  style,
  children,
  Variant,
  Size,
  OnClick,
  onClick,
  ...rest
}: ButtonProps) {
  const handleClick = (OnClick ?? onClick) as ((...args: unknown[]) => void) | undefined;
  return (
    <button
      type="button"
      className={buttonClasses(asString(Variant), asString(Size), className)}
      style={style}
      onClick={typeof handleClick === "function" ? () => handleClick() : undefined}
      {...rest}
    >
      {children}
    </button>
  );
}
