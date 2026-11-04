import * as React from "react";

export type ButtonProps = {
  variant?: "default" | "ghost" | "outline";
  size?: "default" | "sm" | "lg";
  className?: string;
  children?: React.ReactNode;
  onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
};

export function Button({
  variant = "default",
  size = "default",
  className,
  children,
  onClick,
}: ButtonProps) {
  return (
    <button
      type="button"
      data-variant={variant}
      data-size={size}
      className={className}
      onClick={onClick}
    >
      {children}
    </button>
  );
}
