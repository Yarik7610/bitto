package utils

import "slices"

func SortMapKeys(m map[string]any) []string {
	sortedKeys := make([]string, len(m))
	i := 0
	for key := range m {
		sortedKeys[i] = key
		i++
	}
	slices.Sort(sortedKeys)
	return sortedKeys
}
