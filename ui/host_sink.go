package ui

// HostSink receives render trees from the runtime.
// The js build adapts globalThis.formaHost.setTree; tests use a recording fake.
type HostSink interface {
	SetTree(tree *Node)
}
