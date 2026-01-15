package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "0", errors.New("Numero fuera de rango.")
	}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result string
	for i := range values {
		for input >= values[i] {
			result += letters[i]
			input -= values[i]
		}
	}
	return result, nil
}
