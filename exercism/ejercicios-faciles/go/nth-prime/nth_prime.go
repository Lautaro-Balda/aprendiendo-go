package prime

import (
	"errors"
	"math"
)

// esPrimo verifica si un número es primo
func esPrimo(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	limite := int(math.Sqrt(float64(num)))
	for i := 3; i <= limite; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("there is no zeroth prime")
	}

	contadorPrimos := 0
	candidato := 2

	for {
		if esPrimo(candidato) {
			contadorPrimos++
			if contadorPrimos == n {
				return candidato, nil
			}
		}
		candidato++
	}
}
