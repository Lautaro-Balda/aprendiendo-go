package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	result := []string{}

	subjectRunas := []rune(strings.ToLower(subject))
	slices.Sort(subjectRunas)

	for _, candidato := range candidates {
		candidatoRunas := []rune(strings.ToLower(candidato))
		if string(candidatoRunas) == strings.ToLower(subject) {
			continue
		}
		slices.Sort(candidatoRunas)
		if string(subjectRunas) == string(candidatoRunas) {
			result = append(result, candidato)
		}
	}
	return result
}
