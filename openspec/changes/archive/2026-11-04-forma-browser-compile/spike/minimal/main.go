package main

import "syscall/js"

// Minimal setTree crossing: Go object graph → JS host without JSON.
func main() {
	host := js.Global().Get("formaHost")
	if host.IsUndefined() || host.IsNull() {
		panic("formaHost missing")
	}
	tree := map[string]any{
		"Kind": "Button",
		"Props": map[string]any{
			"Text": "hello",
		},
		"Children": []any{},
	}
	host.Call("setTree", tree)
	js.Global().Set("__spikeMinimalOK", true)
}
