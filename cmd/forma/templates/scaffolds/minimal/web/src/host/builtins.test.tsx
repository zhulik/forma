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
    Type: "Card",
    Props: { Padding: 2 },
    Children: [{ Type: "Text", Props: { Text: "card-child" } }],
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
});
