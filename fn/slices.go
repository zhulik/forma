package fn

import (
	"cmp"
	"slices"
)

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

// UniqSorted returns unique values from input sorted in ascending order.
func UniqSorted[T cmp.Ordered](input []T) []T {
	out := Uniq(input)
	slices.Sort(out)

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

// FlatMap maps each input element to a slice and concatenates the results.
func FlatMap[I any, O any](input []I, f func(I) []O) []O {
	result := make([]O, 0, len(input))
	for _, item := range input {
		result = append(result, f(item)...)
	}

	return result
}

// GroupBy groups input elements by key; order within each group matches input order.
func GroupBy[T any, K comparable](input []T, key func(T) K) map[K][]T {
	byKey := make(map[K][]T)
	for _, item := range input {
		k := key(item)
		byKey[k] = append(byKey[k], item)
	}

	return byKey
}

// MapErr maps each element; on the first error it returns nil and that error.
func MapErr[I any, O any](input []I, f func(I) (O, error)) ([]O, error) {
	output := make([]O, len(input))
	for i, item := range input {
		out, err := f(item)
		if err != nil {
			return nil, err
		}
		output[i] = out
	}

	return output, nil
}

// FilterMapErr maps each element with an optional keep flag; on the first error it returns nil and that error.
func FilterMapErr[I any, O any](input []I, f func(I) (O, bool, error)) ([]O, error) {
	result := make([]O, 0, len(input))
	for _, item := range input {
		out, ok, err := f(item)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, out)
		}
	}

	return result, nil
}
