package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_typescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"

	"github.com/zhulik/forma/cmd/forma/pkg/ident"
	"github.com/zhulik/forma/fn"
)

const (
	kindExportStatement      = "export_statement"
	kindExportSpecifier      = "export_specifier"
	kindFunctionDeclaration  = "function_declaration"
	kindClassDeclaration     = "class_declaration"
	kindLexicalDeclaration   = "lexical_declaration"
	kindVariableDeclaration  = "variable_declaration"
	kindVariableDeclarator   = "variable_declarator"
	kindInterfaceDeclaration = "interface_declaration"
	kindTypeAliasDeclaration = "type_alias_declaration"
	kindObjectType           = "object_type"
	kindPropertySignature    = "property_signature"
)

var skipFiles = map[string]struct{}{
	"index.tsx": {},
	"index.ts":  {},
	"utils.ts":  {},
	"utils.tsx": {},
}

// AnalyzeComponentsDir parses eligible TSX files under dir into widgets.
func AnalyzeComponentsDir(componentsDir string) ([]Widget, error) {
	entries, err := listTsx(componentsDir)
	if err != nil {
		return nil, err
	}

	parser := sitter.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(sitter.NewLanguage(tree_sitter_typescript.LanguageTSX())); err != nil {
		return nil, fmt.Errorf("set TSX language: %w", err)
	}

	var widgets []Widget
	for _, path := range entries {
		src, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", path, err)
		}
		tree := parser.Parse(src, nil)
		if tree == nil {
			return nil, fmt.Errorf("parse %s: nil tree", path)
		}
		root := tree.RootNode()
		w, ok := analyzeSource(path, src, root)
		tree.Close()
		if ok {
			widgets = append(widgets, w)
		}
	}
	slices.SortFunc(widgets, func(a, b Widget) int {
		return strings.Compare(a.FileBase, b.FileBase)
	})
	return widgets, nil
}

func listTsx(dir string) ([]string, error) {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", dir)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	out := fn.FilterMap(entries, func(e os.DirEntry) (string, bool) {
		if e.IsDir() {
			return "", false
		}
		name := e.Name()
		if _, skip := skipFiles[name]; skip {
			return "", false
		}
		if !strings.HasSuffix(name, ".tsx") && !strings.HasSuffix(name, ".ts") {
			return "", false
		}
		return filepath.Join(dir, name), true
	})
	slices.Sort(out)
	return out, nil
}

func analyzeSource(path string, src []byte, root *sitter.Node) (Widget, bool) {
	fileBase := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	importPath := "@/components/ui/" + fileBase

	var valueExports []string
	propsByType := map[string][]PropField{}

	for i := uint(0); i < root.NamedChildCount(); i++ {
		child := root.NamedChild(i)
		if child == nil || child.Kind() != kindExportStatement {
			continue
		}
		valueExports = append(valueExports, componentExportNames(findDescendants(child, kindExportSpecifier), src)...)
		decl := child.ChildByFieldName("declaration")
		if decl == nil {
			continue
		}
		switch decl.Kind() {
		case kindFunctionDeclaration, kindClassDeclaration:
			name := identText(decl.ChildByFieldName("name"), src)
			if isComponentExportName(name) {
				valueExports = append(valueExports, name)
			}
		case kindLexicalDeclaration, kindVariableDeclaration:
			valueExports = append(valueExports, componentExportNames(findDescendants(decl, kindVariableDeclarator), src)...)
		case kindInterfaceDeclaration:
			name := identText(decl.ChildByFieldName("name"), src)
			if name != "" {
				propsByType[name] = propsFromObjectType(decl.ChildByFieldName("body"), src)
			}
		case kindTypeAliasDeclaration:
			name := identText(decl.ChildByFieldName("name"), src)
			value := decl.ChildByFieldName("value")
			if name != "" && value != nil && value.Kind() == kindObjectType {
				propsByType[name] = propsFromObjectType(value, src)
			}
		}
	}

	valueExports = fn.UniqSorted(valueExports)
	if len(valueExports) == 0 {
		return Widget{}, false
	}

	rootName := pickRoot(valueExports, fileBase)
	if rootName == "" {
		return Widget{}, false
	}

	prefixed := fn.Filter(valueExports, func(n string) bool {
		return n != rootName && strings.HasPrefix(n, rootName) && len(n) > len(rootName)
	})

	props := propsByType[rootName+"Props"]
	if props == nil {
		props = propsByType["Props"]
	}
	props = mergeProps(props, heuristicProps(string(src)))

	if len(prefixed) == 0 {
		return Widget{
			Kind:       "leaf",
			Name:       rootName,
			FileBase:   fileBase,
			ImportPath: importPath,
			ExportName: rootName,
			Props:      props,
		}, true
	}

	slots := fn.Map(prefixed, func(exportName string) SlotField {
		return SlotField{Field: exportName[len(rootName):], ExportName: exportName}
	})
	slots = fn.Filter(slots, func(s SlotField) bool { return s.Field != "" })
	slices.SortFunc(slots, func(a, b SlotField) int {
		return strings.Compare(a.Field, b.Field)
	})

	return Widget{
		Kind:       "compound",
		Name:       rootName,
		FileBase:   fileBase,
		ImportPath: importPath,
		RootExport: rootName,
		Props:      props,
		Slots:      slots,
	}, true
}

