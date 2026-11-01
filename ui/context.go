package ui

import "context"

type Context interface {
	context.Context
	RefreshUI()
}
