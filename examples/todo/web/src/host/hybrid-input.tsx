import {
  useEffect,
  useRef,
  useState,
  type ComponentType,
  type FocusEventHandler,
  type ReactNode,
} from "react";

import { sanitizeStringValue } from "./sanitize";

type HybridInputProps = {
  component: ComponentType<Record<string, unknown>>;
  value?: unknown;
  Value?: unknown;
  onChange?: unknown;
  OnChange?: unknown;
  onFocus?: FocusEventHandler;
  onBlur?: FocusEventHandler;
  children?: ReactNode;
  [key: string]: unknown;
};

function stringProp(value: unknown): string {
  return typeof value === "string" ? value : "";
}

/**
 * Renderer-policy wrapper for text Input: local state while focused;
 * sync from value prop when unfocused; sanitized string OnChange.
 */
export function HybridInput({
  component: Input,
  value,
  Value,
  onChange: onChangeProp,
  OnChange,
  onFocus,
  onBlur,
  children,
  ...rest
}: HybridInputProps) {
  const propValue = stringProp(value ?? Value);
  const rawOnChange = OnChange ?? onChangeProp;
  const onChange =
    typeof rawOnChange === "function"
      ? sanitizeStringValue(rawOnChange as (v: string) => void)
      : undefined;

  const [focused, setFocused] = useState(false);
  const [local, setLocal] = useState(propValue);
  const lastEmitted = useRef(propValue);

  useEffect(() => {
    if (!focused) {
      setLocal(propValue);
      lastEmitted.current = propValue;
      return;
    }
    // While focused, ignore refreshes that merely echo the last emitted value
    // so the caret stays stable across RefreshUI.
    if (propValue === lastEmitted.current) {
      return;
    }
  }, [propValue, focused]);

  const displayValue = focused ? local : propValue;

  return (
    <Input
      {...rest}
      value={displayValue}
      onChange={(arg: unknown) => {
        let next = "";
        if (typeof arg === "string") {
          next = arg;
        } else if (arg && typeof arg === "object" && "target" in arg) {
          const target = (arg as { target?: { value?: unknown } }).target;
          if (target && typeof target.value === "string") {
            next = target.value;
          }
        }
        setLocal(next);
        lastEmitted.current = next;
        onChange?.(next);
      }}
      onFocus={(event: Parameters<NonNullable<FocusEventHandler>>[0]) => {
        setFocused(true);
        onFocus?.(event);
      }}
      onBlur={(event: Parameters<NonNullable<FocusEventHandler>>[0]) => {
        setFocused(false);
        onBlur?.(event);
      }}
    >
      {children}
    </Input>
  );
}
