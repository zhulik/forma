import { fireEvent, render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";

import { builtinRegistry } from "./builtins";
import { renderNode } from "./render";
import type { FormaNode } from "./types";

const types: FormaNode[] = [
  { Type: "Fragment", Children: [{ Type: "Text", Props: { Text: "frag" } }] },
  { Type: "Text", Props: { Text: "hello-text" } },
  { Type: "Heading", Props: { Level: 2, Text: "hello-heading" } },
  {
    Type: "Row",
    Props: { Gap: 2 },
    Children: [{ Type: "Text", Props: { Text: "row-child" } }],
  },
  {
    Type: "Column",
    Props: { Gap: 1 },
    Children: [{ Type: "Text", Props: { Text: "col-child" } }],
  },
  {
    Type: "Container",
    Children: [{ Type: "Text", Props: { Text: "container-child" } }],
  },
  {
    Type: "Card",
    Props: { Padding: 2 },
    Children: [{ Type: "Text", Props: { Text: "card-child" } }],
  },
  {
    Type: "Button",
    Props: { Variant: "default" },
    Children: [{ Type: "Text", Props: { Text: "click-me" } }],
  },
  {
    Type: "Input",
    Props: { Placeholder: "type here", Value: "" },
  },
  {
    Type: "Checkbox",
    Props: { Checked: false },
  },
];

describe("builtinRegistry", () => {
  it("renders each todo-sized type without a generated registry", () => {
    for (const node of types) {
      const { unmount } = render(<>{renderNode(node, builtinRegistry)}</>);
      expect(builtinRegistry[node.Type]).toBeDefined();
      unmount();
    }
    render(<>{renderNode(types[1], builtinRegistry)}</>);
    expect(screen.getByText("hello-text")).toBeTruthy();
  });

  it("maps interactive widgets with semantic classes and sanitized callbacks", () => {
    const onClick = vi.fn();
    const onChange = vi.fn();
    const onChecked = vi.fn();

    const { container } = render(
      <>
        {renderNode(
          {
            Type: "Column",
            Children: [
              {
                Type: "Button",
                Props: { Variant: "primary", OnClick: onClick },
                Children: [{ Type: "Text", Props: { Text: "Go" } }],
              },
              {
                Type: "Input",
                Props: { Value: "hi", OnChange: onChange },
              },
              {
                Type: "Checkbox",
                Props: { Checked: false, OnCheckedChange: onChecked },
              },
              {
                Type: "Row",
                Props: { Gap: 3, Align: "center" },
                Children: [{ Type: "Text", Props: { Text: "row" } }],
              },
            ],
          },
          builtinRegistry,
        )}
      </>,
    );

    const button = screen.getByRole("button", { name: "Go" });
    expect(button.className).toContain("z-button");
    expect(button.className).toContain("z-button-primary");
    fireEvent.click(button);
    expect(onClick).toHaveBeenCalled();

    const input = screen.getByDisplayValue("hi");
    expect(input.className).toContain("z-input");
    fireEvent.change(input, { target: { value: "bye" } });
    expect(onChange).toHaveBeenCalledWith("bye");

    const checkbox = container.querySelector('input[type="checkbox"]');
    expect(checkbox).toBeTruthy();
    fireEvent.click(checkbox!);
    expect(onChecked).toHaveBeenCalledWith(true);

    const row = screen.getByText("row").parentElement;
    expect(row?.className).toContain("display-flex");
    expect(row?.className).toContain("gap=3");
  });
});
