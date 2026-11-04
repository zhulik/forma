package main

import (
	"reflect"
	"syscall/js"
)

func main() {
	var got string
	author := func(val string) {
		got = val
	}
	fnVal := reflect.ValueOf(author)
	hostType := reflect.FuncOf([]reflect.Type{reflect.TypeFor[string]()}, nil, false)
	wrapped := reflect.MakeFunc(hostType, func(args []reflect.Value) []reflect.Value {
		return fnVal.Call(args)
	})

	jsFn := js.FuncOf(func(_ js.Value, args []js.Value) any {
		in := []reflect.Value{reflect.ValueOf(args[0].String())}
		wrapped.Call(in)

		return nil
	})
	js.Global().Set("spikeCall", jsFn)
	js.Global().Set("spikeRead", js.FuncOf(func(js.Value, []js.Value) any {
		return got
	}))
}
