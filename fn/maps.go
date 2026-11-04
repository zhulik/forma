package fn

// MapValues builds a new map with the same keys and values transformed by f.
func MapValues[K comparable, V any, O any](input map[K]V, f func(V) O) map[K]O {
	out := make(map[K]O, len(input))
	for k, v := range input {
		out[k] = f(v)
	}

	return out
}

// MapEntries builds a new map by transforming each key-value pair; later keys overwrite earlier ones.
func MapEntries[K comparable, V any, K2 comparable, V2 any](
	input map[K]V,
	f func(K, V) (K2, V2),
) map[K2]V2 {
	out := make(map[K2]V2, len(input))
	for k, v := range input {
		k2, v2 := f(k, v)
		out[k2] = v2
	}

	return out
}

// FilterEntries returns a map containing only entries for which f returns true.
func FilterEntries[K comparable, V any](input map[K]V, f func(K, V) bool) map[K]V {
	out := make(map[K]V)
	for k, v := range input {
		if f(k, v) {
			out[k] = v
		}
	}

	return out
}

// FilterMapEntries transforms entries and keeps those where ok is true.
func FilterMapEntries[K comparable, V any, K2 comparable, V2 any](
	input map[K]V,
	f func(K, V) (K2, V2, bool),
) map[K2]V2 {
	out := make(map[K2]V2)
	for k, v := range input {
		k2, v2, ok := f(k, v)
		if ok {
			out[k2] = v2
		}
	}

	return out
}

// MapValuesErr is MapValues with a fallible transform; on the first error it returns nil and that error.
func MapValuesErr[K comparable, V any, O any](
	input map[K]V,
	f func(V) (O, error),
) (map[K]O, error) {
	out := make(map[K]O, len(input))
	for k, v := range input {
		o, err := f(v)
		if err != nil {
			return nil, err
		}
		out[k] = o
	}

	return out, nil
}

// MapEntriesErr is MapEntries with a fallible transform; on the first error it returns nil and that error.
func MapEntriesErr[K comparable, V any, K2 comparable, V2 any](
	input map[K]V,
	f func(K, V) (K2, V2, error),
) (map[K2]V2, error) {
	out := make(map[K2]V2, len(input))
	for k, v := range input {
		k2, v2, err := f(k, v)
		if err != nil {
			return nil, err
		}
		out[k2] = v2
	}

	return out, nil
}

// FilterMapEntriesErr is FilterMapEntries with a fallible transform; on the first error it returns nil and that error.
func FilterMapEntriesErr[K comparable, V any, K2 comparable, V2 any](
	input map[K]V,
	f func(K, V) (K2, V2, bool, error),
) (map[K2]V2, error) {
	out := make(map[K2]V2)
	for k, v := range input {
		k2, v2, ok, err := f(k, v)
		if err != nil {
			return nil, err
		}
		if ok {
			out[k2] = v2
		}
	}

	return out, nil
}
