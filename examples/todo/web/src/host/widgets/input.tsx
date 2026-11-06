import type { CSSProperties, ChangeEvent, ReactNode } from "react";

import { cn } from "@/lib/utils";

type InputProps = Record<string, unknown> & {
  className?: string;
  style?: CSSProperties;
  children?: ReactNode;
  Placeholder?: unknown;
  placeholder?: unknown;
  Value?: unknown;
  value?: unknown;
  OnChange?: unknown;
  onChange?: unknown;
};

function asString(value: unknown): string {
  return typeof value === "string" ? value : "";
}

export function InputBuiltin({
  className,
  style,
  children,
  Placeholder,
  placeholder,
  Value,
  value,
  OnChange,
  onChange,
  ...rest
}: InputProps) {
  const handleChange = (OnChange ?? onChange) as
    | ((...args: unknown[]) => void)
    | undefined;
  const display = asString(value ?? Value);

  return (
    <input
      type="text"
      className={cn("z-input", className)}
      style={style}
      placeholder={asString(placeholder ?? Placeholder)}
      value={display}
      onChange={(event: ChangeEvent<HTMLInputElement>) => {
        handleChange?.(event);
      }}
      {...rest}
    />
  );
}
