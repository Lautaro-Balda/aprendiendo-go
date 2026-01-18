package lsproduct

import (
	"errors"
)

var ErrInvalidSpan = errors.New("span inválido")
var ErrInvalidDigits = errors.New("digitos inválidos")

func LargestSeriesProduct(digits string, span int) (int64, error) {
	// validando el span
	if span > len(digits) || span <= 0 {
		return 0, ErrInvalidSpan
	}
	// validando que digits sean todos numeros.
	for _, val := range digits {
		if val < '0' || val > '9' {
			return 0, ErrInvalidDigits
		}
	}

	var serieMayor int64 = 0
	for i := 0; i <= len(digits)-span; i++ {
		serie := digits[i : i+span]
		producto := int64(1)
		for _, val := range serie {
			producto *= int64(val - '0')
		}
		if serieMayor < producto {
			serieMayor = producto
		}
	}
	return serieMayor, nil
}
