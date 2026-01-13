package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newMap := make(map[string]int)
	for points, letras := range in {
		for _, letter := range letras {
			newMap[strings.ToLower(letter)] = points

		}

	}
	return newMap
}
