package utils

import "slices"

func SortMapKeys(m map[string]any, sortedKeys []string) {
	i := 0
	for key := range m {
		sortedKeys[i] = key
		i++
	}
	slices.Sort(sortedKeys)
}
