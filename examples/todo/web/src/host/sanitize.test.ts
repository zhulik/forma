import { describe, expect, it, vi } from "vitest";

import {
  sanitizeChecked,
  sanitizeClick,
  sanitizeStringValue,
} from "./sanitize";

describe("sanitizeClick", () => {
  it("invokes without forwarding the event object", () => {
    const fn = vi.fn();
    const wrapped = sanitizeClick(fn);
    const fakeEvent = {
      preventDefault: vi.fn(),
      target: document.createElement("button"),
    };
    wrapped(fakeEvent);
    expect(fn).toHaveBeenCalledOnce();
    expect(fn).toHaveBeenCalledWith();
  });
});

describe("sanitizeStringValue", () => {
  it("passes plain strings through", () => {
    const fn = vi.fn();
    sanitizeStringValue(fn)("hello");
    expect(fn).toHaveBeenCalledWith("hello");
  });

  it("extracts target.value from event-like objects", () => {
    const fn = vi.fn();
    const input = document.createElement("input");
    input.value = "typed";
    sanitizeStringValue(fn)({
      preventDefault: () => undefined,
      target: input,
    });
    expect(fn).toHaveBeenCalledWith("typed");
    expect(fn.mock.calls[0][0]).toBeTypeOf("string");
  });
});

describe("sanitizeChecked", () => {
  it("passes plain booleans through", () => {
    const fn = vi.fn();
    sanitizeChecked(fn)(true);
    expect(fn).toHaveBeenCalledWith(true);
  });

  it("extracts target.checked from event-like objects", () => {
    const fn = vi.fn();
    const input = document.createElement("input");
    input.type = "checkbox";
    input.checked = true;
    sanitizeChecked(fn)({
      preventDefault: () => undefined,
      target: input,
    });
    expect(fn).toHaveBeenCalledWith(true);
    expect(fn.mock.calls[0][0]).toBeTypeOf("boolean");
  });
});
