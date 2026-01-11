package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	for _, value := range strings.ToLower(word) {
		if value == ' ' || value == '-' {
			continue
		}
		if seen[value] {
			return false
		}
		seen[value] = true
	}
	return true
}
