import * as React from "react";

export type InputProps = {
  className?: string;
  type?: string;
  value?: string;
  placeholder?: string;
  disabled?: boolean;
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
};

export function Input({
  className,
  type = "text",
  value,
  placeholder,
  disabled,
  onChange,
}: InputProps) {
  return (
    <input
      type={type}
      className={className}
      value={value}
      placeholder={placeholder}
      disabled={disabled}
      onChange={onChange}
    />
  );
}
