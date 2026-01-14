package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var contieneLetras bool
	for _, val := range remark {
		if unicode.IsLetter(val) {
			contieneLetras = true
			break
		}
	}
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case remark[len(remark)-1] == '?' && remark == strings.ToUpper(remark) && contieneLetras:
		return "Calm down, I know what I'm doing!"
	case remark[len(remark)-1] == '?':
		return "Sure."
	case remark == strings.ToUpper(remark) && contieneLetras:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}

}
