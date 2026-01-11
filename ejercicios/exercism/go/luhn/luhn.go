package luhn

import (
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	for _, num := range id {
		if num < '0' || num > '9' {
			return false
		}
	}

	b := []byte(id)

	for i, num := range b {
		b[i] = num - '0'
	}

	if len(b)%2 == 0 {
		for i := 0; i < len(b); i += 2 {
			num := int(b[i]) * 2
			if num > 9 {
				num -= 9
			}
			b[i] = byte(num)
		}
	} else {
		for i := 1; i < len(b); i += 2 {
			num := int(b[i]) * 2
			if num > 9 {
				num -= 9
			}
			b[i] = byte(num)
		}
	}

	var resultado int
	for _, num := range b {
		resultado += int(num)
	}
	return resultado%10 == 0

}
