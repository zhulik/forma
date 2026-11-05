import * as React from "react";

// onSelect would be click-style by heuristic; override forces a string payload.
export type WidgetProps = {
  className?: string;
  onSelect?: (event: React.MouseEvent<HTMLButtonElement>) => void;
};

export function Widget({ className, onSelect }: WidgetProps) {
  return (
    <button type="button" className={className} onClick={onSelect}>
      pick
    </button>
  );
}
