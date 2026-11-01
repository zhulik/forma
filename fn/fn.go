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
