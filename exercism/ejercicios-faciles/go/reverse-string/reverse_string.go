package reverse

// primera solución
/*
	func Reverse(input string) string {
		runes := []rune(input)
		newRunes := make([]rune, len(runes))
		for i := 0; i < len(runes); i++ {
			newRunes[i] = runes[len(runes)-1-i]
		}

		return string(newRunes)
	}
*/
// solución favorita
func Reverse(input string) string {
	resultado := []rune(input)

	for i, j := 0, len(resultado)-1; i < j; i, j = i+1, j-1 {
		resultado[i], resultado[j] = resultado[j], resultado[i]
	}
	return string(resultado)

}
