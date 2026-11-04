import { pathToFileURL } from "node:url";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const dir = dirname(fileURLToPath(import.meta.url));
const mod = await import(
  pathToFileURL(
    join(dir, "output/@goscript/github.com/zhulik/forma/spike/makefunc/main.gs.ts"),
  ).href
);
await mod.main();
await globalThis.spikeCall("hello-makefunc");
const got = await globalThis.spikeRead();
console.log("got:", got);
if (got !== "hello-makefunc") process.exit(1);
console.log("PASS makefunc+js.FuncOf string");
