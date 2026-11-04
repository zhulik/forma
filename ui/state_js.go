//go:build js

package ui

import (
	"fmt"
	"reflect"

	"syscall/js"

	"github.com/zhulik/forma/fn"
)

// UIMount registers root against the host seam and performs an initial refresh.
// The mount element id is owned by the React host; Go only drives formaHost.setTree.
func UIMount(_ string, root Component) { //nolint:revive // exported naming
	sink := jsHostSink{}
	_, _ = MountWithSink(sink, root)
}

type jsHostSink struct{}

func (jsHostSink) SetTree(tree *Node) {
	host := js.Global().Get("formaHost")
	if host.IsUndefined() || host.IsNull() {
		panic("forma: globalThis.formaHost is not registered; mount the host before the Go bundle")
	}
	setTree := host.Get("setTree")
	if setTree.IsUndefined() || setTree.IsNull() {
		panic("forma: formaHost.setTree is missing")
	}
	// GoScript passes Go pointers as VariableRef; convert to a plain JS object
	// graph (maps/slices/js.Func) so the React host can read Type/Props/Children.
	setTree.Invoke(nodeToHost(tree))
}

func nodeToHost(node *Node) any {
	if node == nil {
		return nil
	}
	out := map[string]any{
		"Type":     node.Type,
		"Key":      node.Key,
		"Class":    node.Class,
		"Props":    propsToHost(node.Props),
		"Children": fn.Map(node.Children, nodeToHost),
	}
	if node.Slots != nil {
		out["Slots"] = fn.MapValues(node.Slots, func(nodes []*Node) any {
			return fn.Map(nodes, nodeToHost)
		})
	}

	return out
}

func propsToHost(props any) any {
	if props == nil {
		return nil
	}
	value := reflect.ValueOf(props)
	for value.Kind() == reflect.Interface {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
	}
	if value.Kind() == reflect.Map {
		out := make(map[string]any, value.Len())
		for _, key := range value.MapKeys() {
			keyStr, ok := reflect.TypeAssert[string](key)
			if !ok {
				continue
			}
			field := value.MapIndex(key)
			if !field.IsValid() {
				continue
			}
			for field.Kind() == reflect.Interface {
				if field.IsNil() {
					break
				}
				field = field.Elem()
			}
			if field.Kind() == reflect.Func {
				out[keyStr] = wrapJSFunc(field)

				continue
			}
			out[keyStr] = field.Interface()
		}

		return out
	}
	if value.Kind() == reflect.Func {
		return wrapJSFunc(value)
	}

	return props
}

func wrapJSFunc(fnVal reflect.Value) js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) any {
		in := make([]reflect.Value, fnVal.Type().NumIn())
		for i := range in {
			if i >= len(args) {
				in[i] = reflect.Zero(fnVal.Type().In(i))

				continue
			}
			in[i] = jsArgToReflect(args[i], fnVal.Type().In(i))
		}
		out := fnVal.Call(in)
		if len(out) == 0 {
			return nil
		}

		return out[0].Interface()
	})
}

func jsArgToReflect(arg js.Value, typ reflect.Type) reflect.Value {
	switch typ.Kind() {
	case reflect.String:
		return reflect.ValueOf(arg.String()).Convert(typ)
	case reflect.Bool:
		return reflect.ValueOf(arg.Bool()).Convert(typ)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(int64(arg.Int())).Convert(typ)
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(arg.Float()).Convert(typ)
	default:
		panic(fmt.Sprintf("forma: unsupported js callback arg kind %s", typ.Kind()))
	}
}
