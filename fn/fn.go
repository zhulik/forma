package fn

func Map[I any, O any](input []I, f func(I) O) []O {
	output := make([]O, len(input))
	for i, item := range input {
		output[i] = f(item)
	}

	return output
}

func Filter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, item := range input {
		if f(item) {
			result = append(result, item)
		}
	}

	return result
}

func FilterMap[I any, O any](input []I, f func(I) (O, bool)) []O {
	result := make([]O, 0, len(input))
	for _, item := range input {
		out, ok := f(item)
		if ok {
			result = append(result, out)
		}
	}

	return result
}

// Uniq returns input with duplicate values removed, preserving first occurrence order.
func Uniq[T comparable](input []T) []T {
	seen := make(map[T]struct{}, len(input))
	out := make([]T, 0, len(input))
	for _, item := range input {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}

	return out
}

// Associate builds map[K]V from input; f runs per element. Later keys overwrite earlier ones.
func Associate[T any, K comparable, V any](input []T, f func(T) (K, V)) map[K]V {
	byKey := make(map[K]V, len(input))
	for _, item := range input {
		k, v := f(item)
		byKey[k] = v
	}

	return byKey
}
