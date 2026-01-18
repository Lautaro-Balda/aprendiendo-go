package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	palabras := strings.FieldsFunc(phrase, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\''
	})
	mapaDePalabras := Frequency{}
	for _, val := range palabras {
		limpia := strings.TrimFunc(strings.ToLower(val), func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		if limpia != "" {
			mapaDePalabras[limpia]++
		}
	}
	return mapaDePalabras
}
