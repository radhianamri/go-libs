package maps

func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V, len(m1)) // preallocate memory from m1 (make sure it has biggest length)
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

func Intersect[K comparable, V any](m1, m2 map[K]V) map[K]V {
	set := make(map[K]bool)
	for k, _ := range m1 {
		set[k] = true
	}

	result := make(map[K]V)
	for k, v := range m2 {
		if set[k] {
			result[k] = v
		}
	}
	return result
}

func KeysToSlice[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for key := range m {
		result = append(result, key)
	}
	return result
}

func ValuesToSlice[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, value := range m {
		result = append(result, value)
	}
	return result
}

func FindDuplicates[K comparable, V any](maps ...map[K]V) []K {
	if len(maps) == 0 {
		return nil
	}

	duplicates := make(map[K]bool, len(maps[0]))
	for key := range maps[0] {
		duplicates[key] = true
	}

	for _, slice := range maps[1:] {
		current := make(map[K]bool)
		for key := range slice {
			if duplicates[key] {
				current[key] = true
			}
		}
		duplicates = current
	}

	return KeysToSlice(duplicates)
}
