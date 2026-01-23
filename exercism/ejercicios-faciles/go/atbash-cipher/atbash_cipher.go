package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	var contador int
	var builder strings.Builder
	for _, caracter := range s {
		if unicode.IsLetter(caracter) {
			if contador > 0 && contador%5 == 0 {
				builder.WriteString(" ")
			}
			pos := (caracter - 'a')
			invertPos := 25 - pos
			caracterInvertido := invertPos + 'a'
			builder.WriteString(string(caracterInvertido))
			contador++
		}
		if unicode.IsDigit(caracter) {
			if contador > 0 && contador%5 == 0 {
				builder.WriteString(" ")

			}
			builder.WriteString(string(caracter))

			contador++

		}
	}
	return builder.String()
}
