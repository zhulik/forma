package ui

import (
	"context"
	"unicode"
	"unicode/utf8"
)

// RenderComponentSlot renders one slot field for compound widgets (used from generated *.gen.go).
func RenderComponentSlot(ctx context.Context, c Component) []*Node {
	if c == nil {
		return nil
	}

	node := c.Render(ctx)
	if node == nil {
		return nil
	}

	return []*Node{node}
}

// PutSlot adds nodes to a slot map when non-empty (used from generated *.gen.go).
func PutSlot(slots map[string][]*Node, wireKey string, nodes []*Node) {
	if len(nodes) == 0 {
		return
	}

	slots[wireKey] = nodes
}

// SlotKeyFromField maps an exported Go field name to a wire slot key (Trigger → trigger).
func SlotKeyFromField(fieldName string) string {
	if fieldName == "" {
		return ""
	}

	r, size := utf8.DecodeRuneInString(fieldName)
	if r == utf8.RuneError && size == 0 {
		return ""
	}

	return string(unicode.ToLower(r)) + fieldName[size:]
}
