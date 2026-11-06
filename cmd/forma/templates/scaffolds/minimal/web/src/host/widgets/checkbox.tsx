import type { CSSProperties, ChangeEvent, ReactNode } from "react";

import { cn } from "@/lib/utils";

type CheckboxProps = Record<string, unknown> & {
  className?: string;
  style?: CSSProperties;
  children?: ReactNode;
  Checked?: unknown;
  checked?: unknown;
  OnCheckedChange?: unknown;
  onCheckedChange?: unknown;
};

export function CheckboxBuiltin({
  className,
  style,
  children,
  Checked,
  checked,
  OnCheckedChange,
  onCheckedChange,
  ...rest
}: CheckboxProps) {
  const handleChange = (OnCheckedChange ?? onCheckedChange) as
    | ((...args: unknown[]) => void)
    | undefined;
  const isChecked = Checked === true || checked === true;

  return (
    <label className={cn("z-checkbox", className)} style={style}>
      <input
        type="checkbox"
        checked={isChecked}
        onChange={(event: ChangeEvent<HTMLInputElement>) => {
          handleChange?.(event);
        }}
        {...rest}
      />
      {children}
    </label>
  );
}
