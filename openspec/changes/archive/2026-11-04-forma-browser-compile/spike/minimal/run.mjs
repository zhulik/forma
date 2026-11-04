import { pathToFileURL } from "node:url";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const dir = dirname(fileURLToPath(import.meta.url));
const received = [];

globalThis.formaHost = {
  setTree(tree) {
    received.push(tree);
    console.log("setTree typeof:", typeof tree);
    console.log("constructor:", tree?.constructor?.name);
    if (tree && typeof tree === "object" && !(tree instanceof Map)) {
      console.log("Kind:", tree.Kind);
      console.log("Props.Text:", tree.Props?.Text);
      console.log("Children isArray:", Array.isArray(tree.Children));
    }
    console.log("isJSONString:", typeof tree === "string");
  },
};

const mainMod = await import(
  pathToFileURL(
    join(
      dir,
      "output/@goscript/github.com/zhulik/forma/spike/goscript-minimal/main.gs.ts",
    ),
  ).href
);

await mainMod.main();
console.log("received:", received.length, "ok:", globalThis.__spikeMinimalOK);
if (received.length < 1 || globalThis.__spikeMinimalOK !== true) {
  process.exit(1);
}
console.log("PASS: minimal setTree object graph");
