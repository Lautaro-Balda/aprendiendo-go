package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("El numero es inválido")
	}
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else if n%2 != 0 {
			n = n*3 + 1
		}
		steps++
	}
	return steps, nil
}
