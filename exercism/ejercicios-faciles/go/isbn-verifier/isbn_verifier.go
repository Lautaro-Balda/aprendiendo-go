package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	var sum int
	for i := 0; i < len(isbn); i++ {
		if (isbn[i] < '0' || isbn[i] > '9') && i != len(isbn)-1 {
			return false
		}
		if string(isbn[i]) == "X" {
			sum += 10 * 1
			break
		}

		num, _ := strconv.Atoi(string(isbn[i]))
		sum += num * (10 - i)

	}
	return sum%11 == 0
}
