import { act } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";

import { mountFormaHost } from "./mount";

describe("mountFormaHost", () => {
  it("fails closed when the mount element is missing", () => {
    const err = vi.spyOn(console, "error").mockImplementation(() => undefined);
    document.body.innerHTML = "";
    const root = mountFormaHost("missing-root", <p>nope</p>);
    expect(root).toBeNull();
    expect(document.body.querySelector("#missing-root")).toBeNull();
    expect(err).toHaveBeenCalledWith(
      expect.stringContaining("missing-root"),
    );
    err.mockRestore();
  });

  it("attaches to the configured mount element", () => {
    document.body.innerHTML = `<div id="app-root"></div>`;
    let root: ReturnType<typeof mountFormaHost> = null;
    act(() => {
      root = mountFormaHost("app-root", <p>hosted</p>);
    });
    expect(root).not.toBeNull();
    expect(document.getElementById("app-root")?.textContent).toBe("hosted");
    act(() => {
      root?.unmount();
    });
  });
});
