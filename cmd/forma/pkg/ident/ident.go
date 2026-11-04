package ident

import (
	"strings"
	"unicode"
)

// ToPascal converts kebab-case or snake_case into PascalCase.
func ToPascal(kebabOrSnake string) string {
	parts := strings.FieldsFunc(kebabOrSnake, func(r rune) bool {
		return r == '-' || r == '_'
	})
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		if len(p) > 1 {
			b.WriteString(p[1:])
		}
	}
	return b.String()
}

// ToExportedGoName maps a TS/JS prop name to an exported Go field name.
func ToExportedGoName(name string) string {
	if strings.HasPrefix(name, "on") && len(name) > 2 {
		return "On" + strings.ToUpper(name[2:3]) + name[3:]
	}
	if name == "" {
		return name
	}
	r := []rune(name)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// ReceiverName returns a single-letter value receiver name for a type.
func ReceiverName(typeName string) string {
	if typeName == "" {
		return "v"
	}
	ch := strings.ToLower(typeName[:1])
	if ch == typeName[:1] {
		return "v"
	}
	return ch
}
