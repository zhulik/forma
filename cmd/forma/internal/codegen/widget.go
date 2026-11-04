package codegen

import "github.com/zhulik/forma/ui"

// PropField is a Go struct field derived from a TSX props type.
type PropField struct {
	Name   string
	GoType string
}

// SlotField is a compound widget slot (Go field + React export).
type SlotField struct {
	Field      string
	ExportName string
}

// WireKey returns the render-tree / registry slot key for this field (Trigger → trigger).
func (s SlotField) WireKey() string {
	return ui.SlotKeyFromField(s.Field)
}

// Widget is one analyzed shadcn component file.
type Widget struct {
	Kind       string // "leaf" or "compound"
	Name       string
	FileBase   string
	ImportPath string
	ExportName string // leaf
	RootExport string // compound
	Props      []PropField
	Slots      []SlotField
}
