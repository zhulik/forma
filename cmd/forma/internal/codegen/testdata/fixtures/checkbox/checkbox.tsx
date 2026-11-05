import * as React from "react";

export type CheckboxProps = {
  className?: string;
  checked?: boolean;
  disabled?: boolean;
  onCheckedChange?: (checked: boolean) => void;
};

export function Checkbox({
  className,
  checked,
  disabled,
  onCheckedChange,
}: CheckboxProps) {
  return (
    <button
      type="button"
      role="checkbox"
      aria-checked={checked}
      disabled={disabled}
      className={className}
      onKeyDown={() => onCheckedChange?.(!checked)}
    />
  );
}
