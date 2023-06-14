package functional

func MapSlice[T, K any](slice []T, function func(element T, index int) K) []K {
	var mappedElements []K

	for index, element := range slice {
		mappedElements = append(mappedElements, function(element, index))
	}

	return mappedElements
}
