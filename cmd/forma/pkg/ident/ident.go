package ident

import (
	"strings"
	"unicode"

	"github.com/zhulik/forma/fn"
)

// ToPascal converts kebab-case or snake_case into PascalCase.
func ToPascal(kebabOrSnake string) string {
	parts := strings.FieldsFunc(kebabOrSnake, func(r rune) bool {
		return r == '-' || r == '_'
	})
	return strings.Join(fn.Map(fn.Filter(parts, nonEmpty), capitalizePart), "")
}

func nonEmpty(p string) bool {
	return p != ""
}

func capitalizePart(p string) string {
	if len(p) == 1 {
		return strings.ToUpper(p)
	}
	return strings.ToUpper(p[:1]) + p[1:]
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
