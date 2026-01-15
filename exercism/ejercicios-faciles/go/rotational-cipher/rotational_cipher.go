package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	var nuevoTexto = strings.Map(func(r rune) rune {
		minuscula := r >= 'a' && r <= 'z'
		mayuscula := r >= 'A' && r <= 'Z'
		if minuscula {
			return ((r-'a'+rune(shiftKey))%26 + 'a')
		}
		if mayuscula {
			return ((r-'A'+rune(shiftKey))%26 + 'A')
		}
		return r
	}, plain)
	return nuevoTexto
}
