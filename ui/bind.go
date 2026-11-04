package ui

import (
	"fmt"
	"reflect"

	"github.com/zhulik/forma/fn"
)

func contextIfaceType() reflect.Type {
	return reflect.TypeFor[Context]()
}

// bindTree copies node and rewrites function Props into host-callable closures
// that inject ctx as the first argument.
func bindTree(node *Node, ctx Context) *Node {
	if node == nil {
		return nil
	}
	out := *node
	out.Props = bindProps(node.Props, ctx)
	out.Children = fn.Map(node.Children, func(child *Node) *Node {
		return bindTree(child, ctx)
	})
	if node.Slots != nil {
		out.Slots = fn.MapValues(node.Slots, func(nodes []*Node) []*Node {
			return fn.Map(nodes, func(n *Node) *Node {
				return bindTree(n, ctx)
			})
		})
	}

	return &out
}

func bindProps(props any, ctx Context) any {
	if props == nil {
		return nil
	}
	value := reflect.ValueOf(props)
	for value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return props
		}
		value = value.Elem()
	}
	if value.Kind() == reflect.Map {
		return bindPropsMap(value, ctx)
	}
	if value.Kind() != reflect.Struct {
		return props
	}

	out := make(map[string]any, value.NumField())
	typ := value.Type()
	for i := range value.NumField() {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldVal := value.Field(i)
		if fieldVal.Kind() == reflect.Func {
			if fieldVal.IsNil() {
				continue
			}
			out[field.Name] = wrapHandler(fieldVal, ctx).Interface()

			continue
		}
		out[field.Name] = fieldVal.Interface()
	}

	return out
}

func bindPropsMap(value reflect.Value, ctx Context) any {
	out := make(map[string]any, value.Len())
	for _, key := range value.MapKeys() {
		keyStr, ok := reflect.TypeAssert[string](key)
		if !ok {
			continue
		}
		fieldVal := value.MapIndex(key)
		if !fieldVal.IsValid() {
			continue
		}
		if fieldVal.Kind() == reflect.Interface {
			fieldVal = fieldVal.Elem()
		}
		if fieldVal.Kind() == reflect.Func {
			if fieldVal.IsNil() {
				continue
			}
			out[keyStr] = wrapHandler(fieldVal, ctx).Interface()

			continue
		}
		out[keyStr] = fieldVal.Interface()
	}

	return out
}

func wrapHandler(fnVal reflect.Value, ctx Context) reflect.Value {
	funcType := fnVal.Type()
	if funcType.NumIn() < 1 || !isContextType(funcType.In(0)) {
		panic(fmt.Sprintf("forma: callback %s must take ui.Context as first argument", funcType))
	}

	numPayload := funcType.NumIn() - 1
	inTypes := make([]reflect.Type, numPayload)
	for i := range numPayload {
		inTypes[i] = funcType.In(i + 1)
	}
	outTypes := make([]reflect.Type, funcType.NumOut())
	for i := range funcType.NumOut() {
		outTypes[i] = funcType.Out(i)
	}
	hostType := reflect.FuncOf(inTypes, outTypes, funcType.IsVariadic())

	return reflect.MakeFunc(hostType, func(args []reflect.Value) []reflect.Value {
		callArgs := make([]reflect.Value, 0, len(args)+1)
		callArgs = append(callArgs, reflect.ValueOf(ctx))
		callArgs = append(callArgs, args...)

		return fnVal.Call(callArgs)
	})
}

func isContextType(typ reflect.Type) bool {
	iface := contextIfaceType()
	if typ == iface {
		return true
	}

	return typ.Kind() == reflect.Interface && typ.Implements(iface)
}
