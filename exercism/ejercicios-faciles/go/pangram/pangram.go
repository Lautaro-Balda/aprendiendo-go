package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	// Mapa para EXTRAER las letras del input.
	letras := make(map[rune]bool)

	input = strings.ToLower(input)
	for _, value := range input {
		// comprobamos que NO esté en el mapa y la añadimos.
		if !letras[value] && unicode.IsLetter(value) {
			letras[value] = true
			if len(letras) == 26 {
				return true
			}
		}
	}
	return false
}