func findDescendants(n *sitter.Node, kind string) []*sitter.Node {
	if n == nil {
		return nil
	}
	var out []*sitter.Node
	var walk func(*sitter.Node)
	walk = func(node *sitter.Node) {
		if node.Kind() == kind {
			out = append(out, node)
		}
		for i := uint(0); i < node.NamedChildCount(); i++ {
			child := node.NamedChild(i)
			if child != nil {
				walk(child)
			}
		}
	}
	walk(n)
	return out
}

func pickRoot(names []string, fileBase string) string {
	pascal := ident.ToPascal(fileBase)
	if slices.Contains(names, pascal) {
		return pascal
	}
	sorted := slices.Clone(names)
	slices.SortFunc(sorted, func(a, b string) int {
		if len(a) != len(b) {
			return len(a) - len(b)
		}
		return strings.Compare(a, b)
	})
	for _, candidate := range sorted {
		others := slices.DeleteFunc(slices.Clone(names), func(n string) bool { return n == candidate })
		if len(others) == 0 || !slices.ContainsFunc(others, func(o string) bool {
			return !strings.HasPrefix(o, candidate)
		}) {
			return candidate
		}
	}
	if len(sorted) > 0 {
		return sorted[0]
	}
	return ""
}

func propsFromObjectType(body *sitter.Node, src []byte) []PropField {
	if body == nil {
		return nil
	}
	var out []PropField
	for i := uint(0); i < body.NamedChildCount(); i++ {
		child := body.NamedChild(i)
		if child == nil || child.Kind() != kindPropertySignature {
			continue
		}
		name := identText(child.ChildByFieldName("name"), src)
		if name == "" || name == "className" || name == "children" || name == "asChild" {
			continue
		}
		typeNode := child.ChildByFieldName("type")
		typeText := ""
		if typeNode != nil {
			typeText = typeNode.Utf8Text(src)
			typeText = strings.TrimPrefix(typeText, ":")
			typeText = strings.TrimSpace(typeText)
		}
		goType := mapTSTypeToGo(name, typeText)
		if goType == "" {
			continue
		}
		out = append(out, PropField{Name: ident.ToExportedGoName(name), GoType: goType})
	}
	slices.SortFunc(out, func(a, b PropField) int {
		return strings.Compare(a.Name, b.Name)
	})
	return out
}

func heuristicProps(text string) []PropField {
	var fields []PropField
	if strings.Contains(text, "variant") {
		fields = append(fields, PropField{Name: "Variant", GoType: "string"})
	}
	if strings.Contains(text, "size") {
		fields = append(fields, PropField{Name: "Size", GoType: "string"})
	}
	if strings.Contains(text, "onClick") || strings.Contains(text, "React.MouseEvent") ||
		strings.Contains(text, "ButtonHTMLAttributes") {
		fields = append(fields, PropField{Name: "OnClick", GoType: "func(ctx ui.Context)"})
	}
	return fields
}

func mergeProps(primary, fallback []PropField) []PropField {
	seen := fn.Associate(primary, func(p PropField) (string, struct{}) {
		return p.Name, struct{}{}
	})
	out := slices.Clone(primary)
	out = append(out, fn.Filter(fallback, func(p PropField) bool {
		_, ok := seen[p.Name]
		return !ok
	})...)
	slices.SortFunc(out, func(a, b PropField) int {
		return strings.Compare(a.Name, b.Name)
	})
	return out
}

func isComponentExportName(name string) bool {
	if name == "" || strings.HasPrefix(name, "_") {
		return false
	}
	r := []rune(name)
	return unicode.IsUpper(r[0])
}

func componentExportNames(nodes []*sitter.Node, src []byte) []string {
	return fn.FilterMap(nodes, func(n *sitter.Node) (string, bool) {
		name := identText(n.ChildByFieldName("name"), src)
		return name, isComponentExportName(name)
	})
}

func mapTSTypeToGo(propName, typeText string) string {
	t := strings.Join(strings.Fields(typeText), " ")
	if strings.HasPrefix(propName, "on") || strings.HasPrefix(propName, "On") {
		return "func(ctx ui.Context)"
	}
	if strings.Contains(t, "=>") || strings.Contains(t, "MouseEventHandler") || strings.Contains(t, "EventHandler") {
		return "func(ctx ui.Context)"
	}
	if t == "boolean" || t == "bool" {
		return "bool"
	}
	if t == "number" {
		return "float64"
	}
	if t == "string" || strings.Contains(t, "|") {
		return "string"
	}
	if strings.Contains(t, "ReactNode") {
		return ""
	}
	if !strings.Contains(t, "=>") && t != "" {
		return "string"
	}
	return ""
}

func identText(n *sitter.Node, src []byte) string {
	if n == nil {
		return ""
	}
	return n.Utf8Text(src)
}
