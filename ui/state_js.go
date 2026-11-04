//go:build js

package ui

import "github.com/gopherjs/gopherjs/js"

// UIMount registers root against the host seam and performs an initial refresh.
// The mount element id is owned by the React host; Go only drives formaHost.setTree.
func UIMount(_ string, root Component) { //nolint:revive // exported naming
	sink := jsHostSink{}
	_, _ = MountWithSink(sink, root)
}

type jsHostSink struct{}

func (jsHostSink) SetTree(tree *Node) {
	host := js.Global.Get("formaHost")
	if host == nil || host == js.Undefined {
		panic("forma: globalThis.formaHost is not registered; mount the host before the Go bundle")
	}
	setTree := host.Get("setTree")
	if setTree == nil || setTree == js.Undefined {
		panic("forma: formaHost.setTree is missing")
	}
	setTree.Invoke(tree)
}
