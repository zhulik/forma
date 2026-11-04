//go:build !js

package ui

// UIMount is a no-op on native builds so apps compile and test without a browser host.
func UIMount(string, Component) {} //nolint:revive // exported naming
