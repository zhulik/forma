import { render, screen } from "@testing-library/react";
import type { ReactNode } from "react";
import { describe, expect, it, vi } from "vitest";

import { builtinRegistry } from "./builtins";
import { mergeRegistry } from "./merge";
import { renderNode } from "./render";
import type { ComponentRegistry, FormaNode } from "./types";

describe("renderNode", () => {
  it("renders nested children in order", () => {
    const tree: FormaNode = {
      Type: "Column",
      Children: [
        { Type: "Text", Props: { Text: "first" } },
        { Type: "Text", Props: { Text: "second" } },
      ],
    };
    render(<>{renderNode(tree, builtinRegistry)}</>);
    const texts = screen.getAllByText(/first|second/).map((el) => el.textContent);
    expect(texts).toEqual(["first", "second"]);
  });

  it("Fragment introduces no extra DOM host element", () => {
    const tree: FormaNode = {
      Type: "Row",
      Children: [
        {
          Type: "Fragment",
          Children: [
            { Type: "Text", Props: { Text: "a" } },
            { Type: "Text", Props: { Text: "b" } },
          ],
        },
      ],
    };
    const { container } = render(<>{renderNode(tree, builtinRegistry)}</>);
    const row = container.firstElementChild;
    expect(row?.tagName).toBe("DIV");
    expect(row?.children).toHaveLength(2);
    expect(row?.children[0].textContent).toBe("a");
    expect(row?.children[1].textContent).toBe("b");
  });

  it("passes stable keys to React", () => {
    const warn = vi.spyOn(console, "warn").mockImplementation(() => undefined);
    const tree: FormaNode = {
      Type: "Column",
      Children: [
        { Type: "Text", Key: "k1", Props: { Text: "one" } },
        { Type: "Text", Key: "k2", Props: { Text: "two" } },
      ],
    };
    // Rendering with keys must not throw; React uses them for reconciliation.
    expect(() => render(<>{renderNode(tree, builtinRegistry)}</>)).not.toThrow();
    warn.mockRestore();
  });

  it("skips unknown types with a diagnostic", () => {
    const warn = vi.spyOn(console, "warn").mockImplementation(() => undefined);
    const tree: FormaNode = {
      Type: "Column",
      Children: [
        { Type: "NotARealWidget", Props: { Text: "nope" } },
        { Type: "Text", Props: { Text: "ok" } },
      ],
    };
    render(<>{renderNode(tree, builtinRegistry)}</>);
    expect(screen.getByText("ok")).toBeTruthy();
    expect(screen.queryByText("nope")).toBeNull();
    expect(warn).toHaveBeenCalledWith(
      expect.stringContaining("NotARealWidget"),
    );
    warn.mockRestore();
  });

  it("assembles compound slots with aligned wire keys", () => {
    const Trigger = ({ children }: { children?: ReactNode }) => (
      <button type="button">{children}</button>
    );
    const Content = ({ children }: { children?: ReactNode }) => (
      <div data-testid="content">{children}</div>
    );
    const Root = ({ children }: { children?: ReactNode }) => (
      <div data-testid="root">{children}</div>
    );

    const compounds: ComponentRegistry = {
      Accordion: {
        component: Root,
        slots: {
          trigger: Trigger,
          content: Content,
        },
      },
    };
    const registry = mergeRegistry(builtinRegistry, compounds);
    const tree: FormaNode = {
      Type: "Accordion",
      Slots: {
        trigger: [{ Type: "Text", Props: { Text: "Open" } }],
        content: [{ Type: "Text", Props: { Text: "Body" } }],
      },
    };
    render(<>{renderNode(tree, registry)}</>);
    expect(screen.getByRole("button", { name: "Open" })).toBeTruthy();
    expect(screen.getByTestId("content").textContent).toBe("Body");
  });

  it("sanitizes click callbacks so they never receive events", async () => {
    const onClick = vi.fn();
    const Button = ({
      onClick: handler,
      children,
    }: {
      onClick?: (...args: unknown[]) => void;
      children?: ReactNode;
    }) => (
      <button type="button" onClick={(e) => handler?.(e)}>
        {children}
      </button>
    );
    const registry = mergeRegistry(builtinRegistry, {
      Button: { component: Button },
    });
    const tree: FormaNode = {
      Type: "Button",
      Props: { OnClick: onClick },
      Children: [{ Type: "Text", Props: { Text: "Go" } }],
    };
    render(<>{renderNode(tree, registry)}</>);
    screen.getByRole("button", { name: "Go" }).click();
    expect(onClick).toHaveBeenCalledOnce();
    expect(onClick.mock.calls[0]).toEqual([]);
  });
});
