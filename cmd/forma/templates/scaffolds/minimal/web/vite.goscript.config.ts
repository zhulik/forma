import path from "path";
import { defineConfig } from "vite";

// Bundles GoScript TypeScript output into public/app.js for the host to import.
export default defineConfig({
  resolve: {
    alias: {
      "@goscript": path.resolve(__dirname, ".goscript/@goscript"),
    },
  },
  build: {
    target: "esnext",
    outDir: "public",
    emptyOutDir: false,
    lib: {
      entry: path.resolve(__dirname, "goscript-entry.ts"),
      formats: ["es"],
      fileName: () => "app.js",
    },
  },
});
