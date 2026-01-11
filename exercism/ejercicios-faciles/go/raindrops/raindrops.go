package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	div3 := number%3 == 0
	div5 := number%5 == 0
	div7 := number%7 == 0
	var result string
	if div3 {
		result += "Pling"
	}
	if div5 {
		result += "Plang"
	}
	if div7 {
		result += "Plong"
	}
	if !div3 && !div5 && !div7 {
		result = strconv.Itoa(number)
	}
	return result

}
