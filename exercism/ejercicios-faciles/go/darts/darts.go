package darts

import "math"

func Score(x, y float64) int {
	distanciaConElCentro := math.Sqrt((0-x)*(0-x) + (0-y)*(0-y))
	switch {
	case distanciaConElCentro <= 1:
		return 10
	case distanciaConElCentro <= 5:
		return 5
	case distanciaConElCentro <= 10:
		return 1
	default:
		return 0
	}
}
