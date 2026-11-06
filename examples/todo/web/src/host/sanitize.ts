type UnknownRecord = Record<string, unknown>;

function isEventLike(value: unknown): value is UnknownRecord {
  if (value === null || typeof value !== "object") {
    return false;
  }
  const record = value as UnknownRecord;
  return (
    typeof record.preventDefault === "function" ||
    typeof record.stopPropagation === "function" ||
    "nativeEvent" in record ||
    (typeof Event !== "undefined" && value instanceof Event)
  );
}

function targetValue(arg: unknown): string | undefined {
  if (!arg || typeof arg !== "object") {
    return undefined;
  }
  const target = (arg as UnknownRecord).target;
  if (!target || typeof target !== "object") {
    return undefined;
  }
  const value = (target as UnknownRecord).value;
  return typeof value === "string" ? value : undefined;
}

function targetChecked(arg: unknown): boolean | undefined {
  if (!arg || typeof arg !== "object") {
    return undefined;
  }
  const target = (arg as UnknownRecord).target;
  if (!target || typeof target !== "object") {
    return undefined;
  }
  const checked = (target as UnknownRecord).checked;
  return typeof checked === "boolean" ? checked : undefined;
}

/** Invoke a void click-style callback without passing the event object. */
export function sanitizeClick(fn: (...args: never[]) => void): (...args: unknown[]) => void {
  return () => {
    fn();
  };
}

/** Pass a plain string (or extract `event.target.value`); never the event. */
export function sanitizeStringValue(
  fn: (value: string) => void,
): (...args: unknown[]) => void {
  return (arg?: unknown) => {
    if (typeof arg === "string") {
      fn(arg);
      return;
    }
    const fromTarget = targetValue(arg);
    if (fromTarget !== undefined) {
      fn(fromTarget);
      return;
    }
    if (isEventLike(arg)) {
      fn("");
      return;
    }
    if (arg === undefined || arg === null) {
      fn("");
      return;
    }
    fn(String(arg));
  };
}

/** Pass a plain boolean (or extract `event.target.checked`); never the event. */
export function sanitizeChecked(
  fn: (checked: boolean) => void,
): (...args: unknown[]) => void {
  return (arg?: unknown) => {
    if (typeof arg === "boolean") {
      fn(arg);
      return;
    }
    const fromTarget = targetChecked(arg);
    if (fromTarget !== undefined) {
      fn(fromTarget);
      return;
    }
    if (isEventLike(arg)) {
      fn(false);
      return;
    }
    fn(Boolean(arg));
  };
}

const CLICK_PROP = /^(on)?click$/i;
const CHANGE_PROP = /^(on)?(change|valuechange)$/i;
const CHECKED_PROP = /^(on)?checkedchange$/i;

/** Wrap a function prop with the sanitizer matching its Go/React name. */
export function sanitizeFunctionProp(
  name: string,
  fn: (...args: never[]) => unknown,
): (...args: unknown[]) => unknown {
  if (CHECKED_PROP.test(name)) {
    return sanitizeChecked(fn as (checked: boolean) => void);
  }
  if (CHANGE_PROP.test(name)) {
    return sanitizeStringValue(fn as (value: string) => void);
  }
  if (CLICK_PROP.test(name)) {
    return sanitizeClick(fn as () => void);
  }
  // Default: strip event-like first args; forward plain payloads.
  return (...args: unknown[]) => {
    const plain = args.map((arg) => {
      if (!isEventLike(arg)) {
        return arg;
      }
      const value = targetValue(arg);
      if (value !== undefined) {
        return value;
      }
      const checked = targetChecked(arg);
      if (checked !== undefined) {
        return checked;
      }
      return undefined;
    });
    const filtered = plain.filter((v) => v !== undefined);
    return (fn as (...a: unknown[]) => unknown)(...filtered);
  };
}
