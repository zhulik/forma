import { act, render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";

import {
  createFormaBridge,
  FORMA_HOST_GLOBAL,
  registerFormaHost,
  unregisterFormaHost,
} from "./bridge";
import { FormaBridgeHost } from "./bridge-host";
import { builtinRegistry } from "./builtins";

describe("formaHost bridge", () => {
  it("registers globalThis.formaHost and re-renders FormaHost on setTree", () => {
    const bridge = createFormaBridge();
    registerFormaHost(bridge);

    document.body.innerHTML = `<div id="app-root"></div>`;
    const mount = document.getElementById("app-root")!;

    render(<FormaBridgeHost registry={builtinRegistry} bridge={bridge} />, {
      container: mount,
    });

    expect(
      (globalThis as typeof globalThis & { formaHost?: { setTree: unknown } })[
        FORMA_HOST_GLOBAL
      ]?.setTree,
    ).toBeTypeOf("function");
    expect(mount.id).toBe("app-root");
    expect(screen.queryByText("Hello")).toBeNull();

    act(() => {
      bridge.setTree({
        Type: "Text",
        Props: { Text: "Hello" },
      });
    });

    expect(screen.getByText("Hello")).toBeTruthy();
    expect(document.getElementById("app-root")).toBe(mount);

    act(() => {
      bridge.setTree({
        Type: "Text",
        Props: { Text: "World" },
      });
    });

    expect(screen.getByText("World")).toBeTruthy();
    expect(document.getElementById("app-root")).toBe(mount);

    unregisterFormaHost();
  });
});
