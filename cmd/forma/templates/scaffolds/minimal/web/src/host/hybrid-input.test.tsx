import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useState, type InputHTMLAttributes } from "react";
import { describe, expect, it, vi } from "vitest";

import { HybridInput } from "./hybrid-input";

function NativeInput(props: Record<string, unknown>) {
  return <input {...(props as InputHTMLAttributes<HTMLInputElement>)} />;
}

describe("HybridInput", () => {
  it("keeps caret-stable local text when prop echoes the last emitted value", async () => {
    const user = userEvent.setup();
    const onChange = vi.fn();

    function Harness() {
      const [value, setValue] = useState("ab");
      return (
        <HybridInput
          component={NativeInput}
          value={value}
          OnChange={(next: string) => {
            onChange(next);
            // Simulate RefreshUI echoing the same value.
            setValue(next);
          }}
        />
      );
    }

    render(<Harness />);
    const input = screen.getByRole("textbox") as HTMLInputElement;
    await user.click(input);
    await user.type(input, "c");
    expect(input.value).toBe("abc");
    expect(onChange).toHaveBeenLastCalledWith("abc");
    expect(onChange.mock.calls.at(-1)?.[0]).toBeTypeOf("string");
  });

  it("syncs from value prop when unfocused", () => {
    const { rerender } = render(
      <HybridInput component={NativeInput} value="one" />,
    );
    const input = screen.getByRole("textbox") as HTMLInputElement;
    expect(input.value).toBe("one");
    rerender(<HybridInput component={NativeInput} value="two" />);
    expect(input.value).toBe("two");
  });

  it("notifies OnChange with a sanitized string, not an event", async () => {
    const user = userEvent.setup();
    const onChange = vi.fn();
    render(
      <HybridInput component={NativeInput} value="" OnChange={onChange} />,
    );
    await user.type(screen.getByRole("textbox"), "x");
    expect(onChange).toHaveBeenCalled();
    for (const call of onChange.mock.calls) {
      expect(call).toHaveLength(1);
      expect(call[0]).toBeTypeOf("string");
      expect(call[0]).not.toHaveProperty("target");
    }
  });
});
