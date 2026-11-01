package ui

import "context"

// Context is for event and interaction callbacks only. Tree rendering uses
// context.Context on Component.Render and internal forma lifecycle.
type Context interface {
	context.Context
	RefreshUI()
	Update(fn func())
}
