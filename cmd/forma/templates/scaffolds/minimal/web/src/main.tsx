import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import "./index.css";

const mount = document.getElementById("app-root");
if (mount) {
  createRoot(mount).render(
    <StrictMode>
      <p className="text-muted-foreground text-sm">forma web host (runtime not connected)</p>
    </StrictMode>,
  );
}
