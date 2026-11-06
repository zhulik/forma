import { main } from "@goscript/github.com/zhulik/forma/examples/todo/main.gs.ts";

/** Boot the GoScript-compiled Go main (UIMount). Exported for explicit callers. */
export async function boot(): Promise<void> {
  await main();
}

void boot();
