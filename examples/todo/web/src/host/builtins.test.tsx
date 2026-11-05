import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";

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
    Type: "Panel",
    Props: { Padding: 2 },
    Children: [{ Type: "Text", Props: { Text: "panel-child" } }],
  },
];

describe("builtinRegistry", () => {
  it("renders each framework type without the generated registry", () => {
    for (const node of types) {
      const { unmount } = render(<>{renderNode(node, builtinRegistry)}</>);
      expect(builtinRegistry[node.Type]).toBeDefined();
      unmount();
    }
    render(<>{renderNode(types[1], builtinRegistry)}</>);
    expect(screen.getByText("hello-text")).toBeTruthy();
  });

  it("maps gap and padding through static class literals", () => {
    const { container, unmount } = render(
      <>
        {renderNode(
          {
            Type: "Row",
            Props: { Gap: 2, Padding: 3 },
            Children: [{ Type: "Text", Props: { Text: "spaced" } }],
          },
          builtinRegistry,
        )}
      </>,
    );
    const el = container.firstElementChild as HTMLElement;
    expect(el.className.split(/\s+/)).toEqual(
      expect.arrayContaining(["gap-2", "p-3"]),
    );
    unmount();
  });

  it("maps max-width tokens through static class literals", () => {
    const { container, unmount } = render(
      <>
        {renderNode(
          {
            Type: "Container",
            Props: { MaxWidth: "md", MarginY: 4 },
            Children: [{ Type: "Text", Props: { Text: "box" } }],
          },
          builtinRegistry,
        )}
      </>,
    );
    const el = container.firstElementChild as HTMLElement;
    expect(el.className.split(/\s+/)).toEqual(
      expect.arrayContaining(["max-w-md", "my-4"]),
    );
    unmount();
  });

  it("omits unsupported spacing and max-width values", () => {
    const { container, unmount } = render(
      <>
        {renderNode(
          {
            Type: "Column",
            Props: { Gap: 99 },
            Children: [
              {
                Type: "Container",
                Props: { MaxWidth: "nope", MarginY: 99 },
                Children: [{ Type: "Text", Props: { Text: "skip" } }],
              },
            ],
          },
          builtinRegistry,
        )}
      </>,
    );
    const column = container.firstElementChild as HTMLElement;
    expect(column.className).not.toMatch(/gap-99/);
    const inner = column.firstElementChild as HTMLElement;
    expect(inner.className).not.toMatch(/max-w-nope/);
    expect(inner.className).not.toMatch(/my-99/);
    unmount();
  });

  it("does not register a layout Card builtin", () => {
    expect(builtinRegistry.Card).toBeUndefined();
    expect(builtinRegistry.Panel).toBeDefined();
  });
});
