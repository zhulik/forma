package ui

import (
	"context"
	"sync"
)

// mountSession owns a mounted root, host sink, and refresh-generation cancellation.
type mountSession struct {
	root Component
	sink HostSink

	mu           sync.Mutex
	parentCtx    context.Context //nolint:containedctx // mount lifetime; parent of refresh generations
	parentCancel context.CancelFunc
	genCtx       context.Context //nolint:containedctx // current refresh generation for callback Context
	genCancel    context.CancelFunc
}

func newMountSession(root Component, sink HostSink) *mountSession {
	parentCtx, parentCancel := context.WithCancel(context.Background())

	return &mountSession{
		root:         root,
		sink:         sink,
		parentCtx:    parentCtx,
		parentCancel: parentCancel,
	}
}

func (s *mountSession) callbackContext() *callbackContext {
	return &callbackContext{
		Context: s.genCtx,
		session: s,
	}
}

func (s *mountSession) refresh() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.genCancel != nil {
		s.genCancel()
	}
	genCtx, genCancel := context.WithCancel(s.parentCtx)
	s.genCtx = genCtx
	s.genCancel = genCancel

	tree := s.root.Render(context.Background())
	bound := bindTree(tree, s.callbackContext())
	s.sink.SetTree(bound)
}

func (s *mountSession) update(fn func()) {
	var panicVal any
	func() {
		defer func() { panicVal = recover() }()
		fn()
	}()
	if panicVal != nil {
		panic(panicVal)
	}
	s.refresh()
}

func (s *mountSession) unmount() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.genCancel != nil {
		s.genCancel()
		s.genCancel = nil
	}
	s.genCtx = nil
	if s.parentCancel != nil {
		s.parentCancel()
		s.parentCancel = nil
	}
}

// callbackContext is the concrete ui.Context supplied to author handlers.
type callbackContext struct {
	context.Context //nolint:containedctx // Context interface embeds context.Context by design

	session *mountSession
}

func (c *callbackContext) RefreshUI() {
	c.session.refresh()
}

func (c *callbackContext) Update(fn func()) {
	c.session.update(fn)
}

// MountWithSink mounts root against sink, performs an initial refresh, and
// returns a callback context for the current generation plus an unmount func
// that cancels outstanding callback contexts.
func MountWithSink(sink HostSink, root Component) (Context, func()) {
	session := newMountSession(root, sink)
	session.refresh()

	return session.callbackContext(), session.unmount
}
