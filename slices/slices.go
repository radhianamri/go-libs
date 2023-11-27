package slices

import (
	"github.com/radhianamri/go-libs/generic"
	"github.com/radhianamri/go-libs/maps"
)

// remove multiple values from a given input list
func RemoveElementMulti[T comparable](input []T, elements []T) []T {
	elementsToRemoveMap := make(map[T]bool)
	for _, element := range elements {
		elementsToRemoveMap[element] = true
	}

	var result []T
	for _, element := range input {
		if !elementsToRemoveMap[element] {
			result = append(result, element)
		}
	}
	return result
}

// remove values for multiple indices from a given input list
func RemoveElement[T comparable](input []T, elem T) []T {
	for i := range input {
		if input[i] == elem {
			return RemoveIndex(input, i)
		}
	}
	return input
}

// remove values for multiple indices from a given input list
func RemoveIndexMulti[T comparable](input []T, indices []int) []T {
	indicesToRemoveMap := make(map[int]bool)
	for _, index := range indices {
		indicesToRemoveMap[index] = true
	}

	var result []T
	for index, value := range input {
		if !indicesToRemoveMap[index] {
			result = append(result, value)
		}
	}
	return result
}

// remove value for a single index from a given input list
func RemoveIndex[T comparable](input []T, index int) []T {
	if index >= len(input) {
		return input
	}
	return append(input[:index], input[index+1:]...)
}

func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// return true if elem exists in the given input list
func Contains[T comparable](input []T, elem T) bool {
	for i := range input {
		if input[i] == elem {
			return true
		}
	}
	return false
}

// return true if any of the elem from a list exists in the given input list
func ContainsAny[T comparable](input []T, elements []T) bool {
	elementsMap := make(map[T]bool)
	for _, element := range elements {
		elementsMap[element] = true
	}

	for _, element := range input {
		if elementsMap[element] {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Unique[T comparable](slice []T) []T {
	var result []T
	encountered := map[T]bool{}
	for _, v := range slice {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

func Merge[T comparable](s1, s2 []T) []T {
	set := make(map[T]bool)
	for _, v := range s1 {
		set[v] = true
	}

	for _, v := range s2 {
		if !set[v] {
			s1 = append(s1, v)
		}
	}

	return s1
}

func Intersect[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	duplicates := make(map[T]bool)
	for _, value := range slices[0] {
		duplicates[value] = true
	}

	for _, slice := range slices[1:] {
		current := make(map[T]bool)
		for _, value := range slice {
			if duplicates[value] {
				current[value] = true
			}
		}
		duplicates = current
	}

	return maps.KeysToSlice(duplicates)
}

func RemoveEmpty[T comparable](values []T) []T {
	var filteredValues []T
	for _, value := range values {
		if generic.IsEmpty(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return filteredValues
}
