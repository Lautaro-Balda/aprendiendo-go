package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("El numero no puede ser 0 o negativo. Tampoco mayor a 65.")
	}

	return uint64(math.Pow(2, float64(number)-1)), nil
}

func Total() uint64 {
	var res uint64
	for num := range 64 {
		resultado, _ := Square(num + 1)
		res += resultado
	}
	return res
}
