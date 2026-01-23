package cipher

import (
	"math/rand/v2"
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
type shift int
type vigenere string

// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	var desplazamiento shift = 3
	return desplazamiento
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	result := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return ((r - 'a' + rune(c) + 26) % 26) + 'a'
		}
		return -1

	}, input)

	return result
}

func (c shift) Decode(input string) string {
	result := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return ((r - 'a' - rune(c) + 26) % 26) + 'a'
		}

		return -1

	}, input)

	return result
}

func NewVigenere(key string) Cipher {
	// verificamos que la llave no esté vacía. si lo está, devolvemos una llave random de al menos 100 caracteres.
	if key == "" {
		return nil
	}
	var valorDeLlave int
	// verificacion de si la llave está completamente hecha de 'a'.
	for _, char := range key {
		valorDeLlave += int(char - 'a')
	}
	if valorDeLlave == 0 {
		return nil
	}
	// verificacion de si la llave entera está compuesta de letras válidas.
	for _, rune := range key {
		if !unicode.IsLetter(rune) || (rune >= 'A' && rune <= 'Z') {
			return nil
		}
	}

	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	if v == "" {
		const letras = "abcdefghijklmnopqrstuvwxyz"
		resultado := make([]byte, 100)
		for i := range resultado {
			resultado[i] = letras[rand.IntN(len(letras))]
		}
		v = vigenere(resultado)
	}
	var j int
	resultado := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			desplazamiento := int(v[j%len(v)] - 'a')
			letraCifrada := ((r - 'a' + rune(desplazamiento) + 26) % 26) + 'a'
			j++
			return letraCifrada
		}
		return -1

	}, input)
	return string(resultado)

}

func (v vigenere) Decode(input string) string {
	var resultado []rune
	var j int
	for _, r := range input {
		desplazamiento := int(v[j%len(v)] - 'a')
		letraCifrada := ((r - 'a' - rune(desplazamiento) + 26) % 26) + 'a'
		resultado = append(resultado, letraCifrada)
		j++
	}
	return string(resultado)

}
